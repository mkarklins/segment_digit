package main

import (
	"flag"
	"strconv"
	"time"

	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/all"
)

func main() {
	flag.Parse()

	if err := embd.InitGPIO(); err != nil {
		panic(err)
	}
	defer embd.CloseGPIO()

	for i := 0; i < 10; i++ {
		for gpioPin, isEnabled := range gpioLayoutFor(strconv.Itoa(i)) {

			embd.SetDirection(gpioPin, embd.Out)

			if isEnabled == 1 {
				embd.DigitalWrite(gpioPin, embd.High)
			} else {
				embd.DigitalWrite(gpioPin, embd.Low)
			}
		}

		time.Sleep(1 * time.Second)
	}
}

func gpioLayoutFor(letter string) map[int]int {
	var resultingGpio = make(map[int]int)

	for index, element := range segmentPinLayoutFor(letter) {
		resultingGpio[gpioForSegmentPosition(index)] = element
	}

	return resultingGpio
}

func gpioForSegmentPosition(position int) int {
	return map[int]int{
		0: 12,
		1: 16,
		2: 26,
		3: 25,
		4: 24,
		5: 23,
		6: 18,
	}[position]
}

func segmentPinLayoutFor(letter string) []int {
	return map[string][]int{
		"0": {1, 1, 1, 1, 1, 1, 0},
		"1": {0, 1, 1, 0, 0, 0, 0},
		"2": {1, 1, 0, 1, 1, 0, 1},
		"3": {1, 1, 1, 1, 0, 0, 1},
		"4": {0, 1, 1, 0, 0, 1, 1},
		"5": {1, 0, 1, 1, 0, 1, 1},
		"6": {1, 0, 1, 1, 1, 1, 1},
		"7": {1, 1, 1, 0, 0, 0, 0},
		"8": {1, 1, 1, 1, 1, 1, 1},
		"9": {1, 1, 1, 1, 0, 1, 1},
		" ": {0, 0, 0, 0, 0, 0, 0},
		"_": {0, 0, 0, 1, 0, 0, 0},
		"-": {0, 0, 0, 0, 0, 0, 1},
		"A": {1, 1, 1, 0, 1, 1, 1},
		"B": {0, 0, 1, 1, 1, 1, 1},
		"C": {0, 0, 0, 1, 1, 0, 1},
		"D": {0, 1, 1, 1, 1, 0, 1},
		"E": {1, 0, 0, 1, 1, 1, 1},
		"F": {1, 0, 0, 0, 1, 1, 1},
		"G": {1, 0, 1, 1, 1, 1, 0},
		"H": {0, 0, 1, 0, 1, 1, 1},
		"I": {0, 0, 1, 0, 0, 0, 0},
		"J": {0, 1, 1, 1, 1, 0, 0},
		"K": {1, 0, 1, 0, 1, 1, 1},
		"L": {0, 0, 0, 1, 1, 1, 0},
		"M": {1, 1, 1, 0, 1, 1, 0},
		"N": {0, 0, 1, 0, 1, 0, 1},
		"O": {0, 0, 1, 1, 1, 0, 1},
		"P": {1, 1, 0, 0, 1, 1, 1},
		"Q": {1, 1, 1, 0, 0, 1, 1},
		"R": {0, 0, 0, 0, 1, 0, 1},
		"S": {0, 0, 1, 1, 0, 1, 1},
		"T": {0, 0, 0, 1, 1, 1, 1},
		"U": {0, 0, 1, 1, 1, 0, 0},
		"V": {0, 1, 1, 1, 1, 1, 0},
		"W": {0, 1, 1, 1, 1, 1, 1},
		"X": {0, 1, 1, 0, 1, 1, 1},
		"Y": {0, 1, 1, 1, 0, 1, 1},
		"Z": {1, 1, 0, 1, 1, 0, 0},
	}[letter]
}
