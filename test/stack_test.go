package test

import (
	"testing"

	"github.com/ja-errorpro/CYCS_OSReplacementAlgorithms/lib"
)

func TestStack(t *testing.T) {

	t.Run("Test for stack", func(t *testing.T) {

		t.Run("Test for stack push", func(t *testing.T) {
			s := lib.NewStack()

			for i := 0; i < 10; i++ {
				s.Push(i)
			}

		})

		t.Run("Test for stack pop", func(t *testing.T) {

			s := lib.NewStack()

			for i := 0; i < 10; i++ {
				s.Push(i)
			}

			for i := 0; i < 10; i++ {
				s.Pop()
			}

		})

		t.Run("Test for stack top", func(t *testing.T) {

			s := lib.NewStack()

			for i := 0; i < 10; i++ {
				s.Push(i)
			}

			for i := 0; i < 6; i++ {
				s.Pop()
			}

			if s.Top() != 3 {
				t.Error("Top() should return 3")
			}

			for i := 0; i < 4; i++ {
				s.Pop()
			}

			if s.Pop() != nil {
				t.Error("Pop() should return nil")
			}
		})

		t.Run("Test for stack empty", func(t *testing.T) {

			s := lib.NewStack()

			if !s.Empty() {
				t.Error("Empty() should return true")
			}

			s.Push(1)

			if s.Empty() {
				t.Error("Empty() should return false")
			}
		})
	})

}
