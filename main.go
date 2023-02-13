package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	line,err:=in.ReadString('\n')
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(line)
}