package codederror

import (
	"fmt"
	"strings"
)

var codedErrorMin = 0
var codedErrorMax = 33554431
var codedErrorLength = 5
var codedErrorCharacterCount = 31
var codedErrorCharacters = map[int]string{
	1:  "A",
	2:  "B",
	3:  "C",
	4:  "D",
	5:  "E",
	6:  "F",
	7:  "G",
	8:  "H",
	9:  "J",
	10: "K",
	11: "L",
	12: "M",
	13: "N",
	14: "P",
	15: "Q",
	16: "R",
	17: "S",
	18: "T",
	19: "U",
	20: "V",
	21: "W",
	22: "X",
	23: "Y",
	24: "Z",
	25: "2",
	26: "3",
	27: "4",
	28: "5",
	29: "6",
	30: "7",
	31: "8",
	32: "9",
}

type CodedError struct {
	namespaces []string
	code       uint16
}

func encode(code uint16) string {
	if int(code) <= codedErrorCharacterCount {
		return codedErrorCharacters[int(code)+1]
	}
	places := make([]string, 5)
	num := int(code)
	place := 0
	for {
		remainder := num % codedErrorCharacterCount
		digit := codedErrorCharacters[remainder]
		places[place] = digit
		place++
		num = num / codedErrorCharacterCount
		if num < 1 {
			break
		}
	}
	for i, j := 0, len(places)-1; i < j; i, j = i+1, j-1 {
		places[i], places[j] = places[j], places[i]
	}
	return strings.Join(places, "")
}

func (ce *CodedError) String() string {
	return fmt.Sprintf("%s%s", strings.Join(ce.namespaces, ""), encode(ce.code))
}

func (ce *CodedError) Error() string {
	return ce.String()
}

func (ce *CodedError) Namespaces() []string {
	return ce.namespaces
}

func (ce *CodedError) Code() uint16 {
	return ce.code
}

func NewCodedError(namespaces []string, code uint16) *CodedError {
	ce := new(CodedError)
	ce.namespaces = namespaces
	ce.code = code
	return ce
}
