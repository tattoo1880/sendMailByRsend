package main

import "fmt"

func main() {
	utils := NewUtils()
	err := utils.SendEmail("test in train")
	if err != nil {
		fmt.Println(err)
	}
}
