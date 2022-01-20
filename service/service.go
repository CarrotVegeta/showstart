package service

import (
	"log"
	"showstart/request"
)

func GetActivityDetail() map[string]interface{} {
	m := make(map[string]interface{})
	m["data"] = "RdKnoWKXGxZP0zYw1ohcYvZbIyor3SpU3DeN8H53n1LW5PjbQdqFCFCiM8cBUfFO4clLwzs+dME5t5ajE68fEj+I0EmGeoCrC5bNRlwR4xXmzY0QM8ki0G3S/1dnd1zwIxWL/h5Fx1/P3I8durNbfu5qSMilOkrMqM9M4W6m/M0FEq0rssxosRykT7vCKolDgaIU93Tl980WUUtJ89bfCN78ABBOjEl+5942Xtc2gTGOFd3jLz2B+o0tM1p5G9w4KUBrnPtP3NMDl+fzXuga9S2FlU2301ILeKk8Lab/gz1Cnq4DMZduIbfI1NV2zhsEhU4Kv3sA2TWxXmT6iNFk7n7YCPqTB6TGS0bpA+Ltlzo="
	m["sign"] = "1e45f745caa4989f140ff9d296abcaef"
	m["appid"] = "wap"
	m["terminal"] = "wap"
	m["version"] = "997"
	m, err := request.HttpDo(request.ActivityDetail, "4ea93e5909142dc64f673ed009134cbb", m)
	if err != nil {
		log.Fatal(err.Error())
	}
	return m
}
func Order() map[string]interface{} {
	m := make(map[string]interface{})
	m["data"] = "XZ7oPYA9M5X1aWyQIUsydlM/KogoVPC2F4hclDEl6yOjnnt/0yiEBEHnRWpS722gqt9uGTC3xhuAeg359u0Xq1fwN3ucrSJbirA0lHXTPqKUYzLCPrbfv/2G9jRbVdefdoJskpBS4m4y5TYXDSQnXC1OsCrFSpaD1Ci9Gkue4nqKOttt+vEYSdERj+RlQPD9BPKx4er+Mq0fD9SOxQZvVGuk5izRA+1JXqgee0x1nJIp3W84qiBxI/AeBJWtaTApPBYw0EF+eRyGBaFaCvvyZYjE2vrn623EAgeUbHj+c0LUUUSlIRBFusgUuSXHZiFbiy8kSGiqlFiZknKgNucBjvSl0fck9x4GM3DGSxhz5Syshy4RR4wKMnY87ZEI2lWm"
	m["sign"] = "3307ed926bd720dac578cdaf1566c14f"
	m["appid"] = "wap"
	m["terminal"] = "wap"
	m["version"] = "997"
	m, err := request.HttpDo(request.Order, "4ea93e5909142dc64f673ed009134cbb", m)
	if err != nil {
		log.Fatal(err.Error())
	}
	return m
}
