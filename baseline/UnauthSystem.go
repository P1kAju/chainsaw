package baseline

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

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
			log.Println("Detected Druid unauthorized.", entry)
			return true
		}
	}
	return false
}