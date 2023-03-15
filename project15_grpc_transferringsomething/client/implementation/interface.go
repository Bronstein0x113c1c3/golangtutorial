package implementation

import "client_transfer/proto/file_transfer"

type FT_Service_Interface interface {
	Init()
	GetTheFile(name string) *file_transfer.File_Info
	Transfer() bool
	ReturnSelf() *FT_Service_Information
	Exit()
}
