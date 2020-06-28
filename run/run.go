package run

var (
	dt = 1. / FPS
)

func Run() {
	Scr.Init()
	InitChunks()

	NewStar(0, 0, 0, 0, 100)

	for !Scr.Win.Closed() {
		if Scr.Win.MouseScroll().Y > 0 {
			Scr.ChangeZoom(Scr.Win.MousePosition(), 1/1.1)
		} else if Scr.Win.MouseScroll().Y < 0 {
			Scr.ChangeZoom(Scr.Win.MousePosition(), 1.1)
		}
		for _, S := range Stars.Array {
			S.UpdateSpeed()
		}
		Stars.Array[0].SpeedX = Stars.Array[0].PosX - (Scr.Win.MousePosition().X*Scr.Zoom + Scr.PosX)
		Stars.Array[0].SpeedY = Stars.Array[0].PosY - (Scr.Win.MousePosition().Y*Scr.Zoom + Scr.PosY)

		for _, S := range Stars.Array {
			S.UpdatePos()
			S.CheckChunks()
		}
		for _, S := range Stars.Array {
			Scr.DrawStar(S)
		}
		Scr.Update()
	}
}
