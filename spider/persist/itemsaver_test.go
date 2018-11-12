package persist

import (
	"testing"
	"github.com/xxg3053/go-spider/spider/model"
	"gopkg.in/olivere/elastic.v6"
	"context"
	"encoding/json"
	"github.com/xxg3053/go-spider/spider/engine"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:"aaa",
		Type:"zhenai",
		Id:"aaa",
		Payload:model.Profile{
			Name: "abc",
			Gender: "女",
			Age:34,
			Height:177,
			Weight:57,
			Income:"3000-5000",
			Xinzuo:"处女",
			Occupation:"人事",
			Marriage:"离异",
			Hokou:"深圳",
			Education:"本科",
			Car:"未购车",
		},
	}
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(elUrl),
	)
	if err != nil{
		panic(err)
	}
	err = save(*client, expected)
	if err != nil{
		panic(err)
	}



	resp, err := client.Get().
		Index("data_zhenai_profile").
		Type(expected.Type).Id(expected.Id).Do(context.Background())

	if err != nil{
		panic(err)
	}
	t.Logf("%s", resp.Source)
	var actual engine.Item
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil{
		panic(err)
	}
	actualProfile, err := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if actual != expected{
		t.Errorf("get %v; expected: %v", actual, expected)
	}
}
