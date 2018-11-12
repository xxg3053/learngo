package parser

import (
	"github.com/xxg3053/go-spider/spider/engine"
	"strings"
	"regexp"
	"strconv"
	"github.com/xxg3053/go-spider/spider/model"
)


var userBriefRe  = regexp.MustCompile(`class="des f-cl" data-v-+[0-9a-z]+>([^<]+)</div>`)
var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)
var genderRe = regexp.MustCompile(`"genderString":"(女士|男士)"`)
//个人页面解析
func ParserProfile(contents []byte,url string, name string) engine.ParseResult {

	m := splitUserInfo(extractString(contents, userBriefRe))
	m.Name = name
	m.Gender = extractString(contents, genderRe)
	result := engine.ParseResult{
		Items: []engine.Item{{
			Url: url,
			Type: "zhenai",
			Id: extractString([]byte(url), idUrlRe),
			Payload: m,
		}},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

func splitUserInfo(s string) model.Profile {
	userinfo := model.Profile{}
	m := strings.Split(s, "|")
	if len(m) == 6 {

		if age, err := strconv.Atoi(strings.TrimSuffix(strings.Replace(m[1], " ", "", -1), "岁")); err == nil {

			userinfo.Age = age
		}

		if height, err := strconv.Atoi(strings.TrimSuffix(strings.Replace(m[4], " ", "", -1), "cm")); err == nil {

			userinfo.Height = height
		}

		userinfo.Hokou = strings.Replace(m[0], " ", "", -1)
		userinfo.Education = strings.Replace(m[2], " ", "", -1)
		userinfo.Marriage = strings.Replace(m[3], " ", "", -1)
		userinfo.Income = strings.Replace(m[5], " ", "", -1)
	}

	return userinfo
}
