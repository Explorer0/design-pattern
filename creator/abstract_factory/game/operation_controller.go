package game

import "fmt"


type SymbianOpController struct {}

func (op *SymbianOpController) Operate() {
	fmt.Println("SymbianOpController operating...")
}

type WinPhoneOpController struct {}


func (op *WinPhoneOpController) Operate() {
	fmt.Println("WinPhoneOpController operating...")
}


type AndroidOpController struct {}

func (op *AndroidOpController) Operate() {
	fmt.Println("AndroidOpController operating...")
}


type SymbianDpController struct {}

func (op *SymbianDpController) Display() {
	fmt.Println("SymbianDpController displaying...")
}


type WinPhoneDpController struct {}

func (op *WinPhoneDpController) Display() {
	fmt.Println("WinPhoneDpController displaying...")
}


type AndroidDpController struct {}

func (op *AndroidDpController) Display() {
	fmt.Println("AndroidDpController displaying...")
}
