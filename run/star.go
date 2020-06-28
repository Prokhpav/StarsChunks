package run

import (
	"fmt"
	"github.com/faiface/pixel"
	"math"
)

var (
	Stars = StarQueue{Array: []*Star{}}
)

type MassPoint struct {
	PosX, PosY, Mass float64
}

type Star struct {
	MassPoint
	SpeedX, SpeedY float64
	Radius         float64

	PosInQueue []int
}

func (S *Star) CheckChunks() {
	px, py := (S.PosX+MaxR)/MaxChunkSize, (S.PosY+MaxR)/MaxChunkSize
	npx, npy := S.PosX+S.SpeedX, S.PosY+S.SpeedY
	if npx < -MaxR || npx > MaxR || npy < -MaxR || npy > MaxR {
		return
	}
	npx, npy = (npx+MaxR)/MaxChunkSize, (npy+MaxR)/MaxChunkSize
	yes := false
	for i := 0; i < NumOfSlices; i++ {
		if yes || int(px) != int(npx) || int(py) != int(npy) {
			yes = true
			fmt.Println(Chunks[i][int(px)][int(py)].Stars)
			fmt.Println(Chunks[i][int(npx)][int(npy)].Stars)

			Chunks[i][int(px)][int(py)].DelStar(S)
			Chunks[i][int(npx)][int(npy)].AddStar(S)
		}
		px *= 2
		py *= 2
		npx *= 2
		npy *= 2
	}
}

func (S *Star) UpdateSpeed() {
	var MinX, MinY, MaxX, MaxY, C2X, C2Y, k int

	R := math.Sqrt(S.PosX*S.PosX + S.PosY*S.PosY)
	if R > MaxR {
		S.PosX *= MaxR / R
		S.PosY *= MaxR / R
	}

	PpR2X := S.PosX + MaxR
	PpR2Y := S.PosY + MaxR

	ChunkSize := MaxChunkSize

	Min2X := 0
	Min2Y := 0
	Max2X := 3
	Max2Y := 3

	for j := 0; j < NumOfSlices; j++ {
		k = len(Chunks[j]) - 1

		MinX = Min2X * 2
		MinY = Min2Y * 2
		MaxX = Max2X*2 + 1
		MaxY = Max2Y*2 + 1

		C2X = int(PpR2X/ChunkSize + 0.5)
		C2Y = int(PpR2Y/ChunkSize + 0.5)

		Min2X = Cmp(C2X-2, 0, k)
		Min2Y = Cmp(C2Y-2, 0, k)
		Max2X = Cmp(C2X+1, 0, k)
		Max2Y = Cmp(C2Y+1, 0, k)

		for x := MinX; x <= MaxX; x++ {
			for y := MinY; y < Min2Y; y++ {
				S.GravityTo(Chunks[j][x][y].CenterOfMass)
				Chunks[j][x][y].Draw(x, y, pixel.RGB(1, 0, 0))
			}
			for y := Max2Y + 1; y <= MaxY; y++ {
				S.GravityTo(Chunks[j][x][y].CenterOfMass)
				Chunks[j][x][y].Draw(x, y, pixel.RGB(1, 0, 0))
			}
		}
		for y := Min2Y; y <= Max2Y; y++ {
			for x := MinX; x < Min2X; x++ {
				S.GravityTo(Chunks[j][x][y].CenterOfMass)
				Chunks[j][x][y].Draw(x, y, pixel.RGB(1, 0, 0))
			}
			for x := Max2X + 1; x <= MaxX; x++ {
				S.GravityTo(Chunks[j][x][y].CenterOfMass)
				Chunks[j][x][y].Draw(x, y, pixel.RGB(1, 0, 0))
			}
		}

		ChunkSize /= 2
	}

	for x := Min2X; x <= Max2X; x++ {
		for y := Min2Y; y < Max2Y; y++ {
			for _, S2 := range Chunks[NumOfSlices-1][x][y].Stars.Array {
				if S2.PosInQueue[0] != S.PosInQueue[0] {
					S.GravityToStar(S2)
				}
			}
		}
	}
}

func (S *Star) UpdatePos() {
	S.PosX += S.SpeedX * dt
	S.PosY += S.SpeedY * dt
}

func (S *Star) GravityTo(M MassPoint) {
	if M.Mass != 0 {
		dx, dy := M.PosX-S.PosX, M.PosY-S.PosY
		R := math.Sqrt(dx*dx + dy*dy)
		k := G * M.Mass / (R * R * R)
		S.SpeedX += dx * k
		S.SpeedY += dy * k
	}

}

func (S *Star) GravityToStar(S2 *Star) {
	dx, dy := S2.PosX-S.PosX, S2.PosY-S.PosY
	R := math.Sqrt(dx*dx + dy*dy)
	k := G * dt * dt * S2.Mass / (R * R * R)
	if R < S.Radius+S2.Radius {
		k *= R / (S.Radius + S2.Radius)
	}
	S.SpeedX += dx * k
	S.SpeedY += dy * k
}
