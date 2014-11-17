package main

import (
  "fmt"
  "net/http"
  "math"
  "io/ioutil"
)

type Answer struct {
  id		string
  result    []byte
  err      	error
}

func myfunc() []*Answer{
	var ab = []string{"a","b","c","d","e",}
	key := ""
	addr := "https://www.googleapis.com/urlshortener/v1/url?shortUrl=http://goo.gl/"
	
    ch := make(chan *Answer)
    results := []*Answer{}
	for _, a := range ab{
		for _, b := range ab{
			for _, c := range ab{
				for _, d := range ab{
					id := fmt.Sprintf("%s%s%s%s",a,b,c,d)
					url := fmt.Sprintf("%s%s&key=%s",addr,id,key)
					go func (url string, id string){
						resp, err:= http.Get(url)
						result, _ := ioutil.ReadAll(resp.Body)
					    resp.Body.Close()
						ch <- &Answer{id, result, err}
					}(url,id)	
				}
			}
		}
	}
	for {
	    r := <-ch
	    results = append(results, r)
	    if len(results) == int(math.Pow(5,4)) {
			close(ch)
			return results
		}
	}
	return results
}

func main() {
	results := myfunc()
	for _, r := range results{
		if r.err != nil{
			fmt.Println(r.err)
		}else{
			fmt.Println(r.id)
			fmt.Printf("%s\n",r.result)
		}
	}
}
