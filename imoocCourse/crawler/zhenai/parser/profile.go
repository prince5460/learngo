package parser

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"learngo/course/crawler/engine"
	"learngo/course/crawler/model"
	"log"
	"regexp"
)

var re = regexp.MustCompile(`<script>window.__INITIAL_STATE__=(.+);\(function`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	match := re.FindSubmatch(contents)
	result := engine.ParseResult{}
	if len(match) >= 2 {
		json := match[1]
		//fmt.Printf("json : %s\n", json)
		profile := parseJson(json)
		profile.Name = name
		result.Items = append(result.Items, profile)
		fmt.Println(profile)
	}
	return engine.ParseResult{}
}

//解析json数据
func parseJson(json []byte) model.Profile {
	res, err := simplejson.NewJson(json)
	if err != nil {
		log.Println("解析json失败。。")
	}
	infos, err := res.Get("objectInfo").Get("basicInfo").Array()
	//infos是一个切片，里面的类型是interface{}

	//fmt.Printf("infos:%v,  %T\n", infos, infos) //infos:[离异 47岁 射手座(11.22-12.21) 157cm 55kg 工作地:阿坝汶川 月收入:3-5千 教育/科研 大学本科],  []interface {}

	var profile model.Profile
	//所以我们遍历这个切片，里面使用断言来判断类型
	for k, v := range infos {
		//fmt.Printf("k:%v,%T\n", k, k)
		//fmt.Printf("v:%v,%T\n", v, v)

		/*
		    "basicInfo":[
		       "未婚",
		       "25岁",
		       "魔羯座(12.22-01.19)",
		       "152cm",
		       "42kg",
		       "工作地:阿坝茂县",
		       "月收入:3-5千",
		       "医生",
		       "大专"
		   ],
		*/
		if e, ok := v.(string); ok {
			switch k {
			case 0:
				profile.Marriage = e
			case 1:
				//年龄:47岁，我们可以设置int类型，所以可以通过另一个json字段来获取
				profile.Age = e
			case 2:
				profile.Xingzuo = e
			case 3:
				profile.Height = e
			case 4:
				profile.Weight = e
			case 6:
				profile.Income = e
			case 7:
				profile.Occupation = e
			case 8:
				profile.Education = e
			}
		}

	}

	return profile
}

//var (
//	AgeRe = regexp.MustCompile(`<td><span[^>]*>年龄：</span>([\d]+)岁</td>`)
//	GenderRe = regexp.MustCompile(`<td><span[^>]*>性别：</span><span field="">([^<]+)</span></td>`)
//	MarriageRe = regexp.MustCompile(`<td><span[^>]*>婚况：</span>([^<]+)</td>`)
//	HeightRe = regexp.MustCompile(`<td><span[^>]*>身高：</span><span field="">([\d]+)CM</span></td>`)
//	WeightRe = regexp.MustCompile(`<td><span[^>]*>体重：</span><span field="">([\d]+)</span></td>`)
//	IncomeRe = regexp.MustCompile(`<td><span[^>]*>月收入：</span>([^<]+)</td>`)
//	EducationRe = regexp.MustCompile(`<td><span[^>]*>学历：</span>([^<]+)</td>`)
//	OccupationRe = regexp.MustCompile(`<td><span[^>]*>职业： </span>([^<]+)</td>`)
//	ConstellationRe = regexp.MustCompile(`<td><span[^>]*>星座：</span><span field="">([^<]+)</span></td>`)
//	HouseRe = regexp.MustCompile(`<td><span[^>]*>住房条件：</span><span field="">([^<]+)</span></td>`)
//	CarRe = regexp.MustCompile(`<td><span[^>]*>是否购车：</span><span field="">([^<]+)</span></td>`)
//)
//
//func ParseProfile(contents []byte,name string) engine.ParseResult {
//	profile := model.Profile{}
//
//	profile.Name = name
//	profile.Age = extractInt(contents, AgeRe)
//	profile.Gender = extractString(contents, GenderRe)
//	profile.Marriage = extractString(contents, MarriageRe)
//	profile.Height = extractInt(contents, HeightRe)
//	profile.Weight = extractInt(contents, WeightRe)
//	profile.Income = extractString(contents, IncomeRe)
//	profile.Education = extractString(contents, EducationRe)
//	profile.Occupation = extractString(contents, OccupationRe)
//	profile.Xingzuo = extractString(contents, ConstellationRe)
//	profile.House = extractString(contents, HouseRe)
//	profile.Car = extractString(contents, CarRe)
//
//	result := engine.ParseResult{
//		Items: []interface{}{profile},
//	}
//
//	return result
//
//}
//
//func extractString(contents []byte, re *regexp.Regexp) string {
//	match := re.FindSubmatch(contents)
//
//	if len(match) >= 2 {
//		return string(match[1])
//	} else {
//		return ""
//	}
//}
//
//func extractInt(contents []byte, Re *regexp.Regexp) int {
//	r := extractString(contents, Re)
//
//	num, err := strconv.Atoi(r)
//
//	if err != nil {
//		return 0
//	}
//
//	return num
//}
