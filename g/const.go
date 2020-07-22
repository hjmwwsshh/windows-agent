package g

import (
	"time"
)

// changelog:
// 1.0.0 windows-agent
// 1.0.1 ifstat use ifname instead ifdescription
// 1.0.2 fix net.listen.port bug
// 1.0.3 add default tag ,fix net.listen.port bug
// 1.0.4 merge ctck1995's pull request to fix net.port.listen bug
// 1.0.5 add log rotate;add collected metrics time
const (
	VERSION          = "1.0.5"
	COLLECT_INTERVAL = time.Second
	NET_PORT_LISTEN  = "net.port.listen"
	DU_BS            = "du.bs"
	PROC_NUM         = "proc.num"
)
