package parser

import (
	"testing"
	"io/ioutil"
)

func TestParserCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil{
		panic(err)
	}
	result := ParserCityList(contents)

	const resultSize  = 470
	if len(result.Requests) != resultSize{
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Requests))
	}
}
