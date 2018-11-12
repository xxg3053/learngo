package model

import "encoding/json"

type Profile struct {
	Name       string
	Gender     string
	Age        int
	Height     int
	Weight     int
	Income     string  //收入
	Marriage   string  //婚姻状况
	Education  string
	Occupation string  //
	Hokou      string  //户口
	Xinzuo     string
	House      string  //房子
	Car        string  //车子
}

func FromJsonObj(o interface{})  (Profile, error){
	var prof Profile
	s, err := json.Marshal(o)
	if err != nil{
		return prof, err
	}
	err = json.Unmarshal(s, &prof)
	return prof, err
}