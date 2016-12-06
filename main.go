package main

import (
	"fmt"
	"os"

	"github.com/yuderekyu/expresso-subscription/router"
)

func main() {
	b, err := router.New()
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