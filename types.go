package main

import (
  "fmt"
  "reflect"
  "net"
  "time"
)

//int, int8, int16, int32, int64
//uint, uint8, uint16, uint32, uint64
//float32, float64
//string
//bool

func main() {
  fmt.Println(reflect.TypeOf(1))
  fmt.Println(reflect.TypeOf(1.5))
  fmt.Println(reflect.TypeOf("Hello"))
  fmt.Println(reflect.TypeOf(false))
  fmt.Println(reflect.TypeOf(net.IPv4(127, 9, 0, 1)))
  fmt.Println(reflect.TypeOf(time.Now()))
}
