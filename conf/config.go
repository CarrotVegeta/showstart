package conf

import (
	"flag"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	ConfigPath string     `yaml:"config_path" json:"config_path"`
	Server     *Server    `yaml:"server"`
	ShowStart  *ShowStart `yaml:"show_start"`
	Logger     *Logger    `yaml:"logger"`
	Mysql
}
type Mysql struct {
	Server          string `yaml:"server"`
	Port            int    `yaml:"port"`
	Name            string `yaml:"name"`
	UserName        string `yaml:"username"`
	Password        string `yaml:"password"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`
	MaxOpenConns    int    `yaml:"max_open_conns"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime"`
}
type ShowStart struct {
	//CUUSEREF          string `yaml:"CUUSEREF"`
	//StFlpv            string `yaml:"StFlpv"`
	//CUSUT             string `yaml:"CUSUT"`
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
type Server struct {
	Port int `yaml:"port"`
}
type Logger struct {
	InfoFp   string `yaml:"info_fp"`
	ErrorFn  string `yaml:"error_fn"`
	LogLevel string `yaml:"log_level"`
}

func NewLogger() *Logger {
	return &Logger{}
}

var Conf *Config

func NewConfig() *Config {
	Conf = &Config{}
	return Conf
}
func Init() {
	NewConfig()
	flag.StringVar(&Conf.ConfigPath, "c", "./conf/config.yaml", "config file path")
	flag.Parse()
	f, err := os.ReadFile(Conf.ConfigPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	c := NewConfig()
	if err := yaml.Unmarshal(f, &c); err != nil {
		log.Fatalf("yaml 解析错误:%v", err)
	}
	Conf = c
}
