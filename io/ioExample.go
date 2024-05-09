package main

import (
	"fmt"
	"strings"
	"io"
	"bytes"
)

func printContents(reader io.Reader) {
	contents, _ := io.ReadAll(reader)
	fmt.Println(string(contents))
}

func main() {
	reader := strings.NewReader("HIHI")
	printContents(reader)

	buf := make([]byte, 1024) // make a []byte len 1024
	r := bytes.NewBuffer([]byte("")) 
	// read r into buf
	if _, err := r.Read(buf); err != nil {
		fmt.Println(err) // EOF err
	}

	buf2 := make([]byte, 1024)
	r2 := bytes.NewBuffer([]byte("HIHIHIH"))
	read, err := r2.Read(buf2)  // read r2 into buf2
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(read)
	
	// write
	w := bytes.NewBuffer([]byte("WRITE IT!"))

	res, err := w.Write(buf2); // write buf2 into w, so w will be WRITE IT!HIHIH
	
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(w.String())
	fmt.Println(res) // it will write all buf2 (len=1024)
}

