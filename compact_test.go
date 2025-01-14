package offheap

import (
	"testing"

	cv "github.com/glycerine/goconvey/convey"
)

func TestCompact(t *testing.T) {

	cv.Convey("Given a big table with no values it, Compact() should re-allocate it to be smaller", t, func() {
		h := NewHashTable(128)
		cv.So(h.ArraySize, cv.ShouldEqual, 128)
		cv.So(h.Population, cv.ShouldEqual, 0)

		h.Compact()

		cv.So(h.ArraySize, cv.ShouldEqual, 1)
		cv.So(h.Population, cv.ShouldEqual, 0)
	})

}

func TestCompatAfterDelete(t *testing.T) {

	h := NewHashTable(2)

	h.Clear()
	cv.Convey("after being filled up and then deleted down to just 2 elements, Compact() should reduce table size to 4", t, func() {
		N := 100

		// fill up a table
		cv.So(h.Population, cv.ShouldEqual, 0)
		for i := 0; i < N; i++ {
			h.Insert(uint64(i))
		}

		// now delete from it
		for i := 0; i < N-2; i++ {
			h.DeleteKey(uint64(i))
		}
		cv.So(h.ArraySize, cv.ShouldEqual, 256)
		h.Compact()
		cv.So(h.ArraySize, cv.ShouldEqual, 4)

	})
}
