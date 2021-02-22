package main

import (
	"fmt"
	"sort"
)

func toRoot(set map[int]int, data int) int {
	id := data
	for {
		if idx, ok := set[id]; ok && idx != id {
			id = idx
		} else {
			break
		}
	}
	return id
}

func accountsMerge(accounts [][]string) [][]string {
	idx := 0
	mailToIdx := make(map[string]int)
	userMerge := make(map[int]int)
	users := []string{}
	for _, account := range accounts {
		name := account[0]
		id := -1
		for _, mail := range account[1:] {
			if i, ok := mailToIdx[mail]; ok { // 发现用户，合并
				id = i
				break
			}
		}
		if id == -1 { // 创建新用户
			users = append(users, name)
			id = idx
			idx ++
		}
		for _, mail := range account[1:] {
			if idx, ok := mailToIdx[mail]; ok && idx != id{
				userMerge[toRoot(userMerge, id)] = toRoot(userMerge, idx)
				id = idx
			}
			mailToIdx[mail] = id
		}
	}
	result := [][]string{}
	for _, name := range users {
		result = append(result, []string{name})
	}
	for mail, id := range mailToIdx {
		id = toRoot(userMerge, id)
		result[id] = append(result[id], mail)
	}
	for i := range result {
		sort.Strings(result[i])
	}
	final := [][]string{}
	for _, item := range result {
		if len(item) > 1 {
			final = append(final, item)
		}
	}
	return final
}

func main () {
	fmt.Println(accountsMerge([][]string{{"Hanzo","Hanzo1@m.co","Hanzo2@m.co","Hanzo17@m.co","Hanzo18@m.co","Hanzo19@m.co"},{"Hanzo","Hanzo34@m.co","Hanzo59@m.co"},{"Hanzo","Hanzo7@m.co","Hanzo8@m.co","Hanzo47@m.co","Hanzo48@m.co","Hanzo49@m.co"},{"Hanzo","Hanzo0@m.co","Hanzo1@m.co","Hanzo12@m.co","Hanzo13@m.co","Hanzo14@m.co"},{"Hanzo","Hanzo3@m.co","Hanzo4@m.co","Hanzo27@m.co","Hanzo28@m.co","Hanzo29@m.co"},{"Hanzo","Hanzo9@m.co","Hanzo5@m.co","Hanzo57@m.co","Hanzo58@m.co","Hanzo59@m.co"},{"Hanzo","Hanzo5@m.co","Hanzo6@m.co","Hanzo37@m.co","Hanzo38@m.co","Hanzo39@m.co"},{"Hanzo","Hanzo2@m.co","Hanzo3@m.co","Hanzo22@m.co","Hanzo23@m.co","Hanzo24@m.co"},{"Hanzo","Hanzo8@m.co","Hanzo9@m.co","Hanzo52@m.co","Hanzo53@m.co","Hanzo54@m.co"},{"Hanzo","Hanzo4@m.co","Hanzo0@m.co","Hanzo32@m.co","Hanzo33@m.co","Hanzo34@m.co"},{"Hanzo","Hanzo6@m.co","Hanzo7@m.co","Hanzo42@m.co","Hanzo43@m.co","Hanzo44@m.co"}}))
}