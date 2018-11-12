package parser

import (
	"regexp"
	"github.com/xxg3053/go-spider/spider/engine"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://city.zhenai.com/[^"]+)"`)
)

/**
城市页面解析
 */
func ParserCity(contents []byte) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		name := string(m[2])
		url := string(m[1])
		//result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url:        url,
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParserProfile(c, url, name)
			},
		})
	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParserFunc: ParserCity,
		})
	}

	return result
}


