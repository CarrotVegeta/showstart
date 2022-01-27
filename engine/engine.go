package engine

import (
	"flag"
	"fmt"
	"github.com/CarrotVegeta/showstart/config"
)

func NewEngine() {
	f := &Flag{}
	f.parseFlag()
	config.Init(f.ConfigPath)
	args := flag.Args()
	switch args[0] {
	case "getActivityList":
		getActivityList(f.Keyword, f.PagNo)
	case "getActivityDetail":
		getActivityDetail(fmt.Sprintf("%d", f.ActivityId))
	case "getTicketList":
		getTicketList(fmt.Sprintf("%d", f.ActivityId), "")
	case "getSpectatorList":
		getSpectatorList(0)
	case "order":
		order(fmt.Sprintf("%d", f.ActivityId), f.TicketId)
	}
}

type Flag struct {
	ConfigPath string
	Keyword    string
	PagNo      int
	ActivityId int
	TicketId   string
}

func (f *Flag) parseFlag() {
	configPath := flag.String("config", "", "config path")
	keyword := flag.String("keyword", "", "search keyword")
	pagNo := flag.Int("pagno", 1, "number of pages")
	activityId := flag.Int("activityId", 0, "id of activity")
	ticketId := flag.String("ticketId", "", "ticketId of ticket")
	flag.Parse()
	f.ConfigPath = *configPath
	f.Keyword = *keyword
	f.PagNo = *pagNo
	f.ActivityId = *activityId
	f.TicketId = *ticketId
}
