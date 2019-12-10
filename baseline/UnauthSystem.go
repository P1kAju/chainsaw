package baseline

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func springActuator(u *string) bool {
	list := [...]string{"/autoconfig", "/beans", "/env", "/configprops", "/dump", "/health", "/info", "/mappings", "/metrics", "/shutdown", "/trace"}
	for _, l := range list {
		entry := *u + l
		resp, e := http.Get(entry)
		if e != nil {
			panic(e)
		}
		defer resp.Body.Close()
		if resp.StatusCode == 200 {
			log.Println("[*] Detected Spring Actuator information leak.", entry)
		}
	}
	return false
}

func druid(u *string) bool {
	entry := *u+"/druid/index.html"
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
		if strings.Contains(string(body), "Druid Stat Index")  {
			log.Println("[*] Detected Druid unauthorized.", entry)
			return true
		}
	}
	return false
}

func laravelDebug(u *string) bool {
	resp, e := http.Post(*u, "", nil)
	if e != nil {
		panic(e)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 405 {
		body, e := ioutil.ReadAll(resp.Body)
		if e != nil {
			panic(e)
		}
		if strings.Contains(string(body), "MethodNotAllowedHttpException") {
			log.Println("[*] Detected Laravel debug mode.", *u)
			return true
		}
	}
	return false
}