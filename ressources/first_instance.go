package color

import (
	"fmt"
	"strings"
)

func FirstInstance(sub, str string) []int {
	temp := str
	v,n := 0,0
	first := false
	mok := []int{}
	for i := 0; i < len(str); i++ {
		v = strings.Index(temp, sub)
		n = strings.Index(temp, "\\n")
		fmt.Println(n)
		if v != -1  {
			if !first {
			first = true 
			mok = append(mok, v)
			}else{
				mok = append(mok, v+mok[len(mok)-1]+len(sub))
			}
			
			temp = temp[v+len(sub):]
		}
		if n != -1  {
			mok = append(mok, -1)
			if len(temp) != 0 {
				temp = temp[n+len("\n"):]
			}
			
		}
	}
	fmt.Println(mok)
	return mok
}
