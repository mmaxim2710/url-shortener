package main

import (
	"fmt"
	"github.com/mmaxim2710/url-shortener/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// logger

	// storage

	// router

	// start server
}
