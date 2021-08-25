// Given a year, report if it is a leap year.
//
// In the Gregorian calendar, every year that is exactly divisible by 4
// is a leap year, except for years that are exactly divisible by 100
// unless they are also exactly divisible by 400.
//
// For example, the years 2000 and 2400 are leap years, while 1800,
// 1900, 2100, 2200, 2300 and 2500 are not.
package gocopilot

import "fmt"

func leapYear() {
	fmt.Println(isLeapYear(2000))
	fmt.Println(isLeapYear(2400))
	fmt.Println(isLeapYear(1800))
	fmt.Println(isLeapYear(1900))
	fmt.Println(isLeapYear(2100))
	fmt.Println(isLeapYear(2200))
	fmt.Println(isLeapYear(2300))
	fmt.Println(isLeapYear(2500))
}

func isLeapYear(year int) bool {
	if (year % 4 == 0) && (year % 100 != 0) || (year % 400 == 0) {
		return true
	}
	return false
}