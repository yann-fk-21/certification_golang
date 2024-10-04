package main

import (
	"flag"
	"fmt"

	"training.go/pdf_certif_project/cert"
	htmlcert "training.go/pdf_certif_project/htmlCert"
)


func main() {

	 typeFile := flag.String("type", "html", "Output html file")
	c, err :=cert.New("Golang programing", "Yann Fk", "2024-06-21")
	errorHandler(err)

	var saver cert.Saver

	switch *typeFile {
	case "html": 
	saver, err = htmlcert.New("output")
	case "pdf":
		saver, err = htmlcert.New("output")
	default:
		fmt.Println("This option is not present!")
	}
	
	saver.Save(c)

	errorHandler(err)


}

func errorHandler(err error) {
	if err != nil {
		fmt.Printf("have an error, err=%v", err)
		return
	}
}