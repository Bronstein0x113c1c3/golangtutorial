package implementation

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	pb "project16/file_transfer"
	"runtime"
	"strconv"
	"strings"

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
	//For step 1
	chunk_list := make([]string, 0)
	process := 0

	//For step 2
	data := make([]byte, 0)
	file_name := ""
	//For the status:
	status := &pb.Status{}
	for {
		/*
			This problem is that I could not optimize or split the processes.
				so, after analyzing, i decided to split into 2 processes:
					1.Receive the chunk
						1.1: Split the file name into: name, extension
						1.2: Write the chunk into temporary files
						1.3: Get it a name: name + index
						1.4: Put the name of temp into a slices
					2.Combining:
						2.1: Create the data slice
						2.2: From the slice with name:
							2.2.1: Read the file, get the data from temp
							2.2.2: Combine to data slice
							2.2.3: Delete the temp after doing that
		*/

		// chunk_file = append(chunk_file, file_data.Data...)
		file_info, err := stream.Recv()
		if file_info == nil {
			status.Status = true
			break
		}
		file_name = file_info.Name
		{
			name_of_the_file := fmt.Sprintf("%v/%v", server.Dir, fmt.Sprintf("%v%v", strings.Split(file_info.Name, ".")[0], process))
			os.WriteFile(name_of_the_file, file_info.Data, 0644)
			chunk_list = append(chunk_list, name_of_the_file)
		}

		// server.ProcessTheChunk(file_info, chunk_list, process)
		_ = file_name
		process++
		log.Printf("%v parts received! \n", process)
		if err != nil {
			status.Status = false
			break
		}
		// fmt.Println(file_name)
	}

	log.Println("Phase 1 - Receiving chunk: Done!")
	runtime.GC()
	stream.SendAndClose(status)
	for _, chunk := range chunk_list {
		data = append(data, server.ReadFromChunk(chunk)...)
		os.Remove(chunk)
		log.Printf("Delete %v: Done! \n", chunk)
	}
	log.Println("Phase 2 - Combining chunk: Done!")
	return os.WriteFile(fmt.Sprintf("%v/%v", server.Dir, file_name), data, 0644)

	// return server.Combine(chunk_list, data, file_name)
}
func (server *Server_Struct) ReadFromChunk(chunk_name string) []byte {
	data_from_chunk, err := ioutil.ReadFile(chunk_name)
	if err != nil {
		return make([]byte, 0)
	}
	return data_from_chunk
}

// log.Println("Got problem???")
// return nil

// func (s *Server_Struct) Self() *Server_Struct {
// 	return s
// }

// func (server *Server_Struct) ProcessTheChunk(file_info *pb.File_Info, chunk_list []string, process_index int) error {
// 	name_of_the_file := fmt.Sprintf("%v/%v", server.Dir, fmt.Sprintf("%v%v", strings.Split(file_info.Name, ".")[0], process_index))
// 	err := os.WriteFile(name_of_the_file, file_info.Data, 0644)
// 	if err != nil {
// 		return err
// 	}
// 	chunk_list = append(chunk_list, name_of_the_file)
// 	log.Printf("%v: Downloaded!", name_of_the_file)
// 	return nil
// }
// func (server *Server_Struct) Combine(chunk_list []string, data []byte, file_name string) error {
// 	for _, chunk_name := range chunk_list {
// 		chunk, err := ioutil.ReadFile(chunk_name)
// 		if err != nil {
// 			return err
// 		}
// 		data = append(data, chunk...)
// 		// os.Remove(chunk_name)
// 	}
// 	os.WriteFile(fmt.Sprintf("%v/%v", server.Dir, file_name), data, 0644)
// 	return nil
// }
