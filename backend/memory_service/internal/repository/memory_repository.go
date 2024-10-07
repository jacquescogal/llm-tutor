package repository
// DONE
import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"memory_core/internal/proto/common"
	mpb "memory_core/internal/proto/memory"
	"time"
)

type MemoryRepository struct{}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{}
}

// CreateMemory inserts a new memory into memory_tab
func (repo *MemoryRepository) CreateMemory(ctx context.Context, db *sql.DB, req *mpb.CreateMemoryRequest) error {
	query := `
		INSERT INTO memory_tab (doc_id, user_id, memory_title, memory_content, created_time, updated_time)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	createdTime := time.Now().Unix()
	_, err := db.ExecContext(ctx, query, req.DocId, req.UserId, req.MemoryTitle, req.MemoryContent, createdTime, createdTime)
	if err != nil {
		log.Printf("Error creating memory: %v\n", err)
		return err
	}
	log.Println("Memory created successfully")
	return nil
}

// GetMemoryById retrieves a memory by memory_id
func (repo *MemoryRepository) GetMemoryById(ctx context.Context, db *sql.DB, memoryID uint64) (*mpb.DBMemory, error) {
	query := `
		SELECT memory_id, doc_id, user_id, memory_title, memory_content, created_time, updated_time
		FROM memory_tab 
		WHERE memory_id = ?
	`

	row := db.QueryRowContext(ctx, query, memoryID)
	var dbMemory mpb.DBMemory
	err := row.Scan(&dbMemory.MemoryId, &dbMemory.DocId, &dbMemory.UserId, &dbMemory.MemoryTitle, &dbMemory.MemoryContent, &dbMemory.CreatedTime, &dbMemory.UpdatedTime)
	if err != nil {
		log.Printf("Error retrieving memory: %v\n", err)
		return nil, err
	}
	log.Println("Memory retrieved successfully")
	return &dbMemory, nil
}

// GetMemoriesByDocId retrieves memories for a specific document
func (repo *MemoryRepository) GetMemoriesByDocId(ctx context.Context, db *sql.DB, docID uint64, pageNumber, pageSize uint32, orderByField common.ORDER_BY_FIELD, orderByDirection common.ORDER_BY_DIRECTION) ([]*mpb.DBMemory, error) {
	offset := pageOffset(pageNumber, pageSize)
	sanitisedOrderByString := repo.generateModuleOrderByString(orderByField, orderByDirection)
	query := fmt.Sprintf(`
		SELECT memory_id, doc_id, user_id, memory_title, memory_content, created_time, updated_time
		FROM memory_tab
		WHERE doc_id = ?
		ORDER BY %s
		LIMIT ? OFFSET ?
	`, sanitisedOrderByString)

	rows, err := db.QueryContext(ctx, query, docID, pageSize, offset)
	if err != nil {
		log.Printf("Error retrieving memories: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var memories []*mpb.DBMemory
	for rows.Next() {
		var dbMemory mpb.DBMemory
		if err := rows.Scan(&dbMemory.MemoryId, &dbMemory.DocId, &dbMemory.UserId, &dbMemory.MemoryTitle, &dbMemory.MemoryContent, &dbMemory.CreatedTime, &dbMemory.UpdatedTime); err != nil {
			log.Printf("Error scanning memory row: %v\n", err)
			return nil, err
		}
		memories = append(memories, &dbMemory)
	}

	log.Printf("Retrieved %d memories\n", len(memories))
	return memories, nil
}


// generateModuleOrderByString generates the ORDER BY string for the module query
func (repo *MemoryRepository) generateModuleOrderByString(orderByField common.ORDER_BY_FIELD, orderByDirection common.ORDER_BY_DIRECTION) string {
	var orderByString string

	switch orderByField {
	case common.ORDER_BY_FIELD_ORDER_BY_FIELD_ID:
		orderByString = "memory_id"
	case common.ORDER_BY_FIELD_ORDER_BY_FIELD_TITLE:
		orderByString = "memory_title"
	case common.ORDER_BY_FIELD_ORDER_BY_FIELD_CREATED_TIME:
		orderByString = "created_time"
	case common.ORDER_BY_FIELD_ORDER_BY_FIELD_UPDATED_TIME:
		orderByString = "updated_time"
	default:
		orderByString = "created_time"
	}

	if orderByDirection == common.ORDER_BY_DIRECTION_ORDER_BY_DIRECTION_ASC {
		orderByString += " ASC"
	} else {
		orderByString += " DESC"
	}

	return orderByString
}
