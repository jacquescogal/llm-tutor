package controllers

import (
	"bff/internal/db"
	"bff/internal/proto/common"
	"bff/internal/proto/document"
	"bff/internal/services"

	"github.com/gin-gonic/gin"
)

type DocumentController struct{
	s3Client *db.S3UploadClient
    docService *services.DocumentService
}

func NewDocumentController(s3Client *db.S3UploadClient, docService *services.DocumentService) *DocumentController {
	return &DocumentController{s3Client: s3Client, docService: docService}
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
	
	// upload the document to s3
	s3ObjectKey, err := c.uploadDocument(ctx)
	if err != nil {
		return err
	}
	var req document.CreateDocRequest
	ctx.Bind(&req)
	req.UserId = userId
	req.UploadStatus = common.UploadStatus_UPLOAD_STATUS_UPLOADING
	req.S3ObjectKey = s3ObjectKey

	return c.docService.CreateDocument(ctx, &req)
}

func (c *DocumentController) GetDocumentById(ctx *gin.Context) (*document.GetDocByIdResponse, error) {
	userSession, err := getUserSession(ctx)
	if err != nil {
		return nil, err
	}
	userId := userSession.UserId
	if err != nil {
		return nil, err
	}
	var req document.GetDocByIdRequest
	ctx.Bind(&req)
	req.UserId = userId

	return c.docService.GetDocumentById(ctx, &req)
}

func (c *DocumentController) GetDocumentsByModuleId(ctx *gin.Context) (*document.GetDocsByModuleIdResponse, error) {
	userSession, err := getUserSession(ctx)
	if err != nil {
		return nil, err
	}
	userId := userSession.UserId
	var req document.GetDocsByModuleIdRequest
	ctx.Bind(&req)
	req.UserId = userId
	return c.docService.GetDocumentsByModuleId(ctx, &req)
}

func (c *DocumentController) GetDocumentsByNameSearch(ctx *gin.Context) (*document.GetDocsByNameSearchResponse, error) {
	userSession, err := getUserSession(ctx)
	if err != nil {
		return nil, err
	}
	userId := userSession.UserId
	var req document.GetDocsByNameSearchRequest
	ctx.Bind(&req)
	req.UserId = userId
	return c.docService.GetDocumentsByNameSearch(ctx, &req)
}

func (c *DocumentController) UpdateDocument(ctx *gin.Context) error {
	userSession, err := getUserSession(ctx)
	if err != nil {
		return err
	}
	userId := userSession.UserId
	var req document.UpdateDocRequest
	ctx.Bind(&req)
	req.UserId = userId
	return c.docService.UpdateDocument(ctx, &req)
}

func (c *DocumentController) DeleteDocument(ctx *gin.Context) error {
	userSession, err := getUserSession(ctx)
	if err != nil {
		return err
	}
	userId := userSession.UserId
	var req document.DeleteDocRequest
	ctx.Bind(&req)
	req.UserId = userId
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