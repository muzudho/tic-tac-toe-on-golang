package main

import (
	"fmt"
	"os"
)

type log struct {
	file string
}

func newLog(file string) *log {
	p := log{file: file}
	return &p
}

func (self *log) writeln(contents string) {
	f, err := os.OpenFile(self.file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// ファイルの作成に失敗しても、続行優先で無視しようぜ☆（＾～＾）
		// 大会で 標準入力にエラーメッセージを出すと、大会サーバーへ不要なメッセージを送信することがあるからな☆（＾～＾）
		return
	}
	// ログの書込みに失敗しても、続行優先で無視しようぜ☆（＾～＾）
	f.WriteString(contents + "\r\n")
	// ファイルを閉じるのに失敗しても、続行優先で無視しようぜ☆（＾～＾）
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (self *log) clear() {
	f, err := os.Create(self.file)
	if err != nil {
		// ファイルの作成に失敗しても、続行優先で無視しようぜ☆（＾～＾）
		// 大会で 標準入力にエラーメッセージを出すと、大会サーバーへ不要なメッセージを送信することがあるからな☆（＾～＾）
		return
	}
	// ログの書込みに失敗しても、続行優先で無視しようぜ☆（＾～＾）
	f.WriteString("")
	// ファイルを閉じるのに失敗しても、続行優先で無視しようぜ☆（＾～＾）
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (self *log) println(contents string) {
	fmt.Println(contents)
	self.writeln(contents)
}
