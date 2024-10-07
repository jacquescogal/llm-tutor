package repository

import (
	"context"
	"log"
	mpb "memory_core/internal/proto/memory"
	"time"
)

type DocRepository struct{}

func NewDocRepository() *DocRepository {
	return &DocRepository{}
}

// CreateDoc inserts a new document into doc_tab
func (repo *DocRepository) CreateDoc(ctx context.Context, req *mpb.CreateDocRequest) error {
	query := `
		INSERT INTO doc_tab (topic_id, doc_title, doc_summary, upload_status, s3_object_key, created_time, updated_time)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	createdTime := time.Now().Unix()
	_, err := repo.db.ExecContext(ctx, query, req.TopicId, req.DocTitle, req.DocSummary, req.UploadStatus, req.S3ObjectKey, createdTime, createdTime)
	if err != nil {
		log.Printf("Error creating document: %v\n", err)
		return err
	}
	log.Println("Document created successfully")
	return nil
}

// GetDocById retrieves a document by doc_id
func (repo *DocRepository) GetDocById(ctx context.Context, docID uint64) (*mpb.DBDoc, error) {
	query := `
		SELECT doc_id, topic_id, doc_title, doc_summary, upload_status, s3_object_key, created_time, updated_time 
		FROM doc_tab 
		WHERE doc_id = ?
	`

	row := repo.db.QueryRowContext(ctx, query, docID)
	var dbDoc mpb.DBDoc
	err := row.Scan(&dbDoc.DocId, &dbDoc.TopicId, &dbDoc.DocTitle, &dbDoc.DocSummary, &dbDoc.UploadStatus, &dbDoc.S3ObjectKey, &dbDoc.CreatedTime, &dbDoc.UpdatedTime)
	if err != nil {
		log.Printf("Error retrieving document: %v\n", err)
		return nil, err
	}
	log.Println("Document retrieved successfully")
	return &dbDoc, nil
}

// GetDocsByTopicId retrieves documents for a specific topic
func (repo *DocRepository) GetDocsByTopicId(ctx context.Context, topicID uint64, pageNumber, pageSize uint32) ([]*mpb.DBDoc, error) {
	offset := pageOffset(pageNumber, pageSize)
	query := `
		SELECT doc_id, topic_id, doc_title, doc_summary, upload_status, s3_object_key, created_time, updated_time 
		FROM doc_tab 
		WHERE topic_id = ? 
		ORDER BY created_time DESC
		LIMIT ? OFFSET ?
	`

	rows, err := repo.db.QueryContext(ctx, query, topicID, pageSize, offset)
	if err != nil {
		log.Printf("Error retrieving documents: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var docs []*mpb.DBDoc
	for rows.Next() {
		var dbDoc mpb.DBDoc
		if err := rows.Scan(&dbDoc.DocId, &dbDoc.TopicId, &dbDoc.DocTitle, &dbDoc.DocSummary, &dbDoc.UploadStatus, &dbDoc.S3ObjectKey, &dbDoc.CreatedTime, &dbDoc.UpdatedTime); err != nil {
			log.Printf("Error scanning document row: %v\n", err)
			return nil, err
		}
		docs = append(docs, &dbDoc)
	}

	log.Printf("Retrieved %d documents\n", len(docs))
	return docs, nil
}
