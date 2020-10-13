package parser

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)


func GetPrice(number string) string {
	//fmt.Println(number)

	//url := "https://m.avito.ru/api/14/items/1856218254?key=af0deccbgcgidddjgnvljitntccdduijhdinfgjgfjir"
	//req, _ := http.NewRequest("GET", url, nil)
	//req.Header.Add("cache-control", "no-cache")
	//req.Header.Add("Postman-Token", "b5befdaf-aecd-4fe5-ad38-ec1ef0b6d730")
	//res, _ := http.DefaultClient.Do(req)
	//defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(res)
	//fmt.Println(string(body))


	//proxyString, _ := url.Parse("https://34.80.14.167:3128")
	//client := &http.Client{
	//	Transport: &http.Transport{
	//		Proxy: http.ProxyURL(proxyString),
	//	},
	//	Timeout:   5 * time.Second,
	//}

	//dialer, err := proxy.SOCKS5("tcp", "https://34.80.14.167:3128", nil, proxy.Direct)
	//if err != nil {
	//	fmt.Println("errtcp = ", err)
	//}
	//client := &http.Client{
		//Transport: &http.Transport{
		//	Dial: dialer.Dial,
		//},
		//Timeout: 15 * time.Second,
	//}

	req, err := http.NewRequest("GET", "https://m.avito.ru/api/14/items/1856218254?key=af0deccbgcgidddjgnvljitntccdduijhdinfgjgfjir", nil)
	if err != nil {
	//	fmt.Println("err=", err)
	}
	//req := &http.Request{
	//	Method: http.MethodGet,
	//	Header: http.Header{
	//		"Accept":{"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"},
	//		"User-Agent": {"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:81.0) Gecko/20100101 Firefox/81.0"},
	//		"Upgrade-Insecure-Requests":{"1"},
	//	},
	//}
	//req.Host = "m.avito.ru"
	//req.URL, _ = url.Parse("https://m.avito.ru/api/14/items/1882620810?key=af0deccbgcgidddjgnvljitntccdduijhdinfgjgfjir")


//	req.Header.Add("cache-control", "no-cache")
	//req.Header.Add("Postman-Token", "4075e0b5-1b09-4dfe-841f-f65cf1f386fe")
	//req.URL.Query().Set("key","af0deccbgcgidddjgnvljitntccdduijhdinfgjgfjir")
	req.Header.Add("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	//req.Header.Add("Origin","https://google.com")
	req.Header.Add("Accept-Encoding","gzip, deflate, br")
	req.Header.Add("Accept-Language","ru-RU,ru;q=0.8,en-US;q=0.5,en;q=0.3")
	req.Header.Add("Connection","keep-alive")
	//req.Host = "m.avito.ru"
	req.Header.Add("TE","Trailers")
	req.Header.Add("Upgrade-Insecure-Requests","1")
	req.Header.Set("User-Agent","Mozilla/5.0 (iPhone; CPU iPhone OS 12_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/12.0 Mobile/15A372 Safari/604.1")
	//req.Cookie("u=2k8vufu8.1n4jx70.rcsldcmdxq00; buyer_location_id=639580; _ga=GA1.2.1842486484.1596309987; _ym_uid=159630998721248234; _ym_d=1596309987; _fbp=fb.1.1596309987939.1410309791; __gads=ID=7a91aef661790707:T=1601918577:S=ALNI_MatnHWNYCcRWHy7oN4UA1yAykrIqA; cto_bundle=U1FTYF9xNzVvN2NRV3hOVSUyQnJrbDlSS3dxaHQ0cGRtN2FNJTJCenBxMTd1akhFaGJTNmElMkYwMHklMkZDaFZNOEtITHJvRzFmSjBBTTE2V01QdWhBS09XYlFVUnNRTmVDNHdyV0pkV3dYWGVQTXlrUWJaTEhXNSUyQm15UHBJb0F3STJPWmtodk43bU9ZUjNKMDV0NzI4QjByVkdnbDZKYkZBJTNEJTNE; buyer_laas_location…3de19da9ed218fe23de19da9ed218fe2d50b96489ab264ed3de19da9ed218fe23de19da9ed218fe23de19da9ed218fe27fde300814b1e855f88859c11ff0089524a135baa76198de8732de926882853a8e3a80a29e104a6c2c61f4550df136d8efbe2742df161a37970f60f7d7c2f5061da9f488c469f0759ec7677a13d7c2bc76ab1e125221cbfe53bc326cd5f74c8b53b558646d992dbe8c15ca6e2019742946b8ae4e81acb9fa46b8ae4e81acb9fa02c68186b443a7ac9317cdb1f8e40fc82e8d32736bfac515f0c77052689da50d2da10fb74cac1eab3069315ebaf9ae7f8012e98924060d02; luri=stupino; abp=0; v=1602483156; _ym_isad=2")

	//req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11`)
	//req.URL, _ = url.Parse("https://m.avito.ru/api/14/items/1882620810?key=af0deccbgcgidddjgnvljitntccdduijhdinfgjgfjir")

	response, _ := http.DefaultClient.Do(req)
	defer response.Body.Close()
	//fmt.Println(response.StatusCode)
	//fmt.Println(response.Request)
	//fmt.Println(response.Header)
	data, _ := ioutil.ReadAll(response.Body)
	//fmt.Println("data =",string(data))
	var test string
	if err := json.Unmarshal(data, &test); err != nil {
		//fmt.Println("err2")
	}
	//fmt.Println("price =",test)


	/*
	   	response, err := http.DefaultClient.Get("https://m.avito.ru/api/14/items/1882620810?key=af0deccbgcgidddjgnvljitntccdduijhdinfgjgfjir")
	if err != nil {
		fmt.Println(err)
	}

	if response.StatusCode != http.StatusOK {
		fmt.Println("err1", response.StatusCode)
	}
	fmt.Println(response)
	fmt.Println(response.Body)
	data, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	var test string
	if err = json.Unmarshal(data, &test); err != nil {
		fmt.Println("err2")
	}
	fmt.Println("price =",test)
*/
	return "10"
}




