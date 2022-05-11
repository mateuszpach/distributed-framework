package test

import (
	"github.com/krzysztof-turowski/distributed-framework/consensus/sync_ben_or"
	"github.com/krzysztof-turowski/distributed-framework/consensus/sync_phase_king"
	"github.com/krzysztof-turowski/distributed-framework/consensus/sync_single_bit"
	"github.com/krzysztof-turowski/distributed-framework/lib"
	"io/ioutil"
	"log"
	"testing"
)

func TestBenOr(t *testing.T) {
	checkLogOutput()
	for iteration := 0; iteration < 50; iteration++ {
		v, s := lib.BuildCompleteGraphWithLoops(11, true, lib.GetGenerator())
		sync_ben_or.Run(v, s, 2, []int{0, 1, 1, 0, 1, 0, 1, 0, 1, 1, 0}, sync_ben_or.EachMessageRandom(11, 2))
	}
}

func BenchmarkBenOr(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	for iteration := 0; iteration < b.N; iteration++ {
		v, s := lib.BuildCompleteGraphWithLoops(11, true, lib.GetGenerator())
		sync_ben_or.Run(v, s, 2, []int{0, 1, 1, 0, 1, 0, 1, 0, 1, 1, 0}, sync_ben_or.EachMessageRandom(11, 2))
	}
}

func TestPhaseKing(t *testing.T) {
	checkLogOutput()
	for iteration := 0; iteration < 50; iteration++ {
		v, s := lib.BuildCompleteGraphWithLoops(10, true, lib.GetGenerator())
		sync_phase_king.Run(v, s, 3, []int{0, 1, 1, 0, 1, 0, 1, 0, 1, 1}, sync_phase_king.EachMessageRandom(10, 3))
	}
}

func BenchmarkPhaseKing(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	for iteration := 0; iteration < b.N; iteration++ {
		v, s := lib.BuildCompleteGraphWithLoops(11, true, lib.GetGenerator())
		sync_phase_king.Run(v, s, 2, []int{0, 1, 1, 0, 1, 0, 1, 0, 1, 1, 0}, sync_phase_king.EachMessageRandom(11, 2))
	}
}

func TestSingleBit(t *testing.T) {
	checkLogOutput()
	for iteration := 0; iteration < 50; iteration++ {
		v, s := lib.BuildCompleteGraphWithLoops(9, true, lib.GetGenerator())
		sync_single_bit.Run(v, s, 2, []int{0, 1, 1, 0, 1, 0, 1, 0, 1}, sync_single_bit.EachMessageRandom(9, 2))
	}
}

func BenchmarkSingleBit(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	for iteration := 0; iteration < b.N; iteration++ {
		v, s := lib.BuildCompleteGraphWithLoops(11, true, lib.GetGenerator())
		sync_single_bit.Run(v, s, 2, []int{0, 1, 1, 0, 1, 0, 1, 0, 1, 1, 0}, sync_single_bit.EachMessageRandom(11, 2))
	}
}