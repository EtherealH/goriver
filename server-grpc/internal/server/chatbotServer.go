package server

import (
	"context"
	"langriver_service/internal/logic"
	"langriver_service/internal/svc"
	pb "langriver_service/proto"
)

type ChatbotServer struct {
	pb.UnimplementedChatbotServiceServer
	svcCtx *svc.ServiceContext
}

func NewChatbotServer(svcCtx *svc.ServiceContext) *ChatbotServer {
	return &ChatbotServer{
		svcCtx: svcCtx,
	}
}

func (s *ChatbotServer) Chat(ctx context.Context, req *pb.ChatRequest) (*pb.ChatResponse, error) {
	l := logic.NewChatbotLogic(ctx, s.svcCtx)
	return l.Chat(req)
}
