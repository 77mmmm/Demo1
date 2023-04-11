package main

import (
	"awesomeProject/handler"
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"
)

func TestMyHandler(t *testing.T) {
	values := url.Values{}
	values.Set("username", "testuser")
	values.Set("password", "testpassword")
	values.Set("email", "test@example.com")
	body := bytes.NewBufferString(values.Encode())

	req := httptest.NewRequest("POST", "/my-endpoint", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i <= 4; i++ {
		go func() {
			handler := http.HandlerFunc(handler.Register)
			handler.ServeHTTP(rr, req)

			wg.Done()
		}()
	}

	wg.Wait()
}
