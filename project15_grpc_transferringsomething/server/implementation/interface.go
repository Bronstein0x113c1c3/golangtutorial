package implementation

import (
	pb "server_transfer/proto/file_transfer"
)

type FT_Service_Interface interface {
	GetServerInfo() string
	InitDaConnection()
	ProcessTheFile(dir string)
	pb.File_TransferServer
}
