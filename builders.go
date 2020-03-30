package ismuniapi

import (
	"net/http"
	"net/url"
)

type operationBuilder struct {
	preparedRequest preparedRequest
	operation       string
}

type CourseInfoBuilder struct {
	operationBuilder
}

type CourseStudentsBuilder struct {
	operationBuilder
}

type SeminaryTeachersBuilder struct {
	operationBuilder
}

type SeminaryStudentsBuilder struct {
	operationBuilder
}

type ExamsBuilder struct {
	operationBuilder
}

type NotepadContentBuilder struct {
	operationBuilder
}

type NotepadsBuilder struct {
	operationBuilder
}

type WriteNotepadBuilder struct {
	operationBuilder
}

type CreateNotepadBuilder struct {
	operationBuilder
}

func (builder *operationBuilder) Bytes() ([]byte, error) {
	return builder.preparedRequest.Bytes()
}

func (builder *operationBuilder) Execute() (*http.Response, error) {
	builder.Query().Add("operace", builder.operation)
	return builder.preparedRequest.Execute()
}

func (builder *operationBuilder) Query() *url.Values {
	return builder.preparedRequest.Query
}

func (builder *CourseInfoBuilder) Unmarshal() (entity CourseInfo, err error) {
	data, err := builder.Bytes()
	if err != nil {
		return entity, err
	}
	return UnmarshalCourseInfo(data)
}

func (builder *CourseStudentsBuilder) Unmarshal() (entity CourseStudentsInfo, err error) {
	data, err := builder.Bytes()
	if err != nil {
		return entity, err
	}
	return UnmarshalCourseStudents(data)
}

func (builder *SeminaryTeachersBuilder) Unmarshal() (entity SeminaryTeachersInfo, err error) {
	data, err := builder.Bytes()
	if err != nil {
		return entity, err
	}
	return UnmarshalSeminaryTeachers(data)
}

func (builder *SeminaryStudentsBuilder) Unmarshal() (entity SeminaryStudentsInfo, err error) {
	data, err := builder.Bytes()
	if err != nil {
		return entity, err
	}
	return UnmarshalSeminaryStudents(data)
}

func (builder *ExamsBuilder) Unmarshal() (entity Exams, err error) {
	data, err := builder.Bytes()
	if err != nil {
		return entity, err
	}
	return UnmarshalExams(data)
}

func (builder *NotepadContentBuilder) Unmarshal() (entity NotepadContent, err error) {
	data, err := builder.Bytes()
	if err != nil {
		return entity, err
	}
	return UnmarshalNotepadContent(data)
}

func (builder *NotepadsBuilder) Unmarshal() (entity Notepads, err error) {
	data, err := builder.Bytes()
	if err != nil {
		return entity, err
	}
	return UnmarshalNotepads(data)
}

// COURSE STUDENTS

func (builder *CourseStudentsBuilder) WithRegistered() *CourseStudentsBuilder {
	builder.Query().Set("zareg", "a")
	return builder
}

func (builder *CourseStudentsBuilder) WithEnded(b bool) *CourseStudentsBuilder {
	builder.Query().Set("vcukonc", "a")
	return builder
}
func (builder *CourseStudentsBuilder) WithInactive(b bool) *CourseStudentsBuilder {
	builder.Query().Set("vcneaktiv", "a")
	return builder
}

// SEMINAR STUDENTS

func (builder *SeminaryStudentsBuilder) WithEnded(b bool) *SeminaryStudentsBuilder {
	builder.Query().Set("vcukonc", "a")
	return builder
}
func (builder *SeminaryStudentsBuilder) WithInactive(b bool) *SeminaryStudentsBuilder {
	builder.Query().Set("vcneaktiv", "a")
	return builder
}

// Create notepad

func (builder *CreateNotepadBuilder) WithAvailableToStudents(b bool) *CreateNotepadBuilder {
	builder.Query().Set("nahlizi", toIsBool(b))
	return builder
}

func (builder *CreateNotepadBuilder) WithoutMissingStudies(b bool) *CreateNotepadBuilder {
	builder.Query().Set("nedoplnovat", toIsBool(b))
	return builder
}

func (builder *CreateNotepadBuilder) WithStatistics(b bool) *CreateNotepadBuilder {
	builder.Query().Set("statistika", toIsBool(b))
	return builder
}

// Write notepad

func (builder *WriteNotepadBuilder) WithOverride() *WriteNotepadBuilder {
	builder.Query().Set("prepis", toIsBool(true))
	return builder
}

// WithLastChange - provide the datetime in format: YYYYMMDDHH24MISS, example: 20160901181030
func (builder *WriteNotepadBuilder) WithLastChange(datetime string) *WriteNotepadBuilder {
	builder.Query().Set("poslzmeneno", datetime)
	return builder
}

// Exams

func (builder *ExamsBuilder) WithEnded(b bool) *ExamsBuilder {
	builder.Query().Set("vcukonc", "a")
	return builder
}

func (builder *ExamsBuilder) WithInactive(b bool) *ExamsBuilder {
	builder.Query().Set("vcneaktiv", "a")
	return builder
}

func toIsBool(b bool) string {
	if b {
		return "a"
	} else {
		return "n"
	}
}
