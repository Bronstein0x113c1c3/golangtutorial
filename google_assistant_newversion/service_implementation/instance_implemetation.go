package service_implementation

import (
	"context"
	pb "project_a/embedded"
	"time"

	"google.golang.org/grpc"
)

type Assistant_Interface interface {
	Init()
	NewConnection(ctx context.Context) (*grpc.ClientConn, error)
	NewConversation(timeout time.Duration) (*Conversation, error)
	Close()
}
type Conversation_Interface interface {
	Query(query string) *pb.AssistResponse
}
