package main

import (
	"fmt"
	"os"
)

// Log is print or write to a file.
type Log struct {
	file string
}

func newLog(file string) *Log {
	p := Log{file: file}
	return &p
}

func (log *Log) writeln(contents string) {
	f, err := os.OpenFile(log.file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
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

func (log *Log) clear() {
	f, err := os.Create(log.file)
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

func (log *Log) println(contents string) {
	fmt.Println(contents)
	log.writeln(contents)
}
