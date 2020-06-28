package run

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"time"
)

var (
	Scr Screen
)

type Screen struct {
	Win   *pixelgl.Window
	Imd   *imdraw.IMDraw
	Timer <-chan time.Time

	PosX, PosY, Zoom float64
}

func (Scr *Screen) Init() {
	var err error
	Scr.Win, err = pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "Gravity simulation",
		Bounds: pixel.R(0, 0, float64(WinSizeW), float64(WinSizeH))},
	)
	if err != nil {
		panic(err)
	}
	Scr.Timer = time.Tick(time.Second / FPS)
	Scr.Imd = imdraw.New(nil)
	Scr.Zoom = 1
	Scr.PosX = -Scr.Zoom * WinSizeW / 2.
	Scr.PosY = -Scr.Zoom * WinSizeH / 2.
}

func (Scr *Screen) Update() {
	Scr.Win.Clear(pixel.RGB(0, 0, 0))
	Scr.Imd.Draw(Scr.Win)
	Scr.Imd.Clear()
	Scr.Win.Update()
	select {
	case <-Scr.Timer:
	}
}

func (Scr *Screen) DrawCircle(PosX, PosY, R float64, Color pixel.RGBA) {
	PosX, PosY = (PosX-Scr.PosX)/Scr.Zoom, (PosY-Scr.PosY)/Scr.Zoom
	//fmt.Println(PosX, PosY)
	if PosX <= -R || PosX >= WinSizeW+R || PosY <= -R || PosY >= WinSizeH+R {
		return
	}
	Scr.Imd.Color = Color
	Scr.Imd.Push(pixel.Vec{X: PosX, Y: PosY})
	Scr.Imd.Circle(Max(R/Scr.Zoom, 0.6), 0)
}

func (Scr *Screen) DrawStar(S *Star) {
	Scr.DrawCircle(S.PosX, S.PosY, S.Radius, pixel.RGB(1, 1, 1))
}

func (Scr *Screen) DrawRect(PosX, PosY, MaxPosX, MaxPosY float64, Color pixel.RGBA) {
	//fmt.Println(PosX, PosY, MaxPosX, MaxPosY)
	PosX, PosY, MaxPosX, MaxPosY = (PosX-Scr.PosX)/Scr.Zoom, (PosY-Scr.PosY)/Scr.Zoom, (MaxPosX-Scr.PosX)/Scr.Zoom, (MaxPosY-Scr.PosY)/Scr.Zoom
	if MaxPosX <= 0 || PosX >= WinSizeW || MaxPosY <= 0 || PosY >= WinSizeH {
		return
	}
	Scr.Imd.Color = Color
	Scr.Imd.Push(pixel.Vec{X: PosX, Y: PosY}, pixel.Vec{X: MaxPosX, Y: MaxPosY})
	Scr.Imd.Rectangle(2)
}

func (Scr *Screen) ChangeZoom(focus pixel.Vec, k float64) {
	Scr.PosX += focus.X * Scr.Zoom * (1 - k)
	Scr.PosY += focus.Y * Scr.Zoom * (1 - k)
	Scr.Zoom *= k
}
