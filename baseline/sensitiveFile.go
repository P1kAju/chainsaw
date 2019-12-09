package baseline

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func robots(u *string) bool {
	resp, e := http.Get(*u+"/robots.txt")
	if e != nil {
		panic(e)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, e := ioutil.ReadAll(resp.Body)
		if e != nil {
			panic(e)
		}
		fmt.Println("Detect robots.txt file returned 200.", body)
		return true
	}
	return false
}