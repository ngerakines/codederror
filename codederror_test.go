
package codederror;

import(
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSomething(t *testing.T) {
    Convey("CodedError types can be created", t, func() {
        ce := NewCodedError([]string{"ABC", "DEF", "GHI"}, 14)
	    So(ce.String(), ShouldEqual, "ABCDEFGHI14")
    })
}
