package repository

// DONE
import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"memory_core/internal/cache"
	mpb "memory_core/internal/proto/subject"
	"strings"
)

type SubjectModuleMapRepository struct {
	cacheStore *cache.CacheStore
}

func NewSubjectModuleMapRepository(cacheStore *cache.CacheStore) *SubjectModuleMapRepository {
	return &SubjectModuleMapRepository{cacheStore: cacheStore}
}

// CreateSubjectModuleMapping inserts a new subject module mapping into subject_module_map_tab
func (repo *SubjectModuleMapRepository) CreateSubjectModuleMapping(ctx context.Context, tx *sql.Tx, subjectId, moduleId uint64) error {
	query := `
		INSERT INTO subject_module_map_tab (subject_id, module_id)
		VALUES (?, ?)
	`
	_, err := tx.ExecContext(ctx, query, subjectId, moduleId)
	if err != nil {
		log.Printf("Error creating subject module mapping: %v\n", err)
		return err
	}
	log.Println("Subject module mapping created successfully")
	return nil
}

// CreateSubjectModuleMappings inserts multiple module mappings for a single subject into subject_module_map_tab in a batch
func (repo *SubjectModuleMapRepository) CreateSubjectModuleMappings(ctx context.Context, tx *sql.Tx, subjectId uint64, moduleIds []uint64) error {
	// Base query
	query := `INSERT INTO subject_module_map_tab (subject_id, module_id) VALUES `

	// Prepare placeholders and values
	values := []interface{}{}
	placeholders := []string{}
	for _, moduleId := range moduleIds {
		placeholders = append(placeholders, "(?, ?)")
		values = append(values, subjectId, moduleId)
	}

	// Join the placeholders to form the full query
	query += strings.Join(placeholders, ", ")

	// Execute the batch insert query
	_, err := tx.ExecContext(ctx, query, values...)
	if err != nil {
		log.Printf("Error creating subject module mappings: %v\n", err)
		return err
	}

	log.Println("Batch subject module mappings created successfully")
	return nil
}



// GetSubjectModuleMapping retrieves a subject module mapping by subject_id and module_id
func (repo *SubjectModuleMapRepository) GetSubjectModuleMapping(ctx context.Context, db *sql.DB, subjectId, moduleId uint64) (*mpb.DBSubjectModuleMap, error) {
	// checkCache
	var subjectModuleMapping mpb.DBSubjectModuleMap
	keyString := repo.getCacheKeyString(subjectId, moduleId)
	err := repo.cacheStore.RetrieveData(ctx, keyString, &subjectModuleMapping)
	if err == nil {
		// cache hit
		return &subjectModuleMapping, nil
	}

	query := `
		SELECT subject_id, module_id
		FROM subject_module_map_tab
		WHERE subject_id = ? AND module_id = ?
		LIMIT 1
	`
	row := db.QueryRowContext(ctx, query, subjectId, moduleId)
	
	err = row.Scan(&subjectModuleMapping.SubjectId, &subjectModuleMapping.ModuleId)
	if err != nil {
		// dependency error or ErrNoRows
		log.Printf("Error getting subject module mapping: %v\n", err)
		return nil, err
	}

	// store in cache
	err = repo.cacheStore.StoreData(ctx, keyString, &subjectModuleMapping, 30)
	if err != nil {
		// just log the error
		log.Printf("Error storing subject module mapping in cache: %v\n", err)
	}
	return &subjectModuleMapping, nil
}

// GetSubjectModuleMappingBySubjectId retrieves all module mappings for a single subject by subject_id
func (repo *SubjectModuleMapRepository) GetSubjectModuleMappingBySubjectId(ctx context.Context, db *sql.DB, subjectId uint64) ([]*mpb.DBSubjectModuleMap, error) {
	query := `
		SELECT subject_id, module_id
		FROM subject_module_map_tab
		WHERE subject_id = ?
	`
	rows, err := db.QueryContext(ctx, query, subjectId)
	if err != nil {
		log.Printf("Error getting subject module mappings: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var subjectModuleMappings []*mpb.DBSubjectModuleMap
	for rows.Next() {
		var subjectModuleMapping mpb.DBSubjectModuleMap
		err := rows.Scan(&subjectModuleMapping.SubjectId, &subjectModuleMapping.ModuleId)
		if err != nil {
			log.Printf("Error scanning subject module mapping row: %v\n", err)
			return nil, err
		}
		subjectModuleMappings = append(subjectModuleMappings, &subjectModuleMapping)
	}
	log.Printf("Retrieved %d subject module mappings\n", len(subjectModuleMappings))
	return subjectModuleMappings, nil
}


// DeleteSubjectModuleMapping deletes a subject module mapping by subject_id and module_id
func (repo *SubjectModuleMapRepository) DeleteSubjectModuleMapping(ctx context.Context, tx *sql.Tx, subjectId, moduleId uint64) error {
	// Deletes the mapping between a subject and a module
	query := `
		DELETE FROM subject_module_map_tab
		WHERE subject_id = ? AND module_id = ?
	`
	_, err := tx.ExecContext(ctx, query, subjectId, moduleId)
	if err != nil {
		log.Printf("Error deleting subject module mapping: %v\n", err)
		return err
	}
	log.Println("Subject module mapping deleted successfully")
	return nil
}


// DeleteSubjectModuleMappingsBySubjectId deletes all subject module mappings by subject_id
func (repo *SubjectModuleMapRepository) DeleteSubjectModuleMappingsBySubjectId(ctx context.Context, tx *sql.Tx, subjectId uint64) error {
	// Should not be called as the cascade delete is enabled, and if subject is deleted, all its mappings will be deleted
	fmt.Println(subjectId)
	query := `
		DELETE FROM subject_module_map_tab
		WHERE subject_id = ?
	`
	_, err := tx.ExecContext(ctx, query, subjectId)
	if err != nil {
		log.Printf("Error deleting subject module mappings: %v", err)
		return err
	}
	log.Println("Subject module mappings deleted successfully")
	return nil
}

// DeleteSubjectModuleMappingsByModuleId deletes all subject module mappings by module_id
func (repo *SubjectModuleMapRepository) DeleteSubjectModuleMappingsByModuleId(ctx context.Context, tx *sql.Tx, moduleId uint64) error {
	// Effectively removes the module from all subjects
	// if module becomes private, all subjects will lose access to it
	query := `
		DELETE FROM subject_module_map_tab
		WHERE module_id = ?
	`
	_, err := tx.ExecContext(ctx, query, moduleId)
	if err != nil {
		log.Printf("Error deleting subject module mappings: %v", err)
		return err
	}
	log.Println("Subject module mappings deleted successfully")
	return nil
}

func (repo *SubjectModuleMapRepository) getCacheKeyString(subjectId, moduleId uint64) string {
	return fmt.Sprintf("subject_module_mapping:%d:%d", subjectId, moduleId)
}