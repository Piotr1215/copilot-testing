package gocopilot

// Convert digits (from 0 to 9) to text output given the pattern below. Each digit is represented by a grid of 4 rows and 3 columns with last row always empty.
//         " _ "
//         "| |"
//         "|_|"
//         "   " = 0
//
//         "   "
//         "  |"
//         "  |"
//         "   " = 1
//
//         " _ "
//         " _|"
//         "|_ "
//         "   " = 2
//
//         " _ "
//         " _|"
//         " _|"
//         "   " = 3
//
//         "   "
//         "|_|"
//         "  |"
//         "   " = 4
//
//         " _ "
//         "|_ "
//         " _|"
//         "   " = 5
//
//         " _ "
//         "|_ "
//         "|_|"
//         "   " = 6
//
//         " _ "
//         "  |"
//         "  |"
//         "   " = 7
//
//         " _ "
//         "|_|"
//         "|_|"
//         "   " = 8
//
//         " _ "
//         "|_|"
//         " _|"
//         "   " = 9
//
//         all else = an error
import (
	"fmt"
	"os"
)

func ocrRecognition() {
	if len(os.Args) != 2 {
		fmt.Println("Please enter a number")
		os.Exit(1)
	}
	number := os.Args[1]
	if number == "0" {
		fmt.Println(" _ ")
		fmt.Println("| |")
		fmt.Println("|_|")
		fmt.Println("   ")
	} else if number == "1" {
		fmt.Println("   ")
		fmt.Println("  |")
		fmt.Println("  |")
		fmt.Println("   ")
	} else if number == "2" {
		fmt.Println(" _ ")
		fmt.Println(" _|")
		fmt.Println("|_ ")
		fmt.Println("   ")
	} else if number == "3" {
		fmt.Println(" _ ")
		fmt.Println(" _|")
		fmt.Println(" _|")
		fmt.Println("   ")
	} else if number == "4" {
		fmt.Println("   ")
		fmt.Println("|_|")
		fmt.Println("  |")
		fmt.Println("   ")
	} else if number == "5" {
		fmt.Println(" _ ")
		fmt.Println("|_ ")
		fmt.Println(" _|")
		fmt.Println("   ")
	} else if number == "6" {
		fmt.Println(" _ ")
		fmt.Println("|_ ")
		fmt.Println("|_|")
		fmt.Println("   ")
	} else if number == "7" {
		fmt.Println(" _ ")
		fmt.Println("  |")
		fmt.Println("  |")
		fmt.Println("   ")
	} else if number == "8" {
		fmt.Println(" _ ")
		fmt.Println("|_|")
		fmt.Println("|_|")
		fmt.Println("   ")
	} else if number == "9" {
		fmt.Println(" _ ")
		fmt.Println("|_|")
		fmt.Println(" _|")
		fmt.Println("   ")
	} else {
		fmt.Println("Invalid number")
		os.Exit(1)
	}
}