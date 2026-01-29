package main

import (
	singleflight "myproject/single-flight"
)

func main() {
   	singleflight.NewGroup().Do("Hi", func() (interface{}, error) {
		return struct{}{}, nil
	})
}