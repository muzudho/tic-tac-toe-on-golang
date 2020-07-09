package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	// しょっぱなにプログラムが壊れてないかテストしているぜ☆（＾～＾）
	// こんなとこに書かない方がいいが、テストを毎回するのが めんどくさいんで 実行するたびにテストさせているぜ☆（＾～＾）
	// Step 1.
	log := newLog("test.log")
	// ほんとは logrotateライブラリ使うのがいいんだが、お前らにはまだ早いしな☆ｍ９（＾～＾）！
	log.writeln("これはすぐあとにクリアーされるぜ☆（＾～＾）")
	log.clear()
	log.writeln("Hello, world!!")
	log.println("こんにちわ、世界！！")
	// こんにちわ、世界！！

	// Step 2.
	log.println(fmt.Sprintf("Nought=|%s|", PieceNought))
	// Nought=|o|
	log.println(fmt.Sprintf("Cross =|%s|", PieceCross))
	// Cross =|x|
	log.println(fmt.Sprintf("Win   =|%s|", GameResultWin))
	// Win   =|win|
	log.println(fmt.Sprintf("Draw  =|%s|", GameResultDraw))
	// Draw  =|draw|
	log.println(fmt.Sprintf("Lose  =|%s|", GameResultLose))
	// Lose  =|lose|

	pos := newPosition()
	log.println(pos.pos())
	// [Next 1 move(s) | Go o]
	//
	//         +---+---+---+
	//         |   |   |   | マスを選んでください。例 `do 7`
	//         +---+---+---+
	//         |   |   |   |    7 8 9
	//         +---+---+---+    4 5 6
	//         |   |   |   |    1 2 3
	//         +---+---+---+
	log.println(positionResult(GameResultWin, PieceNought))
	// win o

	search := newSearch(pos.friend, pos.piecesNum, true)
	log.println(fmt.Sprintf("pv=|%s|", search.pv(pos)))
	// pv=||
	log.println(searchInfoHeader(pos))
	// info nps ...... nodes ...... pv O X O X O X O X O
	// 適当な内容を入れて、入れ物として、入れた中身を見せてくれるか、チェックしろだぜ☆（＾～＾）
	log.println(search.infoForward(123, pos, 1, "Hello!"))
	// info nps      0 nodes      0 pv                   | + [1] | ->   to height 1 |       |      | + "Hello!"
	log.println(search.infoForwardLeaf(456, pos, 1, GameResultWin, "Hello!"))
	// info nps      0 nodes      0 pv                   | + [1] | .       height 0 |       | win  | + "Hello!"
	log.println(search.infoBackward(789, pos, 1, GameResultWin, "Hello!"))
	// info nps      0 nodes      0 pv                   |       | <- from height 0 | + [1] | win  | + "Hello!"

	// Step 3.
	pos.doMove(1)
	log.println(pos.pos())
	// [Next 2 move(s) | Go x]
	//
	//         +---+---+---+
	//         |   |   |   | マスを選んでください。例 `do 7`
	//         +---+---+---+
	//         |   |   |   |    7 8 9
	//         +---+---+---+    4 5 6
	//         | o |   |   |    1 2 3
	//         +---+---+---+
	pos.undoMove()
	log.println(pos.pos())
	// [Next 1 move(s) | Go o]
	//
	//         +---+---+---+
	//         |   |   |   | マスを選んでください。例 `do 7`
	//         +---+---+---+
	//         |   |   |   |    7 8 9
	//         +---+---+---+    4 5 6
	//         |   |   |   |    1 2 3
	//         +---+---+---+
	log.println(fmt.Sprintf("opponent=%s", pos.opponent()))

	// Step 4.
	p := newCommandLineParser("Go to the Moon!")
	log.println(fmt.Sprintf("Go to   =|%t|", p.startsWith("Go to")))
	// Go to   =|True|
	log.println(fmt.Sprintf("Goto    =|%t|", p.startsWith("Goto")))
	// Goto    =|False|
	log.println(fmt.Sprintf("p.starts=|%d|", p.starts))
	// p.starts=|0|
	log.println(fmt.Sprintf("p.rest  =|%s|", p.rest()))
	// p.rest  =|Go to the Moon!|
	p.goNextTo("Go to")
	log.println(fmt.Sprintf("p.starts=|%d|", p.starts))
	// p.starts=|5|
	log.println(fmt.Sprintf("p.rest  =|%s|", p.rest()))
	// p.rest  =| the Moon!|

	// Step 5.
	log.println(fmt.Sprintf("xfen=|%s|", pos.toXfen()))
	// xfen=|xfen 3/3/3 o|
	pos.do("2", log)
	log.println(pos.pos())
	// [Next 2 move(s) | Go x]
	//
	// +---+---+---+
	// |   |   |   | マスを選んでください。例 `do 7`
	// +---+---+---+
	// |   |   |   |    7 8 9
	// +---+---+---+    4 5 6
	// |   | o |   |    1 2 3
	// +---+---+---+
	xfen := "xfen xo1/xox/oxo o"
	pos = positionFromXfen(xfen, log)
	log.println(pos.pos())
	// [Next 9 move(s) | Go o]
	//
	// +---+---+---+
	// | x | o |   | マスを選んでください。例 `do 7`
	// +---+---+---+
	// | x | o | x |    7 8 9
	// +---+---+---+    4 5 6
	// | o | x | o |    1 2 3
	// +---+---+---+
	xfen = "xfen 3/3/3 x moves 1 7 4 8 9 3 6 2 5"
	pos = positionFromXfen(xfen, log)
	log.println(pos.pos())
	// win x
	// [Next 10 move(s) | Go o]
	//
	// +---+---+---+
	// | o | o | x | マスを選んでください。例 `do 7`
	// +---+---+---+
	// | x | x | x |    7 8 9
	// +---+---+---+    4 5 6
	// | x | o | o |    1 2 3
	// +---+---+---+
	pos.undo()
	log.println(pos.pos())
	// [Next 9 move(s) | Go x]
	//
	// +---+---+---+
	// | o | o | x | マスを選んでください。例 `do 7`
	// +---+---+---+
	// | x |   | x |    7 8 9
	// +---+---+---+    4 5 6
	// | x | o | o |    1 2 3
	// +---+---+---+

	// Step 6.
	// Step 7.
	xfen = "xfen o2/xox/oxo x"

	pos = positionFromXfen(xfen, log)
	if pos == nil {
		panic(fmt.Sprintf("Invalid xfen=|%s|", xfen))
	}
	log.println(fmt.Sprintf("win=|%t|", pos.isOpponentWin()))
	// win=|True|
	xfen = "xfen xox/oxo/oxo x"
	pos = positionFromXfen(xfen, log)
	if pos == nil {
		panic(fmt.Sprintf("Invalid xfen=|%s|", xfen))
	}
	log.println(fmt.Sprintf("draw=|%t|", pos.isDraw()))
	// draw=|True|

	// Step 8.
	// 探索してないんだから、 nodes も nps も 0 になるはずだよな☆（＾～＾）
	time.Sleep(time.Second * 1)
	log.println(fmt.Sprintf("nodes=%d", search.nodes))
	// nodes=0
	log.println(fmt.Sprintf("sec  =%d", search.sec()))
	// sec  =1.0
	log.println(fmt.Sprintf("nps  =%d", search.nps()))
	// nps  =0.0

	// Step 9.
	xfen = "xfen 3/3/3 o moves 1 5 2 3 7 4"
	pos = positionFromXfen(xfen, log)
	if pos == nil {
		panic(fmt.Sprintf("Invalid xfen=|%s|", xfen))
	}
	search = newSearch(pos.friend, pos.piecesNum, true)
	addr, result := search.goIt(pos, log)
	// info nps ...... nodes ...... pv O X O X O X O X O
	// info nps      1 nodes      1 pv 6                 | - [6] | ->   to height 8 |       |      | - "Search."
	// info nps      2 nodes      2 pv 6 8               | + [8] | ->   to height 9 |       |      | + "Search."
	// info nps      3 nodes      3 pv 6 8 9             | - [9] | .       height 9 |       | draw | - "It's ok."
	// info nps      3 nodes      3 pv 6 8               |       | <- from height 8 | + [9] | draw |
	// info nps      3 nodes      3 pv 6                 |       | <- from height 7 | - [8] | draw | - "Fmmm."
	// info nps      4 nodes      4 pv 6 9               | + [9] | ->   to height 9 |       |      | + "Search."
	// info nps      5 nodes      5 pv 6 9 8             | - [8] | .       height 9 |       | draw | - "It's ok."
	// info nps      5 nodes      5 pv 6 9               |       | <- from height 8 | + [8] | draw |
	// info nps      5 nodes      5 pv 6                 |       | <- from height 7 | - [9] | draw | - "Fmmm."
	// info nps      5 nodes      5 pv                   |       | <- from height 6 | + [6] | draw | + "Fmmm."
	// info nps      6 nodes      6 pv 8                 | - [8] | ->   to height 8 |       |      | - "Search."
	// info nps      7 nodes      7 pv 8 6               | + [6] | .       height 8 |       | win  | + "Hooray!"
	// info nps      7 nodes      7 pv 8                 |       | <- from height 7 | - [6] | win  |
	// info nps      7 nodes      7 pv                   |       | <- from height 6 | + [8] | lose | + "Resign."
	// info nps      8 nodes      8 pv 9                 | - [9] | ->   to height 8 |       |      | - "Search."
	// info nps      9 nodes      9 pv 9 6               | + [6] | .       height 8 |       | win  | + "Hooray!"
	// info nps      9 nodes      9 pv 9                 |       | <- from height 7 | - [6] | win  |
	// info nps      9 nodes      9 pv                   |       | <- from height 6 | + [9] | lose | + "Resign."
	log.println(fmt.Sprintf("result=|%s|", result))
	// result=|draw|
	var bestmove string
	if addr == 0 {
		bestmove = "resign"
	} else {
		bestmove = fmt.Sprintf("%d", addr)
	}
	log.println(fmt.Sprintf("bestmove=|%s|", bestmove))
	// bestmove=|6|

	// End.

	// 説明を出そうぜ☆（＾～＾）
	log.println(`きふわらべの〇×ゲーム

コマンド:
'do 7'     - 手番のプレイヤーが、 7 番地に印を付けます。
'go'       - コンピューターが次の1手を示します。
'info-off' - info出力なし。
'info-on'  - info出力あり(既定)。
'pos'      - 局面表示。
'position xfen 3/3/3 o moves 5 1 2 8 4 6 3 7 9' - 初期局面と棋譜を入力。
'undo'     - 1手戻します。
'uxi'      - 'uxiok tic-tac-toe {protocol-version}' を返します。
'xfen'     - 現局面のxfen文字列表示。
`)

	// 初期局面
	pos = newPosition()
	infoEnable := true

	// [Ctrl]+[C] でループを終了。ログ・ファイルを閉じないのが気になるが……☆（＾～＾）
	for {
		line := ""
		// まず最初に、コマンドライン入力を待機しろだぜ☆（＾～＾）
		var scanner = bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			line = scanner.Text()
		}

		// コマンドライン☆（＾～＾） p は parser の意味で使ってるぜ☆（＾～＾）
		p := newCommandLineParser(line)

		// 本当は よく使うコマンド順に並べた方が高速だが、先に見つけた方が選ばれるので後ろの方を漏らしやすくて むずかしいし、
		// だから、アルファベット順に並べた方が見やすいぜ☆（＾～＾）
		if p.startsWith("do") {
			p.goNextTo("do ")
			rest := p.rest()
			if rest != "" {
				pos.do(rest, log)
			}
		} else if p.startsWith("go") {
			search = newSearch(pos.friend, pos.piecesNum, infoEnable)
			addr, result = search.goIt(pos, log)
			if addr == 0 {
				log.println("resign")
			} else {
				log.println(fmt.Sprintf("info result=%s nps=%d", result, search.nps()))
				log.println(fmt.Sprintf("bestmove %d", addr))
			}
		} else if p.startsWith("info-off") {
			infoEnable = false
		} else if p.startsWith("info-on") {
			infoEnable = true
		} else if p.startsWith("quit") {
			log.close()
			return
		} else if p.startsWith("position") {
			p.goNextTo("position ")
			rest := p.rest()
			if rest != "" {
				temp := positionFromXfen(rest, log)
				if temp != nil {
					pos = temp
				}
			}
		} else if p.startsWith("pos") {
			log.println(pos.pos())
		} else if p.startsWith("undo") {
			pos.undo()
		} else if p.startsWith("uxi") {
			log.println("uxiok tic-tac-toe v20200704.0.0")
		} else if p.startsWith("xfen") {
			log.println(fmt.Sprintf("%s", pos.toXfen()))
		} else {
			log.println(fmt.Sprintf("Debug   | Invalid command=|%s|", line))
		}
	}
}
