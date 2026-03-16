package gocvKit

import (
	"fmt"
	"testing"

	"gocv.io/x/gocv"
)

func TestMatchTemplate(t *testing.T) {
	srcPath := "_source.png"
	templatePath := "_template.png"

	matchVal, matchRect, err := MatchTemplate(srcPath, templatePath, gocv.TmCcoeffNormed, false)
	if err != nil {
		panic(err)
	}
	fmt.Println("matchVal:", matchVal, "matchRect:", matchRect)
}
