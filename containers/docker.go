// Copyright (C) YetiCloud
// This file is part of yeti-discover <https://github.com/yeticloud/yeti-discover>.
//
// yeti-discover is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// yeti-discover is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with yeti-discover.  If not, see <http://www.gnu.org/licenses/>.

package containers

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	units "github.com/docker/go-units"
	"github.com/yeticloud/yeti-discover/data"
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
