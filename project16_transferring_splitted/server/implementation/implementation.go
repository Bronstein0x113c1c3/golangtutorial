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

func New(address string, port int16, dir string) *Server_Struct {
	return &Server_Struct{
		DestIP:   address,
		DestPort: port,
		Dir:      dir,
	}
}
func (server *Server_Struct) InitDaConnection() {
	lis, _ := net.Listen("tcp", server.DestIP+":"+strconv.Itoa(int(server.DestPort)))
	log.Printf("Server created: %v:%v \n", server.DestIP, server.DestPort)
	log.Printf("Working directory: %v \n", server.Dir)
	s := grpc.NewServer()
	pb.RegisterFile_TransferServer(s, server)
	log.Println("Registered!")
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
func (server *Server_Struct) SendFile(stream pb.File_Transfer_SendFileServer) error {
	log.Println("Got request!")
	chunk_file := make([]byte, 0)
	status := &pb.Status{}
	file_name := ""
	for {
		file_data, err := stream.Recv()
		// log.Println(err)
		// log.Println("hmmm")
		if file_data == nil {
			status.Status = true
			break
		}
		chunk_file = append(chunk_file, file_data.Data...)
		file_name = file_data.Name
		if err != nil {
			status.Status = false
			break
		}
	}
	stream.SendAndClose(status)
	return server.ProcessFile(chunk_file, file_name)
}

// log.Println("Got problem???")
// return nil

func (server *Server_Struct) ProcessFile(data []byte, name string) error {
	if err := os.WriteFile(server.Dir+"/"+name, data, 0644); err != nil && err != io.EOF {
		return err
	}
	log.Println("Done!")
	return nil
}

// func (s *Server_Struct) Self() *Server_Struct {
// 	return s
// }
