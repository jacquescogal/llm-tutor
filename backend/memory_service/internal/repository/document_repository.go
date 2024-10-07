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
func (repo *DocRepository) CreateDoc(ctx context.Context, tx *sql.Tx, req *mpb.DBDoc) error {
	query := `
		INSERT INTO doc_tab (module_id, doc_name, doc_description, doc_summary, upload_status, s3_object_key, created_time, updated_time)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	createdTime := time.Now().Unix()
	_, err := tx.ExecContext(ctx, query, req.ModuleId, req.DocName, req.DocDescription, req.DocSummary, req.UploadStatus, req.S3ObjectKey, createdTime, createdTime)
	if err != nil {
		log.Printf("Error creating document: %v\n", err)
		return err
	}
	log.Println("Document created successfully")
	return nil
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
	sanitisedOrderByString := repo.generateModuleOrderByString(orderByField, orderByDirection)
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


// generateModuleOrderByString generates the ORDER BY string for the module query
func (repo *DocRepository) generateModuleOrderByString(orderByField common.ORDER_BY_FIELD, orderByDirection common.ORDER_BY_DIRECTION) string {
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
