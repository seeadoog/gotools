package gopool

import (
	"testing"
)

func TestNewRPS(t *testing.T) {

}

func BenchmarkRPS(b *testing.B) {
	rps := NewRPS(12, func() RoutinePool {
		return NewRPool()
	})
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next(){
			rps.Execute(func() {

			})
		}
	})
}

func BenchmarkRPS2(b *testing.B) {
	rps := NewRPS(12, func() RoutinePool {
		return NewFixedPool(12)
	})
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next(){
			rps.Execute(func() {

			})
		}
	})
}


