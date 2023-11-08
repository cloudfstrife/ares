package utils

import (
	"io"
	"os"
	"testing"
)

var testCase = []struct {
	w       io.Writer
	color   string
	content string
}{
	{
		w:       os.Stdout,
		color:   RED,
		content: "golang",
	},
	{
		w:       os.Stdout,
		color:   GREEN,
		content: "golang",
	},
}

func TestFprintWithColor(t *testing.T) {
	for _, tcase := range testCase {
		FprintWithColor(tcase.w, tcase.color, tcase.content)
	}
}
func TestPrintWithColor(t *testing.T) {
	for _, tcase := range testCase {
		PrintWithColor(tcase.color, tcase.content)
	}
}
func TestPrintlnWithColor(t *testing.T) {
	for _, tcase := range testCase {
		PrintlnWithColor(tcase.color, tcase.content)
	}
}
