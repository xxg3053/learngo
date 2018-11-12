package parser

import (
	"github.com/xxg3053/go-spider/spider/engine"
	"regexp"
)

var cityListRe  = regexp.MustCompile(`<a href="(http://city.zhenai.com/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

/**
城市列表页面解析
 */
func ParserCityList(contents []byte) engine.ParseResult {

	matches := cityListRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}


	for _, m := range matches{
		//result.Items = append(result.Items, string(m[2]),)
		result.Requests = append(
			result.Requests, engine.Request{
				Url: string(m[1]),
				ParserFunc: ParserCity,
			},
		)
	}

	return result
}
