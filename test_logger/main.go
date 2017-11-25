package main

import (
	//"fmt"
	"github.com/op/go-logging"
	"os"
	//"time"
)

var log = logging.MustGetLogger("loader")

func main() {
	var format = logging.MustStringFormatter(
		//"%{color} %{shortfile}: %{level:.4s} %{message}%{color:reset}",
		"%{color}%{time:01/02/2006 15:04:05.000} %{shortfile}: %{level:.4s} %{message}%{color:reset}",
	)
	//logging.SetFormatter(format)
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)
	logging.SetLevel(logging.INFO, "loader")

	log.Info("test")
}
