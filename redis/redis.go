package redis

import (
	"fmt"
	"strings"

	"github.com/rancher/cluster-manager/db"
	"github.com/rancher/cluster-manager/docker"
)

type Redis struct {
	d           *docker.Docker
	ipList      string
	clusterSize int
}

func New(d *docker.Docker, clusterSize int) *Redis {
	return &Redis{
		d:           d,
		clusterSize: clusterSize,
	}
}

func (r *Redis) Update(byIndex map[int]db.Member) error {
	ips := []string{}
	for i := 1; i <= r.clusterSize; i++ {
		ip := byIndex[i].IP
		if ip != "" {
			ips = append(ips, ip)
		}
	}

	ipList := strings.Join(ips, ",")
	if r.ipList != ipList {
		fmt.Println("REDIS CHANGED", ipList)
		r.ipList = ipList
	}

	return nil
}
