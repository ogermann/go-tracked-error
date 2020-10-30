package terr

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
	"testing"
)

func TestEErr_Add_trackCount(t *testing.T) {
	getError := func() error {
		return Track(errors.New("i should occur"))
	}

	err := Of(getError())

	expected := 1
	if len(err.tracks) != expected {
		t.Fatalf("should have %d gauges but had %d", expected, len(err.tracks))
	}
}

func TestEErr_Add_trackCount2(t *testing.T) {
	additionalMsg := "with some additions"
	getError := func() error {
		return Track(errors.New("i should occur"), additionalMsg)
	}

	err := Of(getError())

	expected := 1
	if len(err.tracks) != expected {
		t.Fatalf("should have %d gauges but had %d", expected, len(err.tracks))
	}

	if !strings.HasSuffix(err.Error(), additionalMsg)  {
		t.Fatalf("should end with %s but was %s", additionalMsg, err.Error())
	}
}

func TestEErr_Add_trackText(t *testing.T) {
	text := "i should occur"

	getError := func() error {
		return Track(errors.New(text))
	}

	_, _, line, _ := runtime.Caller(0)

	err := Of(getError())

	expected := fmt.Sprintf("main_test.go:%d: %s", line-3, text)

	if err.Error() != expected {
		t.Fatalf("text should be \"%s\" but was \"%s\"", expected, err.Error())
	}
}

func TestEErr_Error(t *testing.T) {
	text := "i should occur"
	getError := func() error {
		return Track(errors.New(text))
	}
	_, _, line, _ := runtime.Caller(0)

	text2 := "i should occur too, but i'm different"
	getError2 := func() error {
		err := getError()
		return Track(err, text2)
	}
	_, _, line2, _ := runtime.Caller(0)

	err := Of(getError2())

	expected := fmt.Sprintf("main_test.go:%d: %s\n\tmain_test.go:%d: %s", line-2, text, line2-2, text2)

	if err.Error() != expected {
		t.Fatalf("text should be \"%s\" but was \"%s\"", expected, err.Error())
	}
}

func TestEErr_Error2(t *testing.T) {
	text := "i should occur"
	getError := func() error {
		return errors.New(text)
	}

	text2 := "i should occur too, but i'm different"
	getError2 := func() error {
		err := getError()
		return Track(err, text2)
	}
	_, _, line, _ := runtime.Caller(0)

	err := Of(getError2())

	expected := fmt.Sprintf("main_test.go:%d: %s %s", line-2, text, text2)

	if err.Error() != expected {
		t.Fatalf("text should be \"%s\" but was \"%s\"", expected, err.Error())
	}
}


func TestEErr_Error3(t *testing.T) {
	var err error
	err = Track(errors.New("any error"))
	err = Track(err, "add some additional", "text")

	t.Error(err)
}