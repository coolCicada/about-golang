package dict

import (
	"encoding/json"
	"fmt"
)

type userInfo struct {
	Name string
	Age int
	Hobby []string
}

func Solve() {
	a := userInfo{"wang", 18, []string{"Gloang", "TS"}}
	buf, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	fmt.Println(buf)
	fmt.Println(string(buf))

	buf, err = json.MarshalIndent(a, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))

	var b userInfo
	err = json.Unmarshal(buf, &b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", b)
}