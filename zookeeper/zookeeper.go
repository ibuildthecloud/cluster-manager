package zookeeper

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/docker/engine-api/client"
	"github.com/rancher/cluster-manager/db"
	"github.com/rancher/cluster-manager/docker"
)

const (
	rancherZk = "rancher-zookeeper"
	zkId      = "zkId"
)

type Zookeeper struct {
	d           *docker.Docker
	clusterSize int
	cluster     []string
	zkId        int
	uuid        string
}

func New(uuid string, d *docker.Docker, clusterSize int) *Zookeeper {
	return &Zookeeper{
		d:           d,
		uuid:        uuid,
		clusterSize: clusterSize,
	}
}

func (z *Zookeeper) Update(byIndex map[int]db.Member) error {
	newCluster := []string{}
	z.zkId = 0
	for i := 1; i <= z.clusterSize; i++ {
		if byIndex[i].UUID == z.uuid {
			z.zkId = i
		}
		newCluster = append(newCluster, byIndex[i].IP)
	}

	if !reflect.DeepEqual(z.cluster, newCluster) {
		fmt.Println("CHANGED", newCluster)
		z.cluster = newCluster
		return z.configure()
	}
	return nil
}

func (z *Zookeeper) configure() error {
}

func (z *Zookeeper) RequestedIndex() (int, error) {
	c, err := z.d.Cli.ContainerInspect(rancherZk)
	if client.IsErrContainerNotFound(err) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	id := docker.ParseEnv(c.Config.Env)[zkId]
	if id == "" {
		return 0, nil
	}
	idNum, err := strconv.Atoi(id)
	if err != nil {
		return 0, nil
	}
	return idNum, nil
}
