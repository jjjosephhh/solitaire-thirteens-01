package textures

import (
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jjjosephhh/solitaire-thirteens-01/card"
	"github.com/jjjosephhh/solitaire-thirteens-01/constants"
)

type Textures struct {
	Spades01 *rl.Texture2D
	Spades02 *rl.Texture2D
	Spades03 *rl.Texture2D
	Spades04 *rl.Texture2D
	Spades05 *rl.Texture2D
	Spades06 *rl.Texture2D
	Spades07 *rl.Texture2D
	Spades08 *rl.Texture2D
	Spades09 *rl.Texture2D
	Spades10 *rl.Texture2D
	Spades11 *rl.Texture2D
	Spades12 *rl.Texture2D
	Spades13 *rl.Texture2D

	Clubs01 *rl.Texture2D
	Clubs02 *rl.Texture2D
	Clubs03 *rl.Texture2D
	Clubs04 *rl.Texture2D
	Clubs05 *rl.Texture2D
	Clubs06 *rl.Texture2D
	Clubs07 *rl.Texture2D
	Clubs08 *rl.Texture2D
	Clubs09 *rl.Texture2D
	Clubs10 *rl.Texture2D
	Clubs11 *rl.Texture2D
	Clubs12 *rl.Texture2D
	Clubs13 *rl.Texture2D

	Diamonds01 *rl.Texture2D
	Diamonds02 *rl.Texture2D
	Diamonds03 *rl.Texture2D
	Diamonds04 *rl.Texture2D
	Diamonds05 *rl.Texture2D
	Diamonds06 *rl.Texture2D
	Diamonds07 *rl.Texture2D
	Diamonds08 *rl.Texture2D
	Diamonds09 *rl.Texture2D
	Diamonds10 *rl.Texture2D
	Diamonds11 *rl.Texture2D
	Diamonds12 *rl.Texture2D
	Diamonds13 *rl.Texture2D

	Hearts01 *rl.Texture2D
	Hearts02 *rl.Texture2D
	Hearts03 *rl.Texture2D
	Hearts04 *rl.Texture2D
	Hearts05 *rl.Texture2D
	Hearts06 *rl.Texture2D
	Hearts07 *rl.Texture2D
	Hearts08 *rl.Texture2D
	Hearts09 *rl.Texture2D
	Hearts10 *rl.Texture2D
	Hearts11 *rl.Texture2D
	Hearts12 *rl.Texture2D
	Hearts13 *rl.Texture2D

	Back01 *rl.Texture2D
	Back02 *rl.Texture2D
	Back03 *rl.Texture2D
	Back04 *rl.Texture2D

	Blank *rl.Texture2D
}

