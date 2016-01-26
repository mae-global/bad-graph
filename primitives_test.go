package badgraph

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Primitives(t *testing.T) {

	Convey("All Primitives", t, func() {
		
		n0 := NewNode(Identity("0"),"Node 0")
		So(n0,ShouldNotBeNil)
		So(n0.Identity.String(),ShouldEqual,"IDENTITY-0")
		So(n0.Alias,ShouldEqual,"Node 0")		

		n1 := NewNode(Identity("1"),"Node 1")
		So(n1,ShouldNotBeNil)

		So(n1.String(),ShouldEqual,`NODE "Node 1" IDENTITY-1, 0 TAGS`)
		
		n1.Set(Tag("foo"),"bar")
		n1.Set(Tag("bar"),"foo")
		
		So(n1.String(),ShouldEqual,`NODE "Node 1" IDENTITY-1, 2 TAGS`)		

		bar,ok := n1.Get(Tag("foo"))
		So(ok,ShouldBeTrue)
		So(bar,ShouldEqual,"bar")

		arc := NewArc(n0.Identity,n1.Identity,Muted)
		So(arc,ShouldNotBeNil)
		So(arc.Type,ShouldEqual,Muted)

		So(arc.String(),ShouldEqual,`ARC IDENTITY-0 <-x-> IDENTITY-1, 0 TAGS`)
		

	})
}
