package main

import (
	"fmt"
	"os"
)

func main() {
  if len(os.Args) > 1 {
    fmt.Printf("Regular : " + os.Args[1] + "\n")
    fmt.Printf("Bold    : " + "\033[1m%s\033[0m\n", os.Args[1])
    fmt.Printf("Italic  : " + "\033[3m%s\033[0m\n", os.Args[1])
  }else{

    fmt.Println("This is normal text. \033[1mThis is bold text!\033[0m \033[3mThis is italic text!\033[0m")
  }
}
