package repository
// DONE
import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"memory_core/internal/proto/common"
	mpb "memory_core/internal/proto/document"
	"time"
)

type DocRepository struct{}

func NewDocRepository() *DocRepository {
	return &DocRepository{}
}

// CreateDoc inserts a new document into doc_tab
func (repo *DocRepository) CreateDoc(ctx context.Context, tx *sql.Tx, moduleID uint64, docName, docDescription, docSummary string, uploadStatus common.UploadStatus, s3ObjectKey string) (uint64, error) {
	query := `
		INSERT INTO doc_tab (module_id, doc_name, doc_description, doc_summary, upload_status, s3_object_key, created_time, updated_time)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	createdTime := time.Now().Unix()
	result, err := tx.ExecContext(ctx, query, moduleID, docName, docDescription, docSummary, uploadStatus, s3ObjectKey, createdTime, createdTime)
	if err != nil {
		log.Printf("Error creating document: %v\n", err)
		return 0, err
	}
	
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID: %v\n", err)
		return 0, err
	}

	log.Printf("Document created successfully with ID: %d\n", lastInsertID)
	return uint64(lastInsertID), nil
}

// GetDocById retrieves a document by doc_id
func (repo *DocRepository) GetDocById(ctx context.Context, db *sql.DB, docID uint64) (*mpb.DBDoc, error) {
	query := `
		SELECT doc_id, module_id, doc_name, doc_description, doc_summary, upload_status, s3_object_key, created_time, updated_time
		FROM doc_tab 
		WHERE doc_id = ?
	`

	row := db.QueryRowContext(ctx, query, docID)
	var dbDoc mpb.DBDoc
	err := row.Scan(&dbDoc.DocId, &dbDoc.ModuleId, &dbDoc.DocName, &dbDoc.DocDescription, &dbDoc.DocSummary, &dbDoc.UploadStatus, &dbDoc.S3ObjectKey, &dbDoc.CreatedTime, &dbDoc.UpdatedTime)
	if err != nil {
		log.Printf("Error retrieving document: %v\n", err)
		return nil, err
	}
	log.Println("Document retrieved successfully")
	return &dbDoc, nil
}

// GetDocsByModuleId retrieves documents for a specific topic
func (repo *DocRepository) GetDocsByModuleId(ctx context.Context, db *sql.DB, moduleId uint64, pageNumber, pageSize uint32, orderByField common.ORDER_BY_FIELD, orderByDirection common.ORDER_BY_DIRECTION) ([]*mpb.DBDoc, error) {
	offset := pageOffset(pageNumber, pageSize)
	sanitisedOrderByString := repo.generateDocOrderByString(orderByField, orderByDirection)
	query := fmt.Sprintf(`
		SELECT doc_id, module_id, doc_name, doc_description, doc_summary, upload_status, s3_object_key, created_time, updated_time
		FROM doc_tab
		WHERE module_id = ?
		ORDER BY %s
		LIMIT ? OFFSET ?
	`, sanitisedOrderByString)

	rows, err := db.QueryContext(ctx, query, moduleId, pageSize, offset)
	if err != nil {
		log.Printf("Error retrieving documents: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var docs []*mpb.DBDoc
	for rows.Next() {
		var dbDoc mpb.DBDoc
		if err := rows.Scan(&dbDoc.DocId, &dbDoc.ModuleId, &dbDoc.DocName, &dbDoc.DocDescription, &dbDoc.DocSummary, &dbDoc.UploadStatus, &dbDoc.S3ObjectKey, &dbDoc.CreatedTime, &dbDoc.UpdatedTime); err != nil {
			log.Printf("Error scanning document row: %v\n", err)
			return nil, err
		}
		docs = append(docs, &dbDoc)
	}

	log.Printf("Retrieved %d documents\n", len(docs))
	return docs, nil
}

// GetDocsByNameSearch
func (repo *DocRepository) GetDocsByNameSearch(ctx context.Context, db *sql.DB, moduleId uint64, search string, pageNumber, pageSize uint32, orderByField common.ORDER_BY_FIELD, orderByDirection common.ORDER_BY_DIRECTION) ([]*mpb.DBDoc, error) {
	offset := pageOffset(pageNumber, pageSize)
	sanitisedOrderByString := repo.generateDocOrderByString(orderByField, orderByDirection)
	query := fmt.Sprintf(`
		SELECT doc_id, module_id, doc_name, doc_description, doc_summary, upload_status, s3_object_key, created_time, updated_time
		FROM doc_tab
		WHERE module_id = ? AND doc_name LIKE ?
		ORDER BY %s
		LIMIT ? OFFSET ?
	`, sanitisedOrderByString)

	rows, err := db.QueryContext(ctx, query, moduleId, "%"+search+"%", pageSize, offset)
	if err != nil {
		log.Printf("Error retrieving documents: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var docs []*mpb.DBDoc
	for rows.Next() {
		var dbDoc mpb.DBDoc
		if err := rows.Scan(&dbDoc.DocId, &dbDoc.ModuleId, &dbDoc.DocName, &dbDoc.DocDescription, &dbDoc.DocSummary, &dbDoc.UploadStatus, &dbDoc.S3ObjectKey, &dbDoc.CreatedTime, &dbDoc.UpdatedTime); err != nil {
			log.Printf("Error scanning document row: %v\n", err)
			return nil, err
		}
		docs = append(docs, &dbDoc)
	}

	log.Printf("Retrieved %d documents\n", len(docs))
	return docs, nil
}

// UpdateDoc
func (repo *DocRepository) UpdateDoc(ctx context.Context, tx *sql.Tx, docID uint64, docName, docDescription, docSummary string, uploadStatus common.UploadStatus, s3ObjectKey string) error {
	query := `
		UPDATE doc_tab
		SET doc_name = ?, doc_description = ?, doc_summary = ?, upload_status = ?, s3_object_key = ?, updated_time = ?
		WHERE doc_id = ?
	`
	updatedTime := time.Now().Unix()
	_, err := tx.ExecContext(ctx, query, docName, docDescription, docSummary, uploadStatus, s3ObjectKey, updatedTime, docID)
	if err != nil {
		log.Printf("Error updating document: %v\n", err)
		return err
	}
	log.Println("Document updated successfully")
	return nil
}

// UpdateSummary
func (repo *DocRepository) UpdateSummary(ctx context.Context, tx *sql.Tx, docID uint64, docSummary string) error {
	query := `
		UPDATE doc_tab
		SET doc_summary = ?, updated_time = ?
		WHERE doc_id = ?
	`
	updatedTime := time.Now().Unix()
	_, err := tx.ExecContext(ctx, query, docSummary, updatedTime, docID)
	if err != nil {
		log.Printf("Error updating document summary: %v\n", err)
		return err
	}
	log.Println("Document summary updated successfully")
	return nil
}

// UpdateUploadStatus
func (repo *DocRepository) UpdateUploadStatus(ctx context.Context, tx *sql.Tx, docID uint64, uploadStatus common.UploadStatus) error {
	query := `
		UPDATE doc_tab
		SET upload_status = ?, updated_time = ?
		WHERE doc_id = ?
	`
	updatedTime := time.Now().Unix()
	_, err := tx.ExecContext(ctx, query, uploadStatus, updatedTime, docID)
	if err != nil {
		log.Printf("Error updating document upload status: %v\n", err)
		return err
	}
	log.Println("Document upload status updated successfully")
	return nil
}

// DeleteDoc
func (repo *DocRepository) DeleteDoc(ctx context.Context, tx *sql.Tx, docID uint64) error {
	query := `
		DELETE FROM doc_tab
		WHERE doc_id = ?
	`
	_, err := tx.ExecContext(ctx, query, docID)
	if err != nil {
		log.Printf("Error deleting document: %v\n", err)
		return err
	}
	log.Println("Document deleted successfully")
	return nil
}

// generateDocOrderByString generates the ORDER BY string for the doc query
func (repo *DocRepository) generateDocOrderByString(orderByField common.ORDER_BY_FIELD, orderByDirection common.ORDER_BY_DIRECTION) string {
	var orderByString string

	switch orderByField {
	case common.ORDER_BY_FIELD_ORDER_BY_FIELD_ID:
		orderByString = "doc_id"
	case common.ORDER_BY_FIELD_ORDER_BY_FIELD_TITLE:
		orderByString = "doc_name"
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

// helper function to get the parent document
func (repo *DocRepository) GetParentModuleId(ctx context.Context, db *sql.DB, docID uint64) (uint64, error) {
	query := `
		SELECT module_id
		FROM doc_tab
		WHERE doc_id = ?
	`

	row := db.QueryRowContext(ctx, query, docID)
	var moduleId uint64
	err := row.Scan(&moduleId)
	if err != nil {
		log.Printf("Error retrieving module_id: %v\n", err)
		return 0, err
	}
	log.Println("ModuleId retrieved successfully")
	return moduleId, nil
}