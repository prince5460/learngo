package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)

	const resultSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/baicheng1",
		"http://www.zhenai.com/zhenghun/cangzhou",
	}
	expectedCities := []string{
		"City 阿坝", "City 白城", "City 沧州",
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result should hava %d requests,had %d ", resultSize, len(result.Requests))
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d:%s;but was %s", i, url, result.Requests[i].Url)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should hava %d requests,had %d ", resultSize, len(result.Items))
	}

	for i, city := range expectedCities {
		if result.Items[i].(string) != city {
			t.Errorf("expected city #%d:%s;but was %s", i, city, result.Items[i].(string))
		}
	}

}
