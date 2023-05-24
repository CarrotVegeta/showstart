package request

import (
	"bufio"
	"github.com/CarrotVegeta/showstart/api"
	"github.com/CarrotVegeta/showstart/logger"
	"github.com/CarrotVegeta/showstart/pkg"
	jsoniter "github.com/json-iterator/go"
	"io"
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
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.50")
	//todo
	req.Header.Set("CUUSERREF", api.User.CUUSERREF)
	req.Header.Set("CUSUT", api.User.CUSUT)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	logger.FileLog.Info(string(body))
	m := make(map[string]interface{})
	jsoniter.Unmarshal(body, &m)
	if _, ok := m["state"]; !ok {
		if _, ok := m["status"]; !ok {
			log.Println(m)
		}
	}
	return m, nil
}
