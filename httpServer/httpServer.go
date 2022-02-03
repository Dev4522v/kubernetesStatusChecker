package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HTTPIntf interface {
	StartHTTPServerHelper(addr string, k8sStatusMap map[string][]string) error
}

type HTTPRequestObj struct{}

func StartHTTPServer(hro HTTPIntf, k8sStatusMap map[string][]string) error {
	if err := hro.StartHTTPServerHelper(":8080", k8sStatusMap); err != nil {
		return err
	}
	return nil
}

func (hro *HTTPRequestObj) StartHTTPServerHelper(addr string, k8sStatusMap map[string][]string) error {
	http.HandleFunc("/pods", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Names of pods running\n")
		res.Header().Set("Content-Type", "application/json")
		jsonResp, _ := json.Marshal(k8sStatusMap["pods"])
		res.Write(jsonResp)
	})
	http.HandleFunc("/services", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Names of services running\n")
		res.Header().Set("Content-Type", "application/json")
		jsonResp, _ := json.Marshal(k8sStatusMap["services"])
		res.Write(jsonResp)
	})
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello World!")
	})
	fmt.Printf("HTTP Server running on localhost:%v", addr)
	return http.ListenAndServe(addr, nil)
}
