package main

import (
	"fmt"

	"github.com/zh1lu0/gocamp/dao"
)

func main() {
	var id int64 = 5
	name, err := dao.NameById(id)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Name of user no. %d: %s", id, name)
}
