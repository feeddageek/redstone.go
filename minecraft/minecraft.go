//Package minecraft provide way to start minecraft servers and interact with them
package minecraft

import (
	"bufio"
	"io"
	"os/exec"
)

//Jar contain the path to a minecraft server jar and its jvm's arguments
type Jar struct {
	jar     string
	args    []string
}

//Jar constructor
func NewJar(jar string, args ...string) Jar{
	return Jar{jar:jar,args:args}
}

//World contain the path to a mincraft world
type World struct {
	path	string
}

//World constructor
func NewWorld(path string) World{
	return World{path:path}
}

//Instance contain usefull data about a running instance
type Instance struct {
	jar	Jar
	world	World
	cmd     *exec.Cmd
	out	*bufio.Scanner
	in	io.Writer
}

//Start launch a new instance and return its handel
func Start(jar Jar,world World) (i *Instance,err error) {
	i = new(Instance)
	i.cmd = exec.Command("java",append(jar.args,"-jar",jar.jar,"nogui")...)
	i.cmd.Dir = world.path
	out, err := i.cmd.StdoutPipe()
	if err != nil{
		return nil,err
	}
	i.out = bufio.NewScanner(out)
	i.in, err = i.cmd.StdinPipe()
	if err != nil{
		return nil,err
	}
	err = i.cmd.Start()
	if err != nil {
		return nil,err
	}
	return
}
//Stop request the instance to stop and wait until its process return
func (i *Instance) Stop() (err error) {
	_, err = i.in.Write([]byte("stop\n"))
	i.cmd.Wait()
	return
}

func (i *Instance) parse(string) {
	//TODO extract information form the server output and do somthing with it
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
