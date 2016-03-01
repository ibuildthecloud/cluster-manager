package docker

import (
	"strings"

	"github.com/docker/engine-api/client"
)

const (
	parent = "rancher-ha-networking"
)

type Docker struct {
	Cli *client.Client
}

func New() (*Docker, error) {
	defaultHeaders := map[string]string{"User-Agent": "engine-api-cli-1.0"}
	cli, err := client.NewClient("unix:///var/run/docker.sock", "v1.22", nil, defaultHeaders)
	return &Docker{Cli: cli}, err
}

func (d *Docker) Name() (string, error) {
	i, err := d.Cli.Info()
	return i.Name, err
}

func (d *Docker) recreate(name string, command []string, env map[string]string) error {
	c, err := d.Cli.ContainerInspect(name)
	if err != nil && !client.IsErrContainerNotFound(err) {
		return err
	}

	if c != nil {
		return nil
	}

	d.Cli.ContainerCreate(client.Config{
		Im
	}

(config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, name
}

func (d *Docker) Launch(name string, command []string, env map[string]string) error {
}

func ParseEnv(env []string) map[string]string {
	result := map[string]string{}
	for _, v := range env {
		parts := strings.SplitN(v, "=", 2)
		if len(parts) == 2 {
			result[parts[0]] = parts[1]
		} else {
			result[parts[0]] = ""
		}
	}
	return result
}
