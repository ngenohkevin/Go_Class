package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func getOne(i int) []byte {

}

func main() {
	var (
		output io.WriteCloser = os.Stdout
		err    error
		cnt    int
		fails  int
		data   []byte
	)

	if len(os.Args) > 0 {
		output, err = os.Create(os.Args[1])

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		defer output.Close()
	}

	// the output will be in the form of a JSON array,
	// so add the brackets before and after

	fmt.Fprint(output, "[")
	defer fmt.Fprint(output, "]")

	// some code here
	for i := 1; fails < 2; i++ {
		data  = getOne(); data == nil {
			fail++
			continue
		}
		if cnt > 0 {
			fmt.Fprint(output, ",") 
		}
	
		_, err = io.Copy(output, bytes.NewBuffer(data))
		if err != nil{
			fmt.Fprintf(os.Stderr, "stopped: %s\n", err)
			os.Exit(-1)
		}
	}
	


	fmt.Fprintf(os.Stderr, "read %d comics\n", cnt)
}
