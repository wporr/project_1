package main
func (g *Game) Draw(screen *ebiten.Image) {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	game := &Game{}
`	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("project_1")

	// Start game loop
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
