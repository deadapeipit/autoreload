package main

import (
	"autoreload/model"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const jsonPath = "html/weatherstatus.json"

var wstat model.WeatherStatus

func Test_writefile(t *testing.T) {
	t.Log("Test write json file")
	//check is debug true
	wd, _ := os.Getwd()

	if strings.Contains(wd, "cmd") {
		os.Chdir("..")
		os.Chdir("..")
	}

	assert.FileExists(t, jsonPath)
	file, err := ioutil.ReadFile(jsonPath)
	assert.NotNil(t, file)
	assert.NoError(t, err)
	err = json.Unmarshal(file, &wstat)
	assert.NotNil(t, wstat)
	assert.NoError(t, err)
}
