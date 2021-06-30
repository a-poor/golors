package golors

import "fmt"

const ColorReset  = "\033[0m"
const ColorRed    = "\033[31m"
const ColorGreen  = "\033[32m"
const ColorYellow = "\033[33m"
const ColorBlue   = "\033[34m"
const ColorPurple = "\033[35m"
const ColorCyan   = "\033[36m"
const ColorGray   = "\033[37m"
const ColorGrey   = ColorGray
const ColorWhite  = "\033[97m"


func Reset(a ...interface{}) {
	fmt.Print(ColorReset)
}

func Red(a ...interface{}) {
	fmt.Print(ColorRed)
	fmt.Print(a...)
	Reset()
}

func Redln(a ...interface{}) {
	Red(a...)
	fmt.Println()
}

func Green(a ...interface{}) {
	fmt.Print(ColorGreen)
	fmt.Print(a...)
	Reset()
}

func Greenln(a ...interface{}) {
	Green(a...)
	fmt.Println()
}

func Yellow(a ...interface{}) {
	fmt.Print(ColorYellow)
	fmt.Print(a...)
	Reset()
}

func Yellowln(a ...interface{}) {
	Yellow(a...)
	fmt.Println()
}

func Blue(a ...interface{}) {
	fmt.Print(ColorBlue)
	fmt.Print(a...)
	Reset()
}

func Blueln(a ...interface{}) {
	Blue(a...)
	fmt.Println()
}

func Purple(a ...interface{}) {
	fmt.Print(ColorPurple)
	fmt.Print(a...)
	Reset()
}

func Purpleln(a ...interface{}) {
	Red(a...)
	fmt.Println()
}

func Cyan(a ...interface{}) {
	fmt.Print(ColorCyan)
	fmt.Print(a...)
	Reset()
}

func Cyanln(a ...interface{}) {
	Cyan(a...)
	fmt.Println()
}

func Gray(a ...interface{}) {
	fmt.Print(ColorGray)
	fmt.Print(a...)
	Reset()
}

func Grayln(a ...interface{}) {
	Gray(a...)
	fmt.Println()
}

func Grey(a ...interface{}) {
	fmt.Print(ColorGrey)
	fmt.Print(a...)
	Reset()
}

func Greyln(a ...interface{}) {
	Grey(a...)
	fmt.Println()
}

func White(a ...interface{}) {
	fmt.Print(ColorWhite)
	fmt.Print(a...)
	Reset()
}

func Whiteln(a ...interface{}) {
	White(a...)
	fmt.Println()
}

