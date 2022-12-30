package implementation

import (
	pb "server_transfer/proto/file_transfer"
)

type FT_Service_Information struct {
	IP     string
	Port   string
	Output *pb.File_Info
	FT_Service_Interface
}
