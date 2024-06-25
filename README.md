# Simpl Password Generator
간단한 패스워드 생성기로 아래의 조건에 따라 랜덤 패스워드를 생성합니다. 
- 최소 3글자 이상이며, 3글자 이하로 추출할 경우 Default 3자로 적용 됩니다. 
- 2글자까지는 무조건 영문자 대문자 또는 소문자가 혼합되어 순서상관없이 Prefix 로 구성됩니다.
- 숫자 및 특수문자를 포함하는 것은 선택할 수 있으나, 포함하는 수에 대해서는 랜덤하게 적용됩니다. (최소 1자 이상)

## Example
```go
package main 
import (
    gen "github.com/sizzlei/SimplePassGenerator"
    "fmt"
)
func main() {
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

	fmt.Println("Password: ",*pass)
}
```


