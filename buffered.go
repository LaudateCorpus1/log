package log

import (
	"fmt"
	"io"
)

var BufferFull = fmt.Errorf("Async Logging Buffer Is Full")

const BufferSize = 1000

func NewBufferedWriter(writer io.Writer) *BufferedWriter {
	bw := &BufferedWriter{
		Writer: writer,
		queue:  make(chan []byte, BufferSize),
		done:   make(chan struct{}),
	}
	return bw
}

type BufferedWriter struct {
	io.Writer
	queue chan []byte
	done  chan struct{}
}

//Write impl io.Writer
func (this *BufferedWriter) Write(p []byte) (n int, err error) {

	dst := make([]byte, len(p), len(p))
	copy(dst, p)

	select {
	case this.queue <- dst:
		return len(dst), nil
	default:
		return 0, BufferFull
	}
}

//process items and write to the underlying writer
func (this *BufferedWriter) Start() {
	go func() {
		defer close(this.done)
		for {
			select {
			case item, ok := <-this.queue:
				if !ok {
					return
				}
				this.Writer.Write(item)
			}
		}
	}()
}

//stop processing items and close
func (this *BufferedWriter) Close() error {
	close(this.queue)
	<-this.done
	return nil
}
