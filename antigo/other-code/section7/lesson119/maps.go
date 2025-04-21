package main

import "fmt"

func main() {
	webSites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}

	fmt.Println(webSites)
	fmt.Println(webSites["Google"]) //Passando a chave para ter acesso ao valor

	webSites["LinkedIn"] = "https://linkedin.com"

	delete(webSites, "Google")
	fmt.Println(webSites)

}
