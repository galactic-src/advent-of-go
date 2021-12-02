package main

import (
	"fmt"
	"log"
	"reflect"
)

func Map(t interface{}, f func(interface{}) interface{} ) []interface{} {
	switch reflect.TypeOf(t).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(t)
		arr := make([]interface{}, s.Len())
		for i := 0; i < s.Len(); i++ {
			arr[i] = f(s.Index(i).Interface())
		}
		return arr
	default:
		log.Fatalf("Don't know how to Map type %s", reflect.TypeOf(t).Name())
		return nil
	}
}

func ForEach(t interface{}, f func(interface{})) {
	switch reflect.TypeOf(t).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(t)
		for i:= 0; i < s.Len(); i++ {
			f(s.Index(i).Interface())
		}
	}
}

func Zip(t1 interface{}, t2 interface{}) {

}

func main() {
	origin := []int{4,5,3}
	newArray := Map(origin, func(item interface{}) interface{} { return item.(int) + 1})
	fmt.Print(newArray, "\n")
	ForEach(origin, func(item interface{}) {fmt.Print(item, "\n")})
}