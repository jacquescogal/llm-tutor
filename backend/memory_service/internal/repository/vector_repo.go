package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	client "github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/filters"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/graphql"
)


type VectorRepository struct {
	Conn *client.Client
}

func NewVectorRepository(conn *client.Client) *VectorRepository {
	return &VectorRepository{Conn: conn}
}

const (
	CLASS_NAME = "Memory"
	MODULE_ID = "module_id"
	MEMORY_TITLE = "memory_title"
	MEMORY_CONTENT = "memory_content"
	DOCUMENT_ID = "document_id"
)

func (vr *VectorRepository) CreateMemoryVector(ctx context.Context, moduleId, documendId uint64, vectorUuid string, memoryTitle, memoryContent string) error{
	// memoryIdString := fmt.Sprintf("%v", memoryId)
	w, err := vr.Conn.Data().Creator().
	WithClassName(CLASS_NAME).
	WithProperties(map[string]interface{}{
		MEMORY_TITLE: memoryTitle,
		MEMORY_CONTENT:memoryContent,
		MODULE_ID: moduleId,
		DOCUMENT_ID: documendId,
	}).WithID(vectorUuid).
	Do(ctx)
	if err != nil {
		fmt.Printf("Error creating object: %v\n", err)
		return err
	}
	log.Printf("Created object with UUID: %v\n", w.Object.ID)
	return nil
}

func (vr *VectorRepository) SearchMemoryVectorSubjectScope(ctx context.Context, query string, module_filter []uint64, limit int) ([]byte, error) {

	module_filter_int64 := make([]float64, len(module_filter))
	for i, v := range module_filter {
		module_filter_int64[i] = float64(v)
	}
	response, err := vr.Conn.GraphQL().Get().
	WithClassName(CLASS_NAME).
	WithFields(
		graphql.Field{Name: MEMORY_TITLE},
		graphql.Field{Name: MEMORY_CONTENT},
		graphql.Field{Name: MODULE_ID},
		graphql.Field{Name: DOCUMENT_ID},
	).
	WithNearText(vr.Conn.GraphQL().NearTextArgBuilder().
		WithConcepts([]string{query})).
	WithLimit(limit).WithWhere(filters.Where().
    WithPath([]string{MODULE_ID}).
    WithOperator(filters.ContainsAny).
    WithValueNumber(module_filter_int64...)).
	Do(ctx)
	jsonResponse, _ := json.MarshalIndent(response,""," ")
	log.Printf("Search response: %v\n", string(jsonResponse))
	return jsonResponse, err
}

func (vr *VectorRepository) SearchMemoryVectorDocScope(ctx context.Context, query string, documentId uint64, limit int) ([]byte, error) {
	response, err := vr.Conn.GraphQL().Get().
	WithClassName(CLASS_NAME).
	WithFields(
		graphql.Field{Name: MEMORY_TITLE},
		graphql.Field{Name: MEMORY_CONTENT},
		graphql.Field{Name: MODULE_ID},
		graphql.Field{Name: DOCUMENT_ID},
	).
	WithWhere(filters.Where().
WithPath([]string{DOCUMENT_ID}).
WithOperator(filters.Equal).WithValueNumber(float64(documentId))). // risky, but current usage, unlikely to be a problem
	WithNearText(vr.Conn.GraphQL().NearTextArgBuilder().
	WithConcepts([]string{query})).
	WithLimit(limit).
	Do(ctx)
	jsonResponse, _ := json.MarshalIndent(response,""," ")
	log.Printf("Search response: %v\n", string(jsonResponse))
	return jsonResponse, err
}

func (vr *VectorRepository) SearchMemoryVectorModuleScope(ctx context.Context, query string, moduleId uint64, limit int) ([]byte, error) {
	response, err := vr.Conn.GraphQL().Get().
	WithClassName(CLASS_NAME).
	WithFields(
		graphql.Field{Name: MEMORY_TITLE},
		graphql.Field{Name: MEMORY_CONTENT},
		graphql.Field{Name: MODULE_ID},
		graphql.Field{Name: DOCUMENT_ID},
	).
	WithNearText(vr.Conn.GraphQL().NearTextArgBuilder().
		WithConcepts([]string{query})).
	WithLimit(limit).
	WithWhere(filters.Where().
WithPath([]string{MODULE_ID}).
WithOperator(filters.Equal).WithValueNumber(float64(moduleId))). // risky, but current usage, unlikely to be a problem
	WithNearText(vr.Conn.GraphQL().NearTextArgBuilder().
	WithConcepts([]string{query})).
	WithLimit(limit).
	Do(ctx)
	jsonResponse, _ := json.MarshalIndent(response,""," ")
	log.Printf("Search response: %v\n", string(jsonResponse))
	return jsonResponse, err
}