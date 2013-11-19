
package codederror

import(
	"strings"
	"fmt"
)

type CodedError struct {
    namespaces []string
    code uint16
}

func (ce CodedError) String() string {
    return fmt.Sprintf("%s%d", strings.Join(ce.namespaces, ""), ce.code)
}
