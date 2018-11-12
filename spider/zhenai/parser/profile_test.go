package parser

import (
	"testing"
	"github.com/xxg3053/go-spider/spider/fetcher"
	"fmt"
)

func TestParserProfile(t *testing.T) {
	body,err := fetcher.Fetch("http://album.zhenai.com/u/1236870869")
	if err != nil{
		panic(err)
	}
	result := ParserProfile(body, "http://album.zhenai.com/u/1236870869","丫丫丫丫")
	fmt.Printf("result: %v", result.Items)
}
