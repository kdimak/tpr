package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_CalculateFiguresArea(t *testing.T) {
	// Given
	user := User{}

	rectangle := NewRectangle(2, 3)
	rectangle.SetWidth(4)
	rectangle.SetHeight(5)

	square := NewSquare(4)
	square.SetWidth(6)
	square.SetHeight(8)

	// When
	area := user.CalculateFiguresArea([]Figure{
		&rectangle,
		&square,
	})

	// Then
	assert.Equal(t, 36, area)
}

func TestUser_SetRectangleSize(t *testing.T) {
	// Given
	user := User{}

	//rectangle := NewRectangle(2, 3)
	rectangle := NewSquare(3)

	// When
	user.SetRectangleSize(&rectangle, 4, 5)

	// Then
	assert.Equal(t, 20, rectangle.Area())
}
