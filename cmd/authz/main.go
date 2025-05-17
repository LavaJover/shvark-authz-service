package main

import (
	"fmt"

	"github.com/LavaJover/shvark-authz-service/internal/config"
)

func main(){
	cfg := config.MustLoad()
	fmt.Println(cfg)
}