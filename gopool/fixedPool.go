package gopool



type fixedPool struct {
	taskSize int
	taskChan chan Task
	maxQueueSize int
}


func NewFixedPool(size int)RoutinePool{
	f:= &fixedPool{
		taskSize:     size,
		taskChan:     make(chan Task,10000),
		maxQueueSize: 10000,
	}
	f.run()
	return f
}

func (f *fixedPool) run()  {
	for i := 0; i < f.taskSize; i++ {
		go func() {
			for{
				t := <- f.taskChan
				t()
			}
		}()
	}
}



func (f *fixedPool) Execute(task Task) {
	f.taskChan<-task
}

