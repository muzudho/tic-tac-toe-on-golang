// 勝敗判定☆（＾～＾）

package main

// 石を置いてから 勝敗判定をするので、
// 反対側の手番のやつが 石を３つ並べてたかどうかを調べるんだぜ☆（＾～＾）
func (pos *Position) isOpponentWin() bool {
	// 8パターンしかないので、全部チェックしてしまおうぜ☆（＾～＾）

	equals := func(expected Piece, actual *Piece) bool {
		return actual != nil && expected == *actual
	}

	opponent := pos.opponent()

	// xxx
	// ...
	// ...
	return (equals(opponent, pos.board[7]) && equals(opponent, pos.board[8]) && equals(opponent, pos.board[9])) || (
	// ...
	// xxx
	// ...
	equals(opponent, pos.board[4]) && equals(opponent, pos.board[5]) && equals(opponent, pos.board[6])) || (
	// ...
	// ...
	// xxx
	equals(opponent, pos.board[1]) && equals(opponent, pos.board[2]) && equals(opponent, pos.board[3])) || (
	// x..
	// x..
	// x..
	equals(opponent, pos.board[7]) && equals(opponent, pos.board[4]) && equals(opponent, pos.board[1])) || (
	// .x.
	// .x.
	// .x.
	equals(opponent, pos.board[8]) && equals(opponent, pos.board[5]) && equals(opponent, pos.board[2])) || (
	// ..x
	// ..x
	// ..x
	equals(opponent, pos.board[9]) && equals(opponent, pos.board[6]) && equals(opponent, pos.board[3])) || (
	// x..
	// .x.
	// ..x
	equals(opponent, pos.board[7]) && equals(opponent, pos.board[5]) && equals(opponent, pos.board[3])) || (
	// ..x
	// .x.
	// x..
	equals(opponent, pos.board[9]) && equals(opponent, pos.board[5]) && equals(opponent, pos.board[1]))
}

/// 石を置いてから 引き分け判定をするので、
/// 反対側の手番のやつが 勝ってなくて、
/// かつ、全てのマスが埋まってたら引き分けだぜ☆（＾～＾）
func (pos *Position) isDraw() bool {
	if pos.isOpponentWin() {
		return false
	}
	for addr := 1; addr < int(BoardLen); addr++ {
		if pos.board[addr] == nil {
			return false
		}
	}

	return true
}
