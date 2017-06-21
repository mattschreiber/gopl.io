// Package lengthconv converts meters to feet and vice versa
package lengthconv

import "fmt"

type Meter float64
type Feet float64

func (m Meter) String() string {
  return fmt.Sprintf("%gm", m)
}

func (f Feet) String() string {
  return fmt.Sprintf("%gft", f)
}
