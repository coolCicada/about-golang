package main

import "fmt"

func main() {
	/* personSalary := make(map[string]int)
	personSalary["steve"] = 12000
	personSalary["jamie"] = 15000
	personSalary["mike"] = 9000
	fmt.Println("personSalary map contents:", personSalary) */

	// 声明时 初始化 map
	/* personSalary := map[string]int {
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	fmt.Println(personSalary); */

	// 判断key是否存在
	/* personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	newEmp := "joe"
	if v, ok := personSalary[newEmp]; ok == true {
		fmt.Println(newEmp, v)
	} else {
		fmt.Println("null")
	} */
	
	// 遍历
	/* personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	fmt.Println("All items of a map")
	for k, v := range personSalary {
		fmt.Println(k, v)
	} */

	// 删除map中的元素
	/* personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	fmt.Println("map before deletion", personSalary, len(personSalary))
	delete(personSalary, "steve")
	fmt.Println("map after deletion", personSalary, len(personSalary)) */
}