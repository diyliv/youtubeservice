package main

import (
	"log"
	"os"

	"github.com/diyliv/youtubeservice/configs"
	"github.com/diyliv/youtubeservice/internal/server"
	"github.com/diyliv/youtubeservice/pkg/logger"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Error while reading .env :%v\n", err)
	}

	YToken := os.Getenv("YT_TOKEN")

	logger := logger.InitLogger()
	config := configs.ReadConfig()

	server := server.NewServer(logger, config, YToken)
	server.RunGRPC()
}
