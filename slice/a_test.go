package main

import (
	"testing"
)

func TestSlice(t *testing.T) {
	t.Run("With Capacity", func(t *testing.T) {
		baseSlice := make([]string, 0, 1)

		// The values of BaseSlice are the same, but it has a higher capacity,
		// now pointing to elements 0:0 (none) of a larger array
		modSlice := append(baseSlice, "A")

		equal(t, "A", modSlice[0])

		// Appending 'B' to the base slice changes the mod slice.
		_ = append(baseSlice, "B")

		equal(t, "B", modSlice[0])
	})

	t.Run("Without Capacity", func(t *testing.T) {
		baseSlice := make([]string, 0)

		// Appending 'A' creates a new memory range, because it can't fit in the base slice.
		modSlice := append(baseSlice, "A")

		equal(t, "A", modSlice[0])

		// Appending to base has no effect on mod.
		_ = append(baseSlice, "B")
		equal(t, "A", modSlice[0])
	})

	t.Run("With Capacity Clone", func(t *testing.T) {
		baseSlice := make([]string, 0, 1)

		printSlice(t, "Base", baseSlice) // cap:1 len:1 []

		append1 := append(baseSlice, "A")

		// baseSlice's capacity is already sufficient to hold the new element, no new memory
		// range is created, and now the underlying array's 0th element is 'A'.

		printSlice(t, "Base", baseSlice) // cap:1 len:0 []
		printSlice(t, "App1", append1)   // cap:1 len:1 [A]
		equal(t, "A", append1[0])        // passes

		append2 := append(baseSlice, "B")

		// Appending to base again, still base has sufficient capacity to hold
		// the new element, since it is len(0) and cap(1), so again no new
		// memory is allocated, and the 0th element is 'B'

		printSlice(t, "Base", baseSlice) // cap:1 len:0 []
		printSlice(t, "App1", append1)   // cap:1 len:1 [B]
		printSlice(t, "App2", append2)   // cap:1 len:1 [B]

		// Wasn't expecting modSlice to be modified, but it is.
		equal(t, "B", append1[0])

	})

}

func printSlice(t testing.TB, name string, slice []string) {
	t.Logf("%s: cap: %d, len: %d at %p - %q", name, cap(slice), len(slice), slice, slice)
}
func equal(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
