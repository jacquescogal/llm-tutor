package repository

import (
	"context"
	"log"
	mpb "memory_core/internal/proto/memory"
	"time"
)

type QuestionRepository struct{}

func NewQuestionRepository() *QuestionRepository {
	return &QuestionRepository{}
}

// CreateQuestion inserts a new question into question_tab
func (repo *QuestionRepository) CreateQuestion(ctx context.Context, req *mpb.CreateQuestionRequest) error {
	query := `
		INSERT INTO question_tab (doc_id, question_title, question_blob, question_type, created_time, updated_time)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	createdTime := time.Now().Unix()
	_, err := repo.db.ExecContext(ctx, query, req.DocId, req.QuestionTitle, req.QuestionBlob, req.QuestionType, createdTime, createdTime)
	if err != nil {
		log.Printf("Error creating question: %v\n", err)
		return err
	}
	log.Println("Question created successfully")
	return nil
}

// GetQuestionById retrieves a question by question_id
func (repo *QuestionRepository) GetQuestionById(ctx context.Context, questionID uint64) (*mpb.DBQuestion, error) {
	query := `
		SELECT question_id, doc_id, question_title, question_blob, question_type, created_time, updated_time 
		FROM question_tab 
		WHERE question_id = ?
	`

	row := repo.db.QueryRowContext(ctx, query, questionID)
	var dbQuestion mpb.DBQuestion
	err := row.Scan(&dbQuestion.QuestionId, &dbQuestion.DocId, &dbQuestion.QuestionTitle, &dbQuestion.QuestionBlob, &dbQuestion.QuestionType, &dbQuestion.CreatedTime, &dbQuestion.UpdatedTime)
	if err != nil {
		log.Printf("Error retrieving question: %v\n", err)
		return nil, err
	}
	log.Println("Question retrieved successfully")
	return &dbQuestion, nil
}

// GetQuestionsByDocId retrieves questions for a specific document
func (repo *QuestionRepository) GetQuestionsByDocId(ctx context.Context, docID uint64, pageNumber, pageSize uint32) ([]*mpb.DBQuestion, error) {
	offset := pageOffset(pageNumber, pageSize)
	query := `
		SELECT question_id, doc_id, question_title, question_blob, question_type, created_time, updated_time 
		FROM question_tab 
		WHERE doc_id = ? 
		ORDER BY created_time DESC LIMIT ? OFFSET ? `
	rows, err := repo.db.QueryContext(ctx, query, docID, pageSize, offset)
	if err != nil {
		log.Printf("Error retrieving questions: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var questions []*mpb.DBQuestion
	for rows.Next() {
		var dbQuestion mpb.DBQuestion
		if err := rows.Scan(&dbQuestion.QuestionId, &dbQuestion.DocId, &dbQuestion.QuestionTitle, &dbQuestion.QuestionBlob, &dbQuestion.QuestionType, &dbQuestion.CreatedTime, &dbQuestion.UpdatedTime); err != nil {
			log.Printf("Error scanning question row: %v\n", err)
			return nil, err
		}
		questions = append(questions, &dbQuestion)
	}

	log.Printf("Retrieved %d questions\n", len(questions))
	return questions, nil
}
