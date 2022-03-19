package main

import (
	"bufio"
	"log"
	"os"
)

type Job struct {
	Command string
	flag    int
	*log.Logger
}

func NewJob(command string, logger *log.Logger) *Job {
	return &Job{command, 0, logger}
}

func main() {
	job := Job{"123", 0, log.New(os.Stdin, "", 1)}
	println(job.flag)
	println(job.Logger.Flags())
	job.Println()
}

type ReadWriter struct {
	reader *bufio.Reader
	writer *bufio.Writer
}

func (rw *ReadWriter) Read(p []byte) (n int, err error) {
	return rw.reader.Read(p)
}
