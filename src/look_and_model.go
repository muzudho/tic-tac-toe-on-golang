package main

import (
	"fmt"
	"strings"
	"time"
)

// 駒とか、石とかのことだが、〇×は 何なんだろうな、マーク☆（＾～＾）？
type Piece int

const (
	// 〇
	Nought Piece = 1 + iota
	// ×
	Cross
)

var pieces = [...]string{
	"o",
	"x",
}

func (self Piece) String() string { return pieces[self-1] }

// 〇×ゲームは完全解析できるから、評価ではなくて、ゲームの結果が分かるんだよな☆（＾～＾）
type GameResult int

const (
	Win GameResult = 1 + iota
	Draw
	Lose
)

var game_results = [...]string{
	"win",
	"draw",
	"lose",
}

func (self GameResult) String() string { return game_results[self-1] }

// 1スタートで9まで☆（＾～＾） 配列には0番地もあるから、要素数は10だぜ☆（＾～＾）
const BOARD_LEN uint8 = 10

// 盤上に置ける最大の駒数だぜ☆（＾～＾） ９マスしか置くとこないから９だぜ☆（＾～＾）
const SQUARES_NUM uint8 = 9

// 局面☆（＾～＾）ゲームデータをセーブしたり、ロードしたりするときの保存されてる現状だぜ☆（＾～＾）
type Position struct {
	// 次に盤に置く駒☆（＾～＾）
	// 英語では 手番は your turn, 相手版は your opponent's turn なんで 手番という英語は無い☆（＾～＾）
	// 自分という意味の単語はプログラム用語と被りまくるんで、
	// あまり被らない 味方(friend) を手番の意味で たまたま使ってるだけだぜ☆（＾～＾）
	friend Piece

	// 開始局面の盤の各マス☆（＾～＾） [0] は未使用☆（＾～＾）
	starting_board [BOARD_LEN]*Piece
	// 盤の上に最初から駒が何個置いてあったかだぜ☆（＾～＾）
	starting_pieces_num uint8

	// 現状の盤の各マス☆（＾～＾） [0] は未使用☆（＾～＾）
	board [BOARD_LEN]*Piece

	// 棋譜だぜ☆（＾～＾）駒を置いた番地を並べてけだぜ☆（＾～＾）
	history [SQUARES_NUM]uint8

	// 盤の上に駒が何個置いてあるかだぜ☆（＾～＾）
	pieces_num uint8
}

