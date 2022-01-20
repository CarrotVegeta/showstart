package main

//import (
//	"fmt"
//	"io/ioutil"
//	"net/http"
//	"net/url"
//	"showstart/service"
//)
//
//func main() {
//	//m := service.GetActivityDetail()
//	//fmt.Println(m)
//	m := service.Order()
//	fmt.Println(m)
//}

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//
//// http get by proxy
func GetByProxy(url_addr, proxyaddr string) (*http.Response, error) {
	request, _ := http.NewRequest("GET", url_addr, nil)
	proxy, err := url.Parse(proxyaddr)
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
	}
	return client.Do(request)
}

func main() {
	proxy := "http://58.252.56.149:9000/"
	url := "http://www.baidu.com/"
	resp, _ := GetByProxy(url, proxy)
	fmt.Println(resp)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println("sdfsdf")
}
