package ddz

import (
	"sync"

	"github.com/Gahd/DDZServer/src/model/ddz"
)

type HallManager struct {
	halls map[ddz.HallType]*ddz.Hall

	mutex sync.Mutex
}

var hallManager *HallManager

func GetHallManager() *HallManager {
	if hallManager == nil {
		hallManager = &HallManager{}
		hallManager.init()
	}

	return hallManager
}

func (this *HallManager) init() {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	this.halls = make(map[ddz.HallType]*ddz.Hall)

	//创建大厅，TODO:根据配置创建
	this.halls[ddz.DDZ_Normal] = ddz.NewHall(ddz.DDZ_Normal)
}

func (this *HallManager) GetHall(hallType ddz.HallType) *ddz.Hall {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	hall, isExists := this.halls[hallType]
	if !isExists {
		return nil
	}

	return hall
}
