package main

import "fmt"

// １手指すぜ☆（＾～＾）
func (pos *Position) doMove(addr uint8) {
	// 石を置くぜ☆（＾～＾）
	pos.board[addr] = pos.friend
	// 棋譜に記すぜ☆（＾～＾）
	pos.history[pos.piecesNum] = addr
	// 棋譜に記した後にカウンターを増やすぜ☆（＾～＾）
	pos.piecesNum++
	// 手番は交代だぜ☆（＾～＾）
	pos.friend = pos.opponent()
}

// １手戻すぜ☆（＾～＾）
func (pos *Position) undoMove() {
	// 手番は交代だぜ☆（＾～＾）
	pos.friend = pos.opponent()
	// 手数は次の要素を指しているんで、先に戻してから、配列の中身を取り出せだぜ☆（＾～＾）
	pos.piecesNum--
	// 置いたところの石は削除な☆（＾～＾）
	addr := pos.history[pos.piecesNum]
	pos.board[addr] = PieceNone
}

// 相手番☆（＾～＾）
func (pos *Position) opponent() Piece {
	switch pos.friend {
	case PieceNought:
		return PieceCross
	case PieceCross:
		return PieceNought
	default:
		panic(fmt.Sprintf("Invalid friend=|%s|", pos.friend))
	}
}
