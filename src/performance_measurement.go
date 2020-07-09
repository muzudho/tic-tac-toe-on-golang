package main

import "time"

func (search *Search) sec() uint64 {
	t := time.Now()
	return uint64(t.Sub(search.stopwatch).Seconds())
}

// Node per second.
func (search *Search) nps() uint64 {
	sec := search.sec()
	if 0 < sec {
		return uint64(search.nodes) / sec
	}
	// 1秒未満で全部探索してしまった☆（＾～＾） 本当は もっと多いと思うんだが☆（＾～＾）
	return uint64(search.nodes)
}
