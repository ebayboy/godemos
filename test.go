package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func GetTSKeys(hkeys []string, split string, tsNow int, tsWin int) []string {
	keys := make([]string, tsWin)

	for i := 0; i < tsWin; i++ {
		keys[i] = strings.Join(hkeys, split)
		keys[i] += split
		keys[i] = fmt.Sprintf("%s%d", keys[i], tsNow-60*(i+1))
	}

	return keys
}

func RiskMsgParse(redisRes map[string]string, split string, cfgCols []string) (map[string]map[string]int, error) {
	outRes := make(map[string]map[string]int, len(redisRes))
	var err error
	for key, value := range redisRes {
		log.Println("split key:", key)
		m := make(map[string]int)
		outRes[key] = m
		vals := strings.Split(value, split)
		if len(vals) != len(cfgCols) {
			log.Printf("Redis values[%s] not match cols:[%s]\n", value, cfgCols)
			return nil, errors.New("Redis values not match cols number.")
		}

		for i := 0; i < len(cfgCols); i++ {
			m[cfgCols[i]], err = strconv.Atoi(vals[i])
			if err != nil {
				return nil, errors.New("Parse error.")
			}
		}
	}

	log.Println("outRes:", outRes)

	return outRes, nil
}

//@keysRes: map[string]string
//"1.1.1.1" : "11|00|31|41|51|61|71"
//"2.2.2.2" : "11|00|31|41|51|61|71"
//@keys: []string {"grp00_t.com_1590977400", "grp00_t.com_1590977340"}
//@output format:
//map{
//"grp00_t.com_1590977400" : {"1.1.1.1":map{"ip_total":50, "disc_uri": 33, "2.2.2.2":map{"ip_total":50, "disc_uri": 33},
//"grp00_t.com_1590977340" : {"3.3.3.3":map{"ip_total":55, "disc_uri": 22, "4.4.4.4":map{"ip_total":55, "disc_uri": 22}
//}
func keysResConvMap(keys []string, keysRes []map[string]string) (map[string]map[string]map[string]int, error) {
	res := make(map[string]map[string]map[string]int)

	for i := 0; i < len(keys); i++ {
		//map[string]map[string]string -> map[string]map[string]int
		rowRes, err := RiskMsgParse(keysRes[i], "|", []string{"total", "disc_uri"})
		if err != nil {
			log.Println("Error:", err.Error())
			return nil, err
		}
		res[keys[i]] = rowRes
	}

	return res, nil
}

func main() {
	// 1.1.1.1 total:10
	// 1.1.1.1 dist_uri:5
	// 2.2.2.2 total:10
	// all : total:10
	m := make(map[string]map[string]float64)
	m["grpid_host_1.1.1.1"] = map[string]float64{"total": 10, "dist_uri": 5}
	m["grpid_host_2.2.2.2"] = map[string]float64{"total": 10}

	for s2, m2 := range m {
		log.Println(s2, ":", m2)
	}
}
