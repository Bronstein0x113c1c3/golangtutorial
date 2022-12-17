package implementation

import (
	"client_transfer/proto/file_transfer"
	"context"

	"google.golang.org/grpc"
)

type FT_Service_Information struct {

	//initial infomation
	DestIP   string
	DestPort string
	//create the connection
	Connection *grpc.ClientConn
	//Create the operating instance from the connection, use this shit to transfer
	Context            context.Context
	Operating_Instance file_transfer.File_TransferClient
	//File information generated from the name and file content
	Input *file_transfer.File_Info
}
