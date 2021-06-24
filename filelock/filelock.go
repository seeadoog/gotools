package filelock

import (
	"fmt"
	"os"
	"syscall"
)

type FileLock struct {
	dir string
	f   *os.File
}

func New(dir string) (*FileLock, error) {
	lock := &FileLock{
		dir: dir,
	}
	f, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	lock.f = f
	return lock, nil
}

//加锁
func (l *FileLock) Lock() error {

	err := syscall.Flock(int(l.f.Fd()), syscall.LOCK_EX)
	if err != nil {
		return fmt.Errorf("cannot flock directory %s - %s", l.dir, err)
	}
	return nil
}

//释放锁
func (l *FileLock) Unlock() error {
	return syscall.Flock(int(l.f.Fd()), syscall.LOCK_UN)
}

func (l *FileLock) Close() error {
	return l.f.Close()
}