func NewTextures() *Textures {
	var t Textures
	pathCards := path.Join("assets", "Pixel Fantasy Playing Cards", "Playing Cards")
	textureCardBack01 := rl.LoadTexture(path.Join(pathCards, "card-back1.png"))
	t.Back01 = &textureCardBack01
	textureCardBack02 := rl.LoadTexture(path.Join(pathCards, "card-back2.png"))
	t.Back02 = &textureCardBack02
	textureCardBack03 := rl.LoadTexture(path.Join(pathCards, "card-back3.png"))
	t.Back03 = &textureCardBack03
	textureCardBack04 := rl.LoadTexture(path.Join(pathCards, "card-back4.png"))
	t.Back04 = &textureCardBack04
	textureCardBlank := rl.LoadTexture(path.Join(pathCards, "card-blank.png"))
	t.Blank = &textureCardBlank
	textureCardClubs01 := rl.LoadTexture(path.Join(pathCards, "card-clubs-1.png"))
	t.Clubs01 = &textureCardClubs01
	textureCardClubs02 := rl.LoadTexture(path.Join(pathCards, "card-clubs-2.png"))
	t.Clubs02 = &textureCardClubs02
	textureCardClubs03 := rl.LoadTexture(path.Join(pathCards, "card-clubs-3.png"))
	t.Clubs03 = &textureCardClubs03
	textureCardClubs04 := rl.LoadTexture(path.Join(pathCards, "card-clubs-4.png"))
	t.Clubs04 = &textureCardClubs04
	textureCardClubs05 := rl.LoadTexture(path.Join(pathCards, "card-clubs-5.png"))
	t.Clubs05 = &textureCardClubs05
	textureCardClubs06 := rl.LoadTexture(path.Join(pathCards, "card-clubs-6.png"))
	t.Clubs06 = &textureCardClubs06
	textureCardClubs07 := rl.LoadTexture(path.Join(pathCards, "card-clubs-7.png"))
	t.Clubs07 = &textureCardClubs07
	textureCardClubs08 := rl.LoadTexture(path.Join(pathCards, "card-clubs-8.png"))
	t.Clubs08 = &textureCardClubs08
	textureCardClubs09 := rl.LoadTexture(path.Join(pathCards, "card-clubs-9.png"))
	t.Clubs09 = &textureCardClubs09
	textureCardClubs10 := rl.LoadTexture(path.Join(pathCards, "card-clubs-10.png"))
	t.Clubs10 = &textureCardClubs10
	textureCardClubs11 := rl.LoadTexture(path.Join(pathCards, "card-clubs-11.png"))
	t.Clubs11 = &textureCardClubs11
	textureCardClubs12 := rl.LoadTexture(path.Join(pathCards, "card-clubs-12.png"))
	t.Clubs12 = &textureCardClubs12
	textureCardClubs13 := rl.LoadTexture(path.Join(pathCards, "card-clubs-13.png"))
	t.Clubs13 = &textureCardClubs13
	textureCardDiamonds01 := rl.LoadTexture(path.Join(pathCards, "card-diamonds-1.png"))
	t.Diamonds01 = &textureCardDiamonds01
	textureCardDiamonds02 := rl.LoadTexture(path.Join(pathCards, "card-diamonds-2.png"))
	t.Diamonds02 = &textureCardDiamonds02
	textureCardDiamonds03 := rl.LoadTexture(path.Join(pathCards, "card-diamonds-3.png"))
	t.Diamonds03 = &textureCardDiamonds03
	textureCardDiamonds04 := rl.LoadTexture(path.Join(pathCards, "card-diamonds-4.png"))
	t.Diamonds04 = &textureCardDiamonds04
	textureCardDiamonds05 := rl.LoadTexture(path.Join(pathCards, "card-diamonds-5.png"))
	t.Diamonds05 = &textureCardDiamonds05
	textureCardDiamonds06 := rl.LoadTexture(path.Join(pathCards, "card-diamonds-6.png"))
	t.Diamonds06 = &textureCardDiamonds06
	textureCardDiamonds07 := rl.LoadTexture(path.Join(pathCards, "card-diamonds-7.png"))
	t.Diamonds07 = &textureCardDiamonds07
	textureCardDiamonds08 := rl.LoadTexture(path.Join(pathCards, "card-diamonds-8.png"))
	t.Diamonds08 = &textureCardDiamonds08
	textureCardDiamonds09 := rl.LoadTexture(path.Join(pathCards, "card-diamonds-9.png"))
	t.Diamonds09 = &textureCardDiamonds09
	textureCardDiamonds10 := rl.LoadTexture(path.Join(pathCards, "card-diamonds-10.png"))
	t.Diamonds10 = &textureCardDiamonds10
	textureCardDiamonds11 := rl.LoadTexture(path.Join(pathCards, "card-diamonds-11.png"))
	t.Diamonds11 = &textureCardDiamonds11
	textureCardDiamonds12 := rl.LoadTexture(path.Join(pathCards, "card-diamonds-12.png"))
	t.Diamonds12 = &textureCardDiamonds12
	textureCardDiamonds13 := rl.LoadTexture(path.Join(pathCards, "card-diamonds-13.png"))
	t.Diamonds13 = &textureCardDiamonds13
	textureCardHearts01 := rl.LoadTexture(path.Join(pathCards, "card-hearts-1.png"))
	t.Hearts01 = &textureCardHearts01
	textureCardHearts02 := rl.LoadTexture(path.Join(pathCards, "card-hearts-2.png"))
	t.Hearts02 = &textureCardHearts02
	textureCardHearts03 := rl.LoadTexture(path.Join(pathCards, "card-hearts-3.png"))
	t.Hearts03 = &textureCardHearts03
	textureCardHearts04 := rl.LoadTexture(path.Join(pathCards, "card-hearts-4.png"))
	t.Hearts04 = &textureCardHearts04
	textureCardHearts05 := rl.LoadTexture(path.Join(pathCards, "card-hearts-5.png"))
	t.Hearts05 = &textureCardHearts05
	textureCardHearts06 := rl.LoadTexture(path.Join(pathCards, "card-hearts-6.png"))
	t.Hearts06 = &textureCardHearts06
	textureCardHearts07 := rl.LoadTexture(path.Join(pathCards, "card-hearts-7.png"))
	t.Hearts07 = &textureCardHearts07
	textureCardHearts08 := rl.LoadTexture(path.Join(pathCards, "card-hearts-8.png"))
	t.Hearts08 = &textureCardHearts08
	textureCardHearts09 := rl.LoadTexture(path.Join(pathCards, "card-hearts-9.png"))
	t.Hearts09 = &textureCardHearts09
	textureCardHearts10 := rl.LoadTexture(path.Join(pathCards, "card-hearts-10.png"))
	t.Hearts10 = &textureCardHearts10
	textureCardHearts11 := rl.LoadTexture(path.Join(pathCards, "card-hearts-11.png"))
	t.Hearts11 = &textureCardHearts11
	textureCardHearts12 := rl.LoadTexture(path.Join(pathCards, "card-hearts-12.png"))
	t.Hearts12 = &textureCardHearts12
	textureCardHearts13 := rl.LoadTexture(path.Join(pathCards, "card-hearts-13.png"))
	t.Hearts13 = &textureCardHearts13
	textureCardSpades01 := rl.LoadTexture(path.Join(pathCards, "card-spades-1.png"))
	t.Spades01 = &textureCardSpades01
	textureCardSpades02 := rl.LoadTexture(path.Join(pathCards, "card-spades-2.png"))
	t.Spades02 = &textureCardSpades02
	textureCardSpades03 := rl.LoadTexture(path.Join(pathCards, "card-spades-3.png"))
	t.Spades03 = &textureCardSpades03
	textureCardSpades04 := rl.LoadTexture(path.Join(pathCards, "card-spades-4.png"))
	t.Spades04 = &textureCardSpades04
	textureCardSpades05 := rl.LoadTexture(path.Join(pathCards, "card-spades-5.png"))
	t.Spades05 = &textureCardSpades05
	textureCardSpades06 := rl.LoadTexture(path.Join(pathCards, "card-spades-6.png"))
	t.Spades06 = &textureCardSpades06
	textureCardSpades07 := rl.LoadTexture(path.Join(pathCards, "card-spades-7.png"))
	t.Spades07 = &textureCardSpades07
	textureCardSpades08 := rl.LoadTexture(path.Join(pathCards, "card-spades-8.png"))
	t.Spades08 = &textureCardSpades08
	textureCardSpades09 := rl.LoadTexture(path.Join(pathCards, "card-spades-9.png"))
	t.Spades09 = &textureCardSpades09
	textureCardSpades10 := rl.LoadTexture(path.Join(pathCards, "card-spades-10.png"))
	t.Spades10 = &textureCardSpades10
	textureCardSpades11 := rl.LoadTexture(path.Join(pathCards, "card-spades-11.png"))
	t.Spades11 = &textureCardSpades11
	textureCardSpades12 := rl.LoadTexture(path.Join(pathCards, "card-spades-12.png"))
	t.Spades12 = &textureCardSpades12
	textureCardSpades13 := rl.LoadTexture(path.Join(pathCards, "card-spades-13.png"))
	t.Spades13 = &textureCardSpades13
	return &t
}

