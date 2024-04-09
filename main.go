package main

import (
	"fmt"
	"goProjectFamily/utils"

	_ "github.com/gin-gonic/gin"
)

func main() {

	acc := utils.NewFamilyAccount()
	acc.MainMenu()
	fmt.Println("------------------byebye~-----------------")
}
