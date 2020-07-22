package funcs

import (
	"github.com/freedomkk-qfeng/windows-agent/g"
	"time"

	"github.com/open-falcon/common/model"
	"github.com/shirou/gopsutil/mem"
)

func mem_info() (*mem.VirtualMemoryStat, error) {
	meminfo, err := mem.VirtualMemory()
	return meminfo, err
}

func MemMetrics() []*model.MetricValue {
	var startTime,endTime time.Time
	if g.Config().Debug {
		startTime = time.Now()
	}

	meminfo, err := mem_info()
	if err != nil {
		g.Logger().Println(err)
		return []*model.MetricValue{}
	}
	memTotal := meminfo.Total
	memUsed := meminfo.Used
	memFree := meminfo.Available
	pmemUsed := 100 * float64(memUsed) / float64(memTotal)
	pmemFree := 100 * float64(memFree) / float64(memTotal)

	if g.Config().Debug {
		endTime = time.Now()
		g.Logger().Printf("collect MemMetrics complete. Process time %s.", endTime.Sub(startTime))
	}

	return []*model.MetricValue{
		GaugeValue("mem.memtotal", memTotal),
		GaugeValue("mem.memused", memUsed),
		GaugeValue("mem.memfree", memFree),
		GaugeValue("mem.memfree.percent", pmemFree),
		GaugeValue("mem.memused.percent", pmemUsed),
	}

}
