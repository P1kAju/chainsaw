package main

/**
	Chainsaw, a web scanner.
 */

import (
	"bufio"
	"chainsaw/baseline"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
)

var arg_file = flag.String("f", "", "Specify a file path.")

func main() {
	flag.Parse()
	if len(os.Args) <= 1 {
		fmt.Println("[*] Use -help to get help.")
		os.Exit(0)
	}
	if *arg_file != "" {
		file, err := os.Open(*arg_file)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			core(scanner.Text())
		}
		os.Exit(0)
	}
	u := os.Args[1]
	core(u)
}

func core(u string) {
	fmt.Println("[+] Working...")
	entry := parseUrl(u)
	baseline.Start(entry)
	fmt.Println("[+] Done.")
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