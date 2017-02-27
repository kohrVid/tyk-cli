package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"
)

func indexRequest(r *Request, index int) (val interface{}) {
	val = reflect.ValueOf(*r).Field(index).Interface()
	return
}

func TestNew(t *testing.T) {
	expectedResult := &Request{
		"12345",
		"http://www.example.com",
		"3000",
		&http.Client{Timeout: 10 * time.Second},
	}
	result := New("12345", "http://www.example.com", "3000")
	var pass bool
	for i := range []int{0, 1, 2} {
		pass = indexRequest(expectedResult, i) == indexRequest(result, i)
	}
	if !pass {
		t.Fatalf("Expected %v, got %v", expectedResult, result)
	}
}

func setup() {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	//client = NewClient(nil, "foo")
	//url, _ := url.Parse(server.URL)
	//client.BaseURL = url

}

func TestFullRequest(t *testing.T) {
	//res := httptest.NewRecorder()
	//server := httptest.NewRecorder(mux)
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()
	request := New("12345", "http://www.example.com", "3000")
	req, err := request.FullRequest("GET", "http://localhost:3000/api/apis", nil)
	mux.HandleFunc("/url", func(req, r *http.Request) {
		fmt.Fprintln(req, `{}`)
	})
	http.DefaultServeMux.ServeHTTP(server, req)
	if p, err := ioutil.ReadAll(req.Body); err != nil {
		t.Fail()
	} else {
		if strings.Contains(string(p), "Error") {
			t.Errorf("Header response shouldn't return error: %s", p)
		} else if !strings.Contains(string(p), `12345`) {
			t.Errorf("Header response doen't match:\n%s", p)
		}
	}
}

/*
func TestFullRequestAuthorisation(t *testing.T) {
	request := New("12345", "http://www.example.com", "3000")
	var expectedResult = []string{"12345"}
	var result []string
	if err != nil {
		t.Fatalf("Unexpected error")
	} else {
		result = req.Header["Authorization"]
	}
	if result[0] != expectedResult[0] {
		t.Fatalf("Expected %v, got %v", expectedResult, result)
	}
}

func TestFullRequestContentType(t *testing.T) {
	request := New("12345", "http://www.example.com", "3000")
	var expectedResult = []string{"application/json"}
	var result []string
	req, err := request.FullRequest("GET", "www.example.com", nil)
	if err != nil {
		t.Fatalf("Unexpected error")
	} else {
		result = req.Header["Content-Type"]
	}
	if result[0] != expectedResult[0] {
		t.Fatalf("Expected %v, got %v", expectedResult, result)
	}
}*/
