package builders

import ismuniapi "github.com/pestanko/go-is-muni-api"

type CourseInfoRequestBuilder struct {
	base *RequestBuilder
}

func (rb *CourseInfoRequestBuilder) parse() (*ismuniapi.CourseInfo, error) {
	courseInfo := &ismuniapi.CourseInfo{}

	if err := rb.base.ParseInto(courseInfo); err != nil {
		return nil, err
	}

	return courseInfo, nil
}


type CourseStudentsBuilder struct {
	base *RequestBuilder
}

func (rb *CourseStudentsBuilder) registered() *CourseStudentsBuilder {
	rb.base.AddQuery("")
	return rb
}

func (rb *CourseStudentsBuilder) parse() (*ismuniapi.CourseStudents, error) {
	courseStudents := &ismuniapi.CourseStudents{}

	if err := rb.base.ParseInto(courseStudents); err != nil {
		return nil, err
	}

	return courseStudents, nil
}
