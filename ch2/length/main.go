package main

import (
  "fmt"
  "os"
  "strconv"
  "gopl.io/ch2/lengthconv"
)

func main() {

  for _, arg := range os.Args[1:] {
    t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
      fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
    m := lengthconv.Meter(t)
    f := lengthconv.Feet(t)

    fmt.Printf("%s = %s, %s = %s\n", m, lengthconv.MToF(m), f, lengthconv.FToM(f))
  }
}

// fmt.Printf("%s = %s, %s = %s\n",
//   f, tempconv.FToC(f), c, tempconv.CToF(c))
