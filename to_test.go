package expect

import "testing"

// TODO(Ariel): Create mock that implement TB interface
// and stub `Error` and `Fatal`

func TestStartWith(t *testing.T) {
	expect := New(t)
	expect("foo").To.StartWith("f")
	expect("foo").Not.To.StartWith("bar")
}

func TestEndWith(t *testing.T) {
	expect := New(t)
	expect("bar").To.EndWith("ar")
	expect("bar").Not.To.EndWith("az")
}

func TestContains(t *testing.T) {
	expect := New(t)
	expect("foobar").To.Contains("ba")
	expect("foobar").Not.To.Contains("ga")
}

func TestMatch(t *testing.T) {
	expect := New(t)
	expect("Foo").To.Match("(?i)foo")
}

func TestEqual(t *testing.T) {
	expect := New(t)
	expect("a").To.Equal("a")
	expect(1).To.Equal(1)
	expect(false).Not.To.Equal("true")
	expect(map[int]int{}).To.Equal(map[int]int{})
	expect(struct{ X, Y int }{1, 2}).Not.To.Equal(&struct{ X, Y int }{1, 2})
}

func TestPanic(t *testing.T) {
	expect := New(t)
	expect(func() {}).Not.To.Panic()
	expect(func() {
		panic("foo")
	}).To.Panic()
	expect(func() {
		panic("bar")
	}).To.Panic("bar")
}

func TestToChaining(t *testing.T) {
	expect := New(t)
	expect("foobarbaz").To.StartWith("foo").And.EndWith("baz").And.Contains("bar")
	expect("foo").Not.To.StartWith("bar").And.EndWith("baz").And.Contains("bob")
	expect("foo").To.Match("f").And.Match("(?i)F")
}
