package implementation

import (
	"io"
	"log"
	"net"
	"os"
	pb "project16/file_transfer"
	"strconv"

	"google.golang.org/grpc"
)

func New(address string, port int16) *Server_Struct {
	return &Server_Struct{
		DestIP:   address,
		DestPort: port,
	}
}
func (server *Server_Struct) InitDaConnection() {
	lis, _ := net.Listen("tcp", server.DestIP+":"+strconv.Itoa(int(server.DestPort)))
	log.Printf("Server created: %v:%v \n", server.DestIP, server.DestPort)
	s := grpc.NewServer()
	pb.RegisterFile_TransferServer(s, server)
	log.Println("Registered!")
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
func (server *Server_Struct) SendFile(stream pb.File_Transfer_SendFileServer) error {
	for {
		file_data, err := stream.Recv()
		if err == io.EOF {
			server.ProcessFile(file_data, "./workdir")
			status := &pb.Status{
				Status: true,
			}
			stream.SendAndClose(status)
		}
	}
}
func (server *Server_Struct) ProcessFile(data *pb.File_Info, dir string) {
	os.WriteFile(dir+data.GetName(), data.GetData(), 0644)
}

// func (s *Server_Struct) Self() *Server_Struct {
// 	return s
// }
