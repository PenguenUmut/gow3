package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	fname := os.Args[1]
	lname := os.Args[2]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
	str := fmt.Sprint(`
		<!DOCTYPE html>
		<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title>Hello World!</title>
			</head>
			<body>
				<h1>` + fname + ` ` + lname + `</h1>
			</body>
		</html>
	`)

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))
}

/*
go run main.go Penguen Umut
*/