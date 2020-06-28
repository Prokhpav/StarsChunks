package run

import "math"

func GetRadius(Mass float64) float64 {
	return math.Sqrt(Mass)
}

func NewStar(PosX, PosY, SpeedX, SpeedY, Mass float64) {
	S := &Star{
		MassPoint: MassPoint{
			PosX: PosX,
			PosY: PosY,
			Mass: Mass,
		},
		SpeedX:     SpeedX,
		SpeedY:     SpeedY,
		Radius:     GetRadius(Mass),
		PosInQueue: make([]int, NumOfSlices+1),
	}
	Stars.Add(S)

	PpR2X := PosX + MaxR
	PpR2Y := PosX + MaxR
	ChunkSize := MaxChunkSize
	for i := 0; i < len(Chunks); i++ {
		Chunks[i][int(PpR2X/ChunkSize)][int(PpR2Y/ChunkSize)].AddStar(S)
		ChunkSize /= 2
	}
}
