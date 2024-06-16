package logic

import (
	"context"
	"strconv"

	"langriver-client/internal/svc"
	"langriver-client/internal/types"
	pb "langriver-client/proto"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
)

type ChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatLogic {
	return &ChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatLogic) Chat(req *types.ChatRequest) (resp *types.ChatResponse, err error) {
	// todo: add your logic here and delete this line
	port := strconv.Itoa(l.svcCtx.Config.RestConf.Port)
	addr := l.svcCtx.Config.RestConf.Host + ":" + port
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	//调用grpc服务
	client := pb.NewChatbotServiceClient(conn)
	//api数据转为rpc数据
	rpcReq := &pb.ChatRequest{
		UserInput: req.UserInput,
	}
	//返回数据
	rpcRssp, err := client.Chat(l.ctx, rpcReq)

	if err != nil {
		return nil, err
	}
	return &types.ChatResponse{
		BotResponse: rpcRssp.GetBotResponse(),
	}, nil
}
