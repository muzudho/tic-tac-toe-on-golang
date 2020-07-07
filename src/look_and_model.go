package main

import (
	"fmt"
	"strings"
	"time"
)

// Piece 駒とか、石とかのことだが、〇×は 何なんだろうな、マーク☆（＾～＾）？
type Piece int

const (
	// PieceNought 〇
	PieceNought Piece = 1 + iota
	// PieceCross ×
	PieceCross
)

var pieces = [...]string{
	"o",
	"x",
}

func (piece Piece) String() string { return pieces[piece-1] }

// GameResult 〇×ゲームは完全解析できるから、評価ではなくて、ゲームの結果が分かるんだよな☆（＾～＾）
type GameResult int

const (
	// GameResultWin is 勝ち☆（＾～＾）
	GameResultWin GameResult = 1 + iota
	// GameResultDraw is 引き分け☆（＾～＾）
	GameResultDraw
	// GameResultLose is 負け☆（＾～＾）
	GameResultLose
)

var gameResults = [...]string{
	"win",
	"draw",
	"lose",
}

func (result GameResult) String() string { return gameResults[result-1] }

// BoardLen 1スタートで9まで☆（＾～＾） 配列には0番地もあるから、要素数は10だぜ☆（＾～＾）
const BoardLen uint8 = 10

// SquaresNum 盤上に置ける最大の駒数だぜ☆（＾～＾） ９マスしか置くとこないから９だぜ☆（＾～＾）
const SquaresNum uint8 = 9

// Position 局面☆（＾～＾）ゲームデータをセーブしたり、ロードしたりするときの保存されてる現状だぜ☆（＾～＾）
type Position struct {
	// 次に盤に置く駒☆（＾～＾）
	// 英語では 手番は your turn, 相手版は your opponent's turn なんで 手番という英語は無い☆（＾～＾）
	// 自分という意味の単語はプログラム用語と被りまくるんで、
	// あまり被らない 味方(friend) を手番の意味で たまたま使ってるだけだぜ☆（＾～＾）
	friend Piece

	// 開始局面の盤の各マス☆（＾～＾） [0] は未使用☆（＾～＾）
	startingBoard [BoardLen]*Piece
	// 盤の上に最初から駒が何個置いてあったかだぜ☆（＾～＾）
	startingPiecesNum uint8

	// 現状の盤の各マス☆（＾～＾） [0] は未使用☆（＾～＾）
	board [BoardLen]*Piece

	// 棋譜だぜ☆（＾～＾）駒を置いた番地を並べてけだぜ☆（＾～＾）
	history [SquaresNum]uint8

	// 盤の上に駒が何個置いてあるかだぜ☆（＾～＾）
	piecesNum uint8
}

