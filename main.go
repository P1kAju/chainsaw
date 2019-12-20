package main

/**
	Chainsaw, a web audit tool.
 */

import (
	"bufio"
	"chainsaw/baseline"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var arg_file = flag.String("f", "", "Specify a file path.")

type Proxy struct {

}

func (p *Proxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("Received request %s %s %s\n", req.Method, req.Host, req.RemoteAddr)
	transport :=  http.DefaultTransport
	outReq := new(http.Request)
	*outReq = *req
	if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
		if prior, ok := outReq.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(prior, ", ") + ", " + clientIP
		}
		outReq.Header.Set("X-Forwarded-For", clientIP)
	}
	res, err := transport.RoundTrip(outReq)
	if err != nil {
		rw.WriteHeader(http.StatusBadGateway)
		return
	}
	for key, value := range res.Header {
		for _, v := range value {
			rw.Header().Add(key, v)
		}
	}
	rw.WriteHeader(res.StatusCode)
	_, _ = io.Copy(rw, res.Body)
	res.Body.Close()
}

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
	http.Handle("/", &Proxy{})
	_ = http.ListenAndServe("0.0.0.0:1234", nil)
}

func core(u string) {
	fmt.Println("[+] Working on "+ u +"...")
	entry := parseUrl(u)
	if isAlive(entry) {
		baseline.Start(entry)
	} else {
		log.Println("[*] " + u + " not alive!")
	}
	fmt.Println("[+] Done.")
}

func isAlive(u string) bool {
	req, _ := http.NewRequest("HEAD", u, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36")
	resp, e := (&http.Client{}).Do(req)
	if e != nil {
		return false
	}
	defer resp.Body.Close()
	return true
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