package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/rwboyer/ginapi/mappings"
)

func terminate(s *http.Server, w *sync.WaitGroup) {

	defer w.Done()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigs
	log.Printf("Recieved signal %v", sig)
	s.Shutdown(context.TODO())

}

func main() {
	//mappings.Router.LoadHTMLGlob("templates/*.tmpl")
	mappings.CreateUrlMappings()

	srv := &http.Server{
		Addr:    ":1111",
		Handler: mappings.Router,
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go terminate(srv, wg)

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
	wg.Wait()
	log.Println("Exiting - http server terminated")
}
