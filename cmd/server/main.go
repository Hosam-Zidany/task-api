package main

import (
	"fmt"
	"log"

	"github.com/Hosam-Zidany/task-api/internal/server"
)

func main() {

	cfg := server.LoadConfig()
	r := server.SetupRouter()
	server.InintDB(cfg)
	addr := fmt.Sprintf(":%s", cfg.Port)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
