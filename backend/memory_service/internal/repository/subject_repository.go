package repository

// DONE
import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"memory_core/internal/proto/common"
	mpb "memory_core/internal/proto/subject"
	"time"
)

type SubjectRepository struct{}

func NewSubjectRepository() *SubjectRepository {
	return &SubjectRepository{}
}

// CreateSubject inserts a new subject into the subject_tab and returns the subject_id
func (repo *SubjectRepository) CreateSubject(ctx context.Context, tx *sql.Tx, subjectName, subjectDescription string, isPublic bool) (uint64, error) {
	query := `
		INSERT INTO subject_tab (subject_name, subject_description, is_public, created_time, updated_time)
		VALUES (?, ?, ?, ?, ?)
	`
	createdTime := time.Now().Unix()
	result, err := tx.ExecContext(ctx, query, subjectName, subjectDescription, isPublic, createdTime, createdTime)
	if err != nil {
		log.Printf("Error creating subject: %v\n", err)
		return 0, err
	}
	// return the last inserted id
	subjectID, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last inserted id: %v\n", err)
		return 0, err
	}
	return uint64(subjectID), nil
}

// GetPublicSubjects retrieves all public subjects
func (repo *SubjectRepository) GetPublicSubjects(ctx context.Context, db *sql.DB, userId uint64, pageNumber, pageSize uint32, orderByField common.ORDER_BY_FIELD, orderByDirection common.ORDER_BY_DIRECTION) ([]*mpb.FullSubject, error) {
	offset := pageOffset(pageNumber, pageSize)
	sanitisedOrderByString := repo.generateSubjectOrderByString(orderByField, orderByDirection)
	query := fmt.Sprintf(`
		SELECT s.subject_id, s.subject_name, s.subject_description, s.is_public, s.created_time, s.updated_time,
		COALESCE(ma.user_subject_role, 0), COALESCE(ma.is_favourite, 0)
		FROM subject_tab s
		LEFT JOIN user_subject_map_tab ma ON s.subject_id = ma.subject_id AND ma.user_id = ?
		WHERE s.is_public = 1
		ORDER BY s.%s
		LIMIT ? OFFSET ?
		`, sanitisedOrderByString)
	rows, err := db.QueryContext(ctx, query, userId, pageSize, offset)
	if err != nil {
		log.Printf("Error retrieving subjects: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var fullSubjects []*mpb.FullSubject
	for rows.Next() {
		var fullSubject mpb.FullSubject
		fullSubject.Subject = new(mpb.DBSubject)
		if err := rows.Scan(&fullSubject.Subject.SubjectId, &fullSubject.Subject.SubjectName, &fullSubject.Subject.SubjectDescription, &fullSubject.Subject.IsPublic, &fullSubject.Subject.CreatedTime, &fullSubject.Subject.UpdatedTime, &fullSubject.UserSubjectRole, &fullSubject.IsFavourite); err != nil {
			log.Printf("Error scanning subject row: %v\n", err)
			return nil, err
		}
		fullSubjects = append(fullSubjects, &fullSubject)
	}

	log.Printf("Retrieved %d subjects\n", len(fullSubjects))
	return fullSubjects, nil
}

// GetPrivateSubjectsByUserId retrieves all private subjects for a user
func (repo *SubjectRepository) GetPrivateSubjectsByUserId(ctx context.Context, db *sql.DB, userID uint64, pageNumber, pageSize uint32, orderByField common.ORDER_BY_FIELD, orderByDirection common.ORDER_BY_DIRECTION) ([]*mpb.FullSubject, error) {
	offset := pageOffset(pageNumber, pageSize)
	sanitisedOrderByString := repo.generateSubjectOrderByString(orderByField, orderByDirection)

	query := fmt.Sprintf(`
		SELECT s.subject_id, s.subject_name, s.subject_description, s.is_public, s.created_time, s.updated_time,
		COALESCE(ma.user_subject_role, 0), COALESCE(ma.is_favourite, 0)
		FROM subject_tab s
		JOIN user_subject_map_tab ma ON s.subject_id = ma.subject_id AND ma.user_id = ? AND ma.user_subject_role != 0
		ORDER BY s.%s
		LIMIT ? OFFSET ?
		`, sanitisedOrderByString)

	rows, err := db.QueryContext(ctx, query, userID, pageSize, offset)
	if err != nil {
		log.Printf("Error retrieving subjects: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var fullSubjects []*mpb.FullSubject
	for rows.Next() {
		var fullSubject mpb.FullSubject
		fullSubject.Subject = new(mpb.DBSubject)
		if err := rows.Scan(&fullSubject.Subject.SubjectId, &fullSubject.Subject.SubjectName, &fullSubject.Subject.SubjectDescription, &fullSubject.Subject.IsPublic, &fullSubject.Subject.CreatedTime, &fullSubject.Subject.UpdatedTime, &fullSubject.UserSubjectRole, &fullSubject.IsFavourite); err != nil {
			log.Printf("Error scanning subject row: %v\n", err)
			return nil, err
		}
		fullSubjects = append(fullSubjects, &fullSubject)
	}

	log.Printf("Retrieved %d subjects\n", len(fullSubjects))
	return fullSubjects, nil
}

// GetFavouriteSubjectsByUserId retrieves all favorite subjects for a user
func (repo *SubjectRepository) GetFavouriteSubjectsByUserId(ctx context.Context, db *sql.DB, userID uint64, pageNumber, pageSize uint32, orderByField common.ORDER_BY_FIELD, orderByDirection common.ORDER_BY_DIRECTION) ([]*mpb.FullSubject, error) {
	offset := pageOffset(pageNumber, pageSize)
	sanitisedOrderByString := repo.generateSubjectOrderByString(orderByField, orderByDirection)
	query := fmt.Sprintf(`
		SELECT s.subject_id, s.subject_name, s.subject_description, s.is_public, s.created_time, s.updated_time, 
		COALESCE(ma.user_subject_role, 0), COALESCE(ma.is_favourite, 0)
		FROM subject_tab s
		JOIN user_subject_map_tab ma ON s.subject_id = ma.subject_id
		WHERE ma.user_id = ? AND ma.is_favourite = 1
		ORDER BY s.%s
		LIMIT ? OFFSET ?
		`, sanitisedOrderByString)

	rows, err := db.QueryContext(ctx, query, userID, pageSize, offset)
	if err != nil {
		log.Printf("Error retrieving subjects: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var fullSubjects []*mpb.FullSubject
	for rows.Next() {
		var fullSubject mpb.FullSubject
		fullSubject.Subject = new(mpb.DBSubject)
		if err := rows.Scan(&fullSubject.Subject.SubjectId, &fullSubject.Subject.SubjectName, &fullSubject.Subject.SubjectDescription, &fullSubject.Subject.IsPublic, &fullSubject.Subject.CreatedTime, &fullSubject.Subject.UpdatedTime, &fullSubject.UserSubjectRole, &fullSubject.IsFavourite); err != nil {
			log.Printf("Error scanning subject row: %v\n", err)
			return nil, err
		}
		fullSubjects = append(fullSubjects, &fullSubject)
	}

	log.Printf("Retrieved %d subjects\n", len(fullSubjects))
	return fullSubjects, nil
}

// GetSubjectById retrieves a subject by subject_id
func (repo *SubjectRepository) GetSubjectById(ctx context.Context, db *sql.DB, userID, subjectID uint64) (*mpb.FullSubject, error) {
	// get subject by subject_id
	// get user_subject_map_tab by user_id and subject_id
	// if subject is public or user is a member of the subject, return subject
	query := `
		WITH subject AS (
			SELECT subject_id, subject_name, subject_description, is_public, created_time, updated_time
			FROM subject_tab
			WHERE subject_id = ?
			LIMIT 1
			),
		user_subject_map AS (
			SELECT user_id, subject_id, user_subject_role, is_favourite
			FROM user_subject_map_tab
			WHERE user_id = ? AND subject_id = ?
			LIMIT 1
		)
		SELECT s.subject_id, s.subject_name, s.subject_description, s.is_public, s.created_time, s.updated_time, 
		COALESCE(ma.user_subject_role, 0), COALESCE(ma.is_favourite, 0)
		FROM subject s
		LEFT JOIN user_subject_map ma ON s.subject_id = ma.subject_id
		WHERE (s.is_public = 1 OR (ma.user_id IS NOT NULL AND ma.user_subject_role != 0))
		LIMIT 1
	`

	row := db.QueryRowContext(ctx, query, subjectID, userID, subjectID)
	var fullSubject mpb.FullSubject
	fullSubject.Subject = new(mpb.DBSubject)
	err := row.Scan(&fullSubject.Subject.SubjectId, &fullSubject.Subject.SubjectName, &fullSubject.Subject.SubjectDescription, &fullSubject.Subject.IsPublic, &fullSubject.Subject.CreatedTime, &fullSubject.Subject.UpdatedTime, &fullSubject.UserSubjectRole, &fullSubject.IsFavourite)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Subject not found: %v\n", err)
			return nil, nil
		}
		log.Printf("Error retrieving subject: %v\n", err)
		return nil, err
	}

	log.Println("Subject retrieved successfully")
	return &fullSubject, nil
}

// GetSubjectsByUserId retrieves subjects if public or if the user is a member of the subject
func (repo *SubjectRepository) GetSubjectsByUserId(ctx context.Context, db *sql.DB, userID uint64, pageNumber, pageSize uint32, orderByField common.ORDER_BY_FIELD, orderByDirection common.ORDER_BY_DIRECTION) ([]*mpb.FullSubject, error) {
	offset := pageOffset(pageNumber, pageSize)
	sanitisedOrderByString := repo.generateSubjectOrderByString(orderByField, orderByDirection)
	query := fmt.Sprintf(`
		SELECT s.subject_id, s.subject_name, s.subject_description, s.is_public, s.created_time, s.updated_time, 
		COALESCE(ma.user_subject_role, 0), COALESCE(ma.is_favourite, 0)
		FROM subject_tab s
		LEFT JOIN user_subject_map_tab ma ON s.subject_id = ma.subject_id AND ma.user_id = ?
		WHERE s.is_public = 1 OR (ma.user_id IS NOT NULL AND ma.user_subject_role != 0)
		ORDER BY s.%s
		LIMIT ? OFFSET ?
		`, sanitisedOrderByString)

	rows, err := db.QueryContext(ctx, query, userID, pageSize, offset)
	if err != nil {
		log.Printf("Error retrieving subjects: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var fullSubjects []*mpb.FullSubject
	for rows.Next() {
		var fullSubject mpb.FullSubject
		fullSubject.Subject = new(mpb.DBSubject)
		if err := rows.Scan(&fullSubject.Subject.SubjectId, &fullSubject.Subject.SubjectName, &fullSubject.Subject.SubjectDescription, &fullSubject.Subject.IsPublic, &fullSubject.Subject.CreatedTime, &fullSubject.Subject.UpdatedTime, &fullSubject.UserSubjectRole, &fullSubject.IsFavourite); err != nil {
			log.Printf("Error scanning subject row: %v\n", err)
			return nil, err
		}
		fullSubjects = append(fullSubjects, &fullSubject)
	}

	log.Printf("Retrieved %d subjects\n", len(fullSubjects))
	return fullSubjects, nil
}

// GetSubjectsByNameSearch retrieves subjects by name search
func (repo *SubjectRepository) GetSubjectsByNameSearch(ctx context.Context, db *sql.DB, userID uint64, nameSearch string, pageNumber, pageSize uint32, orderByField common.ORDER_BY_FIELD, orderByDirection common.ORDER_BY_DIRECTION) ([]*mpb.FullSubject, error) {
	offset := pageOffset(pageNumber, pageSize)
	sanitisedOrderByString := repo.generateSubjectOrderByString(orderByField, orderByDirection)
	query := fmt.Sprintf(`
		SELECT s.subject_id, s.subject_name, s.subject_description, s.is_public, s.created_time, s.updated_time, 
		COALESCE(ma.user_subject_role, 0), COALESCE(ma.is_favourite, 0)
		FROM subject_tab s
		LEFT JOIN user_subject_map_tab ma ON s.subject_id = ma.subject_id AND ma.user_id = ?
		WHERE (s.is_public = 1 OR (ma.user_id IS NOT NULL AND ma.user_subject_role != 0))
		AND s.subject_name LIKE ?
		ORDER BY s.%s
		LIMIT ? OFFSET ?
		`, sanitisedOrderByString)
	rows, err := db.QueryContext(ctx, query, userID, "%"+nameSearch+"%", pageSize, offset)
	if err != nil {
		log.Printf("Error retrieving subjects: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var fullSubjects []*mpb.FullSubject
	for rows.Next() {
		var fullSubject mpb.FullSubject
		fullSubject.Subject = new(mpb.DBSubject)
		if err := rows.Scan(&fullSubject.Subject.SubjectId, &fullSubject.Subject.SubjectName, &fullSubject.Subject.SubjectDescription, &fullSubject.Subject.IsPublic, &fullSubject.Subject.CreatedTime, &fullSubject.Subject.UpdatedTime, &fullSubject.UserSubjectRole, &fullSubject.IsFavourite); err != nil {
			log.Printf("Error scanning subject row: %v\n", err)
			return nil, err
		}
		fullSubjects = append(fullSubjects, &fullSubject)
	}

	log.Printf("Retrieved %d subjects\n", len(fullSubjects))
	return fullSubjects, nil
}

// UpdateSubject updates a subject by subject_id
func (repo *SubjectRepository) UpdateSubject(ctx context.Context, tx *sql.Tx, subjectID uint64, subjectName, subjectDescription string, isPublic bool) error {
	query := `
		UPDATE subject_tab
		SET subject_name = ?, subject_description = ?, is_public = ?, updated_time = ?
		WHERE subject_id = ?
	`
	updatedTime := time.Now().Unix()
	_, err := tx.ExecContext(ctx, query, subjectName, subjectDescription, isPublic, updatedTime, subjectID)
	if err != nil {
		log.Printf("Error updating subject: %v\n", err)
		return err
	}

	log.Println("Subject updated successfully")
	return nil
}

// DeleteSubject deletes a subject by subject_id
func (repo *SubjectRepository) DeleteSubject(ctx context.Context, tx *sql.Tx, subjectID uint64) error {
	query := `
		DELETE FROM subject_tab
		WHERE subject_id = ?
	`
	_, err := tx.ExecContext(ctx, query, subjectID)
	if err != nil {
		log.Printf("Error deleting subject: %v\n", err)
		return err
	}

	log.Println("Subject deleted successfully")
	return nil
}

func (repo *SubjectRepository) generateSubjectOrderByString(orderByField common.ORDER_BY_FIELD, orderByDirection common.ORDER_BY_DIRECTION) string {
	var orderByString string

	switch orderByField {
	case common.ORDER_BY_FIELD_ORDER_BY_FIELD_ID:
		orderByString = "subject_id"
	case common.ORDER_BY_FIELD_ORDER_BY_FIELD_TITLE:
		orderByString = "subject_name"
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
