package baseline

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func detectFiles(u *string) {
	list := [...]string{"robots.txt", "crossdomain.txt", "README.md", "admin.php", "admin.asp", "admin.jsp", "admin.aspx", "admin/"}
	for _,v := range list {
		entry := *u+ "/" +v
		req, _ := http.NewRequest("GET", entry, nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36")
		resp, e := (&http.Client{}).Do(req)
		if e != nil {
			panic(e)
		}
		defer resp.Body.Close()
		if resp.StatusCode == 200 {
			body, e := ioutil.ReadAll(resp.Body)
			if e != nil {
				panic(e)
			}
			str := string(body)
			if len(str) > 500 {
				str = str[:500]
			}
			fmt.Println("[*] Detected "+ entry)
			fmt.Println(str)
		}
	}
}

func directoryListing(u *string) bool {
	entry := *u
	resp, e := http.Get(entry)
	if e != nil {
		panic(e)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, e := ioutil.ReadAll(resp.Body)
		if e != nil {
			panic(e)
		}
		if strings.Contains(string(body), "Directory listing for")  {
			log.Println("Detected Directory List.", entry)
			return true
		}
	}
	return false
}