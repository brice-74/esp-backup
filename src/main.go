package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func main() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err:%v", err)
	}
	for _, ci := range cpuInfos {
		fmt.Printf("%+#v", ci)
		fmt.Println()
	}
	//CPU utilization
	for {
		percent, err := cpu.Percent(time.Second, false)
		if err != nil {
			fmt.Printf("get cpu percent failed, err:%v", err)
		}
		fmt.Printf("cpu percent:%v\n", percent)
	}
}
