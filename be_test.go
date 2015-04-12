package expect

import "testing"

// TODO(Ariel): Create mock that implement TB interface
// and stub `Error` and `Fatal`

func TestAbove(t *testing.T) {
	expect := New(t)
	expect(10).To.Be.Above(0)
	expect(10).Not.To.Be.Above(20)
}

func TestBelow(t *testing.T) {
	expect := New(t)
	expect(10).To.Be.Below(20)
	expect(10).Not.To.Be.Below(0)
}

func TestEmpty(t *testing.T) {
	expect := New(t)
	expect("").To.Be.Empty()
	expect([]int{1, 2, 3}).Not.To.Be.Empty()
	expect(make(map[string]int)).To.Be.Empty()
	expect([2]int{}).Not.To.Be.Empty()
	expect([]byte{}).To.Be.Empty()
}

func TestOk(t *testing.T) {
	expect := New(t)
	expect("").Not.To.Be.Ok()
	expect(false).Not.To.Be.Ok()
	expect(0).Not.To.Be.Ok()
	expect(nil).Not.To.Be.Ok()

	expect("foo").To.Be.Ok()
	expect(1).To.Be.Ok()
	expect(true).To.Be.Ok()
	expect(struct{}{}).To.Be.Ok()
	expect([]int{}).To.Be.Ok()
}

func TestString(t *testing.T) {
	expect := New(t)
	expect("").To.Be.String()
	expect(1).Not.To.Be.String()
}

func TestInt(t *testing.T) {
	expect := New(t)
	expect(0).To.Be.Int()
	expect("").Not.To.Be.Int()
}

func TestFloat(t *testing.T) {
	expect := New(t)
	expect(0).Not.To.Be.Float()
	expect(1.1).To.Be.Float()
}

func TestBool(t *testing.T) {
	expect := New(t)
	expect(true).To.Be.Bool()
	expect(1).Not.To.Be.Bool()
}

func TestMap(t *testing.T) {
	expect := New(t)
	expect(1).Not.To.Be.Map()
	expect(map[string]int{}).To.Be.Map()
}

func TestArray(t *testing.T) {
	expect := New(t)
	expect([]int{}).Not.To.Be.Array()
	expect([1]int{}).To.Be.Array()
	expect([...]int{1}).To.Be.Array()
}

func TestSlice(t *testing.T) {
	expect := New(t)
	expect([]int{}).To.Be.Slice()
	expect([]string{"a"}).To.Be.Slice()
	expect([1]int{}).Not.To.Be.Slice()
	expect([...]int{1}).Not.To.Be.Slice()
}

func TestChan(t *testing.T) {
	expect := New(t)
	var ch chan string
	expect(ch).To.Be.Chan()
	expect(1).Not.To.Be.Chan()
}

func TestBeChaining(t *testing.T) {
	expect := New(t)
	expect(10).To.Be.Above(0).And.Below(20)
	expect(10).Not.To.Be.Above(20).And.Below(0)
}
