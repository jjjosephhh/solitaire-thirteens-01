package thirteens

import (
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jjjosephhh/solitaire-thirteens-01/card"
	"github.com/jjjosephhh/solitaire-thirteens-01/constants"
	"github.com/jjjosephhh/solitaire-thirteens-01/dimensions"
	"github.com/jjjosephhh/solitaire-thirteens-01/pile"
	"github.com/jjjosephhh/solitaire-thirteens-01/textures"
	"github.com/jjjosephhh/solitaire-thirteens-01/utils"
)

type Game struct {
	Selected1 *card.Card
	Selected2 *card.Card
}

func Run() {
	var screenWidth = 1000
	var screenHeight = 1000
	rl.InitWindow(int32(screenWidth), int32(screenHeight), "Solitaire Thirteens 01")
	defer rl.CloseWindow()

	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()

	rl.SetTargetFPS(60)

	t := textures.NewTextures()
	defer t.Unload()

	pathBackground := path.Join("assets", "cyberpunk-street-files", "Version 1", "PNG", "layers")

	dimensionsCard := dimensions.NewDimensions(t.Spades01.Width, t.Spades01.Height)

	posUnplayed := rl.NewVector2(float32(constants.SPACING_H), float32(constants.SPACING_V))
	posMatched := rl.NewVector2(
		float32(screenWidth)-float32(dimensionsCard.Width)-float32(constants.SPACING_H),
		float32(screenHeight)-float32(dimensionsCard.Height)-float32(constants.SPACING_V),
	)
	unplayed := pile.NewDeck(&posUnplayed, dimensionsCard)
	inPlay := &pile.Pile{Cards: unplayed.InitializeInPlay()}
	matched := &pile.Pile{
		Cards:    make([]*card.Card, 0),
		Position: &posMatched,
	}

	game := &Game{}
	var posMouse rl.Vector2
	var leftMouseDown bool
	canPlay := true
	imageBackground01 := rl.LoadImage(path.Join(pathBackground, "foreground.png"))
	scalingFactor01 := float32(imageBackground01.Width) / float32(screenWidth)
	rl.ImageResize(
		imageBackground01,
		int32(float32(imageBackground01.Width)/scalingFactor01),
		int32(float32(imageBackground01.Height)/scalingFactor01),
	)
	textureBackground01 := rl.LoadTextureFromImage(imageBackground01)
	defer rl.UnloadTexture(textureBackground01)
	rl.UnloadImage(imageBackground01)

	imageBackground02 := rl.LoadImage(path.Join(pathBackground, "back-buildings.png"))
	scalingFactor02 := float32(imageBackground02.Width) / float32(screenWidth)
	rl.ImageResize(
		imageBackground02,
		int32(float32(imageBackground02.Width)/scalingFactor02),
		int32(float32(imageBackground02.Height)/scalingFactor02),
	)
	textureBackground02 := rl.LoadTextureFromImage(imageBackground02)
	defer rl.UnloadTexture(textureBackground02)
	rl.UnloadImage(imageBackground02)

	imageBackground03 := rl.LoadImage(path.Join(pathBackground, "far-buildings.png"))
	scalingFactor03 := float32(imageBackground03.Width) / float32(screenWidth)
	rl.ImageResize(
		imageBackground03,
		int32(float32(imageBackground03.Width)/scalingFactor03),
		int32(float32(imageBackground03.Height)/scalingFactor03),
	)
	textureBackground03 := rl.LoadTextureFromImage(imageBackground03)
	defer rl.UnloadTexture(textureBackground03)
	rl.UnloadImage(imageBackground03)

	music := rl.LoadMusicStream(path.Join("assets", "cyberpunk-street-files", "music", "cyberpunk-street.ogg"))
	defer rl.UnloadMusicStream(music)
	rl.PlayMusicStream(music)

	for !rl.WindowShouldClose() {
		rl.UpdateMusicStream(music)
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		if rl.IsMouseButtonDown(rl.MouseLeftButton) {
			posMouse = rl.GetMousePosition()
			if !leftMouseDown {
				if unplayed.IsEmpty() && inPlay.IsEmpty() || !canPlay {
					// for _, c := range matched.Cards {
					// 	c.TargetPosition = unplayed.Position
					// }
					canPlay = true
					unplayed = pile.NewDeck(&posUnplayed, dimensionsCard)
					inPlay.Cards = unplayed.InitializeInPlay()
					matched.Cards = make([]*card.Card, 0)
					continue
				}
				for _, c := range inPlay.Cards {
					if !c.Clicked(&posMouse, dimensionsCard) {
						continue
					}
					if game.Selected1 == nil && game.Selected2 != c {
						game.Selected1 = c
						ClearPair(game, unplayed, inPlay, matched)
						break
					}

					if c == game.Selected1 {
						game.Selected1 = nil
						break
					}

					if game.Selected2 == nil && game.Selected1 != c {
						game.Selected2 = c
						ClearPair(game, unplayed, inPlay, matched)
						break
					}

					if c == game.Selected2 {
						game.Selected2 = nil
						break
					}
				}
			}
			leftMouseDown = true
		} else {
			leftMouseDown = false
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.NewColor(0, 128, 128, 255))

		for _, c := range unplayed.Cards {
			t.DrawCard(c)
		}

		for _, c := range inPlay.Cards {
			if !c.Move() && c.TargetFlipAngle == 0 {
				c.TargetFlipAngle = 180
			}
			if c == game.Selected1 || c == game.Selected2 {
				c.DrawOutline(dimensionsCard)
			}
			t.DrawCard(c)
		}

		for _, c := range matched.Cards {
			if !c.Show {
				continue
			}
			c.Move()
			t.DrawCard(c)
		}

		if inPlay.IsEmpty() {
			textWidth := rl.MeasureText(constants.TEXT_WIN, constants.TEXT_SIZE_WIN)
			rl.DrawText(
				constants.TEXT_WIN,
				(int32(rl.GetScreenWidth())-textWidth)/2,
				(int32(rl.GetScreenHeight())-constants.TEXT_SIZE_WIN)/2,
				constants.TEXT_SIZE_WIN,
				rl.White,
			)
		}

		if canPlay {
			canPlay = inPlay.MatchExists()
		}
		if !canPlay {
			rl.DrawRectangle(0, 0, int32(screenWidth), int32(screenHeight), constants.COLOR_RESTART)
			textWidth := rl.MeasureText(constants.TEXT_RESTART, constants.TEXT_SIZE_WIN)
			rl.DrawText(
				constants.TEXT_RESTART,
				(int32(rl.GetScreenWidth())-textWidth)/2,
				(int32(rl.GetScreenHeight())-constants.TEXT_SIZE_WIN)/2,
				constants.TEXT_SIZE_WIN,
				rl.White,
			)
		}

		rl.EndDrawing()
	}
}

func ClearPair(game *Game, unplayed, inPlay, matched *pile.Pile) {
	if matches := utils.IsMatch(game.Selected1, game.Selected2); len(matches) > 0 {
		inPlay.MoveTo(matched, matches)
		game.Selected1 = nil
		game.Selected2 = nil
		for _, matchedCard := range matches {
			c, ok := unplayed.Draw()
			if !ok {
				continue
			}
			c.Show = true
			c.TargetPosition.X = matchedCard.Position.X
			c.TargetPosition.Y = matchedCard.Position.Y
			inPlay.Cards = append(inPlay.Cards, c)
		}
	}
}
