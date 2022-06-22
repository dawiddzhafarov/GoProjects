package generics

import "testing"

// generic test functions enables them to work with different types

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})
	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "grace")
	})

}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didnt want %+v", got)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, but wanted true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, but wanted true", got)
	}
}

func TestStack(t *testing.T) {
	t.Run("integers", func(t *testing.T) {
		myStackOfInts := new(Stack[int])

		AssertTrue(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(2)
		AssertFalse(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(34)
		value, _ := myStackOfInts.Pop()
		AssertEqual(t, value, 34)
		value, _ = myStackOfInts.Pop()
		AssertEqual(t, value, 2)
		AssertTrue(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(1)
		myStackOfInts.Push(2)
		firNum, _ := myStackOfInts.Pop()
		secNum, _ := myStackOfInts.Pop()
		AssertEqual(t, firNum+secNum, 3)
	})
	// no need to use this test, as we can create another stack[type] of
	// whatever type
	t.Run("strings", func(t *testing.T) {
		myStackOfStrings := new(StackOfStrings)

		AssertTrue(t, myStackOfStrings.IsEmpty())

		myStackOfStrings.Push("hello")
		AssertFalse(t, myStackOfStrings.IsEmpty())

		myStackOfStrings.Push("foo")
		value, _ := myStackOfStrings.Pop()
		AssertEqual(t, value, "foo")
		value, _ = myStackOfStrings.Pop()
		AssertEqual(t, value, "hello")
		AssertTrue(t, myStackOfStrings.IsEmpty())
	})
}
