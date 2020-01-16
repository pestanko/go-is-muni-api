package ismuniapi

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"google.golang.org/appengine/log"
	"fmt"
	"errors"
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

func (c *IsApiClient) GetCourseInfo() CourseInfoRequestBuilder {

}

func (c *IsApiClient) GetCourseStudents() CourseStudents {

}

func (c *IsApiClient) GetSeminaryStudents(seminary string) SeminaryStudents {

}

func (c *IsApiClient) GetSeminaryStudents(seminaries []string) SeminaryStudents {

}

func (c *IsApiClient) GetSeminaryTeachers(seminaries []String) SeminaryTeachers {

}

func (c *IsApiClient) GetExams() Exams {

}

func (c *IsApiClient) GetNotepadList() NotepadList {

}

func (c *IsApiClient) CreateNotepad(shortcut string, name string) {

}

func (c *IsApiClient) GetNotepadContent(shortcut string, ucos []int) NotepadContent {

}

func (c *IsApiClient) WriteNotepadContent(shortcut string, uco int, content string) {

}

func (c *IsApiClient) prepareRequest() (*preparedRequest, error) {
	req, err := http.NewRequest("GET", c.URL+"/export/pb_blok_api", nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("klic", c.token)
	q.Add("fakulta", strconv.Itoa(c.FacultyID))
	q.Add("kod", c.CourseCode)

	prep := &preparedRequest{req: req, q: q}
	return prep, nil
}

type preparedRequest struct {
	req *http.Request
	q   url.Values
}

func (r *preparedRequest) Request() (*http.Request, error) {
	r.req.URL.RawQuery = r.q.Encode()
	return r.req, nil
}

type RequestBuilder struct {
	client      *IsApiClient
	prepRequest preparedRequest
}

func (rqb *RequestBuilder) Raw() (*http.Response, error) {
	req, err := rqb.prepRequest.Request()

	if err != nil {
		return nil, err
	}

	return rqb.client.Client.Do(req)
}

func (rqb *RequestBuilder) ToString() (string, error) {
	resp, err := rqb.Raw()

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		bodyString := string(bodyBytes)
		return bodyString, nil
	} else {
		return "", errors.New(fmt.Sprintf("Status code [{}]", resp.StatusCode))
	}
}

type CourseInfoRequestBuilder struct {
	rb *RequestBuilder
}
