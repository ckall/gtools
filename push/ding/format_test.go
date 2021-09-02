package ding_test

import (
	"github.com/ckall/gtools/push/ding"
	"testing"
)

func TestFormat_AddImage(t *testing.T) {
	ding.NewConText().AddImage("https://gw.alicdn.com/tfs/TB1ut3xxbsrBKNjSZFpXXcXhFXa-846-786.png")
}

func TestAddBlue(t *testing.T) {
	ding.AddBlue("enen")
}

func TestAddGreen(t *testing.T) {
	ding.AddGreen("enen")
}
func TestAddH1(t *testing.T) {
	ding.AddH1("enen")
}
