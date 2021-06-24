package filelock

import (
	"fmt"
	"os"
	"sync"
	"testing"
	"time"
)

func TestFileLock_Lock(t *testing.T) {
	main()
}


func main() {
	test_file_path, _ := os.Getwd()
	locked_file := test_file_path

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int) {
			for i := 0; i < 2; i++ {
				flock,err := New(locked_file)
				if err != nil{
					fmt.Println("err",err)
					return
				}
				err = flock.Lock()
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Printf("output : %d\n", num)
				flock.Unlock()
			}
			wg.Done()


		}(i)
	}
	wg.Wait()
	time.Sleep(2 * time.Second)

}

func BenchmarkFileLock(b *testing.B) {
	test_file_path, _ := os.Getwd()
	locked_file := test_file_path
	flock,err := New(locked_file)
	if err != nil{
		panic(err)
		return
	}
	for i := 0; i < b.N; i++ {
		flock.Lock()
		flock.Unlock()
	}
}
