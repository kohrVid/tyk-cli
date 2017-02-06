package request

import (
	"net/http"
	"net/http/httptest"
	"reflect"
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

func TestFullRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "www.example.com", nil)
	request := New("12345", "http://www.example.com", "3000")
	var expectedResult = []string{"12345"}
	var result []string
	if result[0] != expectedResult[0] {
		t.Fatalf("Expected %v, got %v", expectedResult, result)
	}
}

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
}