func (t *Textures) Unload() {
	rl.UnloadTexture(*t.Back01)
	rl.UnloadTexture(*t.Back02)
	rl.UnloadTexture(*t.Back03)
	rl.UnloadTexture(*t.Back04)
	rl.UnloadTexture(*t.Blank)
	rl.UnloadTexture(*t.Clubs01)
	rl.UnloadTexture(*t.Clubs02)
	rl.UnloadTexture(*t.Clubs03)
	rl.UnloadTexture(*t.Clubs04)
	rl.UnloadTexture(*t.Clubs05)
	rl.UnloadTexture(*t.Clubs06)
	rl.UnloadTexture(*t.Clubs07)
	rl.UnloadTexture(*t.Clubs08)
	rl.UnloadTexture(*t.Clubs09)
	rl.UnloadTexture(*t.Clubs10)
	rl.UnloadTexture(*t.Clubs11)
	rl.UnloadTexture(*t.Clubs12)
	rl.UnloadTexture(*t.Clubs13)
	rl.UnloadTexture(*t.Diamonds01)
	rl.UnloadTexture(*t.Diamonds02)
	rl.UnloadTexture(*t.Diamonds03)
	rl.UnloadTexture(*t.Diamonds04)
	rl.UnloadTexture(*t.Diamonds05)
	rl.UnloadTexture(*t.Diamonds06)
	rl.UnloadTexture(*t.Diamonds07)
	rl.UnloadTexture(*t.Diamonds08)
	rl.UnloadTexture(*t.Diamonds09)
	rl.UnloadTexture(*t.Diamonds10)
	rl.UnloadTexture(*t.Diamonds11)
	rl.UnloadTexture(*t.Diamonds12)
	rl.UnloadTexture(*t.Diamonds13)
	rl.UnloadTexture(*t.Hearts01)
	rl.UnloadTexture(*t.Hearts02)
	rl.UnloadTexture(*t.Hearts03)
	rl.UnloadTexture(*t.Hearts04)
	rl.UnloadTexture(*t.Hearts05)
	rl.UnloadTexture(*t.Hearts06)
	rl.UnloadTexture(*t.Hearts07)
	rl.UnloadTexture(*t.Hearts08)
	rl.UnloadTexture(*t.Hearts09)
	rl.UnloadTexture(*t.Hearts10)
	rl.UnloadTexture(*t.Hearts11)
	rl.UnloadTexture(*t.Hearts12)
	rl.UnloadTexture(*t.Hearts13)
	rl.UnloadTexture(*t.Spades01)
	rl.UnloadTexture(*t.Spades02)
	rl.UnloadTexture(*t.Spades03)
	rl.UnloadTexture(*t.Spades04)
	rl.UnloadTexture(*t.Spades05)
	rl.UnloadTexture(*t.Spades06)
	rl.UnloadTexture(*t.Spades07)
	rl.UnloadTexture(*t.Spades08)
	rl.UnloadTexture(*t.Spades09)
	rl.UnloadTexture(*t.Spades10)
	rl.UnloadTexture(*t.Spades11)
	rl.UnloadTexture(*t.Spades12)
	rl.UnloadTexture(*t.Spades13)
}

