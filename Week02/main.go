package main

import (
	"fmt"
	"week02/service"
)

func main() {
	username, err := service.FindNameLIke("go")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("id=1,name=%2s;", username)
}
