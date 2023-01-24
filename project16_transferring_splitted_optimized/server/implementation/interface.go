package implementation

import (
	pb "project16/file_transfer"
)

type Server_Interface interface {
	pb.File_TransferServer
	// Self() *Server_Struct
	InitDaConnection()
}
