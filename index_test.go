package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"fmt"
	"github.com/stretchr/testify/assert"
	service "./service"
)

func Test_Index(t *testing.T)  {

	server := service.NewServer()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	server.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, 200)
}
