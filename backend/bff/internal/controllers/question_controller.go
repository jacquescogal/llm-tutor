package controllers

import (
	"bff/internal/proto/common"
	"bff/internal/proto/question"
	"bff/internal/services"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

type QuestionController struct{
	questionService *services.QuestionService
  pageSize uint32
}



func NewQuestionController(questionService *services.QuestionService) *QuestionController {
  pageSizeString := os.Getenv("MEMORY_PAGE_SIZE")
	if pageSizeString == "" {
		// fallback to default value
		pageSizeString = "10"
	}
	pageSize, err := getUint32FromString(pageSizeString)
	if err != nil {
		// fatal error on start up
		panic(err)
	}
	return &QuestionController{questionService: questionService, pageSize: pageSize}
}

func (c *QuestionController) CreateQuestion(ctx *gin.Context) error {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return err
	}
	userId := userSession.UserId
  moduleId,err := c.getModuleIdFromContextParams(ctx)
	if err != nil {
		return err
	}
	documentId,err := c.getDocumentIdFromContextParams(ctx)
	if err != nil {
		return err
	}
	var req question.CreateQuestionRequest
	ctx.Bind(&req)
	req.UserId = userId
  req.ModuleId = moduleId
  req.DocId = documentId
	return c.questionService.CreateQuestion(ctx, &req)
}

func (c *QuestionController) GetQuestionById(ctx *gin.Context) (*question.GetQuestionByIdResponse, error) {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return nil, err
	}
	userId := userSession.UserId
  moduleId,err := c.getModuleIdFromContextParams(ctx)
	if err != nil {
		return nil, err
	}
	documentId,err := c.getDocumentIdFromContextParams(ctx)
	if err != nil {
		return nil, err
	}
  questionId,err := c.getQuestionIdFromContextParams(ctx)
  if err != nil {
    return nil, err
  }
	var req question.GetQuestionByIdRequest
	ctx.Bind(&req)
	req.UserId = userId
  req.ModuleId = moduleId
  req.DocId = documentId
  req.QuestionId = questionId
	return c.questionService.GetQuestionById(ctx, &req)
}

func (c *QuestionController) GetQuestionsByDocId(ctx *gin.Context) ([]question.QuestionReturn, error) {
	userSession, err := getUserSession(ctx)
	if err != nil {
		return nil, err
	}
	userId := userSession.UserId

	moduleId, err := c.getModuleIdFromContextParams(ctx)
	if err != nil {
		return nil, err
	}

	documentId, err := c.getDocumentIdFromContextParams(ctx)
	if err != nil {
		return nil, err
	}

	queryItems, err := NewQueryItems(ctx)
	if err != nil {
		return nil, err
	}

	var req question.GetQuestionsByDocIdRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, err
	}

	req.UserId = userId
	req.ModuleId = moduleId
	req.DocId = documentId
	req.PageNumber = queryItems.PageNumber
	req.PageSize = c.pageSize
	req.OrderByField = queryItems.OrderByField
	req.OrderByDirection = queryItems.OrderByDirection

	getQuestionsRes, err := c.questionService.GetQuestionsByDocId(ctx, &req)
	if err != nil {
		return nil, err
	}

	questionList := []question.QuestionReturn{}
	for _, qstn := range getQuestionsRes.Questions {
		var questionJsonString string

		switch qstn.QuestionType {
		case common.QuestionType_QUESTION_TYPE_MCQ, common.QuestionType_QUESTION_TYPE_MULTI_ANSWER_MCQ:
			// Unmarshal MCQ question
			var mcqQuestion question.MCQQuestion
			if err := proto.Unmarshal(qstn.QuestionBlob, &mcqQuestion); err != nil {
				return nil, err
			}
			// Marshal to JSON
			questionJsonBytes, err := protojson.Marshal(&mcqQuestion)
			if err != nil {
				return nil, err
			}
			questionJsonString = string(questionJsonBytes)

		default:
			// Unmarshal short answer question
			var textInputQuestion question.TextInputQuestion
			if err := proto.Unmarshal(qstn.QuestionBlob, &textInputQuestion); err != nil {
				return nil, err
			}
			// Marshal to JSON
			questionJsonBytes, err := protojson.Marshal(&textInputQuestion)
			if err != nil {
				return nil, err
			}
			questionJsonString = string(questionJsonBytes)
		}

		// Construct the return model
		questionReturnModel := question.QuestionReturn{
			QuestionId:        qstn.QuestionId,
			QuestionType:      qstn.QuestionType,
			QuestionSerialized: questionJsonString,
			CreatedTime:         qstn.CreatedTime,
			UpdatedTime:         qstn.UpdatedTime,
			QuestionTitle:     qstn.QuestionTitle,
		}

		questionList = append(questionList, questionReturnModel)
	}

	// Return the final response=
  fmt.Println("questionList",questionList)
		return questionList, nil
}


