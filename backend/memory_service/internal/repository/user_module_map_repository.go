package repository
// DONE
import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"memory_core/internal/cache"
	mpb "memory_core/internal/proto/module"
	"memory_core/internal/proto/common"
)

type UserModuleMapRepository struct {
	cacheStore *cache.CacheStore
}

func NewUserModuleMapRepository(cacheStore *cache.CacheStore) *UserModuleMapRepository {
	return &UserModuleMapRepository{cacheStore: cacheStore}
}

// PutUserModuleMapping inserts a new user module mapping into user_module_map_tab or updates an existing one
func (repo *UserModuleMapRepository) PutUserModuleMappingFavourite(ctx context.Context, tx *sql.Tx, userId, moduleId uint64, isFavourite bool) error {
	// either one of user_module_role or is_favourite will be put
	ifDuplicateString := "is_favourite = VALUES(is_favourite)"
	fmt.Println("ifDuplicateString", ifDuplicateString, "userId", userId, "moduleId", moduleId, "isFavourite", isFavourite)

	// if role != common.UserModuleRole_USER_MODULE_ROLE_UNDEFINED {
	// 	ifDuplicateString = "user_module_role = VALUES(user_module_role)"
	// }

	query := fmt.Sprintf(`
		INSERT INTO user_module_map_tab (user_id, module_id, is_favourite)
		VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE
		%s`, ifDuplicateString)

	_, err := tx.ExecContext(ctx, query, userId, moduleId, isFavourite)
	if err != nil {
		log.Printf("Error putting user module mapping: %v\n", err)
		return err
	}
	log.Println("Put user module mapping successfully")
	return nil
}

// PutUserModuleMapping inserts a new user module mapping into user_module_map_tab or updates an existing one
func (repo *UserModuleMapRepository) PutUserModuleMappingRole(ctx context.Context, tx *sql.Tx, userId, moduleId uint64, role common.UserModuleRole) error {

	ifDuplicateString := "user_module_role = VALUES(user_module_role)"


	query := fmt.Sprintf(`
		INSERT INTO user_module_map_tab (user_id, module_id, user_module_role)
		VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE
		%s`, ifDuplicateString)

	_, err := tx.ExecContext(ctx, query, userId, moduleId, role)
	if err != nil {
		log.Printf("Error putting user module mapping: %v\n", err)
		return err
	}
	log.Println("Put user module mapping successfully")
	return nil
}

// GetUserModuleMapping retrieves a user module mapping by user_id and module_id
func (repo *UserModuleMapRepository) GetUserModuleMapping(ctx context.Context, db *sql.DB, userId, moduleId uint64) (*mpb.DBUserModuleMap, error) {
	// checkCache
	var userModuleMapping mpb.DBUserModuleMap
	keyString := repo.getCacheKeyString(moduleId, userId)
	err := repo.cacheStore.RetrieveData(ctx, keyString, &userModuleMapping)
	if err == nil {
		// cache hit
		return &userModuleMapping, nil
	}

	query := `
		SELECT user_id, module_id, user_module_role, is_favourite
		FROM user_module_map_tab
		WHERE user_id = ? AND module_id = ?
		LIMIT 1
	`
	row := db.QueryRowContext(ctx, query, userId, moduleId)
	
	err = row.Scan(&userModuleMapping.UserId, &userModuleMapping.ModuleId, &userModuleMapping.UserModuleRole, &userModuleMapping.IsFavourite)
	if err != nil {
		// dependency error or ErrNoRows
		log.Printf("Error getting user module mapping: %v\n", err)
		return nil, err
	}

	// store in cache
	err = repo.cacheStore.StoreData(ctx, keyString, &userModuleMapping, 30)
	if err != nil {
		// just log the error
		log.Printf("Error storing user module mapping in cache: %v\n", err)
	}
	return &userModuleMapping, nil
}

func (repo *UserModuleMapRepository) getCacheKeyString(moduleId, userId uint64) string {
	return fmt.Sprintf("user_module_mapping:%d:%d", moduleId, userId)
}