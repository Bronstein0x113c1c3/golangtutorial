package implementation

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	pb "project16/file_transfer"
	"strconv"
	"sync"

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
	var wg sync.WaitGroup
	chunk_chan := make(chan []byte)
	wg.Add(2)
	go func() {

		data, _ := ioutil.ReadFile(client.Dir + "/" + name)
		chunk_size := 1024 * 1024 * 4
		i := 0
		for {
			if len(data) < chunk_size*(i+1) {
				chunk_chan <- data[chunk_size*i : len(data)]
				close(chunk_chan)
				break
			} else {
				chunk_chan <- data[chunk_size*i : chunk_size*(i+1)]
			}
			i++
		}
		wg.Done()
	}()
	// if err != nil {
	// 	panic(err)
	// }
	go func() {
		for chunk := range chunk_chan {
			instance.Send(
				&pb.File_Info{
					Data: chunk,
					Name: name,
				},
			)
		}
		wg.Done()
	}()
	wg.Wait()
	//Another apporach

	// var wg sync.WaitGroup
	// // chunk_chan := make(chan []byte)
	// data, _ := ioutil.ReadFile(client.Dir + "/" + name)
	// // fmt.Println(data)
	// /*First, I want to read all the file data first, to make sure that
	// I transmit all the data without any damage!

	// "data" is the bunch of bytes that I want to transmit!
	// */

	// chunk_size := 1024 //this means that I want to transmit 1KB chunk of data.
	// /*
	// 	1024B = 1KB, this simple exchange!
	// */

	// data_chunk_slice := make([]byte, 0)
	// /*
	// 	This chunk of data is used to send to client with the chunk size!

	// 	References: Slices in Go

	// */
	// /*
	// 	The game started here!
	// */

	// for i := 0; i < len(data); i++ {
	// 	/*Traverse through the "data" with location!*/
	// 	/* I'm trying to append the "point" of data into the data_chunk_slice */
	// 	data_chunk_slice = append(data_chunk_slice, data[i])
	// 	/*When it is enough, the length of data_chunk is the chunk_size */
	// 	if len(data_chunk_slice) == chunk_size {

	// 		/*Start the monster :))*/
	// 		wg.Add(1)

	// 		/*I will create a goroutine (thread) to send to the client*/

	// 		/*It's the same as multithreading, but lighter and faster!*/
	// 		go func() {
	// 			instance.Send(
	// 				&pb.File_Info{
	// 					Data: data_chunk_slice,
	// 					Name: name,
	// 				},
	// 			)
	// 			//We reset the slice!
	// 			data_chunk_slice = make([]byte, 0)
	// 			wg.Done()
	// 		}()
	// 	} else {
	// 		if i == len(data) {
	// 			/*Start the monster :))*/
	// 			wg.Add(1)

	// 			/*I will create a goroutine (thread) to send to the client*/

	// 			/*It's the same as multithreading, but lighter and faster!*/
	// 			go func() {
	// 				instance.Send(
	// 					&pb.File_Info{
	// 						Data: data_chunk_slice,
	// 						Name: name,
	// 					},
	// 				)
	// 				//We reset the slice!
	// 				data_chunk_slice = make([]byte, 0)
	// 				wg.Done()
	// 			}()
	// 		}
	// 	}
	// }
	// wg.Wait()
	/*Just another approach!*/
	// var wg sync.WaitGroup
	// chunk_size := 1024
	// chunk_slice := make([]byte, chunk_size)
	// file, _ := os.Open(client.Dir + "/" + name)
	// for {
	// 	len, err := file.Read(chunk_slice)
	// 	fmt.Println(chunk_slice[:len])
	// 	wg.Add(1)
	// 	go func() {
	// 		instance.Send(
	// 			&pb.File_Info{
	// 				Data: chunk_slice[:len],
	// 				Name: name,
	// 			},
	// 		)
	// 		wg.Done()
	// 	}()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// }
	// wg.Wait()
	log.Printf("Stage 4: Done!")
	//Now, it's done!
	resp, err := instance.CloseAndRecv()
	fmt.Println(resp)
}

//so let me test with the file, with "super-concurrency" mode in splitting files :)))
//this project is kind of a "goodbye" to 2022, as a boy wants to discover himself again
/*
	Mechanism: I mentioned before!
	It's my homework to myself, as a goodbye to myself of 2022.
	I will post it to GitHub after modifications, so stay tuned!

	Thanks for watching!
*/
