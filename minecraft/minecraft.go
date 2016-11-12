//Package minecraft provide way to start minecraft servers and interact with them
package minecraft

import (
	"bufio"
	"errors"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

//Jar contain the path to a minecraft server jar and its jvm's arguments
type Jar struct {
	Jar   string
	JArgs []string
	MArgs []string
}

//World contain the path to a mincraft world
type World struct {
	Path string
	used bool
}

//Instance contain usefull data about a running instance
type Instance struct {
	jar     Jar
	world   *World
	cmd     *exec.Cmd
	out     *bufio.Scanner
	in      io.Writer
	running bool
}

//Start launch a new instance and return its handel
func Start(jar Jar, world *World) (*Instance, error) {
	var i Instance
	var err error
	if world.used {
		err = errors.New("World already in use")
		return nil, err
	}
	world.used = true
	i.world = world
	i.jar = jar
	var jarpath string
	if jarpath, err = filepath.Abs(jar.Jar); err != nil {
		return nil, err
	}
	i.cmd = exec.Command("java", append(append(jar.JArgs, "-jar", jarpath), jar.MArgs...)...)
	if i.cmd.Dir, err = filepath.Abs(world.Path); err != nil {
		return nil, err
	}
	//if out, err := i.cmd.StdoutPipe();err != nil {
	//	return nil,err
	//}
	i.cmd.Stdout = os.Stdout
	//i.out = bufio.NewScanner(out)
	//if i.in, err = i.cmd.StdinPipe();err != nil {
	//	return nil,err
	//}
	i.cmd.Stdin = os.Stdin
	if err = i.cmd.Start(); err != nil {
		return nil, err
	}
	i.running = true
	go i.atStop()
	return &i, err
}

//Stop request the instance to stop and wait until its process return
func (i *Instance) Stop() (err error) {
	_, err = i.in.Write([]byte("stop\n"))
	if err == nil {
		err = i.cmd.Wait()
	}
	return
}

func (i *Instance) parse(string) {
	//TODO extract information form the server output and do somthing with it
}

//atStop wait for cmd and set running to false
func (i *Instance) atStop() {
	err := i.cmd.Wait()
	i.world.used = false
	i.running = false
	if err != nil {
		//TODO something has to be done with that error
	}
}

//Running return false if the cmd exited, true otherwise
func (i *Instance) Running() bool {
	return i.running
}

//Scan wait until the next line from the standard output is ready and return true
//or until an error occurs and return false
//errors are likely to be EOF after the terminaison of the server
func (i *Instance) Scan() bool {
	return i.out.Scan()
}

//Text retrun a line from the standard output of the running server
func (i *Instance) Text() string {
	return i.out.Text()
}
