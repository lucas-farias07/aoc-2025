package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var startPoint int = 50

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	//ans := solvePwd(lines)
	ans := solvePwd2(lines)
	fmt.Println(ans)
}

func solvePwd2(lines []string) (ans int) { // uses the new 0x smth method in part 2
	ans = 0
	lock := startPoint
	for _, line := range lines {
		nRotations, magnitude := parseRotation(line)
		lock = modBounds(0, magnitude*lock) // mirrors it in case it is negative
		lock += nRotations
		ans += lock / 100
		lock = modBounds(0, magnitude*lock) // returns to normal
	}
	return
}

func solvePwd(lines []string) (ans int) {
	ans = 0
	lock := startPoint
	for _, line := range lines {
		nRotations, magnitude := parseRotation(line)
		lock = modBounds(lock, magnitude*nRotations)
		if lock == 0 {
			ans++
		}
	}
	return
}

func modBounds(start int, rotations int) int {
	return ((start+rotations)%100 + 100) % 100
}

func parseRotation(rotation string) (numerical int, sign int) {
	sign = 1 // made it return the sign to avoid extra code in solvePwd2
	numerical, err := strconv.Atoi(rotation[1:])
	if err != nil {
		panic(err)
	}

	if rotation[:1] == "L" {
		sign = -1
	}

	return
}
