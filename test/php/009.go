package main

import (
	"fmt"
	"os/exec"
	"reflect"
	"strings"
	"time"
)

//数组-创建一个数组
func Array(v ...interface{}) []interface{} {
	return v
}

// InArray in_Array()
// haystack supported types: slice, Array or map
func InArray(needle interface{}, haystack interface{}) bool {
	val := reflect.ValueOf(haystack)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if reflect.DeepEqual(needle, val.Index(i).Interface()) {
				return true
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if reflect.DeepEqual(needle, val.MapIndex(k).Interface()) {
				return true
			}
		}
	default:
		panic("haystack: haystack type muset be slice, Array or map")
	}

	return false
}

var datePatterns = []string{
	// year
	"Y", "2006", // A full numeric representation of a year, 4 digits   Examples: 1999 or 2003
	"y", "06", // A two digit representation of a year   Examples: 99 or 03

	// month
	"m", "01", // Numeric representation of a month, with leading zeros 01 through 12
	"n", "1", // Numeric representation of a month, without leading zeros   1 through 12
	"M", "Jan", // A short textual representation of a month, three letters Jan through Dec
	"F", "January", // A full textual representation of a month, such as January or March   January through December

	// day
	"d", "02", // Day of the month, 2 digits with leading zeros 01 to 31
	"j", "2", // Day of the month without leading zeros 1 to 31

	// week
	"D", "Mon", // A textual representation of a day, three letters Mon through Sun
	"l", "Monday", // A full textual representation of the day of the week  Sunday through Saturday

	// time
	"g", "3", // 12-hour format of an hour without leading zeros    1 through 12
	"G", "15", // 24-hour format of an hour without leading zeros   0 through 23
	"h", "03", // 12-hour format of an hour with leading zeros  01 through 12
	"H", "15", // 24-hour format of an hour with leading zeros  00 through 23

	"a", "pm", // Lowercase Ante meridiem and Post meridiem am or pm
	"A", "PM", // Uppercase Ante meridiem and Post meridiem AM or PM

	"i", "04", // Minutes with leading zeros    00 to 59
	"s", "05", // Seconds, with leading zeros   00 through 59

	// time zone
	"T", "MST",
	"P", "-07:00",
	"O", "-0700",

	// RFC 2822
	"r", time.RFC1123Z,
}

// Date - Format a local time/date
func Date(format string, ts ...time.Time) string {
	replacer := strings.NewReplacer(datePatterns...)
	format = replacer.Replace(format)
	t := time.Now()
	if len(ts) > 0 {
		t = ts[0]
	}
	return t.Format(format)
}

// DateAdd-向DateTime对象添加天，月，年，小时，分钟和秒
func DateAdd(t time.Time, years int, months int, days int) time.Time {
	return t.AddDate(years, months, days)
}

// Substr substr()
func Substr(str string, start uint, length int) string {
	if start < 0 || length < -1 {
		return str
	}
	switch {
	case length == -1:
		return str[start:]
	case length == 0:
		return ""
	}
	end := int(start) + length
	if end > len(str) {
		end = len(str)
	}
	return str[start:end]
}

//explode() 函数把字符串打散为数组。
func Explode(delimiter, str string) []string {
	return strings.Split(str, delimiter)
}

// StrReplace str_replace()
func StrReplace(search, replace, subject string, count int) string {
	return strings.Replace(subject, search, replace, count)
}

// Exec - Execute an external program
func Exec(s string) {

	exec.Command(s).Run()
}

func main() {
	a := Array("d", "sdaf", "df")
	aa := [...]string{"a", "b", "c", "d"}
	people := Array(12, "Steve", "Mark", "David")
	/* 向切片添加一个元素 */
	aa = append(aa, "aa")
	fmt.Println(aa)
	if InArray(1, people) {
		fmt.Println("匹配已找到")
	} else {
		fmt.Println("匹配未找到")
	}

	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}
	fmt.Println(Date("Y"))
	fmt.Println(Substr("Hello world", 0, 1))
	//cmd := `echo aksdfjkasdf > D:\\a.txt`
	cmd := `dir`
	c := exec.Command("cmd.exe", `/c`+cmd)
	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(c.Output())
	cars := Array(Array("Volvo", 22, 18), Array("BMW", 15, 13), Array("Saab", 5, 2), Array("Land Rover", 17, 15))
	fmt.Println(cars[0])
}