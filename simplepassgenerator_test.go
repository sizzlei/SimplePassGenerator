package simplepassgenerator_test

import (
	gen "github.com/sizzlei/SimplePassGenerator"
	"testing"
	"fmt"
)

func TestNewGenertor(t *testing.T) {
	genertor := gen.New(
		gen.PassInOut{
			Upper:		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Lower:		"abcdefghijklmnopqrstuvwxyz",
			Digits:		"0123456789",
			Special:	"!#$^&()><~",
			SpecialSet:	true,
			DigitsSet: 	true,
		},
	)

	pass, err := genertor.GeneratePass(12)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Custom: ",*pass)
}


func TestDefaultGenertor(t *testing.T) {
	genertor := gen.New(gen.PassInOut{})

	pass, err := genertor.GeneratePass(12)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Default: ",*pass)
}

func TestNotDigitsGenertor(t *testing.T) {
	genertor := gen.New(gen.PassInOut{SpecialSet: true})

	pass, err := genertor.GeneratePass(12)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("NotDigits: ",*pass)
}
func TestNotSpecialGenertor(t *testing.T) {
	genertor := gen.New(gen.PassInOut{DigitsSet: true})

	pass, err := genertor.GeneratePass(12)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("NotSpecial: ",*pass)
}