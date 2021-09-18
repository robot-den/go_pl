package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Id           int
	Names        map[string]int `json:"Nicknames"`
	internalData string
}

func main() {
	bool, _ := json.Marshal(true)
	fmt.Println("bool:", string(bool))

	n, _ := json.Marshal(1)
	fmt.Println("int:", string(n))

	f, _ := json.Marshal(2.34)
	fmt.Println("float:", string(f))

	s, _ := json.Marshal("golang")
	fmt.Println("string:", string(s))

	sl := []string{"one", "two"}
	slM, _ := json.Marshal(sl)
	fmt.Println("int:", string(slM))

	m := map[int]string{1: "one", 2: "two"}
	mM, _ := json.Marshal(m)
	fmt.Println("int:", string(mM))

	u := &User{Id: 8, Names: map[string]int{"Denis": 5}}
	mU, _ := json.Marshal(u)
	fmt.Println("struct", string(mU))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"Id": 3, "names": {"Rick": 2}}`
	res := User{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Names["Rick"])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
