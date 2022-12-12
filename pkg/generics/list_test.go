package generics

import "testing"

func TestListPushAndPopInts(t *testing.T) {
	t.Run(
		"test int list push and pop", func(t *testing.T) {
			lst := List[int]{}
			lst.push(1)
			t.Log(lst.slice)

			lst.push(2)
			t.Log(lst.slice)

			lst.push(3)
			t.Log(lst.slice)

			lst.pop()
			t.Log(lst.slice)

			lst.pop()
			t.Log(lst.slice)

			lst.pop()
			t.Log(lst.slice)
		},
	)
}

func TestListPushAndPopStrings(t *testing.T) {
	t.Run(
		"test string list push and pop", func(t *testing.T) {
			lst := List[string]{}
			lst.push("A")
			t.Log(lst.slice)

			lst.push("B")
			t.Log(lst.slice)

			lst.push("C")
			t.Log(lst.slice)

			lst.pop()
			t.Log(lst.slice)

			lst.pop()
			t.Log(lst.slice)

			lst.pop()
			t.Log(lst.slice)
		},
	)
}
