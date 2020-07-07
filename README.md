# tic-tac-toe-on-golang

○×ゲームだぜ☆（＾～＾）

* [Python ver](https://github.com/muzudho/tic-tac-toe-on-python)
* [Rust ver](https://github.com/muzudho/tic-tac-toe)

## Kitting

[Downloads](https://golang.org/dl/)
[VSCodeでGo言語の開発環境を構築する](https://qiita.com/melty_go/items/c977ba594efcffc8b567)

## How to make tic tac toe?

During development, you may need to reproduce the behavior of your computer.  
It is difficult to compare the behavior. Instead, it is useful to get the logs and compare the logs.  

* [x] 'log.go' (You can code in 30 minutes)
  * [x] Writeln - Write to a file.
  * [x] Clear - Log to empty.
  * [x] Println - Write and display.

The first thing you have to create is your motive.  
It is important to start with the appearance.  

* [x] 'look_and_model.go' (You can code in 4 hours)
  * [x] Piece - "O", "X".
  * [x] Game result - Win/Draw/Lose.
  * [x] Position - It's the board.
  * [x] Search - Computer player search info.

If you want to play immediately, you have the talent of a game creator.  
Being able to control your position means being able to play.  

* [ ] 'position.go' (You can code in 15 minutes)
  * [ ] doMove
  * [ ] undoMove
  * [ ] opponent

Let's enter commands into the computer. Create a command line parser.  

* [ ] 'command_line_parser.go' (You can code in 40 minutes)
  * [ ] Input.
  * [ ] Starts with.
  * [ ] Go next to.
  * [ ] Rest.

People who are looking for something 10 minutes a day are looking for something for a week in a year.  
Before creating the game itself, let's first create the replay function. Let's get it for a week.  

* [ ] 'uxi_protocol.go' (You can code in 1.5 hours)
  * [ ] To XFEN.
  * [ ] Do. (Before 'From XFEN') Excludes legal moves and winning/losing decisions.
  * [ ] From XFEN.
  * [ ] Undo.

Let's make a principal command.  

* [ ] 'main.go' (You can code in 1 hours)
  * [ ] position.
  * [ ] pos.
  * [ ] do.
  * [ ] undo.
  * [ ] uxi tic-tac-toe
  * [ ] xfen.

Before you make a computer player, let's judge the outcome. And let's test.  

* [ ] 'win_lose_judgment.go' (You can code in 15 minutes)
  * [ ] Win.
  * [ ] Draw - Not win, not lose, can not play.
  * [ ] Lose. - Not win is lose.
* [ ] 'uxi_protocol.go' (You can code in 1.5 hours)
  * [ ] Do. Winning/losing decisions.

Before creating a computer player, let's create a mechanism to measure performance.  

* [ ] 'performance_measurement.go' (You can code in 15 minutes)
  * [ ] Seconds. - Stopwatch.
  * [ ] Node per second.

Finally, let's make a computer player. (You can code in 1.5 hours)  

* [ ] 'computer_player.go'
  * [ ] Search.
  * [ ] Evaluation. - None.
* [ ] 'main.go'
  * [ ] Create "go" command.
