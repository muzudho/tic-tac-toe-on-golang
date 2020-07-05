package main

import (
	"fmt"
	"os"
)

func main() {
	write("おはよう、世界！！\r\n")
}

func write(contents string) {
	f, err := os.Create("test.log")
	if err != nil {
		// ファイルの作成に失敗しても、続行優先で無視しようぜ☆（＾～＾）
		// 大会で 標準入力にエラーメッセージを出すと、大会サーバーへ不要なメッセージを送信することがあるからな☆（＾～＾）
		return
	}
	// ログの書込みに失敗しても、続行優先で無視しようぜ☆（＾～＾）
	f.WriteString(contents)
	// ファイルを閉じるのに失敗しても、続行優先で無視しようぜ☆（＾～＾）
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
