package fetcher

import (
	"net/http"
	"golang.org/x/text/transform"
	"io/ioutil"
	"bufio"
	"golang.org/x/net/html/charset"
	"fmt"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"log"
	"time"
)

//通过http请求获取url的内容

var reateLimiter = time.Tick(100 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-reateLimiter
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil{
		return nil, err
	}

	req.Close = true
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")
	req.Header.Add("Accept-Encoding", "identity")
	resp, err := client.Do(req)
	if err != nil{
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK{
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	e := determineEncoding(bufio.NewReader(resp.Body))
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

/**
获取编码
 */
func determineEncoding(r *bufio.Reader)  encoding.Encoding{
	bytes, err := r.Peek(1024)
	if err != nil{
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8 //default
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}