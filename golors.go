package golors

import "fmt"

const colorReset  = "\033[0m"
const colorRed    = "\033[31m"
const colorGreen  = "\033[32m"
const colorYellow = "\033[33m"
const colorBlue   = "\033[34m"
const colorPurple = "\033[35m"
const colorCyan   = "\033[36m"
const colorGray   = "\033[37m"
const colorGrey   = colorGray
const colorWhite  = "\033[97m"

// Print the character to reset the bash text color
func Reset() {
	fmt.Print(colorReset)
}

// Print arguments as red and then resets the color.
// Wrapper around `fmt.Print`. Note: Doesn't add
// a newline character at the end.
func Red(a ...interface{}) {
	fmt.Print(colorRed)
	fmt.Print(a...)
	Reset()
}

// Print arguments as red and then resets the color.
// Wrapper around `fmt.Println`. Note: Adds a newline
// character at the end.
func Redln(a ...interface{}) {
	Red(a...)
	fmt.Println()
}

// Print arguments as green and then resets the color.
// Wrapper around `fmt.Print`. Note: Doesn't add
// a newline character at the end.
func Green(a ...interface{}) {
	fmt.Print(colorGreen)
	fmt.Print(a...)
	Reset()
}

// Print arguments as green and then resets the color.
// Wrapper around `fmt.Println`. Note: Adds a newline
// character at the end.
func Greenln(a ...interface{}) {
	Green(a...)
	fmt.Println()
}

// Print arguments as yellow and then resets the color.
// Wrapper around `fmt.Print`. Note: Doesn't add
// a newline character at the end.
func Yellow(a ...interface{}) {
	fmt.Print(colorYellow)
	fmt.Print(a...)
	Reset()
}

// Print arguments as yellow and then resets the color.
// Wrapper around `fmt.Println`. Note: Adds a newline
// character at the end.
func Yellowln(a ...interface{}) {
	Yellow(a...)
	fmt.Println()
}

// Print arguments as blue and then resets the color.
// Wrapper around `fmt.Print`. Note: Doesn't add
// a newline character at the end.
func Blue(a ...interface{}) {
	fmt.Print(colorBlue)
	fmt.Print(a...)
	Reset()
}

// Print arguments as blue and then resets the color.
// Wrapper around `fmt.Println`. Note: Adds a newline
// character at the end.
func Blueln(a ...interface{}) {
	Blue(a...)
	fmt.Println()
}

// Print arguments as purple and then resets the color.
// Wrapper around `fmt.Print`. Note: Doesn't add
// a newline character at the end.
func Purple(a ...interface{}) {
	fmt.Print(colorPurple)
	fmt.Print(a...)
	Reset()
}

// Print arguments as purple and then resets the color.
// Wrapper around `fmt.Println`. Note: Adds a newline
// character at the end.
func Purpleln(a ...interface{}) {
	Red(a...)
	fmt.Println()
}

// Print arguments as cyan and then resets the color.
// Wrapper around `fmt.Print`. Note: Doesn't add
// a newline character at the end.
func Cyan(a ...interface{}) {
	fmt.Print(colorCyan)
	fmt.Print(a...)
	Reset()
}

// Print arguments as cyan and then resets the color.
// Wrapper around `fmt.Println`. Note: Adds a newline
// character at the end.
func Cyanln(a ...interface{}) {
	Cyan(a...)
	fmt.Println()
}

// Print arguments as gray and then resets the color.
// Wrapper around `fmt.Print`. Note: Doesn't add
// a newline character at the end.
// Also note: Has same effect as the function `Grey`
func Gray(a ...interface{}) {
	fmt.Print(colorGray)
	fmt.Print(a...)
	Reset()
}

// Print arguments as gray and then resets the color.
// Wrapper around `fmt.Println`. Note: Adds a newline
// character at the end.
// Also note: Has same effect as the function `Greyln`
func Grayln(a ...interface{}) {
	Gray(a...)
	fmt.Println()
}

// Print arguments as grey and then resets the color.
// Wrapper around `fmt.Print`. Note: Doesn't add
// a newline character at the end.
// Also note: Has same effect as the function `Gray`
func Grey(a ...interface{}) {
	fmt.Print(colorGrey)
	fmt.Print(a...)
	Reset()
}

// Print arguments as grey and then resets the color.
// Wrapper around `fmt.Println`. Note: Adds a newline
// character at the end.
// Also note: Has same effect as the function `Grayln`
func Greyln(a ...interface{}) {
	Grey(a...)
	fmt.Println()
}

// Print arguments as white and then resets the color.
// Wrapper around `fmt.Print`. Note: Doesn't add
// a newline character at the end.
func White(a ...interface{}) {
	fmt.Print(colorWhite)
	fmt.Print(a...)
	Reset()
}

// Print arguments as white and then resets the color.
// Wrapper around `fmt.Println`. Note: Adds a newline
// character at the end.
func Whiteln(a ...interface{}) {
	White(a...)
	fmt.Println()
}

