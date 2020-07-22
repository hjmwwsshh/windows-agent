package funcs

import (
	"strings"
	"time"

	"github.com/freedomkk-qfeng/windows-agent/g"
	"github.com/open-falcon/common/model"
	"github.com/shirou/gopsutil/net"
)

func net_status(ifacePrefix []string) ([]net.IOCountersStat, error) {
	net_iocounter, err := net.IOCounters(true)
	netIfs := []net.IOCountersStat{}
	for _, iface := range ifacePrefix {
		for _, netIf := range net_iocounter {
			if strings.Contains(netIf.Name, iface) {
				netIfs = append(netIfs, netIf)
			}
		}
	}
	return netIfs, err
}

func NetMetrics() []*model.MetricValue {
	return CoreNetMetrics(g.Config().Collector.IfacePrefix)
}

func CoreNetMetrics(ifacePrefix []string) (L []*model.MetricValue) {
	var startTime,endTime time.Time
	if g.Config().Debug {
		startTime = time.Now()
	}

	netIfs, err := net_status(ifacePrefix)
	if err != nil {
		g.Logger().Println(err)
		return []*model.MetricValue{}
	}

	for _, netIf := range netIfs {
		iface := "iface=" + netIf.Name
		L = append(L, CounterValue("net.if.in.bytes", netIf.BytesRecv, iface)) //此处乘以8即为bit的流量
		L = append(L, CounterValue("net.if.in.packets", netIf.PacketsRecv, iface))
		L = append(L, CounterValue("net.if.in.errors", netIf.Errin, iface))
		L = append(L, CounterValue("net.if.in.dropped", netIf.Dropin, iface))
		L = append(L, CounterValue("net.if.in.fifo.errs", netIf.Fifoin, iface))
		L = append(L, CounterValue("net.if.out.bytes", netIf.BytesSent, iface)) //此处乘以8即为bit的流量
		L = append(L, CounterValue("net.if.out.packets", netIf.PacketsSent, iface))
		L = append(L, CounterValue("net.if.out.errors", netIf.Errout, iface))
		L = append(L, CounterValue("net.if.out.dropped", netIf.Dropout, iface))
		L = append(L, CounterValue("net.if.out.fifo.errs", netIf.Fifoout, iface))
	}
	if g.Config().Debug {
		endTime = time.Now()
		g.Logger().Printf("collect NetMetrics complete. Process time %s.", endTime.Sub(startTime))
	}
	return
}
