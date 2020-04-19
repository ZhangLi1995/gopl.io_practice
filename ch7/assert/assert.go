package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)
	fmt.Printf("w: %T, %v\n", w, w)
	fmt.Printf("f: %T, %v\n", f, f)
	//c := w.(*bytes.Buffer)
	//fmt.Printf("%T, %v\n", c, c)

	fmt.Println("=====>")

	rw := w.(io.ReadWriter)
	fmt.Printf("f: %T, %v\n", rw, rw)

	w = rw
	w = rw.(io.Writer)

	fmt.Println("=====>")

	f, ok := w.(*os.File)
	fmt.Println(ok)
	b, ok := w.(*bytes.Buffer)
	fmt.Println(b, ok)

	var e = errors.New("file")
	fmt.Println(e.Error())
}

// writeString 将 s 写入 w
// 如果 w 有 WriteString 方法，那么将直接调用该方法
func writeString(w io.Writer, s string) (n int, err error) {
	type stringWriter interface {
		WriteString(string) (n int, err error)
	}
	if sw, ok := w.(stringWriter); ok {
		return sw.WriteString(s) // 避免了内存复制
	}
	return w.Write([]byte(s))
}

func writeHeader(w io.Writer, contentType string) error {
	if _, err := writeString(w, "Content-Type: "); err != nil {
		return err
	}
	if _, err := writeString(w, contentType); err != nil {
		return err
	}
	// ...
	return nil
}
