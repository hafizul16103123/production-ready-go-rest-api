package main

import(
	"net/http"
	"log"
	"time"
	"os"
	"os/signal"
	"syscall"
	"context"
)

func main(){
	mux:=http.NewServeMux()
	mux.HandleFunc("GET /",homeHandler)

	server:=&http.Server{
		Addr:":8000",
		Handler:mux,
		ReadTimeout: 5*time.Second,
		WriteTimeout: 5* time.Second,
		IdleTimeout:       60 * time.Second,
    	MaxHeaderBytes:    1 << 20,
	}
	go func(){

		if err:=server.ListenAndServe() ; err!=nil{
			log.Fatal(err)
		}

	log.Println("Server started on http://localhost:8080")
	}()
	// stop channel
	stop:=make(chan os.Signal,1)
	signal.Notify(stop,os.Interrupt,syscall.SIGTERM)
	<-stop

	ctx, cancel:=context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	if err:=server.Shutdown(ctx); err!=nil{
		log.Fatal(err)
	}
	log.Println("Server gracefully sutdown")

}

func homeHandler(w http.ResponseWriter, r *http.Request){
	log.Println("Method:", r.Method)
	log.Println("Path:", r.URL.Path)
	log.Println("Remote:", r.RemoteAddr)
	time.Sleep(10*time.Second)
	w.Write([]byte("Welcome to REST API"))
}