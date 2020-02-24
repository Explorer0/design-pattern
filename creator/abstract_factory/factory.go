package abstract_factory

import "github.com/design-pattern/creator/abstract_factory/game"

type OperationController interface {
	Operate()
}

type InterfaceController interface {
	Display()
}

type GameFactory interface {
	CreateOpController() OperationController
	CreateIfController() InterfaceController
}


type SymbianGameFactory struct {}

func (gf *SymbianGameFactory) CreateOpController() OperationController {
	return &game.SymbianOpController{}
}

func (gf *SymbianGameFactory) CreateIfController() InterfaceController {
	return &game.SymbianDpController{}
}


type WinPhoneGameFactory struct {}

func (gf *WinPhoneGameFactory) CreateOpController() OperationController {
	return &game.WinPhoneOpController{}
}

func (gf *WinPhoneGameFactory) CreateIfController() InterfaceController {
	return &game.WinPhoneDpController{}
}


type AndroidGameFactory struct {}

func (gf *AndroidGameFactory) CreateOpController() OperationController {
	return &game.AndroidOpController{}
}

func (gf *AndroidGameFactory) CreateIfController() InterfaceController {
	return &game.AndroidDpController{}
}


func CreateSymbianGameFactory() *SymbianGameFactory {
	return &SymbianGameFactory{}
}

func CreateWinPhoneGameFactory() *WinPhoneGameFactory {
	return &WinPhoneGameFactory{}
}

func CreateAndroidGameFactory() *AndroidGameFactory {
	return &AndroidGameFactory{}
}
