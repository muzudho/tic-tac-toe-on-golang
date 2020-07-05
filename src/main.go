package main

import "fmt"

func main() {
	// しょっぱなにプログラムが壊れてないかテストしているぜ☆（＾～＾）
	// こんなとこに書かない方がいいが、テストを毎回するのが めんどくさいんで 実行するたびにテストさせているぜ☆（＾～＾）
	log := newLog("test.log")
	log.writeln("これはすぐあとにクリアーされるぜ☆（＾～＾）")
	log.clear()
	log.writeln("狂った街、東京！！")
	log.println("おはようさん、世界！！")

	log.println(fmt.Sprintf("Nought=|%s|", Nought))
	log.println(fmt.Sprintf("Cross =|%s|", Cross))
}
