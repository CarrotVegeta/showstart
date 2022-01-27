package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	CUUSEREF string `yaml:"CUUSEREF"`
	StFlpv   string `yaml:"StFlpv"`
	CUSUT    string `yaml:"CUSUT"`
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
