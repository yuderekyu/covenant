package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/yuderekyu/expresso-subscription/config"
	"github.com/yuderekyu/expresso-subscription/router"
)

func main() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	config, err := config.Init(path.Join(dir, "config.json"))

	if err != nil {
		fmt.Printf("ERROR: config initialization error.\n%s\n", err.Error())
		return
	}
	b, err := router.New(config)
	if err != nil {
		fmt.Printf("ERROR: %s", err)
		return
	}

	port := os.Getenv("PORT");
	if port == "" {
		port = "8000"
	}

	fmt.Printf("Subscription running on %s\n", port)
	b.Start(":"+port)
}