package homework

import (
	"fmt"
	"strings"
)

/* FindStr
查找一个字符串中每个子字符串出现的次数
*/
func FindStr(s string) {
	strArr := strings.Split(s, " ")
	m := make(map[string]int)
	for _, value := range strArr {
		if strings.Trim(value, " ") == "" {
			continue
		}
		if m[value] == 0 {
			m[value] = 1
		} else {
			m[value] = m[value] + 1
		}
	}

	for key, value := range m {
		fmt.Printf("key : %s, value : %d \n", key, value)
	}

}

func FindStudent(id int) (rid, age, score, name interface{}) {
	m1 := make(map[int]map[string]interface{})
	m1[1] = map[string]interface{}{
		"id":    1,
		"name":  "zhangsan",
		"age":   21,
		"score": 80,
	}

	m1[2] = map[string]interface{}{
		"id":    2,
		"name":  "lisi",
		"age":   22,
		"score": 90,
	}

	m1[3] = map[string]interface{}{
		"id":    3,
		"name":  "wangwu",
		"age":   23,
		"score": 100,
	}
	if m1[id] == nil {
		return
	}
	rid = m1[id]["id"]
	name = m1[id]["name"]
	age = m1[id]["age"]
	score = m1[id]["score"]
	return

}

//根据条件分配数字
func DispatchCoin() int {
	coins := 50
	users := []string{"Mathew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth"}
	distribution := make(map[string]int, len(users))
	for _, user := range users {
		temp := []byte(user)
		for i := 0; i < len(temp); i++ {
			if temp[i] == 'e' || temp[i] == 'E' {
				distribution[user] += 1
				coins -= 1
			} else if temp[i] == 'O' || temp[i] == 'o' {
				distribution[user] += 2
				coins -= 2
			} else if temp[i] == 'i' || temp[i] == 'I' {
				distribution[user] += 3
				coins -= 3
			} else if temp[i] == 'u' || temp[i] == 'U' {
				distribution[user] += 4
				coins -= 4
			}
		}
	}
	return coins
}
