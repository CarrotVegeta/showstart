package request

import (
	"bufio"
	"github.com/CarrotVegeta/showstart/logger"
	"github.com/CarrotVegeta/showstart/pkg"
	"github.com/CarrotVegeta/showstart/server"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"log"
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
	//todo
	req.Header.Set("CUUSERREF", server.User.CUUSERREF)
	req.Header.Set("CUSUT", server.User.CUSUT)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	logger.FileLog.Info(string(body))
	m := make(map[string]interface{})
	jsoniter.Unmarshal(body, &m)
	if _, ok := m["state"]; !ok {
		log.Fatal(err.Error())
	}
	return m, nil
}
