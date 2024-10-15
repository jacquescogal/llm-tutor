package controller

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"memory_core/internal/proto/common"
	"memory_core/internal/proto/vector"
	"memory_core/internal/repository"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type VectorController struct {

	db *sql.DB
	vectorRepo *repository.VectorRepository
	subjectModuleMapRepo *repository.SubjectModuleMapRepository

}

func NewVectorController(db *sql.DB, vectorRepo *repository.VectorRepository, subjectModuleMapRepo *repository.SubjectModuleMapRepository) *VectorController {
	return &VectorController{db: db, vectorRepo: vectorRepo, subjectModuleMapRepo: subjectModuleMapRepo}
}

func (vc *VectorController) CreateMemoryVector(ctx context.Context, req *vector.CreateMemoryVectorRequest) (*vector.CreateMemoryVectorResponse, error){
	// check if module_id, memory_title and memory_content are not empty
	if req.GetModuleId() == 0 {
		log.Println("Module ID is required")
		return nil, status.Error(codes.InvalidArgument, "Module ID is required")
	}else if req.GetMemoryTitle() == "" {
		log.Println("Memory title is required")
		return nil, status.Error(codes.InvalidArgument, "Memory title is required")
	}else if req.GetMemoryContent() == "" {
		log.Println("Memory content is required")
		return nil, status.Error(codes.InvalidArgument, "Memory content is required")
	} else if req.GetMemoryId() == 0 {
		log.Println("Memory ID is required")
		return nil, status.Error(codes.InvalidArgument, "Memory ID is required")
	}
	memoryUUID := uuid.New().String()
	vc.vectorRepo.CreateMemoryVector(ctx, req.GetModuleId(),0, memoryUUID, req.GetMemoryTitle(), req.GetMemoryContent())
	return &vector.CreateMemoryVectorResponse{}, nil
}

func (vc *VectorController) SearchMemoryVector(ctx context.Context, req *vector.SearchMemoryVectorRequest) (*vector.SearchMemoryVectorResponse, error){
	// check if search_query is not empty
	
	if req.GetSearchQuery() == "" {
		log.Println("Search query is required")
		return nil, status.Error(codes.InvalidArgument, "Search query is required")
	}
	var jsonReponse []byte
	var err error;
	if req.GetIdType() == common.IDType_ID_DOCUMENT {
		if req.GetId() == 0 {
			log.Println("Document ID is required")
			return nil, status.Error(codes.InvalidArgument, "Document ID is required")
		}
		jsonReponse,err = vc.vectorRepo.SearchMemoryVectorDocScope(ctx, req.GetSearchQuery(), req.GetId(), int(req.GetLimit()))
		if err != nil {
			log.Printf("Failed to search memory vector: %v", err)
			return nil, err
		}
	} else if req.GetIdType() == common.IDType_ID_MODULE {
		if req.GetId() == 0 {
			log.Println("Memory ID is required")
			return nil, status.Error(codes.InvalidArgument, "Memory ID is required")
		}
		fmt.Println("module_id", req.GetId(), "search_query", req.GetSearchQuery(), "limit", req.GetLimit())
		jsonReponse,err = vc.vectorRepo.SearchMemoryVectorModuleScope(ctx, req.GetSearchQuery(), req.GetId(), int(req.GetLimit()))
		fmt.Println("jsonResponse", jsonReponse)
		if err != nil {
			log.Printf("Failed to search memory vector: %v", err)
			return nil, err
		}
	} else if req.GetIdType() == common.IDType_ID_SUBJECT {
		if req.GetId() == 0 {
			log.Println("Memory ID is required")
			return nil, status.Error(codes.InvalidArgument, "Memory ID is required")
		}
		// get mapping from subjectModule Map repo
	
		subjectModuleMap,err := vc.subjectModuleMapRepo.GetSubjectModuleMappingBySubjectId(ctx, vc.db, req.GetId())
		if err != nil {
			log.Printf("Failed to get subjectModuleMap by subject ID: %v", err)
			return nil, err
		}
		moduleIDs := []uint64{}
		for _, module := range subjectModuleMap {
			moduleIDs = append(moduleIDs, module.ModuleId)
		}
		jsonReponse,err = vc.vectorRepo.SearchMemoryVectorSubjectScope(ctx, req.GetSearchQuery(), moduleIDs, int(req.GetLimit()))
		if err != nil {
			log.Printf("Failed to search memory vector: %v", err)
			return nil, err
		}
	}

	if err != nil {
		log.Printf("Failed to search memory vector: %v", err)
		return nil, err
	}

	fmt.Println("jsonResponse", jsonReponse)
	return &vector.SearchMemoryVectorResponse{
		JsonResponse: jsonReponse,
	}, nil
}