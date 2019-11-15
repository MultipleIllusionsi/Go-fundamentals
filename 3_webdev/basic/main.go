package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Sergey Zakharov"
	str := fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="en">
		<head>
		<meta charset="UTF-8" />
		<title>Go lang</title>
		</head>
		<body>
			<h1>` + name + `</h1>
		</body>
	</html>
	`)

	nf, err := os.Create("go.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))
}
