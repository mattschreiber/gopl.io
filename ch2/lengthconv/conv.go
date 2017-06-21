package lengthconv

//converts Meters to Feet
func MToF(m Meter) Feet {
  return Feet(m*3.28084)
}

func FToM(f Feet) Meter {
  return Meter(f*.3048)
}
