package parser

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)


func GetPrice(number string) (string, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			MaxVersion: tls.VersionTLS12,
		},
	}
	client := &http.Client{Transport: tr}

	url := "https://m.avito.ru/api/1/rmp/show/"+number+"?key=af0deccbgcgidddjgnvljitntccdduijhdinfgjgfjir"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept","application/json, text/plain, */")
	req.Header.Set("Cache-Control","max-age=0")
	req.Header.Set("Connection","keep-alive")
	req.Header.Set("Content-Type","application/json;charset=utf-8")
	req.Header.Set("Referer","https://m.avito.ru")
	req.Header.Set("User-Agent","Mozilla/5.0 (iPhone; CPU iPhone OS 12_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/12.0 Mobile/15A372 Safari/604.1")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	var convertData map[string]map[string]map[string][]map[string]map[string]int
	var convertStatus map[string]string
	json.Unmarshal(data, &convertStatus)
	if convertStatus["status"]=="ok" {
		json.Unmarshal(data, &convertData)
		price := strconv.Itoa(convertData["result"]["banners"]["ads_mob_credit_btn"][0]["parameters"]["par_price"])
		return price,nil
	}
	return "",nil
}




