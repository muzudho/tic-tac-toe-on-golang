package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// CommandLineParser 入力されたコマンドを、読み取る手伝いをするぜ☆（＾～＾）
type CommandLineParser struct {
	line   string
	len    int
	starts int
}

// 末尾の改行を除こうぜ☆（＾～＾）
// trim すると空白も消えるぜ☆（＾～＾）
func newCommandLineParser(line string) *CommandLineParser {
	// TODO これが要るのか分からないが、文字列の末端の改行を削除しようぜ☆（＾～＾）
	line = strings.TrimRight(line, "\r\n")
	// 文字数を調べようぜ☆（＾～＾）
	len := utf8.RuneCountInString(line)
	return &CommandLineParser{
		line:   line,
		len:    len,
		starts: 0,
	}
}

func (p *CommandLineParser) startsWith(expected string) bool {
	len2 := utf8.RuneCountInString(expected)
	return len2 <= p.len && p.line[p.starts:len2] == expected
}

func (p *CommandLineParser) goNextTo(expected string) {
	p.starts += utf8.RuneCountInString(expected)
}

func (p *CommandLineParser) rest() string {
	if p.starts < utf8.RuneCountInString(p.line) {
		return p.line[p.starts:]
	}
	return ""
}

// 文字列を タテボウで クォートする(挟む)のは わたしの癖で、
// |apple|banana|cherry| のように区切れる☆（＾～＾）
// そのうち めんどくさくなったら お前もこうなる☆ｍ９（＾～＾）
func (p *CommandLineParser) String() string {
	return fmt.Sprintf("line=|%s| len=%d starts=%d",
		p.line, p.len, p.starts)
}
