package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type Long struct {
	LongUrl string `json:longUrl`
}

func crawl(half []string, ab []string, a string, key string) {
	var wg sync.WaitGroup
	addr := "https://www.googleapis.com/urlshortener/v1/url?shortUrl=http://goo.gl/"
	for _, b := range ab {
		wg.Add(1922)
		for _, c := range half {
			for _, d := range ab {
				id := fmt.Sprintf("%s%s%s%s", a, b, c, d)
				url := fmt.Sprintf("%s%s&key=%s", addr, id, key)
				go func(url string, id string) {
					l := Long{}
					resp, err := http.Get(url)
					if err == nil {
						result, _ := ioutil.ReadAll(resp.Body)
						resp.Body.Close()
						e := json.Unmarshal(result, &l)
						if e != nil {
							fmt.Println("Problem")
						} else {
							fmt.Printf("%s -> %s\n", id, l.LongUrl)
						}
					} else {
						fmt.Println("Error")
					}
					wg.Done()
				}(url, id)
			}
		}
		wg.Wait()
	}
}

func main() {

	var ab = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	var half1 = []string{"0", "1", "2", "3", "4", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M"}
	var half2 = []string{"5", "6", "7", "8", "9", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	flag.Parse()
	args := flag.Args()

	if len(args) < 2 {
		log.Fatal("This program will expand 238,328 urls from goo.gl.\nUsage: googlcrawler <startingString> <APIKey>")
	}

	start := args[0]
	if len(start) > 3 {
		log.Fatal("Starting string must be up to 3 characters long")
	}

	key := args[1]
	crawl(half1, ab, start, key)
	time.Sleep(10 * time.Second)
	crawl(half2, ab, start, key)
}
