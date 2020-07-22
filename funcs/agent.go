package funcs

import (
	"github.com/freedomkk-qfeng/windows-agent/g"
	"github.com/open-falcon/common/model"
)

func AgentMetrics() []*model.MetricValue {
	if g.Config().Debug {
		g.Logger().Printf("collect AgentMetrics complete.")
	}

	return []*model.MetricValue{GaugeValue("agent.alive", 1)}
}
