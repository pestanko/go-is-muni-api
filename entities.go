package ismuniapi

import "encoding/xml"

type CourseInfo struct {
	FacultyID                  int             `xml:"FAKULTA_ID"`
	FacultyShortcut            string          `xml:"FAKULTA_ZKRATKA_DOM"`
	CourseID                   int             `xml:"PREDMET_ID"`
	CourseCode                 string          `xml:"KOD_PREDMETU"`
	CourseName                 string          `xml:"NAZEV_PREDMETU"`
	CourseNameShort            string          `xml:"KRATKY_NAZEV_PREDMETU"`
	CourseNameEng              string          `xml:"NAZEV_PREDMETU_ANGL"`
	CourseNameShortEng         string          `xml:"KRATKY_NAZEV_PREDMETU_ANGL"`
	PeriodID                   int             `xml:"OBDOBI_ID"`
	PeriodName                 string          `xml:"OBDOBI_NAZEV"`
	PeriodNameEng              string          `xml:"OBDOBI_NAZEV_ANGL"`
	PeriodUrlCode              string          `xml:"OBDOBI_ZKRATKA_PRO_URL"`
	PeriodUrlCodeEng           string          `xml:"OBDOBI_ZKRATKA_PRO_URL_ANGL"`
	NumberOfStudents           int             `xml:"POCET_ZAPSANYCH_STUDENTU"`
	NumberOfRegisteredStudents int             `xml:"POCET_ZAREG_STUDENTU"`
	NotListed                  int             `xml:"NEVYPISUJE_SE"`
	CourseSeminars             []CourseSeminar `xml:"SEMINARE>SEMINAR"`
	CourseTeachers             []CourseTeacher `xml:"VYUCUJICI_SEZNAM>VYUCUJICI"`
}

type CourseSeminar struct {
	LimitID            int    `xml:"LIMIT_ID"`
	ShareableLimit     string `xml:"LIMIT_LZE_SDILET"` // bool
	StudentsLimit      int    `xml:"MAX_STUDENTU"`
	MultipleEnrollment string `xml:"NASOBNE_PRIHLASENI"`
	Name               string `xml:"OZNACENI"`
	ID                 int    `xml:"SEMINAR_ID"`
	NumberOfStudents   int    `xml:"POCET_STUDENTU_VE_SKUPINE"`
	Note               string `xml:"POZNAMKA"`
	Order              string `xml:"PORADI"`             // No idea what it is
	EnrollTo           string `xml:"PRIHLASIT_DO"`       // datetime
	EnrollFrom         string `xml:"PRIHLASIT_OD"`       // datetime
	LeaveTo            string `xml:"ODHLASIT_DO"`        // datetime
	WithInactive       string `xml:"VCETNE_NEAKTIVNICH"` // bool
	ChangedAt          string `xml:"ZMENENO"`            // datetime
	ChangedBy          string `xml:"ZMENIL"`             // Datetime
}

type CourseTeacher struct {
	Name       string `xml:"CELE_JMENO"`
	FirstName  string `xml:"JMENO"`
	LastName   string `xml:"PRIJMENI"`
	Role       string `xml:"ROLE"`
	UCO        string `xml:"UCO"`
	PersonText string `xml:"OSOBA_TEXTOVE"`
	// TODO: Zastupce
}

type CourseStudentsInfo struct {
	Students []CourseStudent `xml:"STUDENT"`
}

type CourseStudent struct {
	Name                   string   `xml:"CELE_JMENO"`
	FirstName              string   `xml:"JMENO"`
	LastName               string   `xml:"PRIJMENI"`
	UCO                    string   `xml:"UCO"`
	StudiumState           string   `xml:"STAV_STUDIA"`
	EnrollmentState        string   `xml:"STAV_ZAPISU"`
	StudentWithoutSeminary string   `xml:"STUDENT_NEMA_SEMINAR"` // bool (represented by int)
	FieldsOfStudy          []string `xml:"STUDIA>STUDIUM_IDENTIFIKACE"`
	CourseCompletion       string   `xml:"UKONCENI"`
}

type SeminaryStudentsInfo struct {
	Name     string            `xml:"OZNACENI"`
	ID       int               `xml:"SEMINAR_ID"`
	Students []SeminaryStudent `xml:"STUDENT"`
}

