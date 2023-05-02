package stableswap

import (
	"math/rand"
	"testing"

	"github.com/merlins-labs/merlins/furymath"
)

func BenchmarkCFMM(b *testing.B) {
	// Uses solveCfmm
	for i := 0; i < b.N; i++ {
		runCalcCFMM(solveCfmm)
	}
}

func BenchmarkBinarySearchMultiAsset(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runCalcMultiAsset(solveCFMMBinarySearchMulti)
	}
}

func runCalcCFMM(solve func(furymath.BigDec, furymath.BigDec, []furymath.BigDec, furymath.BigDec) furymath.BigDec) {
	xReserve := furymath.NewBigDec(rand.Int63n(100000) + 50000)
	yReserve := furymath.NewBigDec(rand.Int63n(100000) + 50000)
	yIn := furymath.NewBigDec(rand.Int63n(100000))
	solve(xReserve, yReserve, []furymath.BigDec{}, yIn)
}

func runCalcTwoAsset(solve func(furymath.BigDec, furymath.BigDec, furymath.BigDec) furymath.BigDec) {
	xReserve := furymath.NewBigDec(rand.Int63n(100000) + 50000)
	yReserve := furymath.NewBigDec(rand.Int63n(100000) + 50000)
	yIn := furymath.NewBigDec(rand.Int63n(100000))
	solve(xReserve, yReserve, yIn)
}

func runCalcMultiAsset(solve func(furymath.BigDec, furymath.BigDec, furymath.BigDec, furymath.BigDec) furymath.BigDec) {
	xReserve := furymath.NewBigDec(rand.Int63n(100000) + 50000)
	yReserve := furymath.NewBigDec(rand.Int63n(100000) + 50000)
	mReserve := furymath.NewBigDec(rand.Int63n(100000) + 50000)
	nReserve := furymath.NewBigDec(rand.Int63n(100000) + 50000)
	w := mReserve.Mul(mReserve).Add(nReserve.Mul(nReserve))
	yIn := furymath.NewBigDec(rand.Int63n(100000))
	solve(xReserve, yReserve, w, yIn)
}
