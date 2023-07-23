package card

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type CardSuit int

const (
	Clubs CardSuit = iota
	Diamonds
	Hearts
	Spades
)

type Card struct {
	Position       *rl.Vector2
	TargetPosition *rl.Vector2
	Suit           CardSuit
	Num            int
	Show           bool
	Width          int32
	Height         int32

	RotationAngle   float32
	FlipAngle       float32
	TargetFlipAngle float32
}

func NewCard(suit CardSuit, num int, cardWidth, cardHeight int32) *Card {
	return &Card{
		Suit:            suit,
		Num:             num,
		Show:            true,
		Width:           cardWidth,
		Height:          cardHeight,
		FlipAngle:       0,
		TargetFlipAngle: 0,
		RotationAngle:   Randf32(-3, 3),
	}
}

func (c *Card) InMotion() bool {
	return c.Position.X != c.TargetPosition.X || c.Position.Y != c.TargetPosition.Y
}

func Randf32(min, max float32) float32 {
	if min >= max {
		panic("Invalid input: min must be less than max")
	}

	rand.Seed(time.Now().UnixNano())
	return min + rand.Float32()*(max-min)
}
