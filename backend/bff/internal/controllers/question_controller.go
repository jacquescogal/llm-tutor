package controllers

import (
	"bff/internal/proto/question"
	"bff/internal/services"

	"github.com/gin-gonic/gin"
)

type QuestionController struct{
	questionService *services.QuestionService
}

func NewQuestionController(questionService *services.QuestionService) *QuestionController {
	return &QuestionController{questionService: questionService}
}

func (c *QuestionController) CreateQuestion(ctx *gin.Context) error {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return err
	}
	userId := userSession.UserId
	var req question.CreateQuestionRequest
	ctx.Bind(&req)
	req.UserId = userId
	return c.questionService.CreateQuestion(ctx, &req)
}

func (c *QuestionController) GetQuestionById(ctx *gin.Context) (*question.GetQuestionByIdResponse, error) {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return nil, err
	}
	userId := userSession.UserId
	var req question.GetQuestionByIdRequest
	ctx.Bind(&req)
	req.UserId = userId
	return c.questionService.GetQuestionById(ctx, &req)
}

func (c *QuestionController) GetQuestionsByDocId(ctx *gin.Context) (*question.GetQuestionsByDocIdResponse, error) {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return nil, err
	}
	userId := userSession.UserId
	var req question.GetQuestionsByDocIdRequest
	ctx.Bind(&req)
	req.UserId = userId
	return c.questionService.GetQuestionsByDocId(ctx, &req)
}

func (c *QuestionController) GetQuestionsByQuestionTitleSearch(ctx *gin.Context) (*question.GetQuestionsByQuestionTitleSearchResponse, error) {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return nil, err
	}
	userId := userSession.UserId
	var req question.GetQuestionsByQuestionTitleSearchRequest
	ctx.Bind(&req)
	req.UserId = userId
	return c.questionService.GetQuestionsByQuestionTitleSearch(ctx, &req)
}

func (c *QuestionController) UpdateQuestion(ctx *gin.Context) error {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return err
	}
	userId := userSession.UserId
	var req question.UpdateQuestionRequest
	ctx.Bind(&req)
	req.UserId = userId
	return c.questionService.UpdateQuestion(ctx, &req)
}

func (c *QuestionController) DeleteQuestion(ctx *gin.Context) error {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return err
	}
	userId := userSession.UserId
	var req question.DeleteQuestionRequest
	ctx.Bind(&req)
	req.UserId = userId
	return c.questionService.DeleteQuestion(ctx, &req)
}



/*
// question.proto

syntax = "proto3";
import "common.proto";
package question;

option go_package = "bff/internal/proto/question";

service QuestionService {
    rpc CreateQuestion (CreateQuestionRequest) returns (CreateQuestionResponse);
    rpc GetQuestionById (GetQuestionByIdRequest) returns (GetQuestionByIdResponse);
    rpc GetQuestionsByDocId (GetQuestionsByDocIdRequest) returns (GetQuestionsByDocIdResponse);
    rpc GetQuestionsByQuestionTitleSearch (GetQuestionsByQuestionTitleSearchRequest) returns (GetQuestionsByQuestionTitleSearchResponse);
    rpc UpdateQuestion (UpdateQuestionRequest) returns (UpdateQuestionResponse);
    rpc DeleteQuestion (DeleteQuestionRequest) returns (DeleteQuestionResponse);
  }


// QuestionService
message CreateQuestionRequest {
    uint64 user_id = 1;
    uint64 doc_id = 2;
    string question_title = 3;
    bytes question_blob = 4;
    common.QuestionType question_type = 5;
  }
  
  message CreateQuestionResponse {
  }
  
  message GetQuestionByIdRequest {
    uint64 user_id = 1;
    uint64 question_id = 2;
  }
  
  message GetQuestionByIdResponse {
    DBQuestion question = 1;
  }
  
  message GetQuestionsByDocIdRequest {
    uint64 user_id = 1;
    uint64 doc_id = 2;
    uint32 page_number = 3;
    uint32 page_size = 4;
    common.ORDER_BY_FIELD order_by_field = 5;
    common.ORDER_BY_DIRECTION order_by_direction = 6;
  }
  
  message GetQuestionsByDocIdResponse {
    repeated DBQuestion questions = 1;
  }
  
  message GetQuestionsByQuestionTitleSearchRequest {
    uint64 user_id = 1;
    string search_query = 2;
    uint32 page_number = 3;
    uint32 page_size = 4;
    common.ORDER_BY_FIELD order_by_field = 5;
    common.ORDER_BY_DIRECTION order_by_direction = 6;
  }
  
  message GetQuestionsByQuestionTitleSearchResponse {
    repeated DBQuestion questions = 1;
  }
  
  message UpdateQuestionRequest {
    uint64 user_id = 1;
    uint64 question_id = 2;
    string question_title = 3;
    bytes question_blob = 4;
    common.QuestionType question_type = 5;
  }
  
  message UpdateQuestionResponse {
  }
  
  message DeleteQuestionRequest {
    uint64 user_id = 1;
    uint64 question_id = 2;
  }
  
  message DeleteQuestionResponse {
  }
  


message DBQuestion {
    uint64 question_id = 1;
    uint64 doc_id = 2;
    string question_title = 3;
    bytes question_blob = 4;
    common.QuestionType question_type = 5;
    uint64 created_time = 6;
    uint64 updated_time = 7;
  }
  
  // Question Types
  message MCQChoice {
    string choice = 1;
    bool is_correct = 2;
  }
  
  message MCQQuestion {
    repeated MCQChoice choices = 1;
  }
  
  message TextInputQuestion {
    string answer = 1;
  }
*/