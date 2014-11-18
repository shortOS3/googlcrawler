package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//"sync"
)

// type Answer struct {
// 	short string
// 	long  string
// 	err   error
// }

type Long struct {
	LongUrl string `json:longUrl`
}

func myFunc() {
	var ab = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	key := ""
	addr := "https://www.googleapis.com/urlshortener/v1/url?shortUrl=http://goo.gl/"

//	results := []*Answer{}

	for _, c := range ab {
//		ch := make(chan *Answer)
		for _, d := range ab {
			id := fmt.Sprintf("aa%s%s", c, d)
			url := fmt.Sprintf("%s%s&key=%s", addr, id, key)
			go func(url string) {
				l := Long{}
				resp, _ := http.Get(url)
				result, _ := ioutil.ReadAll(resp.Body)
				resp.Body.Close()
				e := json.Unmarshal(result, &l)
				if e != nil {
					fmt.Println("Problem")
				} else {
					fmt.Println(l.LongUrl)
					//ch <- &Answer{id, l.LongUrl, err}
				}
			}(url)
		}
		// for {
// 			r := <-ch
// 			results = append(results, r)
// 			if len(results) == 62 {
// 				close(ch)
// 				break
// 			}
// 		}
	}
	//return results
}

func main() {
	myFunc()
	// results := myfunc()
// 	for _, r := range results {
// 		if r.err != nil {
// 			fmt.Println(r.err)
// 		} else {
// 			fmt.Println(r.long)
// 		}
// 	}
}
