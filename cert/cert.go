package cert

import (
	"fmt"
	"strings"
	"time"
)

const MAX_LENGTH_COURSE = 20

type Cert struct {
	Course string
	Name string
	Date time.Time

	LabelTitle string
	LabelCompletion string
	LabelPresented string
	LabelParticipation string
	LabelDate string
}

type Saver interface {
	Save(c *Cert) error
}

func New(course, name, date string) (*Cert, error) {
	c, err := validateCourse(course)
     
	if err != nil {
		return nil, err
	}
	n, err := validateName(name)

	if err != nil {
		return nil, err
	}
	d, err := parseDate(date)

	if err != nil {
		return nil, err
	}

	cert := &Cert{
		Course: c,
		Name: n,
	    
		LabelTitle: fmt.Sprintf("%v -Certifaction- %v", c, n),
		LabelCompletion: "Certifaction of Completion",
		LabelPresented: "This Certificate is prenseted to",
		LabelParticipation: fmt.Sprintf("For participation in the %v",c),
		LabelDate: fmt.Sprintf("Date: %v", d.Format("02/01/2006")),
	}

	return cert, nil
}

func validateCourse(c string) (string, error) {
     course, err := validateStr(c) 

	 if err != nil {
		return "", err
	 }

	if !strings.HasSuffix(course, "course") {
		course = course + " course"
	}

	return strings.ToTitle(course), nil
}

func validateStr(str string) (string, error) {
	strNoSpace := strings.TrimSpace(str)


	if len(strNoSpace) <= 0 {
		return "", fmt.Errorf("length of string is zero")
	}

	if len(strNoSpace) > MAX_LENGTH_COURSE {
		return "", fmt.Errorf("course is too long")
	}

	return str, nil
}

func validateName(n string) (string, error) {

	 name, err := validateStr(n)

	 if err != nil {
		return "", err
	 }

	if name != strings.ToTitle(name) {
		name = strings.ToTitle(name)
	}

	return name, nil
}

func parseDate(date string) (time.Time, error) {
    t, err := time.Parse("2006-01-02", date)

	if err != nil {
		return t, fmt.Errorf("could have an error date parsing")
	}

	return t, nil
}