func newPosition() *Position {
	p := Position{
		friend:            PieceNought,
		startingBoard:     [...]*Piece{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
		startingPiecesNum: 0,
		board:             [...]*Piece{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
		history:           [SquaresNum]uint8{0},
		piecesNum:         0,
	}
	return &p
}

func (pos *Position) cell(index uint8) string {
	if pos.board[index] != nil {
		return fmt.Sprintf(" %s ", pos.board[index])
	}
	return "   "
}

func (pos *Position) pos() string {
	s := fmt.Sprintf(
		`[Next %d move(s) | Go %s]
`,
		pos.piecesNum+1,
		pos.friend)
	// 書式指定子は cell関数の方に任せるぜ☆（＾～＾）
	s += fmt.Sprintf(`
+---+---+---+
|%1s|%1s|%1s| マスを選んでください。例 'do 7'
+---+---+---+
|%1s|%1s|%1s|    7 8 9
+---+---+---+    4 5 6
|%1s|%1s|%1s|    1 2 3
+---+---+---+
`,
		pos.cell(7),
		pos.cell(8),
		pos.cell(9),
		pos.cell(4),
		pos.cell(5),
		pos.cell(6),
		pos.cell(1),
		pos.cell(2),
		pos.cell(3))
	return s
}

func positionResult(result GameResult, winner Piece) string {
	if result == GameResultWin {
		return fmt.Sprintf("win %s", winner)
	} else if result == GameResultDraw {
		return "draw"
	} else {
		return ""
	}
}

// Search は探索部だぜ☆（＾～＾）
type Search struct {
	/// 探索部☆（＾～＾）
	/// この探索を始めたのはどっち側か☆（＾～＾）
	startFriend Piece
	/// この探索を始めたときに石はいくつ置いてあったか☆（＾～＾）
	startPiecesNum uint8
	/// 探索した状態ノード数☆（＾～＾）
	nodes uint32
	/// この構造体を生成した時点からストップ・ウォッチを開始するぜ☆（＾～＾）
	stopwatch time.Time
	/// info の出力の有無。
	infoEnable bool
}

// 初期値だぜ☆（＾～＾）
func newSearch(friend Piece, startPiecesNum uint8, infoEnable bool) *Search {
	p := Search{
		startFriend:    friend,
		startPiecesNum: startPiecesNum,
		nodes:          0,
		stopwatch:      time.Now(),
		infoEnable:     infoEnable,
	}
	return &p
}

// Principal variation. 今読んでる読み筋☆（＾～＾）
func (search *Search) pv(pos *Position) string {
	pv := ""
	for i := search.startPiecesNum; i < pos.piecesNum; i++ {
		pv += fmt.Sprintf("%d ", pos.history[i])
	}
	return strings.TrimRight(pv, "")
}

// 見出しだぜ☆（＾～＾）
func searchInfoHeader(pos *Position) string {
	switch pos.friend {
	case PieceNought:
		return "info nps ...... nodes ...... pv O X O X O X O X O"
	case PieceCross:
		return "info nps ...... nodes ...... pv X O X O X O X O X"
	default:
		panic(fmt.Sprintf("Invalid friend=|%s|", pos.friend))
	}
}

// 前向き探索中だぜ☆（＾～＾）
func (search *Search) infoForward(nps uint64, pos *Position, addr uint8, comment string) string {
	var friendStr string
	if pos.friend == search.startFriend {
		friendStr = "+"
	} else {
		friendStr = "-"
	}

	var height string
	if SquaresNum < pos.piecesNum+1 {
		height = "none    "
	} else {
		height = fmt.Sprintf("height %d", pos.piecesNum+1)
	}

	var commentStr string
	if comment != "" {
		commentStr = fmt.Sprintf(" %s \"%s\"", friendStr, comment)
	} else {
		commentStr = ""
	}

	return fmt.Sprintf("info nps %6d nodes %6d pv %-17s | %s [%d] | ->   to %s |       |      |%s",
		nps,
		search.nodes,
		search.pv(pos),
		friendStr,
		addr,
		height,
		commentStr)
}

// 前向き探索で葉に着いたぜ☆（＾～＾）
func (search *Search) infoForwardLeaf(
	nps uint64,
	pos *Position,
	addr uint8,
	result GameResult,
	comment string,
) string {
	var friendStr string
	if pos.friend == search.startFriend {
		friendStr = "+"
	} else {
		friendStr = "-"
	}

	var height string
	if SquaresNum < pos.piecesNum {
		height = "none    "
	} else {
		height = fmt.Sprintf("height %d", pos.piecesNum)
	}

	var commentStr string
	if comment != "" {
		commentStr = fmt.Sprintf(" %s \"%s\"", friendStr, comment)
	} else {
		commentStr = ""
	}

	return fmt.Sprintf(
		"info nps %6d nodes %6d pv %-17s | %s [%d] | .       %s |       | %4s |%s",
		nps,
		search.nodes,
		search.pv(pos),
		friendStr,
		addr,
		height,
		result,
		commentStr,
	)
}

// 後ろ向き探索のときの表示だぜ☆（＾～＾）
func (search *Search) infoBackward(
	nps uint64,
	pos *Position,
	addr uint8,
	result GameResult,
	comment string,
) string {
	var friendStr string
	if pos.friend == search.startFriend {
		friendStr = "+"
	} else {
		friendStr = "-"
	}

	var height string
	if SquaresNum < pos.piecesNum+1 {
		height = "none    "
	} else {
		height = fmt.Sprintf("height %d", pos.piecesNum+1)
	}

	var commentStr string
	if comment != "" {
		commentStr = fmt.Sprintf(" %s \"%s\"", friendStr, comment)
	} else {
		commentStr = ""
	}

	return fmt.Sprintf(
		"info nps %6d nodes %6d pv %-17s |       | <- from %s | %s [%d] | %4s |%s",
		nps,
		search.nodes,
		search.pv(pos),
		height,
		friendStr,
		addr,
		result,
		commentStr)
}
