package system

import (
	"context"
	"io"

	"github.com/alexellis/go-execute/v2"
)

type Command struct {
	task execute.ExecTask
}

func (c Command) Cmd(cmd string) Command {
	c.task.Command = cmd
	return c
}

func (c Command) Args(args []string) Command {
	c.task.Args = args
	return c
}

func (c Command) Cwd(cwd string) Command {
	c.task.Cwd = cwd
	return c
}

func (c Command) Stdin(stdin io.Reader) Command {
	c.task.Stdin = stdin
	return c
}

func (c Command) StreamStdio(stream bool) Command {
	c.task.StreamStdio = stream
	return c
}

func (c Command) Exec(ctx context.Context) (execute.ExecResult, error) {
	return c.task.Execute(ctx)
}

func Call() Command {
	return Command{}
}
