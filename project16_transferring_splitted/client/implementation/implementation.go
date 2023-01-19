package implementation

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	pb "project16/file_transfer"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func New(ip string, port int16, dir string) *Client_Struct {
	return &Client_Struct{
		DestIP:   ip,
		DestPort: port,
		Dir:      dir,
	}
}
func (client *Client_Struct) InitDaConnection() {
	conn, err := grpc.Dial(client.DestIP+":"+strconv.Itoa(int(client.DestPort)), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client.Connection = conn
}
func (client *Client_Struct) Transfer(name string) {
	log.Println("Stage 1: Done!")
	cc := pb.NewFile_TransferClient(client.Connection)
	ctx := context.Background()
	instance, err := cc.SendFile(ctx)
	if instance == nil {
		return
	}
	if err != nil {
		panic(err)
	}
	log.Println("Stage 2: Done!")
	max_chunk := make([]byte, 2048) //maximum of chunk is 2MB
	file_data, err := os.Open(client.Dir + "/" + name)
	// if file_data == nil {
	// 	fmt.Println("We got problem")
	// 	return
	// } else {
	// 	io.Copy(os.Stdin, file_data)
	// }
	// if err != nil {
	// 	panic(err)
	// }
	log.Println("Stage 3: Done!")
	for {
		len, err := file_data.Read(max_chunk)
		if err == io.EOF {
			fmt.Println("Read done!")
			break
		}
		// if err != nil {
		// 	panic(err)
		// }
		instance.Send(
			&pb.File_Info{
				Data: max_chunk[:len],
				Name: name,
			},
		)
		fmt.Println(max_chunk[:len])
	}
	log.Printf("Stage 4: Done!")
	resp, err := instance.CloseAndRecv()
	fmt.Println(resp)
}
