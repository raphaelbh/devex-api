package usecase

import (
	"context"
	"io"
	"os"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

type TriggerPipelineCommand struct {
	Commands []string
}

func TriggerPipeline(command TriggerPipelineCommand) {
	dockerImage := "docker.io/library/alpine"

	go func() {
		ctx := context.Background()
		cli, _ := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		defer cli.Close()

		reader, _ := cli.ImagePull(ctx, dockerImage, image.PullOptions{})
		defer reader.Close()
		io.Copy(os.Stdout, reader)

		resp, _ := cli.ContainerCreate(ctx, &container.Config{
			Image: dockerImage,
			Cmd:   []string{"tail", "-f", "/dev/null"},
			Tty:   false,
		}, &container.HostConfig{
			Privileged: true,
			Binds:      []string{"/var/run/docker.sock:/var/run/docker.sock"},
		}, nil, nil, "")

		cli.ContainerStart(ctx, resp.ID, container.StartOptions{})
		execResp, _ := cli.ContainerExecCreate(ctx, resp.ID, container.ExecOptions{
			Cmd:          []string{"sh", "-c", strings.Join(command.Commands, " && ")},
			AttachStdout: true,
			AttachStderr: true,
		})

		hijackedResp, _ := cli.ContainerExecAttach(ctx, execResp.ID, container.ExecAttachOptions{})
		defer hijackedResp.Close()

		stdcopy.StdCopy(os.Stdout, os.Stderr, hijackedResp.Reader)

		cli.ContainerStop(ctx, resp.ID, container.StopOptions{})
		cli.ContainerRemove(ctx, resp.ID, container.RemoveOptions{Force: true})
	}()
}
