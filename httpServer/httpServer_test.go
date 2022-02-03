package httpserver_test

import (
	"fmt"
	"testing"

	httpserver "practice/k8sjobs/httpServer"
)

type fakeHTTPHandler struct {
	err error
}

func (hro *fakeHTTPHandler) StartHTTPServerHelper(addr string, fakeK8sStatusMap map[string][]string) error {
	return hro.err
}

func TestStartHTTPServer(t *testing.T) {
	fhhObj := fakeHTTPHandler{
		err: nil,
	}

	var fakeK8sStatusMap = make(map[string][]string)
	if err := httpserver.StartHTTPServer(&fhhObj, fakeK8sStatusMap); err != nil {
		t.Fatalf("Unexpected error:%v", err)
	}
	fhhObj.err = fmt.Errorf("Error found")
	if err := httpserver.StartHTTPServer(&fhhObj, fakeK8sStatusMap); err == nil {
		t.Fatalf("Expected error: %v, but code did not throw any", err)
	}
}
