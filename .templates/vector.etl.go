package src

import (
	"github.com/GabeCordo/etl/components/channel"
	"log"
	"time"
)

type Vector struct {
	x int
	y int
}

func (m Vector) ExtractFunc(output channel.OutputChannel) {
	v := Vector{1, 5} // simulate pulling data from a source
	for i := 0; i < 10; i++ {
		output <- v // send data to the TransformFunc
	}
	close(output)
}

func (m Vector) TransformFunc(input channel.InputChannel, output channel.OutputChannel) {
	for request := range input {
		if v, ok := (request).(Vector); ok {
			v.x = v.x * 5
			v.y = v.y * 5

			output <- v // send data to the LoadFunc
		}
		time.Sleep(500 * time.Millisecond)
	}
	close(output)
}

func (m Vector) LoadFunc(input channel.InputChannel) {
	for request := range input {
		if v, ok := (request).(Vector); ok {
			log.Printf("Vector(%d, %d)", v.x, v.y)
		}
	}
}
