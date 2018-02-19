package main
import "fmt"

func main() {

  i := 0

  fmt.Println("START")

  mylable:
    i = i + 1
    fmt.Println(i)

    if i < 10 {
      goto mylable
    }

  fmt.Println("END")

}
