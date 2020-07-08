// 局面データを文字列にしたり、文字列を局面データに復元するのに使うぜ☆（＾～＾）
package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// 現局面を xfen に変換するぜ☆（＾～＾）
func (pos *Position) toXfen() string {
	xfen := ""
	xfen += "xfen "

	// StartingBoard
	spaces := 0
	addresses := [...]uint8{7, 8, 9, 4, 5, 6, 1, 2, 3}
	for _, addr := range addresses {
		pPiece := pos.startingBoard[addr]
		if pPiece == nil {
			spaces++
		} else {
			if 0 < spaces {
				xfen += strconv.Itoa(spaces)
				spaces = 0
			}
			switch *pPiece {
			case PieceNought:
				xfen += "o"
			case PieceCross:
				xfen += "x"
			default:
				panic("Invalid piece.")
			}
		}

		if addr == 9 || addr == 6 {
			if 0 < spaces {
				xfen += strconv.Itoa(spaces)
				spaces = 0
			}
			xfen += "/"
		}
	}

	// 残ってるスペースの flush を忘れないぜ☆（＾～＾）
	if 0 < spaces {
		xfen += strconv.Itoa(spaces)
	}

	// Phase
	switch pos.friend {
	case PieceNought:
		xfen += " o"
	case PieceCross:
		xfen += " x"
	default:
		panic("Invalid pos.friend.")
	}

	// Moves
	if 0 < pos.piecesNum-pos.startingPiecesNum {
		xfen += " moves"
		for i := pos.startingPiecesNum; i < pos.piecesNum; i++ {
			xfen += fmt.Sprintf(" %d", pos.history[i])
		}
	}

	return xfen
}

// xfen を board に変換するぜ☆（＾～＾）
func fromXfen(xfen string, log Log) *Position {
	if !strings.HasPrefix(xfen, "xfen ") {
		return nil
	}

	pos := newPosition()

	// 文字数☆（＾～＾）
	starts := 0
	// 番地☆（＾～＾） 0 は未使用☆（＾～＾）
	// 7 8 9
	// 4 5 6
	// 1 2 3
	addr := 7

	const (
		/// 最初☆（＾～＾）
		MachineStateStart Piece = 1 + iota
		/// 初期局面の盤上を解析中☆（＾～＾）
		MachineStateStartingBoard
		/// 手番の解析中☆（＾～＾）
		MachineStatePhase
		/// ` moves ` 読取中☆（＾～＾）
		MachineStateMovesLabel
		/// 棋譜の解析中☆（＾～＾）
		MachineStateMoves
	)
	machineState := MachineStateStart
	// Rust言語では文字列に配列のインデックスを使ったアクセスはできないので、
	// 一手間かけるぜ☆（＾～＾）
	for i, ch := range xfen {
		switch machineState {
		case MachineStateStart:
			if i+1 == utf8.RuneCountInString("xfen ") {
				// 先頭のキーワードを読み飛ばしたら次へ☆（＾～＾）
				machineState = MachineStateStartingBoard
			}

		case MachineStateStartingBoard:
			switch ch {
			case 'x':
				// 手番の順ではないので、手番は分からないぜ☆（＾～＾）
				temp := PieceCross
				pos.startingBoard[addr] = &temp
				pos.piecesNum++
				addr++
			case 'o':
				temp := PieceNought
				pos.startingBoard[addr] = &temp
				pos.piecesNum++
				addr++
			case '1':
				addr++
			case '2':
				addr += 2
			case '3':
				addr += 3
			case '/':
				addr -= 6
			case ' ':
				// 明示的にクローン☆（＾～＾）
				for i = 0; i < int(BoardLen); i++ {
					pos.board[i] = pos.startingBoard[i]
				}
				pos.startingPiecesNum = pos.piecesNum
				machineState = MachineStatePhase
			default:
				log.println(fmt.Sprintf("Error   | xfen startingBoard error: %c", ch))
				return nil
			}
		case MachineStatePhase:
			switch ch {
			case 'x':
				pos.friend = PieceCross
			case 'o':
				pos.friend = PieceNought
			default:
				log.println(fmt.Sprintf("Error   | xfen phase error: %c", ch))
				return nil
			}
			// 一時記憶。
			starts = i
			machineState = MachineStateMovesLabel
		case MachineStateMovesLabel:
			if starts+utf8.RuneCountInString(" moves ") <= i {
				machineState = MachineStateMoves
			}
		case MachineStateMoves:
			if ch == ' ' {
			} else {
				pos.do(string(ch), log)
			}
		default:
			panic("(Err.162) Invalid state")
		}
	}

	return pos
}

// 未来へ駒を置く
// 最初は、合法手判定や勝敗判定をせずに　とりあえず動かせだぜ☆（＾～＾）
//
// # Arguments
//
// * `argStr` - コマンドラインの残り。ここでは駒を置く場所。 `1` とか `7` など。
func (pos *Position) do(argStr string, log Log) {
	addr, err := strconv.Atoi(argStr)
	if err != nil {
		log.println(fmt.Sprintf(
			"Error   | `do 数字` で入力してくれだぜ☆（＾～＾） 引数=|%s|",
			argStr))
		return
	}

	// 合法手判定☆（＾～＾）
	// 移動先のマスに駒があってはダメ☆（＾～＾）
	if addr < 1 || 9 < addr {
		log.println(fmt.Sprintf(
			"Error   | 1～9 で指定してくれだぜ☆（＾～＾） 番地=%d",
			addr))
		return
	} else if pos.board[addr] != nil {
		log.println(fmt.Sprintf(
			"Error   | 移動先のマスに駒があってはダメだぜ☆（＾～＾） 番地=%d",
			addr))
		return
	}

	pos.doMove(uint8(addr))

	/*
		// TODO 勝ち負け判定☆（*＾～＾*）
		// これは PositionHelper, WinLoseJudgment を作ってから実装しろだぜ☆（＾～＾）
		if self.is_opponent_win() {
			if let Some(result) = Position::result(GameResult::Win, Some(self.opponent())) {
				log::println(&result);
			}
		} else if self.is_draw() {
			if let Some(result) = Position::result(GameResult::Draw, None) {
				log::println(&result);
			}
		}
	*/
}

// 未来の駒を１つ戻す
func (pos *Position) undo() {
	pos.undoMove()
}
