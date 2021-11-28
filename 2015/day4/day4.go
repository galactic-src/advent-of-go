package main

import (
"crypto/md5"
"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func md5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

const input = "ckczppom"

func part1() {
	for i:=0; ;i++ {
		result := md5Hash(input + strconv.Itoa(i))
		if strings.HasPrefix(result, "00000") {
			fmt.Printf("Part 1: %d\n", i)
			break
		}
	}
}

func part2() {
	for i:=0; ;i++ {
		result := md5Hash(input + strconv.Itoa(i))
		if strings.HasPrefix(result, "000000") {
			fmt.Printf("Part 2: %d\n", i)
			break
		}
	}
}

func main() {
	part1()
	part2()
}