func newPosition() *Position {
	p := Position{
		friend:              Nought,
		starting_board:      [...]*Piece{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
		starting_pieces_num: 0,
		board:               [...]*Piece{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
		history:             [SQUARES_NUM]uint8{0},
		pieces_num:          0,
	}
	return &p
}

func (self *Position) cell(index uint8) string {
	if self.board[index] != nil {
		return fmt.Sprintf(" %s ", self.board[index])
	} else {
		return "   "
	}
}

func (self *Position) pos() string {
	s := fmt.Sprintf(
		`[Next %d move(s) | Go %s]
`,
		self.pieces_num+1,
		self.friend)
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
		self.cell(7),
		self.cell(8),
		self.cell(9),
		self.cell(4),
		self.cell(5),
		self.cell(6),
		self.cell(1),
		self.cell(2),
		self.cell(3))
	return s
}

func position_result(result GameResult, winner Piece) string {
	if result == Win {
		return fmt.Sprintf("win %s", winner)
	} else if result == Draw {
		return "draw"
	} else {
		return ""
	}
}

type Search struct {
	/// 探索部☆（＾～＾）
	/// この探索を始めたのはどっち側か☆（＾～＾）
	start_friend Piece
	/// この探索を始めたときに石はいくつ置いてあったか☆（＾～＾）
	start_pieces_num uint8
	/// 探索した状態ノード数☆（＾～＾）
	nodes uint32
	/// この構造体を生成した時点からストップ・ウォッチを開始するぜ☆（＾～＾）
	stopwatch time.Time
	/// info の出力の有無。
	info_enable bool
}

// 初期値だぜ☆（＾～＾）
func newSearch(friend Piece, start_pieces_num uint8, info_enable bool) *Search {
	p := Search{
		start_friend:     friend,
		start_pieces_num: start_pieces_num,
		nodes:            0,
		stopwatch:        time.Now(),
		info_enable:      info_enable,
	}
	return &p
}

// Principal variation. 今読んでる読み筋☆（＾～＾）
func (self *Search) pv(pos *Position) string {
	pv := ""
	for i := self.start_pieces_num; i < pos.pieces_num; i++ {
		pv += fmt.Sprintf("%d ", pos.history[i])
	}
	return strings.TrimRight(pv, "")
}

func search_info_header(pos *Position) string {
	switch pos.friend {
	case Nought:
		return "info nps ...... nodes ...... pv O X O X O X O X O"
	case Cross:
		return "info nps ...... nodes ...... pv X O X O X O X O X"
	default:
		panic(fmt.Sprintf("Invalid friend=|%s|", pos.friend))
	}
}

/// 前向き探索中だぜ☆（＾～＾）
func (self *Search) info_forward(nps uint64, pos *Position, addr uint8, comment string) string {
	var friend_str string
	if pos.friend == self.start_friend {
		friend_str = "+"
	} else {
		friend_str = "-"
	}

	var height string
	if SQUARES_NUM < pos.pieces_num+1 {
		height = "none    "
	} else {
		height = fmt.Sprintf("height %d", pos.pieces_num+1)
	}

	var comment_str string
	if comment != "" {
		comment_str = fmt.Sprintf(" %s \"%s\"", friend_str, comment)
	} else {
		comment_str = ""
	}

	return fmt.Sprintf("info nps %6d nodes %6d pv %-17s | %s [%d] | ->   to %s |       |      |%s",
		nps,
		self.nodes,
		self.pv(pos),
		friend_str,
		addr,
		height,
		comment_str)
}

/*
impl Search {


    /// 前向き探索で葉に着いたぜ☆（＾～＾）
    pub fn info_forward_leaf(
        &self,
        pos: &mut Position,
        addr: usize,
        result: GameResult,
        comment: Option<String>,
    ) -> String {
        let friend_str = if pos.friend == self.start_friend {
            "+".to_string()
        } else {
            "-".to_string()
        };
        format!(
            "info nps {: >6} nodes {: >6} pv {: <17} | {} [{}] | .       {} |       | {:4} |{}",
            self.nps(),
            self.nodes,
            self.pv(pos),
            friend_str,
            addr,
            if SQUARES_NUM < pos.pieces_num {
                "none    ".to_string()
            } else {
                format!("height {}", pos.pieces_num)
            },
            result.to_string(),
            if let Some(comment) = comment {
                format!(" {} \"{}\"", friend_str, comment)
            } else {
                "".to_string()
            },
        )
        .to_string()
    }
    /// 後ろ向き探索のときの表示だぜ☆（＾～＾）
    pub fn info_backward(
        &self,
        pos: &mut Position,
        addr: usize,
        result: GameResult,
        comment: Option<String>,
    ) -> String {
        let friend_str = if pos.friend == self.start_friend {
            "+".to_string()
        } else {
            "-".to_string()
        };
        return format!(
            "info nps {: >6} nodes {: >6} pv {: <17} |       | <- from {} | {} [{}] | {:4} |{}",
            self.nps(),
            self.nodes,
            self.pv(pos),
            if SQUARES_NUM < pos.pieces_num + 1 {
                "none    ".to_string()
            } else {
                format!("height {}", pos.pieces_num + 1)
            },
            friend_str,
            addr,
            result.to_string(),
            if let Some(comment) = comment {
                format!(" {} \"{}\"", friend_str, comment)
            } else {
                "".to_string()
            }
        )
        .to_string();
    }
}
*/
