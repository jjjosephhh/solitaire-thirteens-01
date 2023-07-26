package card

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jjjosephhh/solitaire-thirteens-01/constants"
	"github.com/jjjosephhh/solitaire-thirteens-01/dimensions"
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
	position := rl.NewVector2(0, 0)
	targetPosition := rl.NewVector2(0, 0)
	return &Card{
		Position:        &position,
		TargetPosition:  &targetPosition,
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

func (c *Card) Clicked(mousePosition *rl.Vector2, d *dimensions.Dimensions) bool {
	return c.Position.X <= mousePosition.X &&
		mousePosition.X <= c.Position.X+float32(d.Width) &&
		c.Position.Y <= mousePosition.Y &&
		mousePosition.Y <= c.Position.Y+float32(d.Height)
}

func Randf32(min, max float32) float32 {
	if min >= max {
		panic("Invalid input: min must be less than max")
	}

	rand.Seed(time.Now().UnixNano())
	return min + rand.Float32()*(max-min)
}

func (c *Card) Move() bool {
	if !c.InMotion() {
		return false
	}
	direction := rl.Vector2Subtract(*c.TargetPosition, *c.Position)
	direction = rl.Vector2Normalize(direction)
	amount := rl.NewVector2(direction.X*constants.SPEED_CARD, direction.Y*constants.SPEED_CARD)
	newPosition := rl.Vector2Add(*c.Position, amount)
	c.Position.X = newPosition.X
	c.Position.Y = newPosition.Y
	distance := rl.Vector2Distance(*c.Position, *c.TargetPosition)
	if distance <= constants.SPEED_CARD {
		c.Position.X = c.TargetPosition.X
		c.Position.Y = c.TargetPosition.Y
	}
	return true
}

func (c *Card) DrawOutline(d *dimensions.Dimensions) {
	rl.DrawRectanglePro(
		rl.NewRectangle(
			c.Position.X+float32(d.Width)/2-2,
			c.Position.Y+float32(d.Height)/2-2,
			float32(d.Width)+4,
			float32(d.Height)+4,
		),
		rl.NewVector2(
			float32(d.Width)/2,
			float32(d.Height)/2,
		),
		c.RotationAngle,
		rl.Red,
	)
}
