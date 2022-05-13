package containers

import (
	"context"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	units "github.com/docker/go-units"
	"github.com/perlogix/cmon/data"
)

// DockerContainers fetches all docker containers on system
func DockerContainers(d *data.DiscoverJSON) {
	dockerContainers := []data.DockerContainersInfo{}

	cli, err := client.NewClientWithOpts(client.WithVersion("1.38"))
	if err != nil {
		d.DockerContainers = dockerContainers
		return
	}

	cts, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		d.DockerContainers = dockerContainers
		return
	}

	for _, e := range cts {
		dockerContainers = append(dockerContainers, data.DockerContainersInfo{
			Name:    strings.Split(e.Names[0], "/")[1],
			Image:   e.Image,
			Command: e.Command,
			Ports:   e.Ports,
			State:   e.State,
			Status:  e.Status,
		})
	}

	d.DockerContainers = dockerContainers
}

// DockerServer grabs docker server information
func DockerServer(d *data.DiscoverJSON) {
	cli, err := client.NewClientWithOpts(client.WithVersion("1.38"))
	if err != nil {
		d.DockerLabels = []string{}
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	info, err := cli.Info(ctx)
	if err != nil {
		d.DockerLabels = []string{}
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
	dockerImages := []data.DockerImagesInfo{}

	cli, err := client.NewClientWithOpts(client.WithVersion("1.38"))
	if err != nil {
		d.DockerImages = dockerImages
		return
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{All: true})
	if err != nil {
		d.DockerImages = dockerImages
		return
	}

	for _, v := range images {
		dockerImages = append(dockerImages, data.DockerImagesInfo{
			Name:    strings.Join(v.RepoTags, " "),
			Size:    units.HumanSize(float64(v.Size)),
			Created: time.Unix(v.Created, 0).Format(time.RFC3339),
		})
	}

	d.DockerImages = dockerImages
}
