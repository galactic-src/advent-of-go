package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
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
	default:
		log.Fatalf("Don't know how to Map type %s", reflect.TypeOf(t).Name())
	}
}

func parseInt(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		log.Fatalf("urgh")
	}
	return i
}

/*
func Zip(t1 interface{}, t2 interface{}) [][2]interface{} {
	assertSlice("t1", "Zip", t1)
	assertSlice("t2", "Zip", t2)

	t1.len()
}*/

func assertSlice(variable, function string, slice interface{}) {
	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		log.Fatalf("Expected a slice, but got a %s for variable %s in %s", reflect.TypeOf(slice).Name(), variable, function)
	}
}

func main() {
	origin := []int{4,5,3}
	newArray := Map(origin, func(item interface{}) interface{} { return item.(int) + 1})
	fmt.Print(newArray, "\n")
	ForEach(origin, func(item interface{}) {fmt.Print(item, "\n")})
}