package config

import (
    "os"
)

const (
	API_URL = "http://localhost:1000/"
	HOME_URL = "http://localhost:1000/"
)

var (
	PATH = ""

	HTML_WEB = ""
	HTML_PATH = ""
	TPL_PATH = ""

	IMAGE_WEB = ""
	IMAGE_PATH = ""
)

func SetConstants() {
	path, _ := os.Getwd()

	PATH = path + "/../src/api/"

	HTML_WEB = HOME_URL + "html/"
	HTML_PATH = PATH + "html/"
	TPL_PATH = HTML_PATH + "templates/"
}