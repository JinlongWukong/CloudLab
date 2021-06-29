package scheduler

import (
	"log"
	"math/rand"

	"github.com/JinlongWukong/CloudLab/config"
	"github.com/JinlongWukong/CloudLab/node"
)

var allocationRatio = 2
var scheduleAlgorithm = "random"

//initialize configuration
func Initialize() {
	if config.Schedule.AllocationRatio > 0 {
		allocationRatio = config.Schedule.AllocationRatio
	}
	if config.Schedule.ScheduleAlgorithm != "" {
		scheduleAlgorithm = config.Schedule.ScheduleAlgorithm
	}
}

//apply for a node
func Schedule(role node.NodeRole, reqCpu, reqMem, reqDisk int32) *node.Node {
	//filter all nodes
	winerNodes := make([]*node.Node, 0)
	for n := range node.NodeDB.Iter() {
		if role != n.Value.Role {
			continue
		} else if n.Value.State != node.NodeStateEnable ||
			n.Value.Status != node.NodeStatusReady ||
			(n.Value.CPU*int32(allocationRatio)-n.Value.GetCpuUsed()) < reqCpu ||
			(n.Value.Memory*int32(allocationRatio)-n.Value.GetMemUsed()) < reqMem ||
			(n.Value.Disk*int32(allocationRatio)-n.Value.GetDiskUsed()) < reqDisk {
			continue
		} else {
			winerNodes = append(winerNodes, n.Value)
		}
	}
	if len(winerNodes) == 0 {
		log.Println("No available node left")
		return nil
	}
	//Select one node based on scheduleAlgorithm
	if scheduleAlgorithm == "random" {
		return winerNodes[rand.Intn(len(winerNodes))]
	} else {
		return winerNodes[rand.Intn(len(winerNodes))]
	}
}
