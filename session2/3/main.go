package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	name := os.Args[1]
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[0])

	str := fmt.Sprint(`
  <!DOCTYPE html>
  <html>
  <html lang="en">
  <head>
  <meta charset="UTF-8">
  <title>Hello world!</title>
  </head>
  <body>
  <h1>` + name + `</h1>
  </body>
  </html>

    `)

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Error creating file", err)

	}

	defer nf.Close()
	io.Copy(nf, strings.NewReader(str))

}
