package p1

import "strconv"

func StrToIntSlice(s *[]string) []int {
	slc := *s
	retSlc := []int{}
	for _, str := range slc {
		conv, _ := strconv.Atoi(str)
		retSlc = append(retSlc, conv)
	}
	return retSlc
}
