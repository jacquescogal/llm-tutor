package repository

import (
	"context"
	"log"
	mpb "memory_core/internal/proto/memory"
	"time"
)

type MemoryRepository struct{}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{}
}

// CreateMemory inserts a new memory into memory_tab
func (repo *MemoryRepository) CreateMemory(ctx context.Context, req *mpb.CreateMemoryRequest) error {
	query := `
		INSERT INTO memory_tab (doc_id, memory_title, memory_content, created_time, updated_time)
		VALUES (?, ?, ?, ?, ?)
	`
	createdTime := time.Now().Unix()
	_, err := repo.db.ExecContext(ctx, query, req.DocId, req.MemoryTitle, req.MemoryContent, createdTime, createdTime)
	if err != nil {
		log.Printf("Error creating memory: %v\n", err)
		return err
	}
	log.Println("Memory created successfully")
	return nil
}

// GetMemoryById retrieves a memory by memory_id
func (repo *MemoryRepository) GetMemoryById(ctx context.Context, memoryID uint64) (*mpb.DBMemory, error) {
	query := `
		SELECT memory_id, doc_id, memory_title, memory_content, created_time, updated_time 
		FROM memory_tab 
		WHERE memory_id = ?
	`

	row := repo.db.QueryRowContext(ctx, query, memoryID)
	var dbMemory mpb.DBMemory
	err := row.Scan(&dbMemory.MemoryId, &dbMemory.DocId, &dbMemory.MemoryTitle, &dbMemory.MemoryContent, &dbMemory.CreatedTime, &dbMemory.UpdatedTime)
	if err != nil {
		log.Printf("Error retrieving memory: %v\n", err)
		return nil, err
	}
	log.Println("Memory retrieved successfully")
	return &dbMemory, nil
}

// GetMemoriesByDocId retrieves memories for a specific document
func (repo *MemoryRepository) GetMemoriesByDocId(ctx context.Context, docID uint64, pageNumber, pageSize uint32) ([]*mpb.DBMemory, error) {
	offset := pageOffset(pageNumber, pageSize)
	query := `
		SELECT memory_id, doc_id, memory_title, memory_content, created_time, updated_time 
		FROM memory_tab 
		WHERE doc_id = ? 
		ORDER BY created_time DESC
		LIMIT ? OFFSET ?
	`

	rows, err := repo.db.QueryContext(ctx, query, docID, pageSize, offset)
	if err != nil {
		log.Printf("Error retrieving memories: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var memories []*mpb.DBMemory
	for rows.Next() {
		var dbMemory mpb.DBMemory
		if err := rows.Scan(&dbMemory.MemoryId, &dbMemory.DocId, &dbMemory.MemoryTitle, &dbMemory.MemoryContent, &dbMemory.CreatedTime, &dbMemory.UpdatedTime); err != nil {
			log.Printf("Error scanning memory row: %v\n", err)
			return nil, err
		}
		memories = append(memories, &dbMemory)
	}

	log.Printf("Retrieved %d memories\n", len(memories))
	return memories, nil
}
