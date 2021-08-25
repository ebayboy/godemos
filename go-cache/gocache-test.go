package main

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

type Stat struct {
	HostStat map[string]float64
}

func main() {
	c := cache.New(5*time.Minute, 10*time.Minute)

	key := "waf-ins_cn-north-1_9ccaa8af6f95#wdkj-isv.isvjcloud.com#1629108540"
	stat := Stat{HostStat: make(map[string]float64)}
	stat.HostStat["h_total"] = 50
	stat.HostStat["h_disc_uri"] = 30

	//1 .set
	//&{map[h_disc_uri:30 h_total:50]}
	c.Set(key, &stat, cache.DefaultExpiration)

	//2. after set, mod stat
	stat.HostStat["t_4xx"] = 5
	stat.HostStat["t_5xx"] = 1

	//3. Get key
	//&{map[h_disc_uri:30 h_total:50 t_4xx:5 t_5xx:1]}
	s, found := c.Get(key)
	if found {
		fmt.Println(s)

	}

}
