```go
package main

import (
	"fmt"
)

func main() {
	Demo()
	Demo(1)
  fmt.Println(Demo(1, 2, 3))
  numbers := []int{4, 5, 6}
  fmt.Println(Demo(numbers))
}

func Demo(numbers ...int) {
  sum := 0
	for _, number := range numbers {
		sum += number
	}
  return sum
}
```
