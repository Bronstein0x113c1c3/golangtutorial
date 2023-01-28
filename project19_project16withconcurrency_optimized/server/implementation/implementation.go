package implementation

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	pb "project16/file_transfer"
	"strconv"
	"strings"
	"sync"

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

// func (server *Server_Struct) SendFile(stream pb.File_Transfer_SendFileServer) error {
// 	log.Println("Got request!")
// 	chunk_file := make([]byte, 0)
// 	status := &pb.Status{}
// 	file_name := ""
// 	chunk_chan := make(chan []byte)
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go func() {
// 		for {
// 			file_data, err := stream.Recv()
// 			// log.Println(err)
// 			// log.Println("hmmm")

// 			if file_data == nil {
// 				close(chunk_chan)
// 				status.Status = true
// 				break
// 			}
// 			// fmt.Println(file_data.Data)
// 			if len(file_data.Data) != 0 {
// 				// chunk_file = append(chunk_file, file_data.Data...)
// 				chunk_chan <- file_data.Data
// 			}
// 			file_name = file_data.Name
// 			if err != nil {
// 				close(chunk_chan)
// 				status.Status = false
// 				break
// 			}
// 		}
// 		stream.SendAndClose(status)
// 		wg.Done()
// 	}()
// 	go func() {
// 		var i = 0
// 		for chunk := range chunk_chan {
// 			i++
// 			// server.ProcessFile(chunk, fmt.Sprintf("%v%v", file_name, i))
// 			// chunk_file = append(chunk, chunk...)

// 		}
// 		wg.Done()
// 	}()
// 	wg.Wait()
// 	// return server.ProcessFile(chunk_file, file_name)
// }

// log.Println("Got problem???")
// return nil

//	func (s *Server_Struct) Self() *Server_Struct {
//		return s
//	}
func (server *Server_Struct) SendFile(stream pb.File_Transfer_SendFileServer) error {
	log.Println("Got request!")
	packet_chan := make(chan *pb.File_Info)
	status := &pb.Status{}
	list_of_chunk := make([]string, 0)
	data := make([]byte, 0)
	name_of_file := ""
	done_status := make(chan bool)
	wg := &sync.WaitGroup{}
	var err error
	/*
		1.Receiving chunk
			1.1. Extracting from the stream: stream.Recv()
			1.2. Sending the extracted data through the channel
			1.3. Another goroutine: Writing the temp files.
		2.Combining chunk: For temp file, append the data chunk and delete the temp after then.
	*/
	wg.Add(1)
	go func() {
		var i = 0
		for packet_data := range packet_chan {
			name_of_the_chunk := fmt.Sprintf("%v/%v", server.Dir, fmt.Sprintf("%v%v", strings.Split(packet_data.Name, ".")[0], i))
			name_of_file = fmt.Sprintf("%v/%v", server.Dir, packet_data.Name)
			err := os.WriteFile(name_of_the_chunk, packet_data.Data, 0644)
			if err != nil {
				break
			}
			list_of_chunk = append(list_of_chunk, name_of_the_chunk)
			i++
		}
		done_status <- true
		wg.Done()
	}()
	for {
		packet_data, err := stream.Recv()
		if packet_data == nil {
			close(packet_chan)
			status.Status = true
			break
		}
		packet_chan <- packet_data
		if err != nil {
			close(packet_chan)
			status.Status = false
			break
		}
	}
	select {
	case <-done_status:
		for _, chunk_name := range list_of_chunk {
			chunk, _ := ioutil.ReadFile(chunk_name)
			data = append(data, chunk...)
			os.Remove(chunk_name)
			log.Printf("%v: Done!", chunk_name)
		}
		err = os.WriteFile(name_of_file, data, 0644)
	}
	return err
}
