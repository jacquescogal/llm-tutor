package repository
// DONE
import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"memory_core/internal/cache"
	mpb "memory_core/internal/proto/subject"
	"memory_core/internal/proto/common"
)

type UserSubjectMapRepository struct {
	cacheStore *cache.CacheStore
}

func NewUserSubjectMapRepository(cacheStore *cache.CacheStore) *UserSubjectMapRepository {
	return &UserSubjectMapRepository{cacheStore: cacheStore}
}

// PutUserSubjectMappingFavourite inserts a new user subject mapping into user_subject_map_tab or updates an existing one
func (repo *UserSubjectMapRepository) PutUserSubjectMappingFavourite(ctx context.Context, tx *sql.Tx, userId, subjectId uint64, isFavourite bool) error {
	// either one of user_subject_role or is_favourite will be put
	ifDuplicateString := "is_favourite = VALUES(is_favourite)"
	query := fmt.Sprintf(`
		INSERT INTO user_subject_map_tab (user_id, subject_id, is_favourite)
		VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE
		%s`, ifDuplicateString)

	_, err := tx.ExecContext(ctx, query, userId, subjectId, isFavourite)
	if err != nil {
		log.Printf("Error putting user subject mapping: %v\n", err)
		return err
	}
	log.Println("Put user subject mapping successfully")
	return nil
}

func (repo *UserSubjectMapRepository) PutUserSubjectMappingRole(ctx context.Context, tx *sql.Tx, userId, subjectId uint64, role common.UserSubjectRole) error {
	ifDuplicateString := "user_subject_role = VALUES(user_subject_role)"


	query := fmt.Sprintf(`
		INSERT INTO user_subject_map_tab (user_id, subject_id, user_subject_role)
		VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE
		%s`, ifDuplicateString)

	_, err := tx.ExecContext(ctx, query, userId, subjectId, role)
	if err != nil {
		log.Printf("Error putting user subject mapping: %v\n", err)
		return err
	}
	log.Println("Put user subject mapping successfully")
	return nil
}


// GetUserSubjectMapping retrieves a user subject mapping by user_id and subject_id
func (repo *UserSubjectMapRepository) GetUserSubjectMapping(ctx context.Context, db *sql.DB, userId, subjectId uint64) (*mpb.DBUserSubjectMap, error) {
	// checkCache
	var userSubjectMapping mpb.DBUserSubjectMap
	keyString := repo.getCacheKeyString(subjectId, userId)
	err := repo.cacheStore.RetrieveData(ctx, keyString, &userSubjectMapping)
	if err == nil {
		// cache hit
		return &userSubjectMapping, nil
	}

	query := `
		SELECT user_id, subject_id, user_subject_role, is_favourite
		FROM user_subject_map_tab
		WHERE user_id = ? AND subject_id = ?
		LIMIT 1
	`
	row := db.QueryRowContext(ctx, query, userId, subjectId)
	
	err = row.Scan(&userSubjectMapping.UserId, &userSubjectMapping.SubjectId, &userSubjectMapping.UserSubjectRole, &userSubjectMapping.IsFavourite)
	if err != nil {
		// dependency error or ErrNoRows
		log.Printf("Error getting user subject mapping: %v\n", err)
		return nil, err
	}

	// store in cache
	err = repo.cacheStore.StoreData(ctx, keyString, &userSubjectMapping, 30)
	if err != nil {
		// just log the error
		log.Printf("Error storing user subject mapping in cache: %v\n", err)
	}
	return &userSubjectMapping, nil
}

func (repo *UserSubjectMapRepository) getCacheKeyString(subjectId, userId uint64) string {
	return fmt.Sprintf("user_subject_mapping:%d:%d", subjectId, userId)
}