func (t *Textures) DrawCard(c *card.Card) {
	if c == nil {
		return
	}
	if !c.Show {
		return
	}

	var texture rl.Texture2D
	switch c.Suit {
	case card.Clubs:
		switch c.Num {
		case 1:
			texture = *t.Clubs01
		case 2:
			texture = *t.Clubs02
		case 3:
			texture = *t.Clubs03
		case 4:
			texture = *t.Clubs04
		case 5:
			texture = *t.Clubs05
		case 6:
			texture = *t.Clubs06
		case 7:
			texture = *t.Clubs07
		case 8:
			texture = *t.Clubs08
		case 9:
			texture = *t.Clubs09
		case 10:
			texture = *t.Clubs10
		case 11:
			texture = *t.Clubs11
		case 12:
			texture = *t.Clubs12
		case 13:
			texture = *t.Clubs13
		default:
			texture = *t.Blank
		}
	case card.Diamonds:
		switch c.Num {
		case 1:
			texture = *t.Diamonds01
		case 2:
			texture = *t.Diamonds02
		case 3:
			texture = *t.Diamonds03
		case 4:
			texture = *t.Diamonds04
		case 5:
			texture = *t.Diamonds05
		case 6:
			texture = *t.Diamonds06
		case 7:
			texture = *t.Diamonds07
		case 8:
			texture = *t.Diamonds08
		case 9:
			texture = *t.Diamonds09
		case 10:
			texture = *t.Diamonds10
		case 11:
			texture = *t.Diamonds11
		case 12:
			texture = *t.Diamonds12
		case 13:
			texture = *t.Diamonds13
		default:
			texture = *t.Blank
		}
	case card.Hearts:
		switch c.Num {
		case 1:
			texture = *t.Hearts01
		case 2:
			texture = *t.Hearts02
		case 3:
			texture = *t.Hearts03
		case 4:
			texture = *t.Hearts04
		case 5:
			texture = *t.Hearts05
		case 6:
			texture = *t.Hearts06
		case 7:
			texture = *t.Hearts07
		case 8:
			texture = *t.Hearts08
		case 9:
			texture = *t.Hearts09
		case 10:
			texture = *t.Hearts10
		case 11:
			texture = *t.Hearts11
		case 12:
			texture = *t.Hearts12
		case 13:
			texture = *t.Hearts13
		default:
			texture = *t.Blank
		}
	case card.Spades:
		switch c.Num {
		case 1:
			texture = *t.Spades01
		case 2:
			texture = *t.Spades02
		case 3:
			texture = *t.Spades03
		case 4:
			texture = *t.Spades04
		case 5:
			texture = *t.Spades05
		case 6:
			texture = *t.Spades06
		case 7:
			texture = *t.Spades07
		case 8:
			texture = *t.Spades08
		case 9:
			texture = *t.Spades09
		case 10:
			texture = *t.Spades10
		case 11:
			texture = *t.Spades11
		case 12:
			texture = *t.Spades12
		case 13:
			texture = *t.Spades13
		default:
			texture = *t.Blank
		}
	default:
		texture = *t.Blank
	}

	if c.FlipAngle < c.TargetFlipAngle {
		c.FlipAngle += constants.SPEED_FLIP
	}

	if c.FlipAngle >= c.TargetFlipAngle {
		c.FlipAngle = c.TargetFlipAngle
	}

	if c.FlipAngle < 90 {
		rl.DrawTexturePro(
			*t.Back01,
			rl.NewRectangle(0, 0, float32(t.Back01.Width), float32(t.Back01.Height)),
			rl.NewRectangle(
				c.Position.X+float32(t.Back01.Width)/2+(float32(t.Back01.Width/2)*c.FlipAngle/90),
				c.Position.Y+float32(t.Back01.Height)/2,
				float32(t.Back01.Width)*(90-c.FlipAngle)/90,
				float32(t.Back01.Height),
			),
			rl.NewVector2(float32(t.Back01.Width)/2, float32(t.Back01.Height)/2),
			c.RotationAngle,
			rl.White,
		)
	} else {
		rl.DrawTexturePro(
			texture,
			rl.NewRectangle(0, 0, float32(t.Back01.Width), float32(t.Back01.Height)),
			rl.NewRectangle(
				c.Position.X+float32(t.Back01.Width)/2+(float32(t.Back01.Width/2)*(180-c.FlipAngle)/90),
				c.Position.Y+float32(t.Back01.Height)/2,
				float32(t.Back01.Width)*(c.FlipAngle-90)/90,
				float32(t.Back01.Height),
			),
			rl.NewVector2(float32(t.Back01.Width)/2, float32(t.Back01.Height)/2),
			c.RotationAngle,
			rl.White,
		)
	}
}
