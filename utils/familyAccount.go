package utils

import (
	"fmt"
)

type FamilyAccount struct {
	balance float64
	detail  string
}

func NewFamilyAccount() *FamilyAccount { //工厂模式
	return &FamilyAccount{
		balance: 10000,
		detail:  "收支\t账户余额\t收支金额\t说  明",
	}
}

func (F *FamilyAccount) MainMenu() {
	for {
		var chance string
		var confire bool
		fmt.Println("\n\n----------------家庭收支记账系统----------------")
		fmt.Println("                 1.收支明细                 ")
		fmt.Println("                 2.登记收入                 ")
		fmt.Println("                 3.登记支出                 ")
		fmt.Println("                 4.退出软件                 ")
		fmt.Println("                 请选择(1-4)                 ")
		fmt.Printf("你的选择(1-4): ")

		fmt.Scanln(&chance)

		switch chance {
		case "1":
			F.showDetails()
		case "2":
			F.addIncome()
		case "3":
			F.addOutcome()
		case "4":
			confire = F.exit()

		default:
			fmt.Println("------------------请输入正确的选项-----------------")
		}
		if confire {
			break
		}
	}
}

func (F *FamilyAccount) showDetails() {
	fmt.Println("------------------当前收支明细记录-----------------")
	fmt.Println(F.detail)
}

func (F *FamilyAccount) addIncome() {
	var money float64 = 0.0
	var note string
	fmt.Printf("本次收入金额 : ")
	fmt.Scanln(&money)
	F.balance += money
	fmt.Printf("本次收入说明 : ")
	fmt.Scanln(&note)
	F.detail += fmt.Sprintf("\n收入\t%v\t%v\t%v", F.balance, money, note)
}

func (F *FamilyAccount) addOutcome() {
	var money float64 = 0.0
	var note string
	fmt.Printf("本次支出金额 : ")
	fmt.Scanln(&money)
	if F.balance < money {
		fmt.Println("\n余额不足")
		return
	}
	F.balance -= money
	fmt.Printf("本次收入说明 : ")
	fmt.Scanln(&note)
	F.detail += fmt.Sprintf("\n支出\t%v\t%v\t%v", F.balance, money, note)
}

func (F *FamilyAccount) exit() bool {
	confirChoice := ""
	fmt.Printf("你确定要退出吗(y/n)? ")
	fmt.Scanln(&confirChoice)
	if confirChoice == "y" {
		return true
	} else if confirChoice != "n" {
		fmt.Println("------------------请重新输入-----------------")
	}
	return false
}
