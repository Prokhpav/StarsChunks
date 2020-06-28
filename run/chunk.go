package run

import (
	"github.com/faiface/pixel"
)

var (
	Chunks [][][]*Chunk

	NumOfSlices  int
	MaxChunkSize float64
	ThoPowK      []float64
)

func InitChunks() {
	MaxChunkSize = MaxR / 4
	ChunkSize := MaxChunkSize
	NumOfSlices = 0
	ThoPowK = []float64{1}
	for ChunkSize >= MinChunkSize {
		ChunkSize /= 2
		ThoPowK = append(ThoPowK, ThoPowK[NumOfSlices]*2)
		NumOfSlices++
	}
	MinChunkSize = ChunkSize * 2

	Chunks = make([][][]*Chunk, NumOfSlices)
	k := 8
	for i1 := 0; i1 < NumOfSlices; i1++ {
		Chunks[i1] = make([][]*Chunk, k)
		for i2 := 0; i2 < k; i2++ {
			Chunks[i1][i2] = make([]*Chunk, k)
			for i3 := 0; i3 < k; i3++ {
				Chunks[i1][i2][i3] = &Chunk{
					CenterOfMass: MassPoint{},
					Stars: &StarQueue{
						Array: []*Star{},
						Depth: i1 + 1,
					},
				}
			}
		}
		k *= 2
	}
}

type Chunk struct {
	CenterOfMass MassPoint
	Stars        *StarQueue
}

func (Ch *Chunk) Update() { // Обновляет положение центра массы. Звёзды, улетающие за пределы чанка, должны быть удалены заранее
	if len(Ch.Stars.Array) == 0 {
		return
	} else if len(Ch.Stars.Array) == 1 {
		Ch.CenterOfMass.PosX += Ch.Stars.Array[0].SpeedX
		Ch.CenterOfMass.PosY += Ch.Stars.Array[0].SpeedY
		return
	}
	var k float64
	for _, S := range Ch.Stars.Array {
		k = S.Mass / Ch.CenterOfMass.Mass
		Ch.CenterOfMass.PosX += S.SpeedX * k
		Ch.CenterOfMass.PosY += S.SpeedY * k
	}
}

func (Ch *Chunk) DelStar(S *Star) {
	if len(Ch.Stars.Array) == 1 {
		Ch.CenterOfMass.PosX = 0
		Ch.CenterOfMass.PosY = 0
		Ch.CenterOfMass.Mass = 0
	} else {
		Ch.CenterOfMass.Mass -= S.Mass
		k := S.Mass / Ch.CenterOfMass.Mass
		Ch.CenterOfMass.PosX += (Ch.CenterOfMass.PosX - S.PosX) * k
		Ch.CenterOfMass.PosY += (Ch.CenterOfMass.PosY - S.PosY) * k
	}
	Ch.Stars.DelStar(S)
}

func (Ch *Chunk) AddStar(S *Star) {
	if len(Ch.Stars.Array) == 0 {
		Ch.CenterOfMass.PosX = S.PosX
		Ch.CenterOfMass.PosY = S.PosY
		Ch.CenterOfMass.Mass = S.Mass
	} else {
		Ch.CenterOfMass.Mass += S.Mass
		k := S.Mass / Ch.CenterOfMass.Mass
		Ch.CenterOfMass.PosX -= (Ch.CenterOfMass.PosX - S.PosX) * k
		Ch.CenterOfMass.PosY -= (Ch.CenterOfMass.PosY - S.PosY) * k
	}
	Ch.Stars.Add(S)
}

func (Ch *Chunk) Draw(x, y int, color pixel.RGBA) {
	PosX := 2*MaxChunkSize/ThoPowK[Ch.Stars.Depth]*float64(x) - MaxR + 2
	PosY := 2*MaxChunkSize/ThoPowK[Ch.Stars.Depth]*float64(y) - MaxR + 2
	MaxPosX := 2*MaxChunkSize/ThoPowK[Ch.Stars.Depth]*float64(x+1) - MaxR - 2
	MaxPosY := 2*MaxChunkSize/ThoPowK[Ch.Stars.Depth]*float64(y+1) - MaxR - 2
	Scr.DrawRect(PosX, PosY, MaxPosX, MaxPosY, color)
}
