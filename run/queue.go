package run

type StarQueue struct {
	Array []*Star
	Depth int
}

func (Q *StarQueue) Add(S *Star) {
	S.PosInQueue[Q.Depth] = len(Q.Array)
	Q.Array = append(Q.Array, S)
}

func (Q *StarQueue) Del(i int) {
	l := len(Q.Array) - 1
	Q.Array[i].PosInQueue[Q.Depth] = -1
	Q.Array[i] = Q.Array[l]
	Q.Array = Q.Array[:l]
}

func (Q *StarQueue) DelStar(S *Star) {
	Q.Del(S.PosInQueue[Q.Depth])
}
