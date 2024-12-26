package test

import (
	"testing"

	"github.com/ja-errorpro/CYCS_OSReplacementAlgorithms/lib"
)

func TestQueue(t *testing.T) {
	t.Run("Test Queue", func(t *testing.T) {

		t.Run("Test for queue Push", func(t *testing.T) {
			q := lib.NewQueue()
			q.Push(1)
			q.Push(2)
			q.Push(3)
		})

		t.Run("Test for queue Pop", func(t *testing.T) {
			q := lib.NewQueue()
			q.Push(1)
			q.Push(2)
			q.Push(3)
			if q.Pop() != 1 {
				t.Errorf("Queue Pop is not 1")
			}
			if q.Pop() != 2 {
				t.Errorf("Queue Pop is not 2")
			}
			if q.Pop() != 3 {
				t.Errorf("Queue Pop is not 3")
			}
		})

		t.Run("Test for queue Len", func(t *testing.T) {
			q := lib.NewQueue()
			q.Push(1)
			q.Push(2)
			q.Push(3)
			if q.Len() != 3 {
				t.Errorf("Queue length is not 3")
			}
		})

		t.Run("Test for queue Front", func(t *testing.T) {
			q := lib.NewQueue()
			q.Push(1)
			q.Push(2)
			q.Push(3)
			if q.Front() != 1 {
				t.Errorf("Queue Front is not 1")
			}
		})

		t.Run("Test for queue Contains", func(t *testing.T) {
			q := lib.NewQueue()
			q.Push(1)
			q.Push(2)
			q.Push(3)
			if !q.Contains(2) {
				t.Errorf("Queue Contains is not true")
			}
		})

		t.Run("Test for queue Traverse", func(t *testing.T) {
			q := lib.NewQueue()
			q.Push(1)
			q.Push(2)
			q.Push(3)
			if len(q.Traverse()) != 3 {
				t.Errorf("Queue Traverse length is not 3")
			}
		})

		t.Run("Test for queue Empty", func(t *testing.T) {
			q := lib.NewQueue()
			if !q.Empty() {
				t.Errorf("Queue is not empty")
			}
		})
	})
}
