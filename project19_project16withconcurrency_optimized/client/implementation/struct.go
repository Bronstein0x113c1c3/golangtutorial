package implementation

import (
	"google.golang.org/grpc"
	// pb "project16/file_transfer"
)

type Client_Struct struct {
	DestIP     string
	DestPort   int16
	Dir        string
	Connection *grpc.ClientConn
	Client_Interface
}
