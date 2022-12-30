package implementation

import (
	"context"
	"net"
	"os"
	pb "server_transfer/proto/file_transfer"

	"google.golang.org/grpc"
)

func (ftsi *FT_Service_Information) GetServerInfo() string {
	return ftsi.IP + ":" + ftsi.Port
}
func New(IP string, Port string) *FT_Service_Information {
	return &FT_Service_Information{
		IP:   IP,
		Port: Port,
	}
}
func (ftsi *FT_Service_Information) InitDaConnection() {
	lis, _ := net.Listen("tcp", ftsi.GetServerInfo())
	s := grpc.NewServer()
	pb.RegisterFile_TransferServer(s, ftsi)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
func (ftsi *FT_Service_Information) SendFile(ctx context.Context, file_info *pb.File_Info) (*pb.Status, error) {
	data := file_info.Data
	err := os.WriteFile(file_info.Name, data, 0644)
	if err != nil {
		return &pb.Status{
			Status: false,
		}, err
	}
	return &pb.Status{
		Status: true,
	}, nil
}
