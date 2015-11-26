//Package minecraft provide way to start minecraft server ans interact with it
package minecraft

import (
	"bufio"
	"errors"
	"io"
	"os/exec"
)

//Minecraft contain the exec.Cmd and I/O of the current running server
type Minecraft struct {
	cmd     *exec.Cmd
	jar     string
	path    string
	args    []string
	out     *bufio.Scanner
	in      io.Writer
	running bool
}

//New Minecraft constructor
func New(path string, jar string, args ...string) *Minecraft {
	return &Minecraft{path: path, jar: jar, args: args, running: false}
}

func (m *Minecraft) initCmd() {
	m.cmd = exec.Command("java", append(m.args, "-jar", m.jar, "nogui")...)
	m.cmd.Dir = m.path
	out, _ := m.cmd.StdoutPipe()
	m.out = bufio.NewScanner(out)
	m.in, _ = m.cmd.StdinPipe()
}

//Scan wait until the next line from the standard output is ready and return true
//or until an error occurs and return false
//errors are likely to be EOF after the terminaison of the server
func (m *Minecraft) Scan() bool {
	return m.out.Scan()
}

//Text retrun a line from the standard output of the running server
func (m *Minecraft) Text() string {
	return m.out.Text()
}

//Start initialise input and output for the server and start it
func (m *Minecraft) Start() (err error) {
	if m.running {
		err = errors.New("Could not start mincraft : already running")
		return
	}
	m.initCmd()
	err = m.cmd.Start()
	if err == nil {
		m.running = true
	}
	return
}

//Stop request the server to stop and wait until its process return
func (m *Minecraft) Stop() (err error) {
	if !m.running {
		err = errors.New("Could not stop mincraft : not running")
		return
	}
	_, err = m.in.Write([]byte("stop\n"))
	m.cmd.Wait()
	m.running = false
	return
}
