package terr

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

type terr struct {
	tracks []string
}

func Of(originError error) *terr {
	e, ok := originError.(*terr)
	if !ok && originError == nil {
		e = &terr{tracks: []string{"nil"}}
	} else if !ok {
		e = &terr{tracks: []string{originError.Error()}}
	}
	return e
}

func getLastSegment(text string, delimiter string) string {
	i := strings.LastIndex(text, delimiter)
	return text[i+1:]
}

func Track(originError error, msg ...string) *terr {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		log.Fatalf("cannot get the caller to create the gauge from %v", originError)
	}
	e, ok := originError.(*terr)
	if !ok && originError == nil {
		e = &terr{}
		msg = append([]string{"nil"}, msg...)
	} else if !ok {
		e = &terr{}
		msg = append([]string{originError.Error()}, msg...)
	}
	return e.add(file, line, msg...)
}

func (e *terr) add(file string, line int, msg ...string) *terr {
	gauge := fmt.Sprintf("%s:%d: %s", getLastSegment(file, "/"), line, strings.Join(msg, " "))
	e.tracks = append(e.tracks, gauge)
	return e
}

func (e *terr) Error() string {
	return strings.Join(e.tracks, "\n\t")
}
