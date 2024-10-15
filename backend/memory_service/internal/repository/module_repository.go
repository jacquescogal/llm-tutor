package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"memory_core/internal/proto/common"
	mpb "memory_core/internal/proto/module"
	"time"
)

type ModuleRepository struct {
}

func NewModuleRepository() *ModuleRepository {
	return &ModuleRepository{}
}

// CreateModule inserts a new module into the module_tab and returns the module_id
func (repo *ModuleRepository) CreateModule(ctx context.Context, tx *sql.Tx, moduleName, moduleDescription string, isPublic bool) (uint64, error) {
	query := `
		INSERT INTO module_tab (module_name, module_description, is_public, created_time, updated_time)
		VALUES (?, ?, ?, ?, ?)
	`

	createdTime := time.Now().Unix()
	result, err := tx.ExecContext(ctx, query, moduleName, moduleDescription, isPublic, createdTime, createdTime)
	if err != nil {
		log.Printf("Error creating module: %v\n", err)
		return 0, err
	}
	moduleId, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last inserted id: %v\n", err)
		return 0, err
	}
	log.Println("Module created successfully")
	return uint64(moduleId), nil
}

// GetPublicModules retrieves all public modules
func (repo *ModuleRepository) GetPublicModules(ctx context.Context, db *sql.DB, userId uint64, pageNumber, pageSize uint32, orderByField common.ORDER_BY_FIELD, orderByDirection common.ORDER_BY_DIRECTION) ([]*mpb.FullModule, error) {
	offset := pageOffset(pageNumber, pageSize)
	sanitisedOrderByString := repo.generateModuleOrderByString(orderByField, orderByDirection)
	query := fmt.Sprintf(`
		SELECT mt.module_id, mt.module_name, mt.module_description, mt.is_public, mt.created_time, mt.updated_time, 
		COALESCE(um.user_module_role, 0), COALESCE(um.is_favourite, 0)
		FROM module_tab mt
		LEFT JOIN user_module_map_tab um ON mt.module_id = um.module_id AND um.user_id = ?
		WHERE mt.is_public = 1
		ORDER BY mt.%s
		LIMIT ? OFFSET ?
	`, sanitisedOrderByString)

	rows, err := db.QueryContext(ctx, query, userId, pageSize, offset)
	if err != nil {
		log.Printf("Error retrieving modules: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var fullModules []*mpb.FullModule
	for rows.Next() {
		var fullModule mpb.FullModule
		fullModule.Module = &mpb.DBModule{}
		err := rows.Scan(&fullModule.Module.ModuleId, &fullModule.Module.ModuleName, &fullModule.Module.ModuleDescription, &fullModule.Module.IsPublic, &fullModule.Module.CreatedTime, &fullModule.Module.UpdatedTime, &fullModule.UserModuleRole, &fullModule.IsFavourite)
		if err != nil {
			log.Printf("Error scanning module: %v\n", err)
			return nil, err
		}
		fullModules = append(fullModules, &fullModule)
	}

	log.Println("Modules retrieved successfully")
	return fullModules, nil
}

// GetPrivateModulesByUserId retrieves all private modules for a user
func (repo *ModuleRepository) GetPrivateModulesByUserId(ctx context.Context, db *sql.DB, userID uint64, pageNumber, pageSize uint32, orderByField common.ORDER_BY_FIELD, orderByDirection common.ORDER_BY_DIRECTION) ([]*mpb.FullModule, error) {
	offset := pageOffset(pageNumber, pageSize)
	sanitisedOrderByString := repo.generateModuleOrderByString(orderByField, orderByDirection)
	query := fmt.Sprintf(`
		SELECT m.module_id, m.module_name, m.module_description, m.is_public, m.created_time, m.updated_time, 
		COALESCE(ma.user_module_role, 0), COALESCE(ma.is_favourite, 0)
		FROM module_tab m
		JOIN user_module_map_tab ma ON m.module_id = ma.module_id
		WHERE ma.user_id = ? AND ma.user_module_role != 0
		ORDER BY m.%s
		LIMIT ? OFFSET ?
	`, sanitisedOrderByString)

	rows, err := db.QueryContext(ctx, query, userID, pageSize, offset)
	if err != nil {
		log.Printf("Error retrieving modules: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var fullModules []*mpb.FullModule
	for rows.Next() {
		var fullModule mpb.FullModule
		fullModule.Module = &mpb.DBModule{}
		err := rows.Scan(&fullModule.Module.ModuleId, &fullModule.Module.ModuleName, &fullModule.Module.ModuleDescription, &fullModule.Module.IsPublic, &fullModule.Module.CreatedTime, &fullModule.Module.UpdatedTime, &fullModule.UserModuleRole, &fullModule.IsFavourite)
		if err != nil {
			log.Printf("Error scanning module: %v\n", err)
			return nil, err
		}
		fullModules = append(fullModules, &fullModule)
	}

	log.Println("Modules retrieved successfully")
	return fullModules, nil
}

// GetFavouriteModulesByUserId retrieves all favorite modules for a user
func (repo *ModuleRepository) GetFavouriteModulesByUserId(ctx context.Context, db *sql.DB, userID uint64, pageNumber, pageSize uint32, orderByField common.ORDER_BY_FIELD, orderByDirection common.ORDER_BY_DIRECTION) ([]*mpb.FullModule, error) {
	offset := pageOffset(pageNumber, pageSize)
	sanitisedOrderByString := repo.generateModuleOrderByString(orderByField, orderByDirection)

	query := fmt.Sprintf(`
		SELECT m.module_id, m.module_name, m.module_description, m.is_public, m.created_time, m.updated_time,
		COALESCE(ma.user_module_role, 0), COALESCE(ma.is_favourite, 0)
		FROM module_tab m
		JOIN user_module_map_tab ma ON m.module_id = ma.module_id
		WHERE ma.user_id = ? AND ma.is_favourite = 1
		ORDER BY m.%s
		LIMIT ? OFFSET ?
	`, sanitisedOrderByString)

	rows, err := db.QueryContext(ctx, query, userID, pageSize, offset)
	if err != nil {
		log.Printf("Error retrieving modules: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var fullModules []*mpb.FullModule
	for rows.Next() {
		var fullModule mpb.FullModule
		fullModule.Module = &mpb.DBModule{}
		err := rows.Scan(&fullModule.Module.ModuleId, &fullModule.Module.ModuleName, &fullModule.Module.ModuleDescription, &fullModule.Module.IsPublic, &fullModule.Module.CreatedTime, &fullModule.Module.UpdatedTime, &fullModule.UserModuleRole, &fullModule.IsFavourite)
		if err != nil {
			log.Printf("Error scanning module: %v\n", err)
			return nil, err
		}
		fullModules = append(fullModules, &fullModule)
	}

	log.Println("Modules retrieved successfully")
	return fullModules, nil
}

// GetModuleById retrieves a module by module_id
func (repo *ModuleRepository) GetModuleById(ctx context.Context, db *sql.DB, userID, moduleID uint64) (*mpb.FullModule, error) {
	// returns the module if it is public or the user is a member of the module
	// TODO: cache this
	log.Println("GetModuleById HELLO")
	query := `
	SELECT mt.module_id, mt.module_name, mt.module_description, mt.is_public, mt.created_time, mt.updated_time, 
	COALESCE(um.user_module_role, 0), COALESCE(um.is_favourite, 0)
	FROM module_tab mt
	LEFT JOIN user_module_map_tab um ON mt.module_id = um.module_id AND um.user_id = ?
	WHERE mt.module_id = ?
	`

	row := db.QueryRowContext(ctx, query, userID, moduleID)
	log.Println("Query executed")
	var fullModule mpb.FullModule
	fullModule.Module = &mpb.DBModule{}
	log.Println("FullModule created")
	err := row.Scan(&fullModule.Module.ModuleId, &fullModule.Module.ModuleName, &fullModule.Module.ModuleDescription, &fullModule.Module.IsPublic, &fullModule.Module.CreatedTime, &fullModule.Module.UpdatedTime, &fullModule.UserModuleRole, &fullModule.IsFavourite)
	log.Println("Row scanned")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Module not found: %v\n", err)
			return nil, nil
		}
		log.Printf("Error retrieving module: %v\n", err)
		return nil, err
	}

	log.Println("Module retrieved successfully")
	return &fullModule, nil
}

// GetModulesBySubjectId retrieves all modules by subject_id
func (repo *ModuleRepository) GetModulesBySubjectId(ctx context.Context, db *sql.DB, userID, subjectID uint64, pageNumber, pageSize uint32, orderByField common.ORDER_BY_FIELD, orderByDirection common.ORDER_BY_DIRECTION) ([]*mpb.FullModule, error) {
	// if subject owns the module, return the module
	// if module is private, the mapper will not exist
	offset := pageOffset(pageNumber, pageSize)
	sanitisedOrderByString := repo.generateModuleOrderByString(orderByField, orderByDirection)
	query := fmt.Sprintf(`
		WITH stm as (
			SELECT subject_id, module_id
			FROM subject_module_map_tab
			WHERE subject_id = ?
		)
		SELECT t.module_id, t.module_name, t.module_description, is_public, t.created_time, t.updated_time, 
		COALESCE(um.user_module_role, 0), COALESCE(um.is_favourite, 0)
		FROM module_tab t
		JOIN stm ON t.module_id = stm.module_id
		LEFT JOIN user_module_map_tab um ON t.module_id = um.module_id AND um.user_id = ?
		ORDER BY t.%s
		LIMIT ? OFFSET ?
	`, sanitisedOrderByString)

	rows, err := db.QueryContext(ctx, query, subjectID, userID, pageSize, offset)
	if err != nil {
		log.Printf("Error retrieving modules: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var fullModules []*mpb.FullModule
	for rows.Next() {
		var fullModule mpb.FullModule
		fullModule.Module = &mpb.DBModule{}
		err := rows.Scan(&fullModule.Module.ModuleId, &fullModule.Module.ModuleName, &fullModule.Module.ModuleDescription, &fullModule.Module.IsPublic, &fullModule.Module.CreatedTime, &fullModule.Module.UpdatedTime, &fullModule.UserModuleRole, &fullModule.IsFavourite)
		if err != nil {
			log.Printf("Error scanning module: %v\n", err)
			return nil, err
		}
		fullModules = append(fullModules, &fullModule)
	}

	log.Println("Modules retrieved successfully")
	return fullModules, nil
}

func (repo *ModuleRepository) GetPublicModulesByNameSearch(ctx context.Context, db *sql.DB, userId uint64, moduleNameSearch string, pageNumber, pageSize uint32, orderByField common.ORDER_BY_FIELD, orderByDirection common.ORDER_BY_DIRECTION) ([]*mpb.FullModule, error) {
	offset := pageOffset(pageNumber, pageSize)
	sanitisedOrderByString := repo.generateModuleOrderByString(orderByField, orderByDirection)
	query := fmt.Sprintf(`
		SELECT t.module_id, t.module_name, t.module_description, t.is_public, t.created_time, t.updated_time, 
		COALESCE(um.user_module_role, 0), COALESCE(um.is_favourite, 0)
		FROM module_tab t
		LEFT JOIN user_module_map_tab um ON t.module_id = um.module_id AND um.user_id = ?
		WHERE t.module_name LIKE ? AND t.is_public = 1
		ORDER BY t.%s
		LIMIT ? OFFSET ?
	`, sanitisedOrderByString)

	rows, err := db.QueryContext(ctx, query, userId, "%"+moduleNameSearch+"%", pageSize, offset)
	if err != nil {
		log.Printf("Error retrieving modules: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var fullModules []*mpb.FullModule
	for rows.Next() {
		var fullModule mpb.FullModule
		fullModule.Module = &mpb.DBModule{}
		err := rows.Scan(&fullModule.Module.ModuleId, &fullModule.Module.ModuleName, &fullModule.Module.ModuleDescription, &fullModule.Module.IsPublic, &fullModule.Module.CreatedTime, &fullModule.Module.UpdatedTime, &fullModule.UserModuleRole, &fullModule.IsFavourite)
		if err != nil {
			log.Printf("Error scanning module: %v\n", err)
			return nil, err
		}
		fullModules = append(fullModules, &fullModule)
	}

	log.Println("Modules retrieved successfully")
	return fullModules, nil
}

// UpdateModule updates a module by module_id
func (repo *ModuleRepository) UpdateModule(ctx context.Context, tx *sql.Tx, moduleID uint64, moduleName, moduleDescription string, isPublic bool) error {
	query := `
		UPDATE module_tab
		SET module_name = ?, module_description = ?, is_public = ?, updated_time = ?
		WHERE module_id = ?
	`

	updatedTime := time.Now().Unix()
	_, err := tx.ExecContext(ctx, query, moduleName, moduleDescription, isPublic, updatedTime, moduleID)
	if err != nil {
		log.Printf("Error updating module: %v\n", err)
		return err
	}

	log.Println("Module updated successfully")
	return nil
}

// Update module name by module_id
func (repo *ModuleRepository) UpdateModuleName(ctx context.Context, tx *sql.Tx, moduleID uint64, moduleName string) error {
	query := `
		UPDATE module_tab
		SET module_name = ?, updated_time = ?
		WHERE module_id = ?
	`

	updatedTime := time.Now().Unix()
	_, err := tx.ExecContext(ctx, query, moduleName, updatedTime, moduleID)
	if err != nil {
		log.Printf("Error updating module name: %v\n", err)
		return err
	}

	log.Println("Module name updated successfully")
	return nil
}

// Update module description by module_id
func (repo *ModuleRepository) UpdateModuleDescription(ctx context.Context, tx *sql.Tx, moduleID uint64, moduleDescription string) error {
	query := `
		UPDATE module_tab
		SET module_description = ?, updated_time = ?
		WHERE module_id = ?
	`

	updatedTime := time.Now().Unix()
	_, err := tx.ExecContext(ctx, query, moduleDescription, updatedTime, moduleID)
	if err != nil {
		log.Printf("Error updating module description: %v\n", err)
		return err
	}

	log.Println("Module description updated successfully")
	return nil
}

// Update module is_public by module_id
func (repo *ModuleRepository) UpdateModuleIsPublic(ctx context.Context, tx *sql.Tx, moduleID uint64, isPublic bool) error {
	query := `
		UPDATE module_tab
		SET is_public = ?, updated_time = ?
		WHERE module_id = ?
	`

	updatedTime := time.Now().Unix()
	_, err := tx.ExecContext(ctx, query, isPublic, updatedTime, moduleID)
	if err != nil {
		log.Printf("Error updating module is_public: %v\n", err)
		return err
	}

	log.Println("Module is_public updated successfully")
	return nil
}

// DeleteModule deletes a module by module_id
func (repo *ModuleRepository) DeleteModule(ctx context.Context, tx *sql.Tx, moduleID uint64) error {
	query := `
		DELETE FROM module_tab
		WHERE module_id = ?
	`

	_, err := tx.ExecContext(ctx, query, moduleID)
	if err != nil {
		log.Printf("Error deleting module: %v\n", err)
		return err
	}

	log.Println("Module deleted successfully")
	return nil
}


// UserModuleMap

func (repo *ModuleRepository) GetUserModuleMapping(ctx context.Context, tx *sql.Tx, userID, moduleID uint64) error {
	query := `
		SELECT user_module_role, is_favourite
		FROM user_module_map_tab
		WHERE user_id = ? AND module_id = ?
	`

	row := tx.QueryRowContext(ctx, query, userID, moduleID)
	var userModuleRole uint32
	var isFavourite bool
	err := row.Scan(&userModuleRole, &isFavourite)
	if err != nil {
		log.Printf("Error retrieving user module mapping: %v\n", err)
		return err
	}

	log.Println("User module mapping retrieved successfully")
	return nil
}

// generateModuleOrderByString generates the ORDER BY string for the module query
func (repo *ModuleRepository) generateModuleOrderByString(orderByField common.ORDER_BY_FIELD, orderByDirection common.ORDER_BY_DIRECTION) string {
	var orderByString string

	switch orderByField {
	case common.ORDER_BY_FIELD_ORDER_BY_FIELD_ID:
		orderByString = "module_id"
	case common.ORDER_BY_FIELD_ORDER_BY_FIELD_TITLE:
		orderByString = "module_name"
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
