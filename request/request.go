package request

import (
	"bufio"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"strings"
)

const BaseUrl = "https://wap.showstart.com/api"

const ActivityDetail = "/hw/000I8900ggre"
const GetListOfPerformers = "/hw/00000000ggre" //获取观演人列表
const Order = "/appnj/02cCBf00ggre"

func HttpDo(url, CUUSERREF string, param map[string]interface{}) (map[string]interface{}, error) {
	url = BaseUrl + url
	client := &http.Client{}
	bs, _ := jsoniter.Marshal(param)
	read := strings.NewReader(string(bs))
	req, err := http.NewRequest("POST", url, bufio.NewReader(read))
	if err != nil {
		return nil, err
	}
	req.Header.Set("HOST", "wap.showstart.com")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("CUUSERREF", CUUSERREF)
	req.Header.Set("CUSUT", "611344d559eebc9fd940dee43e778f4c")
	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	jsoniter.Unmarshal(body, &m)
	return m, nil
}
