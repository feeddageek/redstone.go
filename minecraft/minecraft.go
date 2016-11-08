//Package minecraft provide way to start minecraft servers and interact with them
package minecraft

import (
	"bufio"
	"io"
	"os"
	"os/exec"
)

//Jar contain the path to a minecraft server jar and its jvm's arguments
type Jar struct {
	Jar     string
	Args    []string
}

//World contain the path to a mincraft world
type World struct {
	Path	string
}

//Instance contain usefull data about a running instance
type Instance struct {
	jar	Jar
	world	World
	cmd     *exec.Cmd
	out	*bufio.Scanner
	in	io.Writer
	running	bool
}

//Start launch a new instance and return its handel
func Start(jar Jar,world World) (*Instance, error) {
	var i Instance
	var err error
	i.cmd = exec.Command("java",append(jar.Args,"-jar",jar.Jar,"nogui")...)
	i.cmd.Dir = world.Path
	//out, err := i.cmd.StdoutPipe()
	i.cmd.Stdout = os.Stdout
	if err == nil{
		//i.out = bufio.NewScanner(out)
		//i.in, err = i.cmd.StdinPipe()
		i.cmd.Stdin = os.Stdin
		if err == nil{
			err = i.cmd.Start()
			if err == nil {
				i.running = true
				go i.atStop()
				return &i,err
			}
		}

	}
	return nil,err
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
func (i *Instance) atStop(){
	err := i.cmd.Wait()
	i.running = false
	if err != nil{
		//TODO something has to be done with that error
	}
}

//Running return false if the cmd exited, true otherwise
func (i *Instance) Running() bool{
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
