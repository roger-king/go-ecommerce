package server

import (
	"os"
)

type configStruct struct {
	PORT string
}

var config = configStruct{
	os.Getenv("PORT"),
}
