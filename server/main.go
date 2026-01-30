package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/pawannn/juzchat/config"
	"github.com/pawannn/juzchat/controllers"
	"github.com/pawannn/juzchat/middlewares"
	"github.com/pawannn/juzchat/service"
)

func main() {
	cfg := config.LoadConfig()

	rdb := config.NewRedis(cfg)

	chatHub := service.NewChatHub(rdb)
	go chatHub.Run()

	controller := controllers.InitControllers(chatHub)

	mux := http.NewServeMux()
	mux.HandleFunc("/chat", controller.HandleConnection)
	mux.HandleFunc("/chats", controller.FetchAvailableChats)

	handler := middlewares.CorsMiddleware(mux)

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		log.Println("Shutting down JuzChat server...")
		os.Exit(0)
	}()

	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Println("Listening on port:", cfg.Port)

	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal(err)
	}

}
