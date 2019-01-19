package main

import "fmt"

type Retriever interface{
	Get(url string) string
}

func download(r Retriever) string{
	return  r.Get("www.baidu.com")
}

func main(){
	var r="sss"

	r+="sssfffr"
	fmt.Print(r)
}