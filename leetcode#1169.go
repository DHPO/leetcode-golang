package main

import (
	"sort"
	"strings"
	"strconv"
)

type Trans struct {
    idx int
    name string
    time int
    amount int
    city string
}

func parse(idx int, transStr string) *Trans {
    trans := new(Trans)
    trans.idx = idx
	splits := strings.Split(transStr, ",")
	trans.name = splits[0]
	trans.city = splits[3]
	trans.time, _ = strconv.Atoi(splits[1])
	trans.amount, _ = strconv.Atoi(splits[2])

    return trans
}

func invalidTransactions(transactions []string) []string {
    transObjs := make([]*Trans, len(transactions))
    nameCityMap := make(map[string][]*Trans)
    results := []string{}

    for i, v := range transactions {
        transObj := parse(i, v)
        transObjs[i] = transObj
        name := transObj.name
        if sls, ok := nameCityMap[name]; ok {
            nameCityMap[name] = append(sls, transObj)
        } else {
            nameCityMap[name] = []*Trans{transObj}
        }
    }

    for _, v := range nameCityMap {
        sort.Slice(v, func (i, j int) bool {
            return v[i].time < v[j].time
        })
		inValid := make(map[int]bool)
		dangerZone := 0
        for i := range v {
			if v[i].amount > 1000 {
				inValid[i] = true
            }
			if i == 0 {
				continue
			}
			for ; v[i].time - v[dangerZone].time > 60 && dangerZone < i; dangerZone ++ {}
			if v[dangerZone].city == v[i].city {
				continue
			}
			for j := dangerZone; j <= i; j ++ {
				inValid[j] = true
			}
		}
		for i := range inValid {
			results = append(results, transactions[v[i].idx])
		}
    }

    return results
}
