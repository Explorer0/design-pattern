package singleton

import (
	"fmt"
	"sync"
	"time"
)
// 一个单例实现的连接池
var (
	connPoolObj	*connectionPool
	fetchMutex	sync.Mutex
	connObjCount int
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


type connectionPool struct {
	connObjList	[]*ConnObj
	usedMap map[*ConnObj]bool
	sync.Mutex
}

func (cp *connectionPool) GetConnObj() *ConnObj {
	cp.Lock()
	defer cp.Unlock()

	for idx:=0; idx<len(cp.connObjList); idx++ {
		if  val, ok :=cp.usedMap[cp.connObjList[idx]]; !ok || !val {
			cp.usedMap[cp.connObjList[idx]] = true
			return cp.connObjList[idx]
		}
	}
	return nil
}

func (cp *connectionPool) FreeConObj(obj *ConnObj) {
	cp.Lock()
	defer cp.Unlock()

	if val, ok := cp.usedMap[obj]; ok && val {
		cp.usedMap[obj] = false
	}
}

func GetConnectionPool() *connectionPool {
	if connPoolObj != nil {
		return connPoolObj
	}

	fetchMutex.Lock()
	// 双重判断
	if connPoolObj != nil {
		return connPoolObj
	}
	defer fetchMutex.Unlock()

	connPoolObj = new(connectionPool)
	connPoolObj.connObjList = make([]*ConnObj, 0)
	connPoolObj.usedMap = make(map[*ConnObj]bool)
	for idx:=0; idx<connObjCount; idx++ {
		connPoolObj.connObjList = append(connPoolObj.connObjList, new(ConnObj))
	}
	return connPoolObj
}
