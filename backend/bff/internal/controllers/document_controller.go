package controllers

import (
	"bff/internal/db"
	"bff/internal/proto/common"
	"bff/internal/proto/document"
	"bff/internal/services"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type DocumentController struct{
	s3Client *db.S3UploadClient
    docService *services.DocumentService
	pagseSize uint32
}

func NewDocumentController(s3Client *db.S3UploadClient, docService *services.DocumentService) *DocumentController {
	pageSizeString := os.Getenv("DOCUMENT_PAGE_SIZE")
	if pageSizeString == "" {
		// fallback to default value
		pageSizeString = "10"
	}
	pageSize, err := getUint32FromString(pageSizeString)
	if err != nil {
		// fatal error on start up
		panic(err)
	}
	return &DocumentController{s3Client: s3Client, docService: docService, pagseSize: pageSize}
}

func (c *DocumentController) CreateDocument(ctx *gin.Context) error {
	// on upload, the document will be uploaded to s3 
	// we send to document service to schedule a processing job
	// the document will be saved in database -> upload status = uploading
	// the document will be kafka processed
	userSession,err := getUserSession(ctx)
	if err != nil {
		return err
	}
	userId := userSession.UserId
	moduleId, err := c.getModuleIdFromContextParams(ctx)
	if err != nil {
		return err
	}
	fmt.Println("uploading document for user", userId)
	// upload the document to s3
	// s3ObjectKey, err := c.uploadDocument(ctx)
	// if err != nil {
	// 	return err
	// }
	s3ObjectKey := "test"
	fmt.Println("uploaded document to s3", s3ObjectKey)
	var req document.CreateDocRequest
	ctx.Bind(&req)
	req.UserId = userId
	req.UploadStatus = common.UploadStatus_UPLOAD_STATUS_UPLOADING
	req.S3ObjectKey = s3ObjectKey
	// get postform
	// unmarshal string to req
	jsonBody := ctx.PostForm("json")
	// req := document.CreateDocRequest{}
	err = json.Unmarshal([]byte(jsonBody), &req)
	if err != nil {
		return err
	}
	req.ModuleId = moduleId

	return c.docService.CreateDocument(ctx, &req)
}

func (c *DocumentController) GetDocumentById(ctx *gin.Context) (*document.GetDocByIdResponse, error) {
	userSession, err := getUserSession(ctx)
	if err != nil {
		return nil, err
	}
	userId := userSession.UserId
	documentId, err := c.getDocumentIdIdFromContextParams(ctx)
	if err != nil {
		return nil, err
	}
	moduleId, err := c.getModuleIdFromContextParams(ctx)
	if err != nil {
		return nil, err
	}
	req := document.GetDocByIdRequest{
		UserId: userId,
		DocId: documentId,
		ModuleId: moduleId,
	}

	return c.docService.GetDocumentById(ctx, &req)
}

func (c *DocumentController) GetDocumentsByModuleId(ctx *gin.Context) (*document.GetDocsByModuleIdResponse, error) {
	userSession, err := getUserSession(ctx)
	if err != nil {
		return nil, err
	}
	userId := userSession.UserId
	moduleId, err := c.getModuleIdFromContextParams(ctx)
	if err != nil {
		return nil, err
	}
	queryItems,err := NewQueryItems(ctx)
	if err != nil {
		return nil, err
	}

	var req document.GetDocsByModuleIdRequest
	ctx.Bind(&req)
	req.UserId = userId
	req.ModuleId = moduleId
	req.PageNumber = queryItems.PageNumber
	req.PageSize = c.pagseSize
	req.OrderByField = queryItems.OrderByField
	req.OrderByDirection = queryItems.OrderByDirection
	return c.docService.GetDocumentsByModuleId(ctx, &req)
}

func (c *DocumentController) GetDocumentsByNameSearch(ctx *gin.Context) (*document.GetDocsByNameSearchResponse, error) {
	userSession, err := getUserSession(ctx)
	if err != nil {
		return nil, err
	}
	userId := userSession.UserId
	moduleId, err := c.getModuleIdFromContextParams(ctx)
	if err != nil {
		return nil, err
	}
	queryItems, err := NewQueryItems(ctx)
	if err != nil {
		return nil, err
	}
	var req document.GetDocsByNameSearchRequest
	ctx.Bind(&req)
	req.UserId = userId
	req.ModuleId = moduleId
	req.PageNumber = queryItems.PageNumber
	req.PageSize = c.pagseSize
	req.OrderByField = queryItems.OrderByField
	req.OrderByDirection = queryItems.OrderByDirection
	return c.docService.GetDocumentsByNameSearch(ctx, &req)
}

func (c *DocumentController) UpdateDocument(ctx *gin.Context) error {
	userSession, err := getUserSession(ctx)
	if err != nil {
		return err
	}
	userId := userSession.UserId
	documentId, err := c.getDocumentIdIdFromContextParams(ctx)
	if err != nil {
		return err
	}
	moduleId, err := c.getModuleIdFromContextParams(ctx)
	if err != nil {
		return err
	}
	var req document.UpdateDocRequest
	ctx.Bind(&req)
	req.UserId = userId
	req.DocId = documentId
	req.ModuleId = moduleId
	return c.docService.UpdateDocument(ctx, &req)
}

func (c *DocumentController) DeleteDocument(ctx *gin.Context) error {
	userSession, err := getUserSession(ctx)
	if err != nil {
		return err
	}
	userId := userSession.UserId
	documentId, err := c.getDocumentIdIdFromContextParams(ctx)
	if err != nil {
		return err
	}
	moduleId, err := c.getModuleIdFromContextParams(ctx)
	if err != nil {
		return err
	}
	req := document.DeleteDocRequest{
		UserId: userId,
		DocId: documentId,
		ModuleId: moduleId,
	}
	return c.docService.DeleteDocument(ctx, &req)
}

//UploadDocument returns uuid upon successful upload else error
func (c *DocumentController) uploadDocument(ctx *gin.Context) (string, error) {
	file, fileHeader, err := ctx.Request.FormFile("file")
        if err != nil {
			return "", err
        }

        // Upload the file to S3
        uuid, err := c.s3Client.UploadFileToS3(file, fileHeader)
        if err != nil {
            return "", err
        }
		
		return uuid, nil
}

func (c *DocumentController) getModuleIdFromContextParams(ctx *gin.Context) (uint64, error) {
	moduleId, err := getUint64FromString(ctx.Param("module_id"))
	if err != nil {
		return 0, err
	}
	return moduleId, nil
}

func (c *DocumentController) getDocumentIdIdFromContextParams(ctx *gin.Context) (uint64, error) {
	documentId, err := getUint64FromString(ctx.Param("document_id"))
	if err != nil {
		return 0, err
	}
	return documentId, nil
}