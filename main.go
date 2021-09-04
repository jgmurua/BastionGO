package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"github.com/aws/aws-lambda-go/lambda"
)

type Req struct {
	Url string `json:"url"`
	Method string `json:"method"`
	Payload string `json:"payload"`
}
 

func HandleLambdaEvent(rq Req) (string, error) {

	client := &http.Client {
	}
	
	payloadRead:= strings.NewReader(string(rq.Payload))

	req, error := http.NewRequest(rq.Method, rq.Url, payloadRead)
	if error != nil {
		fmt.Println(error)
	}
	req.Header.Add("Content-Type", "application/json")
  
	res, error := client.Do(req)
	if error != nil {
	  fmt.Println(error)
	}
	defer res.Body.Close()
  
	body, error := ioutil.ReadAll(res.Body)
	if error != nil {
	  fmt.Println(error)
	}
	return fmt.Sprintf(string(body)), nil
}
 
func main() {
        lambda.Start(HandleLambdaEvent)
}