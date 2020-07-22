package funcs

import (
	"github.com/freedomkk-qfeng/windows-agent/g"
	"sync"
	"time"

	"github.com/open-falcon/common/model"
)

const (
	historyCount int = 2
)

var (
	procStatHistory [historyCount]*CPUTimesStat
	psLock          = new(sync.RWMutex)
)

func UpdateCpuStat() error {
	ps, err := CPUTimes(false)
	if err != nil {
		return err
	}

	psLock.Lock()
	defer psLock.Unlock()
	for i := historyCount - 1; i > 0; i-- {
		procStatHistory[i] = procStatHistory[i-1]
	}

	procStatHistory[0] = &ps[0]

	return nil
}

func deltaTotal() float64 {
	if procStatHistory[1] == nil {
		return 0
	}
	return procStatHistory[0].Total - procStatHistory[1].Total
}

func CpuIdle() float64 {
	psLock.RLock()
	defer psLock.RUnlock()
	dt := deltaTotal()
	if dt == 0 {
		return 0.0
	}
	invQuotient := 100.00 / float64(dt)
	return float64(procStatHistory[0].Idle-procStatHistory[1].Idle) * invQuotient
}

func CpuUser() float64 {
	psLock.RLock()
	defer psLock.RUnlock()
	dt := deltaTotal()
	if dt == 0 {
		return 0.0
	}
	invQuotient := 100.00 / float64(dt)
	return float64(procStatHistory[0].User-procStatHistory[1].User) * invQuotient
}

func CpuSystem() float64 {
	psLock.RLock()
	defer psLock.RUnlock()
	dt := deltaTotal()
	if dt == 0 {
		return 0.0
	}
	invQuotient := 100.00 / float64(dt)
	return float64(procStatHistory[0].System-procStatHistory[1].System) * invQuotient
}

func CpuPrepared() bool {
	psLock.RLock()
	defer psLock.RUnlock()
	return procStatHistory[1] != nil
}

func CpuMetrics() []*model.MetricValue {
	var startTime,endTime time.Time
	if g.Config().Debug {
		startTime = time.Now()
	}
	if !CpuPrepared() {
		return []*model.MetricValue{}
	}

	cpuIdleVal := CpuIdle()
	idle := GaugeValue("cpu.idle", cpuIdleVal)
	busy := GaugeValue("cpu.busy", 100.0-cpuIdleVal)
	user := GaugeValue("cpu.user", CpuUser())
	system := GaugeValue("cpu.system", CpuSystem())

	if g.Config().Debug {
		endTime = time.Now()
		g.Logger().Printf("collect CpuMetrics complete. Process time %s.", endTime.Sub(startTime))
	}

	return []*model.MetricValue{idle, user, busy, system}
}
