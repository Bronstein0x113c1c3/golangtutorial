// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: embedded_assistant.proto

package embedded

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EmbeddedAssistantClient is the client API for EmbeddedAssistant service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmbeddedAssistantClient interface {
	// Initiates or continues a conversation with the embedded Assistant Service.
	// Each call performs one round-trip, sending an audio request to the service
	// and receiving the audio response. Uses bidirectional streaming to receive
	// results, such as the `END_OF_UTTERANCE` event, while sending audio.
	//
	// A conversation is one or more gRPC connections, each consisting of several
	// streamed requests and responses.
	// For example, the user says *Add to my shopping list* and the Assistant
	// responds *What do you want to add?*. The sequence of streamed requests and
	// responses in the first gRPC message could be:
	//
	// *   AssistRequest.config
	// *   AssistRequest.audio_in
	// *   AssistRequest.audio_in
	// *   AssistRequest.audio_in
	// *   AssistRequest.audio_in
	// *   AssistResponse.event_type.END_OF_UTTERANCE
	// *   AssistResponse.speech_results.transcript "add to my shopping list"
	// *   AssistResponse.dialog_state_out.microphone_mode.DIALOG_FOLLOW_ON
	// *   AssistResponse.audio_out
	// *   AssistResponse.audio_out
	// *   AssistResponse.audio_out
	//
	// The user then says *bagels* and the Assistant responds
	// *OK, I've added bagels to your shopping list*. This is sent as another gRPC
	// connection call to the `Assist` method, again with streamed requests and
	// responses, such as:
	//
	// *   AssistRequest.config
	// *   AssistRequest.audio_in
	// *   AssistRequest.audio_in
	// *   AssistRequest.audio_in
	// *   AssistResponse.event_type.END_OF_UTTERANCE
	// *   AssistResponse.dialog_state_out.microphone_mode.CLOSE_MICROPHONE
	// *   AssistResponse.audio_out
	// *   AssistResponse.audio_out
	// *   AssistResponse.audio_out
	// *   AssistResponse.audio_out
	//
	// Although the precise order of responses is not guaranteed, sequential
	// `AssistResponse.audio_out` messages will always contain sequential portions
	// of audio.
	Assist(ctx context.Context, opts ...grpc.CallOption) (EmbeddedAssistant_AssistClient, error)
}

type embeddedAssistantClient struct {
	cc grpc.ClientConnInterface
}

func NewEmbeddedAssistantClient(cc grpc.ClientConnInterface) EmbeddedAssistantClient {
	return &embeddedAssistantClient{cc}
}

func (c *embeddedAssistantClient) Assist(ctx context.Context, opts ...grpc.CallOption) (EmbeddedAssistant_AssistClient, error) {
	stream, err := c.cc.NewStream(ctx, &EmbeddedAssistant_ServiceDesc.Streams[0], "/google.assistant.embedded.v1alpha2.EmbeddedAssistant/Assist", opts...)
	if err != nil {
		return nil, err
	}
	x := &embeddedAssistantAssistClient{stream}
	return x, nil
}

type EmbeddedAssistant_AssistClient interface {
	Send(*AssistRequest) error
	Recv() (*AssistResponse, error)
	grpc.ClientStream
}

type embeddedAssistantAssistClient struct {
	grpc.ClientStream
}

func (x *embeddedAssistantAssistClient) Send(m *AssistRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *embeddedAssistantAssistClient) Recv() (*AssistResponse, error) {
	m := new(AssistResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EmbeddedAssistantServer is the server API for EmbeddedAssistant service.
// All implementations must embed UnimplementedEmbeddedAssistantServer
// for forward compatibility
type EmbeddedAssistantServer interface {
	// Initiates or continues a conversation with the embedded Assistant Service.
	// Each call performs one round-trip, sending an audio request to the service
	// and receiving the audio response. Uses bidirectional streaming to receive
	// results, such as the `END_OF_UTTERANCE` event, while sending audio.
	//
	// A conversation is one or more gRPC connections, each consisting of several
	// streamed requests and responses.
	// For example, the user says *Add to my shopping list* and the Assistant
	// responds *What do you want to add?*. The sequence of streamed requests and
	// responses in the first gRPC message could be:
	//
	// *   AssistRequest.config
	// *   AssistRequest.audio_in
	// *   AssistRequest.audio_in
	// *   AssistRequest.audio_in
	// *   AssistRequest.audio_in
	// *   AssistResponse.event_type.END_OF_UTTERANCE
	// *   AssistResponse.speech_results.transcript "add to my shopping list"
	// *   AssistResponse.dialog_state_out.microphone_mode.DIALOG_FOLLOW_ON
	// *   AssistResponse.audio_out
	// *   AssistResponse.audio_out
	// *   AssistResponse.audio_out
	//
	// The user then says *bagels* and the Assistant responds
	// *OK, I've added bagels to your shopping list*. This is sent as another gRPC
	// connection call to the `Assist` method, again with streamed requests and
	// responses, such as:
	//
	// *   AssistRequest.config
	// *   AssistRequest.audio_in
	// *   AssistRequest.audio_in
	// *   AssistRequest.audio_in
	// *   AssistResponse.event_type.END_OF_UTTERANCE
	// *   AssistResponse.dialog_state_out.microphone_mode.CLOSE_MICROPHONE
	// *   AssistResponse.audio_out
	// *   AssistResponse.audio_out
	// *   AssistResponse.audio_out
	// *   AssistResponse.audio_out
	//
	// Although the precise order of responses is not guaranteed, sequential
	// `AssistResponse.audio_out` messages will always contain sequential portions
	// of audio.
	Assist(EmbeddedAssistant_AssistServer) error
	mustEmbedUnimplementedEmbeddedAssistantServer()
}

// UnimplementedEmbeddedAssistantServer must be embedded to have forward compatible implementations.
type UnimplementedEmbeddedAssistantServer struct {
}

func (UnimplementedEmbeddedAssistantServer) Assist(EmbeddedAssistant_AssistServer) error {
	return status.Errorf(codes.Unimplemented, "method Assist not implemented")
}
func (UnimplementedEmbeddedAssistantServer) mustEmbedUnimplementedEmbeddedAssistantServer() {}

// UnsafeEmbeddedAssistantServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmbeddedAssistantServer will
// result in compilation errors.
type UnsafeEmbeddedAssistantServer interface {
	mustEmbedUnimplementedEmbeddedAssistantServer()
}

func RegisterEmbeddedAssistantServer(s grpc.ServiceRegistrar, srv EmbeddedAssistantServer) {
	s.RegisterService(&EmbeddedAssistant_ServiceDesc, srv)
}

func _EmbeddedAssistant_Assist_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EmbeddedAssistantServer).Assist(&embeddedAssistantAssistServer{stream})
}

type EmbeddedAssistant_AssistServer interface {
	Send(*AssistResponse) error
	Recv() (*AssistRequest, error)
	grpc.ServerStream
}

type embeddedAssistantAssistServer struct {
	grpc.ServerStream
}

func (x *embeddedAssistantAssistServer) Send(m *AssistResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *embeddedAssistantAssistServer) Recv() (*AssistRequest, error) {
	m := new(AssistRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EmbeddedAssistant_ServiceDesc is the grpc.ServiceDesc for EmbeddedAssistant service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmbeddedAssistant_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.assistant.embedded.v1alpha2.EmbeddedAssistant",
	HandlerType: (*EmbeddedAssistantServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Assist",
			Handler:       _EmbeddedAssistant_Assist_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "embedded_assistant.proto",
}
