package export

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"time"
)

func Apis(args []string) {
	authorisation := args[0]
	domain := checkDomain(args[1])
	port := args[2]
	client := &http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf("%s:%s/api/apis", domain, port)
	req, err := httpRequest("GET", url, authorisation)
	resp, err := client.Do(req)
	output_file := args[3]
	exportResponse(resp, err, output_file)
}

func checkDomain(inputString string) string {
	if !isProtocolPresent(inputString) {
		fmt.Println("Please add a protocol to your domain")
		os.Exit(-1)
	}
	return inputString
}

func isProtocolPresent(arg string) bool {
	matched, _ := regexp.MatchString("(http|https)://", arg)
	return matched
}

func httpRequest(requestType string, url string, authorisation string) (*http.Request, error) {
	req, err := http.NewRequest(requestType, url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authorisation)

	return req, err
}

type ExportJson struct {
	Body string
}

func exportResponse(resp *http.Response, err error, file string) {
	if err != nil {
		fmt.Println(err)
	} else {
		defer resp.Body.Close()
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		jsonString := buf.String()
		f, err := os.Create(file)
		if err != nil {
			return
		}
		defer f.Close()
		f.WriteString(jsonString)
	}
}
