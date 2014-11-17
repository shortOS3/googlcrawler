package main

import (
  "fmt"
  "net/http"
  "math"
  "io/ioutil"
)

func myfunc()[]*http.Response {
	var ab = []string{"a","b","c","d","e",}
	key := ""
	addr := "https://www.googleapis.com/urlshortener/v1/url?shortUrl=http://goo.gl/"
   
    ch := make(chan *http.Response)
    responses := []*http.Response{}
	for _, a := range ab{
		for _, b := range ab{
			for _, c := range ab{
				for _, d := range ab{
					id := fmt.Sprintf("%s%s%s%s",a,b,c,d)
					url := fmt.Sprintf("%s%s&key=%s",addr,id,key)
					go func (url string){
						resp, _:= http.Get(url)
						ch <- resp
					}(url)	
				}
			}
		}
	}
	for {
	    r := <-ch
	    responses = append(responses, r)
	    if len(responses) == int(math.Pow(5,4)) {
			return responses
		}
	}
	return responses
}

func main() {
	res := myfunc()
	for _, r := range res{
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Printf("%s",body)
	}
}