func (c *QuestionController) GetQuestionsByQuestionTitleSearch(ctx *gin.Context) (*question.GetQuestionsByQuestionTitleSearchResponse, error) {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return nil, err
	}
	userId := userSession.UserId
  moduleId,err := c.getModuleIdFromContextParams(ctx)
	if err != nil {
		return nil, err
	}
	documentId,err := c.getDocumentIdFromContextParams(ctx)
	if err != nil {
		return nil, err
	}
  queryItems,err := NewQueryItems(ctx)
	if err != nil {
		return nil, err
	}
	var req question.GetQuestionsByQuestionTitleSearchRequest
	ctx.Bind(&req)
	req.UserId = userId
  req.ModuleId = moduleId
  req.DocId = documentId
  req.PageNumber = queryItems.PageNumber
  req.PageSize = c.pageSize
  req.OrderByField = queryItems.OrderByField
  req.OrderByDirection = queryItems.OrderByDirection
	return c.questionService.GetQuestionsByQuestionTitleSearch(ctx, &req)
}

func (c *QuestionController) UpdateQuestion(ctx *gin.Context) error {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return err
	}
	userId := userSession.UserId
  moduleId,err := c.getModuleIdFromContextParams(ctx)
	if err != nil {
		return err
	}
	documentId,err := c.getDocumentIdFromContextParams(ctx)
	if err != nil {
		return err
	}
  questionId,err := c.getQuestionIdFromContextParams(ctx)
  if err != nil {
    return err
  }
	var req question.UpdateQuestionRequest
	ctx.Bind(&req)
	req.UserId = userId
  req.ModuleId = moduleId
  req.DocId = documentId
  req.QuestionId = questionId
	return c.questionService.UpdateQuestion(ctx, &req)
}

func (c *QuestionController) DeleteQuestion(ctx *gin.Context) error {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return err
	}
	userId := userSession.UserId
  moduleId,err := c.getModuleIdFromContextParams(ctx)
	if err != nil {
		return err
	}
	documentId,err := c.getDocumentIdFromContextParams(ctx)
	if err != nil {
		return err
	}
  questionId,err := c.getQuestionIdFromContextParams(ctx)
  if err != nil {
    return err
  }
	var req question.DeleteQuestionRequest
	ctx.Bind(&req)
	req.UserId = userId
  req.ModuleId = moduleId
  req.DocId = documentId
  req.QuestionId = questionId
	return c.questionService.DeleteQuestion(ctx, &req)
}


func (c *QuestionController) getModuleIdFromContextParams(ctx *gin.Context) (uint64, error) {
	moduleId, err := getUint64FromString(ctx.Param("module_id"))
	if err != nil {
		return 0, err
	}
	return moduleId, nil
}

func (c *QuestionController) getDocumentIdFromContextParams(ctx *gin.Context) (uint64, error) {
	documentId, err := getUint64FromString(ctx.Param("document_id"))
	if err != nil {
		return 0, err
	}
	return documentId, nil
}

func (c *QuestionController) getQuestionIdFromContextParams(ctx *gin.Context) (uint64, error) {
	documentId, err := getUint64FromString(ctx.Param("question_id"))
	if err != nil {
		return 0, err
	}
	return documentId, nil
}