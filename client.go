package ismuniapi

import (
	"github.com/pestanko/go-is-muni-api/builders"
	"net/http"
	"net/url"
	"strconv"
)

type IsApiClient struct {
	token      string
	FacultyID  int
	URL        string
	CourseCode string
	Client     *http.Client
}

func NewClient(url string, token string, facultyId int, courseCode string) IsApiClient {
	return IsApiClient{URL: url, token: token, FacultyID: facultyId, CourseCode: courseCode, Client: &http.Client{}}
}

func (c *IsApiClient) GetCourseInfo() builders.CourseInfoRequestBuilder {

}

func (c *IsApiClient) GetCourseStudents() builders.CourseStudentsBuilder  {

}

func (c *IsApiClient) GetSeminaryStudents(seminaries []string) builders.SeminaryStudentsBuilder {

}

func (c *IsApiClient) GetSeminaryTeachers(seminaries []String) builders.SeminaryTeachersBuilder {

}

func (c *IsApiClient) GetExams() builders.ExamsBuilders {

}

func (c *IsApiClient) GetNotepadList() builders.NotepadListBuilder {

}

func (c *IsApiClient) CreateNotepad(shortcut string, name string) builders.CreateNotepadBuilder {

}

func (c *IsApiClient) GetNotepadContent(shortcut string, ucos []int) builders.NotepadContentBuilder {

}

func (c *IsApiClient) WriteNotepadContent(shortcut string, uco int, content string) builders.WriteNotepadBuilder {

}

func (c *IsApiClient) prepareRequest() (*builders.RequestBuilder, error) {
	req, err := http.NewRequest("GET", c.URL+"/export/pb_blok_api", nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("klic", c.token)
	q.Add("fakulta", strconv.Itoa(c.FacultyID))
	q.Add("kod", c.CourseCode)

	prep := &builders.RequestBuilder{Req: req, Q: q}
	return prep, nil
}

type PreparedRequest struct {
	Req *http.Request
	Q   url.Values
}



func (r *PreparedRequest) Request() (*http.Request, error) {
	r.Req.URL.RawQuery = r.Q.Encode()
	return r.Req, nil
}
