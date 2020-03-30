# Go Is Muni API
Go wrapper for the IS MUNI API (https://is.muni.cz/)

Official API documentation can be found [here](https://is.muni.cz/napoveda/technicka/bloky_api).

## Install

To install the package use the `go get`

```bash
go get github.com/pestanko/goismuniapi
```

### Usage


```go
package main

import (
	"github.com/pestanko/goismuniapi"
)

func main() {
	client := goismuniapi.NewClient("https://is.muni.cz", "TOKEN", 1433, "example")

	info, err := client.GetCourseInfo().Unmarshal()

	courseStudents, err := client.GetCourseStudents().WithEnded().WithInactive().Unmarshal()

	seminaryStudents, err := client.GetSeminaryStudents("01", "02").Unmarshal()
	seminaryTeachers, err := client.GetSeminaryTeachers("01", "02").Unmarshal()

	exams, err := client.GetExams().WithInactive().Unmarshal()

	notepads, err := client.GetNotepadList().Unmarshal() // will return Notepads

	content, err := client.GetNotepadContent("hw01").WithUco(12356).WithUcos(12, 45, 5).Unmarshal()

	response, err := client.WriteNotepadContent("hw01", 123465, "Great work! *100").Execute()

	response, err := client.CreateNotepad("hw01", "Homework 01").WithStatistics(true).Execute()
}

```



