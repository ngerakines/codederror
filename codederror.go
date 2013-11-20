package codederror

import (
	"fmt"
	"strings"
)

var codedErrorMin = 0
var codedErrorMax = 33554431
var codedErrorLength = 5
var codedErrorCharacterCount = 33
var codedErrorCharacters = map[int]string{
    0: "A",
    1: "B",
    2: "C",
    3: "D",
    4: "E",
    5: "F",
    6: "G",
    7: "H",
    8: "J",
    9: "K",
    10: "L",
    11: "M",
    12: "N",
    13: "P",
    14: "Q",
    15: "R",
    16: "S",
    17: "T",
    18: "U",
    19: "V",
    20: "W",
    21: "X",
    22: "Y",
    23: "Z",
    24: "2",
    25: "3",
    26: "4",
    27: "5",
    28: "6",
    29: "7",
    30: "8",
    31: "9",
}

type CodedError struct {
	namespaces []string
	code       uint16
}

func encode(code uint16) string {
    if int(code) < codedErrorCharacterCount {
        return codedErrorCharacters[int(code)]
    }
    places := make([]string, 5)
    num := int(code)
    place := 0
    for {
        remainder := num % codedErrorCharacterCount
        places[place] = codedErrorCharacters[remainder]
        num = num / codedErrorCharacterCount;
        place++
        if num <= 0 {
            break
        }
    }
    
    return strings.Join(places, "")
}

func (ce *CodedError) String() string {
	return fmt.Sprintf("%s%s", strings.Join(ce.namespaces, ""), encode(ce.code))
}

func (ce *CodedError) Error() string {
	return ce.String()
}

func NewCodedError(namespaces []string, code uint16) *CodedError {
	ce := new(CodedError)
	ce.namespaces = namespaces
	ce.code = code
	return ce
}
