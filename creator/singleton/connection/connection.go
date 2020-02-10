package connection

import (
	"fmt"
	"time"
)

type ConnObj struct {}

func (cb *ConnObj) Connect() bool {
	fmt.Println("Connecting database...")
	time.Sleep(time.Second*1)
	fmt.Println("Connecting successfully!")
	return true
}

func(cb *ConnObj) Disconnect() bool {
	fmt.Println("Disconnecting database...")
	time.Sleep(time.Second*1)
	fmt.Println("Disconnecting successfully!")
	return true
}