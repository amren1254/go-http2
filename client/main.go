package main

import (
	"fmt"
	"net/http"
)

const url = "https://localhost:4777"

func main() {
	_, err := http.Get(url)
	fmt.Println(err)
}
