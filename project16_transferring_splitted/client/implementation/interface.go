package implementation

// pb "project16/file_transfer"

type Client_Interface interface {
	InitDaConnection()
	// pb.File_TransferClient
	Transfer(name string)
}
