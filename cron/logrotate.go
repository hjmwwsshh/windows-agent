package cron

import (
	"github.com/freedomkk-qfeng/windows-agent/g"
	tcron "github.com/toolkits/cron"
)
var (
	logRotateCron = tcron.New()
)

func StartLogRotateCron() {
	//
	logRotateCronSpec := g.Config().LogRotate
	//step := 30
	//  if step < 10 || step > 60 {
	//  	step = 20
	//  }
	//logRotateCronSpec := fmt.Sprintf("0 0 0 */%d ?", step)
	//
	logRotateCron.AddFuncCC(logRotateCronSpec,func(){
		g.Logger().Printf("start reinit log.")
		g.InitLog()
	},1)
	logRotateCron.Start()
	g.Logger().Printf("start log rotate cron.")
}