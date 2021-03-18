// proc offers control over arbitrary processes.
// proc depends heavily on os/exec.
// If/when proc needs to do resource control & isolation
// its probable that os/exec will need to be replaced with
// either the os package itself or syscall.
package proc

// Fingers crossed that its easier than rewriting os.ForkExec

import (
	"io/ioutil"
	"os"
	"os/exec"
)

// Proc is a BitBox process.
// Stderr and Stdout are dumped to temp files on disk.
type Proc struct {
	// TODO: there is 99% chance we need a synchronization primitive at this level
	stdout *os.File
	stderr *os.File
	cmd    *exec.Cmd
}

// NewProc constructs and begins process.
// Stdout & Stderr of the new process are pointed at temp files.
// The tempfiles are acessable through the coresponding members.
func NewProc(cmdName string, args ...string) (Proc, error) {

	var err error
	var cmdPath string
	if cmdPath, err = exec.LookPath(cmdName); err != nil {
		return Proc{}, err
	}
	cmd := exec.Command(cmdPath, args...)

	var stdout *os.File
	if stdout, err = ioutil.TempFile("", ""); err != nil {
		return Proc{}, err
	}
	cmd.Stdout = stdout

	var stderr *os.File
	if stderr, err = ioutil.TempFile("", ""); err != nil {
		return Proc{}, err
	}
	cmd.Stderr = stderr

	if err = cmd.Start(); err != nil {
		return Proc{}, err
	}

	go func() {
		cmd.Wait()
		// When cmd.Wait returns cmd.ProcessState will be set.
		// TODO: send exit code to query subscribers
	}()

	return Proc{
		stdout: stdout,
		stderr: stderr,
		cmd:    cmd,
	}, nil
}

// Kill causes the running process to exit and closes the related resources.
// There is no guarantee that the process has actually exited when Kill returns.
// See documentation for os.Process.Kill()
func (p Proc) Kill() error {
	if err := p.cmd.Process.Kill(); err != nil {
		// processes which have ended return errors
		return err
	}
	// Only release resources if no err has been returned.
	// TODO: should we return the error or try to close both and return some conglomarate error?
	if err := p.stderr.Close(); err != nil {
		return err
	}
	if err := p.stdout.Close(); err != nil {
		return err
	}
	return nil
}

// Status returns the status of the
func (p Proc) Status() ProcStatus {
	if p.cmd.ProcessState == nil {
		return Running
	}

	if p.cmd.ProcessState.ExitCode() != 0 {
		return Stopped
	}
	return Exited
}

// ProcStatus is the status of a process.
type ProcStatus int

const (
	// Running indicates that the process is running.
	Running ProcStatus = iota
	// Stopped indicates that the process returned no exit code.
	Stopped
	// Exited indicates that the process returned a non-zero exit code.
	Exited
)

func (ps ProcStatus) String() string {
	return [...]string{"Stopped", "Running"}[ps]
}
