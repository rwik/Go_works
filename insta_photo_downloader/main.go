package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
    "log"
	"image/png"
	"image/jpeg"
	"flag"
)

var userName = flag.String("user", "", "The user name")
var numOfWorkers = flag.Int("workers", 2, "The number of workers")


func init() {
}

func main() {
	flag.Parse()
	if *userName == "" {
		fmt.Println("Please provide a user name")
		os.Exit(1)
	}
	var userHolder = *userName

	//

}