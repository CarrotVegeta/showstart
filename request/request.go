package request

import (
	"bufio"
	"github.com/CarrotVegeta/showstart/pkg"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpDo(url string, param interface{}) (map[string]interface{}, error) {
	url = pkg.BaseUrl + url
	client := &http.Client{}
	bs, _ := jsoniter.Marshal(param)
	read := strings.NewReader(string(bs))
	req, err := http.NewRequest("POST", url, bufio.NewReader(read))
	if err != nil {
		return nil, err
	}
	req.Header.Set("HOST", pkg.HOST)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("CUUSERREF", pkg.CUUSEREF)
	req.Header.Set("CUSUT", pkg.CUSUT)
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
