package iteration

import "strings"

//const repeatCnt = 5

func Repeat(str string, times int) string {
	//var res string
	//for i := 0; i < 5; i++ {
	//	res += str
	//}
	//return res
	s_builder := strings.Builder{}
	for i := 0; i < times; i++ {
		s_builder.WriteString(str)
	}
	return s_builder.String()
}
