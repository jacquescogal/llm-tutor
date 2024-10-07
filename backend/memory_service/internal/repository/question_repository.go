package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"memory_core/internal/proto/common"
	mpb "memory_core/internal/proto/question"
	"time"
)

type QuestionRepository struct{}

func NewQuestionRepository() *QuestionRepository {
	return &QuestionRepository{}
}

// CreateQuestion inserts a new question into question_tab
func (repo *QuestionRepository) CreateQuestion(ctx context.Context, tx *sql.Tx, req *mpb.DBQuestion) error {
	query := `
		INSERT INTO question_tab (doc_id, user_id, question_title, question_blob, question_type, created_time, updated_time)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	createdTime := time.Now().Unix()
	_, err := tx.ExecContext(ctx, query, req.DocId, req.UserId, req.QuestionTitle, req.QuestionBlob, req.QuestionType, createdTime, createdTime)
	if err != nil {
		log.Printf("Error creating question: %v\n", err)
		return err
	}
	log.Println("Question created successfully")
	return nil
}

// GetQuestionById retrieves a question by question_id
func (repo *QuestionRepository) GetQuestionById(ctx context.Context, db *sql.DB, questionID uint64) (*mpb.DBQuestion, error) {
	query := `
		SELECT question_id, doc_id, user_id, question_title, question_blob, question_type, created_time, updated_time
		FROM question_tab 
		WHERE question_id = ?
	`

	row := db.QueryRowContext(ctx, query, questionID)
	var dbQuestion mpb.DBQuestion
	err := row.Scan(&dbQuestion.QuestionId, &dbQuestion.DocId, &dbQuestion.UserId, &dbQuestion.QuestionTitle, &dbQuestion.QuestionBlob, &dbQuestion.QuestionType, &dbQuestion.CreatedTime, &dbQuestion.UpdatedTime)
	if err != nil {
		log.Printf("Error retrieving question: %v\n", err)
		return nil, err
	}
	log.Println("Question retrieved successfully")
	return &dbQuestion, nil
}

// GetQuestionsByDocId retrieves questions for a specific document
func (repo *QuestionRepository) GetQuestionsByDocId(ctx context.Context, db *sql.DB, docID uint64, pageNumber, pageSize uint32, orderByField common.ORDER_BY_FIELD, orderByDirection common.ORDER_BY_DIRECTION) ([]*mpb.DBQuestion, error) {
	offset := pageOffset(pageNumber, pageSize)
	sanitisedOrderByString := repo.generateSubjectOrderByString(orderByField, orderByDirection)
	query := fmt.Sprintf(`
		SELECT question_id, doc_id, user_id, question_title, question_blob, question_type, created_time, updated_time
		FROM question_tab
		WHERE doc_id = ?
		ORDER BY %s
		LIMIT ? OFFSET ?`, sanitisedOrderByString)

	rows, err := db.QueryContext(ctx, query, docID, pageSize, offset)
	if err != nil {
		log.Printf("Error retrieving questions: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var questions []*mpb.DBQuestion
	for rows.Next() {
		var dbQuestion mpb.DBQuestion
		if err := rows.Scan(&dbQuestion.QuestionId, &dbQuestion.DocId, &dbQuestion.UserId, &dbQuestion.QuestionTitle, &dbQuestion.QuestionBlob, &dbQuestion.QuestionType, &dbQuestion.CreatedTime, &dbQuestion.UpdatedTime); err != nil {
			log.Printf("Error scanning question row: %v\n", err)
			return nil, err
		}
		questions = append(questions, &dbQuestion)
	}

	log.Printf("Retrieved %d questions\n", len(questions))
	return questions, nil
}

func (repo *QuestionRepository) generateSubjectOrderByString(orderByField common.ORDER_BY_FIELD, orderByDirection common.ORDER_BY_DIRECTION) string {
	var orderByString string

	switch orderByField {
	case common.ORDER_BY_FIELD_ORDER_BY_FIELD_ID:
		orderByString = "doc_id"
	case common.ORDER_BY_FIELD_ORDER_BY_FIELD_TITLE:
		orderByString = "question_title"
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