type SeminaryStudent struct {
	Name             string   `xml:"CELE_JMENO"`
	FirstName        string   `xml:"JMENO"`
	LastName         string   `xml:"PRIJMENI"`
	UCO              string   `xml:"UCO"`
	StudiumState     string   `xml:"STAV_STUDIA"`
	FieldsOfStudy    []string `xml:"STUDIA>STUDIUM_IDENTIFIKACE"`
	CourseCompletion string   `xml:"UKONCENI"`
}

type SeminaryTeachersInfo struct {
	Name     string            `xml:"OZNACENI"`
	ID       int               `xml:"SEMINAR_ID"`
	Teachers []SeminaryTeacher `xml:"CVICICI"`
}

type SeminaryTeacher struct {
	Name      string `xml:"CELE_JMENO"`
	FirstName string `xml:"JMENO"`
	LastName  string `xml:"PRIJMENI"`
	UCO       string `xml:"UCO"`
}

type Termin struct {
	ID           int    `xml:"ID"`
	CAPACITY     int    `xml:"KAPACITA"`
	DateTime     string `xml:"KONANI"`
	Reevaluation string `xml:"OPRAVNY"`
	Standard     string `xml:"RADNY"`
	Room         string `xml:"UCEBNA"`
	// TODO - there are some issues
}

type Exams struct {
	Name    string   `xml:"NAZEV"`
	ID      int      `xml:"ID"`
	Termins []Termin `xml:"TERMIN"`
}

type Notepads struct {
	Notepads []Notepad `xml:"POZN_BLOK"`
}

type Notepad struct {
	ID                        int    `xml:"BLOK_ID"`
	Name                      string `xml:"JMENO"`
	CodeName                  string `xml:"ZKRATKA"`
	TypeID                    int    `xml:"TYP_ID"`
	TypeName                  string `xml:"TYP_NAZEV"`
	ChangedAt                 string `xml:"ZMENENO"`                        // datetime
	ChangedBy                 string `xml:"ZMENIL"`                         // Datetime
	NotCompleteMissingStudies string `xml:"NEDOPLNOVAT_CHYBEJICI_STUDIA"`   // bool (a/n)
	StatisticsForStudent      string `xml:"STUDENTOVI_ZOBRAZIT_STATISTIKU"` // bool (a/n)
	AvailableForStudent       string `xml:"STUDENT_SMI_NAHLIZET"`           // bool (a/n)
}

type NotepadContent struct {
	StudentsContent []StudentContent `xml:"STUDENT" json:"students"`
}

type StudentContent struct {
	Content   string `xml:"OBSAH" json:"content"`
	Uco       string `xml:"UCO" json:"uco"`
	ChangedBy string `xml:"ZMENIL" json:"changed_by"`
	ChangedAt string `xml:"ZMENENO" json:"changed_at"`
}

func UnmarshalCourseInfo(content []byte) (courseInfo CourseInfo, err error) {
	if err := xml.Unmarshal(content, &courseInfo); err != nil {
		return courseInfo, err
	}

	return courseInfo, nil
}

func UnmarshalCourseStudents(content []byte) (data CourseStudentsInfo, err error) {
	if err := xml.Unmarshal(content, &data); err != nil {
		return data, err
	}

	return data, nil
}

func UnmarshalSeminaryStudents(content []byte) (data SeminaryStudentsInfo, err error) {
	if err := xml.Unmarshal(content, &data); err != nil {
		return data, err
	}

	return data, nil
}

func UnmarshalSeminaryTeachers(content []byte) (data SeminaryTeachersInfo, err error) {
	if err := xml.Unmarshal(content, &data); err != nil {
		return data, err
	}

	return data, nil
}

func UnmarshalNotepads(content []byte) (data Notepads, err error) {
	if err := xml.Unmarshal(content, &data); err != nil {
		return data, err
	}

	return data, nil
}

// UnmarshalNotepadContent - unmarshal the notepad content
func UnmarshalNotepadContent(data []byte) (content NotepadContent, err error) {
	if err := xml.Unmarshal(data, &content); err != nil {
		return content, err
	}

	return content, nil
}

func UnmarshalExams(data []byte) (content Exams, err error) {
	if err := xml.Unmarshal(data, &content); err != nil {
		return content, err
	}

	return content, nil
}

