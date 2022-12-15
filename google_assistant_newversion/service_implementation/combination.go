package service_implementation

import (
	"context"
	authentication "project_a/authentication"
	"time"

	pb "project_a/embedded"

	"google.golang.org/api/option"
	"google.golang.org/api/transport"
	"google.golang.org/grpc"
)

func (a *Assistant) Init() {
	authentication_data := authentication.New()
	authentication_data.Load_Info(a.Client_ID, a.Client_Secret)
	authentication_data.StartAuthenticateServer()

	a.Authentication_Info = authentication_data.Return_Info()
}
func (a *Assistant) NewConnection(ctx context.Context) (*grpc.ClientConn, error) {
	tokenSource := a.Authentication_Info.Config.TokenSource(ctx, a.Authentication_Info.Token)
	return transport.DialGRPC(ctx,
		option.WithTokenSource(tokenSource),
		option.WithEndpoint("embeddedassistant.googleapis.com:443"),
		option.WithScopes("https://www.googleapis.com/auth/assistant-sdk-prototype"),
	)
}
func (a *Assistant) NewConversation(timeout time.Duration) (*Conversation, error) {
	a.Contexted, a.Canceler = context.WithDeadline(context.Background(), time.Now().Add(timeout))
	clientconn, err := a.NewConnection(a.Contexted)
	a.Connection = clientconn
	if err != nil {
		return nil, err
	}
	assistant_client, err := pb.NewEmbeddedAssistantClient(clientconn).Assist(a.Contexted)
	if err != nil {
		return nil, err
	}
	return &Conversation{
		a:             a,
		Assist_Client: assistant_client,
	}, nil
}

func (a *Assistant) Close() {
	if a.Contexted != nil {
		a.Contexted.Done()
	}
	if a.Canceler != nil {
		a.Canceler()
	}
	if a.Connection != nil {
		a.Connection.Close()
	}
}

func (c *Conversation) Query(query string) *pb.AssistResponse {
	err := c.Assist_Client.Send(&pb.AssistRequest{
		Type: &pb.AssistRequest_Config{
			Config: &pb.AssistConfig{
				Type: &pb.AssistConfig_TextQuery{
					TextQuery: query,
				},
				AudioOutConfig: &pb.AudioOutConfig{
					Encoding:         pb.AudioOutConfig_LINEAR16,
					SampleRateHertz:  16000,
					VolumePercentage: 100,
				},
				DeviceConfig: &pb.DeviceConfig{
					DeviceId:      "mynewlaptop",
					DeviceModelId: "Inspiron7610",
				},
				DialogStateIn: &pb.DialogStateIn{
					ConversationState: make([]byte, 0),
					LanguageCode:      "en-US",
					IsNewConversation: true,
				},
			},
		},
	})
	if err != nil {
		return nil
	}
	resp, err := c.Assist_Client.Recv()
	return resp
}
