package pile

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jjjosephhh/solitaire-thirteens-01/card"
	"github.com/jjjosephhh/solitaire-thirteens-01/constants"
	"github.com/jjjosephhh/solitaire-thirteens-01/dimensions"
)

type Pile struct {
	Cards    []*card.Card
	Position *rl.Vector2
}

func NewDeck(position *rl.Vector2, dimensionsCard *dimensions.Dimensions) *Pile {
	p := &Pile{
		Cards:    make([]*card.Card, 0),
		Position: position,
	}
	for i := 1; i <= 13; i++ {
		p.Cards = append(
			p.Cards,
			card.NewCard(card.Clubs, i, dimensionsCard.Width, dimensionsCard.Height),
			card.NewCard(card.Diamonds, i, dimensionsCard.Width, dimensionsCard.Height),
			card.NewCard(card.Hearts, i, dimensionsCard.Width, dimensionsCard.Height),
			card.NewCard(card.Spades, i, dimensionsCard.Width, dimensionsCard.Height),
		)
	}
	for _, card := range p.Cards {
		card.Position.X = p.Position.X
		card.Position.Y = p.Position.Y
	}
	p.Shuffle()
	return p
}

func (p *Pile) Size() int {
	return len(p.Cards)
}

func (p *Pile) IsEmpty() bool {
	return p.Size() == 0
}

func (p *Pile) InitializeInPlay() []*card.Card {
	var basePosition *rl.Vector2
	var drawn []*card.Card
	for i := 0; i < 10; i++ {
		c, ok := p.Draw()
		if !ok {
			continue
		}

		if basePosition == nil {
			position := rl.NewVector2(
				float32(rl.GetScreenWidth())/2-2.5*float32(c.Width),
				float32(rl.GetScreenHeight())/2-1*float32(c.Height),
			)
			basePosition = &position
		}

		c.Show = true
		newPosition := rl.NewVector2(
			basePosition.X+float32(i%5)*float32(c.Width)+float32(i%5)*float32(constants.SPACING_H),
			basePosition.Y+float32(i/5)*float32(c.Height+constants.SPACING_V)+constants.TOP_OFFSET,
		)
		c.TargetPosition.X = newPosition.X
		c.TargetPosition.Y = newPosition.Y
		drawn = append(drawn, c)
	}
	return drawn
}

func (p *Pile) Shuffle() {
	// Fisher-Yates shuffle algorithm
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(p.Cards), func(i, j int) {
		p.Cards[i], p.Cards[j] = p.Cards[j], p.Cards[i]
	})
}

func (p *Pile) Draw() (*card.Card, bool) {
	if p.IsEmpty() {
		return nil, false
	}
	end := len(p.Cards) - 1
	c := p.Cards[end]
	p.Cards = p.Cards[:end]
	return c, true
}

func (p *Pile) MoveTo(targetPile *Pile, cards []*card.Card) []*card.Card {
	var moved []*card.Card
	if len(cards) == 0 {
		return moved
	}
	cardMap := make(map[*card.Card]bool)
	for _, c := range cards {
		cardMap[c] = true
	}

	var i int
	for _, cardInPlay := range p.Cards {
		if _, ok := cardMap[cardInPlay]; ok {
			continue
		}
		p.Cards[i] = cardInPlay
		i++
	}
	p.Cards = p.Cards[:len(p.Cards)-len(cards)]
	targetPile.Cards = append(targetPile.Cards, cards...)

	for _, c := range cards {
		c.TargetPosition.X = targetPile.Position.X
		c.TargetPosition.Y = targetPile.Position.Y
		moved = append(moved, c)
	}
	return moved
}

func (p *Pile) MatchExists() bool {
	if p.IsEmpty() {
		return true
	}
	seen := make(map[int]bool)
	for _, c := range p.Cards {
		dif := constants.TARGET_NUM - c.Num
		if dif == 0 {
			return true
		}
		if _, ok := seen[dif]; ok {
			return true
		}
		seen[c.Num] = true
	}
	return false
}
