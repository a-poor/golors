package golors

import (
	"fmt"
	"testing"
)

func TestPlaceholder(t *testing.T) {
	if false {
		t.Error("Hello world!")
	}
}

func TestPrint(t *testing.T) {
	fmt.Println("The following should all be on the same line...")
	Red("red")
	Green("green")
	Yellow("yellow")
	Blue("blue")
	Purple("purple")
	Cyan("cyan")
	Gray("gray")
	Grey("grey")
	White("white")
	fmt.Println("\nDone.")
}

func TestAllColors(t *testing.T) {
	fmt.Println("The following should all be on different lines...")
	Redln("This should be red")
	Greenln("This should be green")
	Yellowln("This should be yellow")
	Blueln("This should be blue")
	Purpleln("This should be purple")
	Cyanln("This should be cyan")
	Grayln("This should be gray")
	Greyln("This should be grey")
	Whiteln("This should be white")
	fmt.Println("\nDone.")
}
