# tic-tac-toe-on-golang

○×ゲームだぜ☆（＾～＾）

* [Python ver](https://github.com/muzudho/tic-tac-toe-on-python)
* [Rust ver](https://github.com/muzudho/tic-tac-toe)

## Kitting

Golang は build して .exe ファイルを叩けだぜ☆ｍ９（＾～＾） `go run` はファイル１個で実行するときにしとけだぜ☆ｍ９（＾～＾）  

[Downloads](https://golang.org/dl/)
[VSCodeでGo言語の開発環境を構築する](https://qiita.com/melty_go/items/c977ba594efcffc8b567)

## Run

ソース・ファイルが１つのときは、  

Input:  

```shell
cd src
go run main.go
```

でいいんだが、複数あるときは、ビルドすると ディレクトリの名前で実行ファイルが作られるから、その実行ファイルの名前を叩けだぜ☆（＾～＾）  

Input:  

```shell
cd src
go build
src
```

参考:  
[Go言語で同じディレクトリ内の関数呼び出すのにエラーがでる（#command-line-arguments ./main.go: undefined: funcName）](https://qiita.com/kosukeKK/items/abb208fd0bbd3744ddfb)

## How to make tic tac toe?

During development, you may need to reproduce the behavior of your computer.  
It is difficult to compare the behavior. Instead, it is useful to get the logs and compare the logs.  

* [x] Step 1. 'log.go' (You can code in 30 minutes)
  * [x] Writeln - Write to a file.
  * [x] Clear - Log to empty.
  * [x] Println - Write and display.

The first thing you have to create is your motive.  
It is important to start with the appearance.  

* [x] Step 2. 'look_and_model.go' (You can code in 4 hours)
  * [x] Piece - "O", "X".
  * [x] Game result - Win/Draw/Lose.
  * [x] Position - It's the board.
  * [x] Search - Computer player search info.

If you want to play immediately, you have the talent of a game creator.  
Being able to control your position means being able to play.  

* [x] Step 3. 'position.go' (You can code in 15 minutes)
  * [x] doMove
  * [x] undoMove
  * [x] opponent

Let's enter commands into the computer. Create a command line parser.  

* [x] Step 4. 'command_line_parser.go' (You can code in 40 minutes)
  * [x] Input.
  * [x] Starts with.
  * [x] Go next to.
  * [x] Rest.

People who are looking for something 10 minutes a day are looking for something for a week in a year.  
Before creating the game itself, let's first create the replay function. Let's get it for a week.  

* [x] Step 5. 'uxi_protocol.go' (You can code in 1.5 hours)
  * [x] To XFEN.
  * [x] Do. (Before 'From XFEN') Excludes legal moves and winning/losing decisions.
  * [x] From XFEN.
  * [x] Undo.

Let's make a principal command.  

* [x] Step 6. 'main.go' (You can code in 1 hours)
  * [x] position.
  * [x] pos.
  * [x] do.
  * [x] undo.
  * [x] uxi tic-tac-toe
  * [x] xfen.

Before you make a computer player, let's judge the outcome. And let's test.  

* [x] Step 7. 'win_lose_judgment.go' (You can code in 15 minutes)
  * [x] Win.
  * [x] Draw - Not win, not lose, can not play.
  * [ ] Lose. - Not win is lose.
* [x] 'uxi_protocol.go' (You can code in 1.5 hours)
  * [x] Do. Winning/losing decisions.

Before creating a computer player, let's create a mechanism to measure performance.  

* [x] Step 8. 'performance_measurement.go' (You can code in 15 minutes)
  * [x] Seconds. - Stopwatch.
  * [x] Node per second.

Finally, let's make a computer player. (You can code in 1.5 hours)  

* [x] Step 9. 'computer_player.go'
  * [x] Search.
  * [ ] Evaluation. - None.
* [x] 'main.go'
  * [x] Create "go" command.
* [x] Remeve all 'TODO' tasks. Examples: '// TODO Write a code here.'
