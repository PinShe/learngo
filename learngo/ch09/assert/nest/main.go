package main

import "fmt"

type MyWriter interface {
	Write(string2 string)
}

type MyReader interface {
	Read() string
}

type MyReadrWriter interface {
	MyWriter
	MyReader
	ReadWrite()
}

type SreadWriter struct{}

func (s *SreadWriter) Write(string) {
	//TODO implement me
	fmt.Println("write")
}

func (s *SreadWriter) Read() string {
	//TODO implement me
	fmt.Println("read")
	return ""
}

func (s *SreadWriter) ReadWrite() {
	//TODO implement me
	fmt.Println("read write")
}

func main() {
	var mrw MyReadrWriter = &SreadWriter{}
	mrw.Read()
}
