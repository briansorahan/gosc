package sc

import (
	"github.com/briansorahan/go-osc/osc"
	. "github.com/briansorahan/sc/types"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
)

const (
	scsynth     = "/usr/bin/scsynth"
	DefaultPort = 57117
	listenPort  = 57118
	listenAddr  = "127.0.0.1"
)

type server struct {
	ErrChan    chan error
	addr       NetAddr
	statusChan chan *osc.OscMessage
	oscClient  *osc.OscClient
	oscServer  *osc.OscServer
	scsynth    *exec.Cmd
}

func (self *server) Addr() NetAddr {
	return self.addr
}

// Status gets the status of scsynth
func (self *server) Status() (ServerStatus, error) {
	statusReq := osc.NewOscMessage("/status")
	err := self.oscClient.Send(statusReq)
	if err != nil {
		log.Println("Failed to send /status message")
		return nil, err
	}
	msg := <-self.statusChan
	return newStatus(msg)
}

// Start starts scsynth
func (self *server) Start() error {
	port := strconv.Itoa(self.addr.Port)
	self.scsynth = exec.Command(scsynth, "-u", port)
	go func() {
		self.ErrChan <-self.scsynth.Run()
	}()
	// stop scsynth on interrupts and kills
	c := make(chan os.Signal)
	go func() {
		<-c
		err := self.stopScsynth()
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(1)
	}()
	signal.Notify(c, os.Interrupt, os.Kill)
	return nil
}

func (self *server) Close() error {
	var oscErr error
	if self.oscServer != nil {
		oscErr = self.oscServer.Close()
	}
	select {
	case err := <-self.ErrChan:
		return err
	default:
		if oscErr == nil {
			return self.stopScsynth()
		}
		return oscErr
	}
}

func (self *server) stopScsynth() error {
	if self.scsynth != nil {
		return syscall.Kill(self.scsynth.Process.Pid, syscall.SIGKILL)
	}
	return nil
}

func NewServer(addr NetAddr) Server {
	oscClient := osc.NewOscClient(addr.Addr, addr.Port)
	oscServer := osc.NewOscServer(listenAddr, listenPort)
	statusChan := make(chan *osc.OscMessage)
	oscClient.SetLocalAddr(listenAddr, listenPort)
	oscServer.AddMsgHandler(StatusOscAddress, func(msg *osc.OscMessage) {
		statusChan <- msg
	})
	errChan := make(chan error)
	go func() {
		errChan <- oscServer.ListenAndDispatch()
	}()
	s := server{
		errChan,
		addr,
		statusChan,
		oscClient,
		oscServer,
		nil,
	}
	return &s
}
