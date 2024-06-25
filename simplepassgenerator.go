package simplepassgenerator

import (
	"crypto/rand"
	"math/big"
)

type PassInOut struct {
	Upper		string 
	Lower 		string 
	Digits 		string 
	Special		string
	All 		string
	DigitsSet 	bool
	SpecialSet 	bool
}

func New(p PassInOut) *PassInOut {
	var in PassInOut 

	if p.Upper == "" {
		in.Upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	} else {
		in.Upper = p.Upper
	}

	if p.Lower == "" {
		in.Lower = "abcdefghijklmnopqrstuvwxyz"
	} else {
		in.Lower = p.Lower
	}
	if p.DigitsSet == true {
		if p.Digits == "" {
			in.Digits = "0123456789"
		} else {
			in.Digits = p.Digits
		}
	}

	if p.SpecialSet == true {
		if p.Special == "" {
			in.Special = "!@#$%^&*()-_=+[]{}|;:,.<>?/~`"
		} else {
			in.Special = p.Special
		}
	}

	in.All = in.Upper + in.Lower + in.Digits + in.Special

	return &in
}

func randChar(c string) (byte, error) {
	if len(c) == 0 {
		return 0, nil
	}
    m := big.NewInt(int64(len(c)))
    n, err := rand.Int(rand.Reader, m)
    if err != nil {
        return 0, err
    }
    return c[n.Int64()], nil
}

func shuffle(slice []byte) {
    for i := range slice {
        j, _ := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
        slice[i], slice[j.Int64()] = slice[j.Int64()], slice[i]
    }
}

func (p *PassInOut) GeneratePass(l int) (*string, error) {
	if l < 3 {
		l = 3
	}
	var outPass []byte 
	var prefix []byte

	upper, err := randChar(p.Upper)
	if err != nil {
		return nil, err
	}
	lower, err := randChar(p.Lower)
	if err != nil {
		return nil, err
	}
	digits, err := randChar(p.Digits)
	if err != nil {
		return nil, err
	}
	special, err := randChar(p.Special)
	if err != nil {
		return nil, err
	}
	
	prefix = append(prefix,upper)
	prefix = append(prefix,lower)
	outPass = append(outPass,digits)
	outPass = append(outPass,special)

	
	for len(outPass) < (l-2) {
		rc, err := randChar(p.All)
		if err != nil {
			return nil, err
		}
		outPass = append(outPass,rc)
	}
	shuffle(prefix)
	shuffle(outPass)

	mergePass := append(prefix,outPass...)

	mergeString := string(mergePass)

	return &mergeString, nil
}