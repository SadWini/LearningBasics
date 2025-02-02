package counterwriter

import "io"

type CounterWriter struct {
	counter int64
	writer  io.Writer
}

func (cw *CounterWriter) Write(p []byte) (int, error) {
	cw.counter += int64(len(p))
	return cw.writer.Write(p)
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := CounterWriter{0, w}
	return &cw, &cw.counter
}
