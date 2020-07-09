package main

import (
	"fmt"
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
	pos.do("2", *log)
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
	pos = positionFromXfen(xfen, *log)
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
	pos = positionFromXfen(xfen, *log)
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

	pos = positionFromXfen(xfen, *log)
	if pos == nil {
		panic(fmt.Sprintf("Invalid xfen=|%s|", xfen))
	}
	log.println(fmt.Sprintf("win=|%t|", pos.isOpponentWin()))
	// win=|True|
	xfen = "xfen xox/oxo/oxo x"
	pos = positionFromXfen(xfen, *log)
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

	/*
		// TODO 標準入力の練習☆（＾～＾）
		var s string
		print("スキャンするぜ☆（＾～＾）")
		var scanner = bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			print("スキャンしたか☆（＾～＾）？")
			s = scanner.Text()
		}
		print("どうだぜ☆（＾～＾）？")
		// 文字数を調べようぜ☆（＾～＾）
		len := utf8.RuneCountInString(s)
	*/
}
