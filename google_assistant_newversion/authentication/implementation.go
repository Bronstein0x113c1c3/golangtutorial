package authentication

import (
	"context"
	_ "context"
	"fmt"

	"net/http"
	_ "net/http"

	"golang.org/x/oauth2"
	_ "golang.org/x/oauth2"
)

func New() Authentication {
	return &Authentication_data{}
}

func (v *Authentication_data) StartAuthenticateServer() {
	//var wg sync.WaitGroup
	// wg.Add(1)
	var ctx context.Context = context.Background()
	url := v.Config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Please use this address to authenticate: %v \n", url)
	var code string
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	log.Println("server is running")
	// 	code = r.URL.Query().Get("code")
	// 	w.Write([]byte("Authentication Done!"))
	// 	// wg.Done()
	// 	return
	// })
	// _ = http.ListenAndServe(Host_Authen, nil)

	// wg.Wait()
	oauthSrv := &http.Server{
		Addr:    ":8090",
		Handler: http.DefaultServeMux,
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		code = r.URL.Query().Get("code")
		w.Write([]byte(fmt.Sprintf("<h1>Your code is: %s</h1>", code)))
		fmt.Println("done!")
		oauthSrv.Shutdown(context.Background())
	})
	oauthSrv.ListenAndServe()
	//fmt.Println(code)
	var err error
	v.Token, err = v.Config.Exchange(ctx, code)
	if err != nil {
		panic(err)
	}
	return
}
func (v *Authentication_data) Load_Info(clid string, clisec string) {
	v.Config = &oauth2.Config{
		ClientID:     clid,
		ClientSecret: clisec,
		RedirectURL:  "http://localhost:8090",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://accounts.google.com/o/oauth2/token",
		},
		Scopes: []string{"https://www.googleapis.com/auth/assistant-sdk-prototype"},
	}
}
func (v *Authentication_data) Return_Info() *Authentication_data {
	return v
}
