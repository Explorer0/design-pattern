package hunger_singleton

import (
	"github.com/design-pattern/creator/singleton/connection"
	"sync"
)

// 一个单例实现的连接池
var (
	connPoolObj	*connectionPool
)

type connectionPool struct {
	connObjList	[]*connection.ConnObj
	usedMap map[*connection.ConnObj]bool
	sync.Mutex
}

const (
	connObjCount = 8
)

// 通过package 隐式init函数保证仅一次初始化
func init() {
	connPoolObj = &connectionPool{connObjList:make([]*connection.ConnObj, 0),
		usedMap:make(map[*connection.ConnObj]bool),
		Mutex:sync.Mutex{}}

	for idx:=0; idx<connObjCount; idx++ {
		connPoolObj.connObjList = append(connPoolObj.connObjList, new(connection.ConnObj))
	}
}

func (cp *connectionPool) GetConnObj() *connection.ConnObj {
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

func (cp *connectionPool) FreeConObj(obj *connection.ConnObj) {
	cp.Lock()
	defer cp.Unlock()

	if val, ok := cp.usedMap[obj]; ok && val {
		cp.usedMap[obj] = false
	}
}

func GetConnectionPool() *connectionPool {
	return connPoolObj
}