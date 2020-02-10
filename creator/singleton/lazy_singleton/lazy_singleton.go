package lazy_singleton

import (
	"github.com/design-pattern/creator/singleton/connection"
	"sync"
)

// 一个单例实现的连接池
var (
	connPoolObj	*connectionPool
	fetchMutex	sync.Mutex
	once		sync.Once
)

const (
	connObjCount = 8
)

type connectionPool struct {
	connObjList	[]*connection.ConnObj
	usedMap map[*connection.ConnObj]bool
	sync.Mutex
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

// 通过双重检验方式保证一次初始化
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

	connPoolObj = &connectionPool{connObjList:make([]*connection.ConnObj, 0),
		 						  usedMap:make(map[*connection.ConnObj]bool),
		 						  Mutex:sync.Mutex{}}

	for idx:=0; idx<connObjCount; idx++ {
		connPoolObj.connObjList = append(connPoolObj.connObjList, new(connection.ConnObj))
	}
	return connPoolObj
}

// 通过sync.Once保证一次初始化
func GetConnectionPool2() *connectionPool {
	once.Do(func() {
		connPoolObj = &connectionPool{connObjList:make([]*connection.ConnObj, 0),
			usedMap:make(map[*connection.ConnObj]bool),
			Mutex:sync.Mutex{}}

		for idx:=0; idx<connObjCount; idx++ {
			connPoolObj.connObjList = append(connPoolObj.connObjList, new(connection.ConnObj))
		}
	})
	return connPoolObj
}
