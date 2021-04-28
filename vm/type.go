package vm

import (
	"time"

	"github.com/JinlongWukong/CloudLab/node"
)

var flavorDetails = map[string]map[string]int{
	"small": {
		"cpu":    2,
		"memory": 2048,
		"disk":   30,
	},
	"middle": {
		"cpu":    4,
		"memory": 4096,
		"disk":   64,
	},
	"large": {
		"cpu":    6,
		"memory": 8192,
		"disk":   80,
	},
}

type VirtualMachine struct {
	Name      string
	CPU       int
	Memory    int
	Disk      int
	IpAddress string
	Status    string
	Type      string
	Host      node.ComputeNode
	Lifetime  time.Duration
}

type VmRequest struct {
	Account  string `form:"cecid"`
	Type     string `form:"os_type"`
	Flavor   string `form:"os_flavor"`
	Number   int    `form:"os_numbers"`
	Duration int    `form:"os_duration"`
}

type VmLiveStatus struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Address string `json:"address"`
}
