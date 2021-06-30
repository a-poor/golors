package main

import "fmt"

const reset  = "\033[0m"
const red    = "\033[31m"
const green  = "\033[32m"
const yellow = "\033[33m"
const blue   = "\033[34m"
const purple = "\033[35m"
const cyan   = "\033[36m"
const gray   = "\033[37m"
const grey   = gray
const white  = "\033[97m"


func Reset(a ...interface{}) {
	fmt.Print(reset)
}

func Red(a ...interface{}) {
	fmt.Print(red)
	fmt.Print(a...)
	Reset()
}

func Redln(a ...interface{}) {
	Red(a...)
	fmt.Println()
}

func Green(a ...interface{}) {
	fmt.Print(green)
	fmt.Print(a...)
	Reset()
}

func Greenln(a ...interface{}) {
	Green(a...)
	fmt.Println()
}

func Yellow(a ...interface{}) {
	fmt.Print(yellow)
	fmt.Print(a...)
	Reset()
}

func Yellowln(a ...interface{}) {
	Yellow(a...)
	fmt.Println()
}

func Blue(a ...interface{}) {
	fmt.Print(blue)
	fmt.Print(a...)
	Reset()
}

func Blueln(a ...interface{}) {
	Blue(a...)
	fmt.Println()
}

func Purple(a ...interface{}) {
	fmt.Print(purple)
	fmt.Print(a...)
	Reset()
}

func Purpleln(a ...interface{}) {
	Red(a...)
	fmt.Println()
}

func Cyan(a ...interface{}) {
	fmt.Print(cyan)
	fmt.Print(a...)
	Reset()
}

func Cyanln(a ...interface{}) {
	Cyan(a...)
	fmt.Println()
}

func Gray(a ...interface{}) {
	fmt.Print(gray)
	fmt.Print(a...)
	Reset()
}

func Grayln(a ...interface{}) {
	Gray(a...)
	fmt.Println()
}

func Grey(a ...interface{}) {
	Gray(a...)
}

func Greyln(a ...interface{}) {
	Grayln(a...)
}

func White(a ...interface{}) {
	fmt.Print(white)
	fmt.Print(a...)
	Reset()
}

func Whiteln(a ...interface{}) {
	White(a...)
	fmt.Println()
}

