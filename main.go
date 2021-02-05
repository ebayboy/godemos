package main

import (
	"encoding/json"
	"fmt"
)

type DnRiskVar struct {
	Id            int64              `xorm:"id" json:"id"`
	WafInstanceId string             `xorm:"waf_instance_id" json:"wafInstanceId"`
	Domain        string             `xorm:"domain" json:"domain"`
	Name          string             `xorm:"name" json:"name"`
	Code          string             `xorm:"code" json:"code"`
	Desc          string             `xorm:"desc" json:"desc"`
	Disable       int8               `xorm:"disable" json:"disable"`
	EventCode     string             `xorm:"event_code" json:"eventCode"`
	Type          string             `xorm:"type" json:"varType"`
	Cost          string             `xorm:"cost" json:"cost"`
	Output        string             `xorm:"output" json:"output"`
	Property      string             `xorm:"property" json:"property"`
	Dura          int64              `xorm:"dura" json:"dura"`
	Logic         string             `xorm:"logic" json:"logic"`
	Rules         []DnRiskVarRuleCfg `xorm:"rules" json:"rules"`
	Pos           string             `xorm:"pos" json:"pos"`
	Key           string             `xorm:"key" json:"key"`
	CreateTime    int64              `xorm:"create_time created" json:"-"`
	UpdateTime    int64              `xorm:"update_time updated" json:"updateTime"`
	DeleteTime    time.Time          `xorm:"delete_time deleted" json:"-"`
}

type DnRiskVarRuleCfg struct {
	Left        string `json:"left,omitempty"`
	Operator    string `json:"operator,omitempty"`
	Right       string `json:"right,omitempty"`
	ResultOpt   string `json:"resultOpt,omitempty"`
	ResultRight string `json:"resultRight,omitempty"`
}

func main() {
	s := `[{
		"left": "vGID#vHST#vRIP#SUM#total#5m",
		"operator": "/",
		"right": "vGID#vHST#SUM#total#5m",
		"resultOpt": ">",
		"resultRight": "0.5"
	}, {
		"left": "vGID#vHST#vRIP#SUM#total#5m",
		"resultOpt": ">",
		"resultRight": "100"
	}]`
}
