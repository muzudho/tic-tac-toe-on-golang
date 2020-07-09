package main

import (
	"fmt"
	"os"
)

// Log is print or write to a file.
type Log struct {
	file string
	// 追加書込み用☆（＾～＾）
	fileAppendHandle *os.File
}

func newLog(file string) *Log {
	p := Log{file: file, fileAppendHandle: nil}
	return &p
}

func (log *Log) writeln(contents string) {
	if log.fileAppendHandle == nil {
		f, err := os.OpenFile(log.file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			// ファイルの作成に失敗しても、続行優先で無視しようぜ☆（＾～＾）
			// 大会で 標準入力にエラーメッセージを出すと、大会サーバーへ不要なメッセージを送信することがあるからな☆（＾～＾）
			return
		}
		log.fileAppendHandle = f
	}
	// ログの書込みに失敗しても、続行優先で無視しようぜ☆（＾～＾）
	log.fileAppendHandle.WriteString(contents + "\r\n")
	/*
		// ファイルを閉じるのに失敗しても、続行優先で無視しようぜ☆（＾～＾）
		err = f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	*/
}

func (log *Log) close() {
	// ファイルを閉じるのに失敗しても、続行優先で無視しようぜ☆（＾～＾）
	err := log.fileAppendHandle.Close()
	if err != nil {
		// 大会で 標準入力にエラーメッセージを出すと、大会サーバーへ不要なメッセージを送信することがあるからな☆（＾～＾）
		// fmt.Println(err)
		return
	}
	log.fileAppendHandle = nil
}

func (log *Log) clear() {
	// ファイル・ハンドルをつかんでいたら、放そうぜ☆（＾～＾）？
	if log.fileAppendHandle != nil {
		log.close()
	}

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
