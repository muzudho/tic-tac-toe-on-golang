package main

func main() {
	// しょっぱなにプログラムが壊れてないかテストしているぜ☆（＾～＾）
	// こんなとこに書かない方がいいが、テストを毎回するのが めんどくさいんで 実行するたびにテストさせているぜ☆（＾～＾）
	log := newLog("test.log")
	log.writeln("これはすぐあとにクリアーされるぜ☆（＾～＾）")
	log.clear()
	log.writeln("おはようさん、世界！！")
	log.println("狂った街、東京！！")
}
