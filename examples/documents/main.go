package main

import (
	"github.com/BANKEX/go-primetrust"
	"log"
	"os"
)

func main() {
	primetrust.Init(true, os.Getenv("PRIMETRUST_LOGIN"), os.Getenv("PRIMETRUST_PASSWORD"))
	log.Println("Done")
}
