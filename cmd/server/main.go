package main

import (
	"autoreload/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func writefile() {
	data := model.RandomValueStatus()
	file, _ := json.MarshalIndent(data, "", "")
	err := ioutil.WriteFile("html/weatherstatus.json", file, 0644)
	if err != nil {
		fmt.Printf("%v \n", err)
	}
}

func main() {
	ticker := time.NewTicker(14 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				writefile()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	http.Handle("/", http.FileServer(http.Dir("./html")))
	http.ListenAndServe(":8080", nil)

}
