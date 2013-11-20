package codederror

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSomething(t *testing.T) {
	Convey("CodedError types can be created", t, func() {
		ce := NewCodedError([]string{"ABC", "DEF", "GHI"}, 0)
		So(ce.String(), ShouldEqual, "ABCDEFGHIA")
	})
	Convey("CodedError types can be created", t, func() {
		ce := NewCodedError([]string{"ABC", "DEF", "GHI"}, 1)
		So(ce.String(), ShouldEqual, "ABCDEFGHIB")
	})
	Convey("CodedError types can be created", t, func() {
		ce := NewCodedError([]string{"ABC", "DEF", "GHI"}, 31)
		So(ce.String(), ShouldEqual, "ABCDEFGHI9")
	})
	Convey("CodedError types can be created", t, func() {
		ce := NewCodedError([]string{"ABC", "DEF", "GHI"}, 33)
		So(ce.String(), ShouldEqual, "ABCDEFGHIAB")
	})
}
