package greetings

import (
	"regexp"
	"testing"
)

func TestHelloNmae(t *testing.T) {
	name := "Vlad"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Vlad")
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T){
	msg,err := Hello("")
	if msg != "" || err == nil {
        t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}
