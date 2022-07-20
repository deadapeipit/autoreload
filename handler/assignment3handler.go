package handler

import (
	"autoreload/model"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"text/template"
)

var wstat model.WeatherStatus

const htmlPath = "html/index.html"
const jsonPath = "html/weatherstatus.json"

func Assignment3Handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content Type", "text/html")
	// read from json file and write to webData
	file, _ := ioutil.ReadFile(jsonPath)
	json.Unmarshal(file, &wstat)
	templates, _ := template.ParseFiles(htmlPath)
	context := model.CompiledWeatherStatus{
		Status: model.Status{
			Water: wstat.Status.Water,
			Wind:  wstat.Status.Wind,
		},
		StatusCompiled: wstat.Status.CheckStatus(),
	}
	templates.Execute(w, context)
}
