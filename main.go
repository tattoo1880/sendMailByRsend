package main

import (
	"fmt"
	"github.com/tattoo1880/sendMailByRsend/sendmail"
)

func main() {
	utils := sendmail.NewUtils()
	err := utils.SendEmail("test in train")
	if err != nil {
		fmt.Println(err)
	}
}
