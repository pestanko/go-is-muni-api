package goismuniapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type IsApiClient struct {
	token      string
	FacultyID  int
	URL        string
	CourseCode string
	client     *http.Client
}

func NewClient(url string, token string, facultyId int, courseCode string) IsApiClient {
	return IsApiClient{URL: url, token: token, FacultyID: facultyId, CourseCode: courseCode, client: http.DefaultClient}
}

func (c *IsApiClient) GetCourseInfo() CourseInfoBuilder {
	builder := CourseInfoBuilder{c.newOperationBuilder("predmet-info")}
	return builder
}

func (c *IsApiClient) GetCourseStudents() CourseStudentsBuilder {
	builder := CourseStudentsBuilder{c.newOperationBuilder("predmet-seznam")}
	return builder
}

func (c *IsApiClient) GetSeminaryStudents(seminaries ...string) SeminaryStudentsBuilder {
	builder := SeminaryStudentsBuilder{c.newOperationBuilder("seminar-seznam")}
	for _, seminar := range seminaries {
		builder.Query().Add("seminar", seminar)
	}
	return builder
}

func (c *IsApiClient) GetSeminaryTeachers(seminaries ...string) SeminaryTeachersBuilder {
	builder := SeminaryTeachersBuilder{c.newOperationBuilder("seminar-cvicici-seznam")}
	for _, seminar := range seminaries {
		builder.Query().Add("seminar", seminar)
	}
	return builder

}

func (c *IsApiClient) GetExams() ExamsBuilder {
	return ExamsBuilder{c.newOperationBuilder("terminy-seznam")}

}

func (c *IsApiClient) GetNotepadList() NotepadsBuilder {
	return NotepadsBuilder{c.newOperationBuilder("bloky-seznam")}

}

func (c *IsApiClient) CreateNotepad(shortcut string, name string) CreateNotepadBuilder {
	builder := CreateNotepadBuilder{c.newOperationBuilder("blok-novy")}
	builder.Query().Set("zkratka", shortcut)
	builder.Query().Set("jmeno", name)
	return builder

}

func (c *IsApiClient) GetNotepadContent(shortcut string) NotepadContentBuilder {
	builder := NotepadContentBuilder{c.newOperationBuilder("blok-dej-obsah")}
	builder.Query().Set("zkratka", shortcut)
	return builder

}

func (c *IsApiClient) WriteNotepadContent(shortcut string, uco int, content string) WriteNotepadBuilder {
	builder := WriteNotepadBuilder{c.newOperationBuilder("blok-pis-student-obsah")}
	builder.Query().Add("uco", strconv.Itoa(uco))
	builder.Query().Set("zkratka", shortcut)
	builder.Query().Set("obsah", content)
	return builder

}

func (c *IsApiClient) newPreparedRequest() preparedRequest {
	values := &url.Values{}
	values.Add("klic", c.token)
	values.Add("fakulta", strconv.Itoa(c.FacultyID))
	values.Add("kod", c.CourseCode)
	return preparedRequest{URL: c.notesUrl(), httpClient: c.client, Query: values}
}

func (c *IsApiClient) newOperationBuilder(operation string) operationBuilder {
	return operationBuilder{c.newPreparedRequest(), operation}
}

func (c *IsApiClient) notesUrl() string {
	return c.URL + "/export/pb_blok_api"
}

type PreparedRequest struct {
	Req *http.Request
	Q   url.Values
}

func (r *PreparedRequest) Request() (*http.Request, error) {
	r.Req.URL.RawQuery = r.Q.Encode()
	return r.Req, nil
}

type preparedRequest struct {
	httpClient *http.Client
	URL        string
	Query      *url.Values
}

// Fetch - fetches XML data
func (request *preparedRequest) Bytes() ([]byte, error) {
	resp, err := request.Execute()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("Status error: %d", resp.StatusCode)
		return nil, fmt.Errorf(msg)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body: %v", err)
	}
	return data, nil
}

func (request *preparedRequest) Execute() (*http.Response, error) {
	req, err := request.request()

	if err != nil {
		return nil, err
	}
	resp, err := request.httpClient.Do(req)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (request *preparedRequest) request() (*http.Request, error) {
	builtUrl, err := url.Parse(request.URL)
	if err != nil {
		return nil, err
	}
	builtUrl.RawQuery = request.Query.Encode()
	return http.NewRequest("GET", builtUrl.String(), nil)
}
