package htmlrenderingbasics

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var sampleHtml = `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	</head>
	<body>
		<h1>
			%s
		</h1>
		<input type="text" width="20px">
	</body>
	</html>
`

func RenderHtml() {
	str := fmt.Sprintf(sampleHtml, "Pranav")
	nf, err := os.Create("index.html")

	if err != nil {
		log.Fatalln("Error occured while creating file:", err)
	}

	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))
}
