package containers

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	units "github.com/docker/go-units"
	"github.com/perlogix/cmon/data"
)

// DockerContainers fetches all docker containers on system
func DockerContainers(d *data.DiscoverJSON) {
	cli, err := client.NewClientWithOpts(client.WithVersion("1.38"))
	if err != nil {
		return
	}

	cts, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		return
	}

	var ctrSlice []string
	for _, e := range cts {
		ctrSlice = append(ctrSlice, fmt.Sprintf("name=%s image=%s command=%s ports=%v state=%s status=%s", strings.Split(e.Names[0], "/")[1], e.Image, e.Command, e.Ports, e.State, e.Status))
	}

	d.DockerContainers = ctrSlice
}

// DockerServer grabs docker server information
func DockerServer(d *data.DiscoverJSON) {
	cli, err := client.NewClientWithOpts(client.WithVersion("1.38"))
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	info, err := cli.Info(ctx)
	if err != nil {
		return
	}

	d.DockerRunning = info.ContainersRunning
	d.DockerPaused = info.ContainersPaused
	d.DockerStopped = info.ContainersStopped
	d.DockerImagesCount = info.Images
	d.DockerLabels = info.Labels
}

// DockerImages grabs all docker images
func DockerImages(d *data.DiscoverJSON) {
	cli, err := client.NewClientWithOpts(client.WithVersion("1.38"))
	if err != nil {
		return
	}
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{All: true})
	if err != nil {
		return
	}
	var imgs []string
	for _, v := range images {
		imgs = append(imgs, fmt.Sprintf("name=%s size=%s created=%s", strings.Join(v.RepoTags, " "), units.HumanSize(float64(v.Size)), time.Unix(v.Created, 0).Format(time.RFC3339)))
	}
	d.DockerImages = imgs
}
