package cert

import (
	"testing"
)

func TestValidCert(t *testing.T) {
	cert, err := New("Golang", "Yann", "2024-06-08")

	if err != nil {
		t.Errorf("Cert data should be valid err=%v",err)
	}

	if cert == nil {
		t.Errorf("cert should be a valid cert=%v", cert)
	}

	if cert.Course != "GOLANG COURSE" {
		t.Errorf("Course name is invalid")
	}
}

func TestCourseEmpty(t *testing.T) {
	_, err := New("", "Yann", "2024-06-08")

	if err == nil {
		t.Errorf("course should be name err=%v", err)
	}
}

func TestLongerOfCourse(t *testing.T) {
     course := "azertyuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuu"
	 _, err := New(course, "Yann", "2024-06-08")

	 if err == nil {
		t.Errorf("course name should short")
	 }
}

func TestValidateName(t *testing.T) {

	cert, _ := New("Golang", "yann", "2012-06-09")

	if cert.Name != "YANN" {
       t.Errorf("cert name should be to upper case")
	}
}

func TestEmptyName(t *testing.T) {
	_, err := New("Go", "", "2012-06-09")

	if err == nil {
		t.Errorf("name field should be not empty")
	}
}

func TestLengthOfName(t *testing.T) {
	name := "eaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	_, err := New("Go", name, "2012-06-09")
	
	if err == nil {
		t.Errorf("length of name should be less than 20")
	}
}

func TestParseDate(t *testing.T) {
	_, err := parseDate("2024-01-20")
	if err != nil {
		t.Errorf("could have problem with parsing date, err=%v", err)
	}
}

func TestEmptyDate(t *testing.T) {
	_, err := parseDate("")

	if err == nil {
		t.Errorf("Should be have error date parsing")
	}
}