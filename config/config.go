package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	//CUUSEREF          string `yaml:"CUUSEREF"`
	StFlpv            string `yaml:"StFlpv"`
	CUSUT             string `yaml:"CUSUT"`
	CustomerName      string `yaml:"CustomerName"`
	Province          string `yaml:"Province"`
	CityName          string `yaml:"CityName"`
	Address           string `yaml:"Address"`
	CommonPerformerID int    `yaml:"CommonPerformerID"`
	Keyword           string `yaml:"Keyword"`
	Pagno             int    `yaml:"Pagno"`
	ActivityId        int    `yaml:"ActivityId"`
	TicketId          string `yaml:"TicketId"`
	Telephone         string `yaml:"Telephone"`
	Ticker            int64  `yaml:"Ticker"`
	CronTime          string `yaml:"CronTime"`
	TicketNum         int    `yaml:"TicketNum"`
	OrderNum          int    `yaml:"OrderNum"`
}

var Conf Config

func Init(path string) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err.Error())
	}
	type ss struct {
		ShowStart Config `yaml:"showstart"`
	}
	s := &ss{}
	if err := yaml.Unmarshal(f, &s); err != nil {
		log.Fatalf("yaml 解析错误:%v", err)
	}
	Conf = s.ShowStart
}
