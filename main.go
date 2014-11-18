package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type Long struct {
	LongUrl string `json:longUrl`
}

func crawl(half []string, ab []string) {
	var wg sync.WaitGroup
	key := ""
	addr := "https://www.googleapis.com/urlshortener/v1/url?shortUrl=http://goo.gl/"
	for _, b := range ab {
		wg.Add(1922)
		for _, c := range half {
			for _, d := range ab {
				id := fmt.Sprintf("0%s%s%s", b, c, d)
				url := fmt.Sprintf("%s%s&key=%s", addr, id, key)
				go func(url string) {
					l := Long{}
					resp, err := http.Get(url)
					if err == nil {
						result, _ := ioutil.ReadAll(resp.Body)
						resp.Body.Close()
						e := json.Unmarshal(result, &l)
						wg.Done()
						if e != nil {
							fmt.Println("Problem")
						} else {
							fmt.Printf("%s\n", l.LongUrl)
						}
					} else {
						fmt.Println("Error")
					}
				}(url)
			}
		}
		wg.Wait()
	}
}

func main() {
	var ab = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	var half1 = []string{"0", "1", "2", "3", "4", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M"}
	var half2 = []string{"5", "6", "7", "8", "9", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	crawl(half1, ab)
	fmt.Println("RETURNED TO MAIN")
	time.Sleep(10 * time.Second)
	crawl(half2, ab)
}
