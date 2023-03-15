package implementation

import (
	"client_transfer/proto/file_transfer"
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func New(DestIP string, DestPort string) FT_Service_Interface {
	return &FT_Service_Information{
		DestIP:   DestIP,
		DestPort: DestPort,
	}
}
func (ftsi *FT_Service_Information) Exit() {
	ftsi.Connection.Close()
}
func (ftsi *FT_Service_Information) ReturnSelf() *FT_Service_Information {
	return ftsi
}
func (ftsi *FT_Service_Information) Init() {
	//Create the connection
	conn, err := grpc.Dial(ftsi.DestIP+":"+ftsi.DestPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	ftsi.Connection = conn
	//Create the operating client instance
	ftsi.Operating_Instance = file_transfer.NewFile_TransferClient(ftsi.Connection)
}
func (ftsi *FT_Service_Information) GetTheFile(name string) *file_transfer.File_Info {
	//get the file information
	data, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println("err encountered")
	}
	//return the file data
	return &file_transfer.File_Info{
		Name: name,
		Data: data,
	}
}

func (ftsi *FT_Service_Information) Transfer() bool {
	/*
		Transfer file including:
			Operating_Instance: to send the file info to the server
			Input: Use the GetTheFile data generated from the file
			Context: context to do the method calls
	*/
	ctx := context.Background()
	_, err := ftsi.Operating_Instance.SendFile(ctx, ftsi.Input)
	if err != nil {
		log.Fatalln(err)
		return false
	}
	fmt.Println("transfer successful")
	return true
}
