package main

import (
	"chainsaw/baseline"
	"fmt"
	"log"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		panic("You must specific a url.")
	}
	u := os.Args[1]
	fmt.Println("Working...")
	entry := parseUrl(u)
	baseline.Start(entry)
}

func parseUrl(u string) string {
	res, err := url.Parse(u)
	if err != nil {
		panic(err)
	}
	if res.Scheme!="http" && res.Scheme!="https" {
		panic("Protocol missing.")
	}
	if res.Host == "" {
		panic("Host missing.")
	}
	if res.Port() == "80" {
		log.Println("Chainsaw suggest you to remove default port 80, because this feature may affect the result.")
	}
	return res.Scheme+"://"+res.Host
}