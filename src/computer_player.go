// サーチ☆（＾～＾）探索部とか言われてるやつだぜ☆（＾～＾）

package main

// 最善の番地を返すぜ☆（＾～＾）
func (search *Search) goIt(pos *Position, log *Log) (uint8, GameResult) {
	if search.infoEnable {
		log.println(searchInfoHeader(pos))
	}
	return search.node(pos, log)
}

/// 手番が来たぜ☆（＾～＾）いわゆる search だぜ☆（＾～＾）
func (search *Search) node(pos *Position, log *Log) (uint8, GameResult) {
	// 投了は 0 で☆（＾～＾）
	var bestAddr uint8
	bestAddr = 0
	bestResult := GameResultLose

	var addr uint8
	for addr = 1; addr < BoardLen; addr++ {
		// 空きマスがあれば
		if pos.board[addr] == PieceNone {
			// とりあえず置いてみようぜ☆（＾～＾）
			pos.doMove(addr)
			search.nodes++

			// 前向き探索というのは、葉っぱの方に進んでるとき☆（＾～＾）
			// 後ろ向き探索というのは、根っこの方に戻ってるとき☆（＾～＾）
			//
			// 勝ったかどうか判定しようぜ☆（＾～＾）？
			if pos.isOpponentWin() {
				// 勝ったなら☆（＾～＾）
				// 前向き探索情報を出して、置いた石は戻して、後ろ向き探索情報を出して、探索終了だぜ☆（＾～＾）
				if search.infoEnable {
					log.println(search.infoForwardLeaf(
						search.nps(),
						pos,
						addr,
						GameResultWin,
						"Hooray!",
					))
				}
				pos.undoMove()
				if search.infoEnable {
					log.println(search.infoBackward(
						search.nps(),
						pos,
						addr,
						GameResultWin,
						"",
					))
				}
				return addr, GameResultWin
			} else if SquaresNum <= pos.piecesNum {
				// 勝っていなくて、深さ上限に達したら、〇×ゲームでは 他に置く場所もないから引き分け確定だぜ☆（＾～＾）
				// 前向き探索情報を出して、置いた石は戻して、後ろ向き探索情報を出して、探索終了だぜ☆（＾～＾）
				if search.infoEnable {
					log.println(search.infoForwardLeaf(
						search.nps(),
						pos,
						addr,
						GameResultDraw,
						"It's ok.",
					))
				}
				pos.undoMove()
				if search.infoEnable {
					log.println(search.infoBackward(
						search.nps(),
						pos,
						addr,
						GameResultDraw,
						"",
					))
				}
				return addr, GameResultDraw
			} else {
				// まだ続きがあるぜ☆（＾～＾）
				if search.infoEnable {
					log.println(search.infoForward(search.nps(), pos, addr, "Search."))
				}
			}

			// 相手の番だぜ☆（＾～＾）
			_, opponentGameResult := search.node(pos, log)

			// 自分が置いたところを戻そうぜ☆（＾～＾）？
			pos.undoMove()

			switch opponentGameResult {
			// 相手の負けなら、この手で勝ちだぜ☆（＾～＾）後ろ向き探索情報を出して、探索終わり☆（＾～＾）
			case GameResultLose:
				if search.infoEnable {
					log.println(search.infoBackward(
						search.nps(),
						pos,
						addr,
						GameResultWin,
						"Ok.",
					))
				}
				return addr, GameResultWin
			// 勝ち負けがずっと見えてないなら☆（＾～＾）後ろ向き探索情報を出して、探索を続けるぜ☆（＾～＾）
			case GameResultDraw:
				if search.infoEnable {
					log.println(search.infoBackward(
						search.nps(),
						pos,
						addr,
						GameResultDraw,
						"Fmmm.",
					))
				}
				switch bestResult {
				case GameResultLose:
					// 更新
					bestAddr = addr
					bestResult = GameResultDraw
				default:
				}

			// 相手が勝つ手を選んではダメだぜ☆（＾～＾）後ろ向き探索情報を出して、探索を続けるぜ☆（＾～＾）
			case GameResultWin:
				if search.infoEnable {
					log.println(search.infoBackward(
						search.nps(),
						pos,
						addr,
						GameResultLose,
						"Resign.",
					))
				}
			}
		}
	}

	return bestAddr, bestResult
}
