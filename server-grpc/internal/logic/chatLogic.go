package logic

import (
	"context"
	"encoding/json"

	"langriver_service/internal/svc"

	pb "langriver_service/proto"

	llms "github.com/tmc/langchaingo/llms"
	chain "github.com/tmc/langchaingo/llms/openai"
	"github.com/zeromicro/go-zero/core/logx"
)

type ChatbotLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	model  *chain.LLM
}

func NewChatbotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatbotLogic {

	llm, err := chain.New(chain.WithToken(svcCtx.Config.OpenaiClient.Token), chain.WithModel(svcCtx.Config.OpenaiClient.Model), chain.WithBaseURL(svcCtx.Config.OpenaiClient.BaseURL))
	if err != nil {
		logx.Errorf("large model import error: %v", err)
		return nil
	}
	return &ChatbotLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		model:  llm,
	}
}

// 对话逻辑
func (c *ChatbotLogic) Chat(req *pb.ChatRequest) (*pb.ChatResponse, error) {
	//chat role and part
	message := make([]llms.MessageContent, 10)
	var index int = 0
	if index == 9 {
		index = 0
	}
	//create role instance
	message[index].Role = "human"
	//create text instance
	text := llms.TextContent{Text: req.UserInput}
	message[index].Parts = []llms.ContentPart{text}
	//input session
	respose, err := c.model.GenerateContent(c.ctx, message)
	if err != nil {
		logx.Errorf("Session error: %v", err)
		return nil, err

	}
	jsonRespponse, err := json.Marshal(respose)
	if err != nil {
		logx.Errorf("struct parse error: %v", err)
		result := respose.Choices[0].Content
		return &pb.ChatResponse{
			BotResponse: result,
		}, err
	}
	jsonRes := string(jsonRespponse)
	return &pb.ChatResponse{
		BotResponse: jsonRes,
	}, nil
}
