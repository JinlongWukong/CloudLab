package vm

import (
	"sync"
	"time"
)

const (
	VmStatusInit      = "init"
	VmStatusScheduled = "scheduled"
	VmStatusCreating  = "creating"
	VmStatusRunning   = "running"
	VmStatusDeleting  = "deleting"
	VmStatusDeleted   = "deleted"
)

var flavorDetails = map[string]map[string]int32{
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

type VncInfo struct {
	Port string `json:"port"`
	Pass string `json:"passwd"`
}

type VirtualMachine struct {
	Name         string         `json:"name"`
	Hostname     string         `json:"hostname"`
	CPU          int32          `json:"cpu"`
	Memory       int32          `json:"mem"`
	Disk         int32          `json:"disk"`
	IpAddress    string         `json:"address"`
	Status       string         `json:"status"`
	Vnc          VncInfo        `json:"vnc"`
	Type         string         `json:"type"`
	Node         string         `json:"node"`
	Lifetime     time.Duration  `json:"lifeTime"`
	PortMap      map[int]string `json:"portMap"`
	RootPass     string         `json:"rootPass"`
	sync.RWMutex `json:"-" gob:"-"`
	lifeMutex    sync.RWMutex `json:"-"`
}

type VmRequest struct {
	Account  string `form:"account" json:"account" binding:"required"`
	Hostname string `form:"hostname" json:"hostname"`
	RootPass string `form:"rootPass" json:"rootPass"`
	Type     string `form:"type" json:"type" binding:"required"`
	Flavor   string `form:"flavor" json:"flavor"`
	CPU      int32  `form:"cpu" json:"cpu"`
	Memory   int32  `form:"mem" json:"memory"`
	Disk     int32  `form:"disk" json:"disk"`
	Number   int32  `form:"numbers" json:"number" binding:"required,min=1,max=5"`
	Duration int    `form:"duration" json:"duration" binding:"required"`
}

type VmRequestGetVm struct {
	Account string `form:"account" json:"account" binding:"required"`
	Name    string `form:"name,omitempty" json:"name,omitempty"`
}

type VmRequestPostAction struct {
	Account string `form:"account" json:"account" binding:"required"`
	Name    string `form:"name" json:"name" binding:"required"`
	Action  string `form:"action" json:"action" binding:"required"`
}

type VmRequestPortExpose struct {
	Account  string `form:"account" json:"account" binding:"required"`
	Name     string `form:"name" json:"name" binding:"required"`
	Port     int    `form:"port" json:"port" binding:"required,min=1"`
	Protocol string `form:"protocol,default=tcp" json:"protocol,default=tcp" binding:"required"`
}

type VmLiveStatus struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Address string `json:"address"`
	VncPort string `json:"vncPort"`
}
