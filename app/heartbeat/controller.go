package heartbeat

import (
	"net/http"
	"encoding/json"
	"log"
)

type HealthCheck struct {
	Msg string `json:"msg"`
	Version string `json:"version"`
}

func HealthCheckController(w http.ResponseWriter, req *http.Request) {
	healthCheck := HealthCheck{"OK", "local"}
	data, err := json.Marshal(healthCheck)

	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}
