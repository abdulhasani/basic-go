package main

import (
	"fmt"
	"time"
)

func main() {
	//waktu sekarnag
	var time1 = time.Now()
	fmt.Printf("time1 %+v\n", time1)
	//format manual time
	var time2 = time.Date(2011, 12, 24, 10, 20, 0, 0, time.Local)
	fmt.Printf("time2 %v\n", time2)
	//parsing time
	var layoutFormat, value string
	var date time.Time
	layoutFormat = "2006-01-02 15:04:05" //unix format in go-lang
	value = "2017-09-02 08:04:00"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t->", date.String())
	layoutFormat = "01-02-2006"
	value = "04-08-2017"
	fmt.Println(value, "\t->", date.String())

	var cal, _ = time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")
	fmt.Println(cal.String())

	//format time.Time
	var tanggal, _ = time.Parse(time.RFC822, "08 Oct 94 08:00 WIB")
	var dateS1 = tanggal.Format("02 January 2006 15:04 MST")
	fmt.Println(dateS1)
}
