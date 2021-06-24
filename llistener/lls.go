package llistener

import (
	"github.com/seeadoog/goutils/filelock"
	"net"
)

type llistener struct {
	ls   net.Listener
	lock *filelock.FileLock
}

func NewFlockedListener(ls net.Listener, lockDir string) (net.Listener, error) {
	lock, err := filelock.New(lockDir)
	if err != nil {
		return nil, err
	}
	return &llistener{
		ls:   ls,
		lock: lock,
	}, nil
}

func (l llistener) Accept() (conn net.Conn, err error) {
	err = l.lock.Lock()
	if err != nil {
		return nil, err
	}
	conn, err = l.ls.Accept()

	if err = l.lock.Unlock(); err != nil {
		return nil, err
	}
	return
}

func (l llistener) Close() error {
	return l.ls.Close()
}

func (l llistener) Addr() net.Addr {
	return l.ls.Addr()
}
