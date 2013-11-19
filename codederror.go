
package codederror

import(
	"strings"
	"fmt"
)

type CodedError struct {
    namespaces []string
    code uint16
}

func (ce *CodedError) String() string {
    return fmt.Sprintf("%s%d", strings.Join(ce.namespaces, ""), ce.code)
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
