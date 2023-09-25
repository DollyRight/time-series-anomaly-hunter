package core

import (
	"fmt"
	"net/http"
	"server/global"
	"server/initialize"
	"time"
)

func StartServer() {
	s := &http.Server{
		Addr:           global.Gva.GvaConfigs.Server.GetAddr(),
		Handler:        initialize.Router(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		panic(fmt.Errorf("fatal start server: %w", err))
	}
	fmt.Println("start server successfully")
}
