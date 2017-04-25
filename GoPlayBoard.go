package main


type GoPlayBoard struct{
	board [3][3]string
	turn string
}

func NewGoPlayBoard() (game *GoPlayBoard) {
	game.turn = "X"
	game = &GoPlayBoard{}
	return
}

func (game *GoPlayBoard) SwitchTurn(){
	if game.turn == "x"{
		game.turn = "o"
	}else{
		game.turn = "x"
	}

}

func (game *GoPlayBoard) WhoTurn() string {
	return  game.turn
}
func (game *GoPlayBoard) SetSquare(i int, y int) {
	game.board[i][y] = game.turn

}

func (game *GoPlayBoard) GameOver() string {
	if game.board[0][0] != "" &&
		game.board[0][0] == game.board[0][1] && game.board[0][0] == game.board[0][2]{
		return game.turn
	}else if game.board[1][0] != "" &&
		game.board[1][0] == game.board[1][1] && game.board[1][0] == game.board[1][2]{
		return game.turn
	}else if game.board[2][0] != "" &&
		game.board[2][0] == game.board[2][1] && game.board[2][0] == game.board[2][2]{
		return game.turn
	}else if game.board[0][0] != "" &&
		game.board[0][0] == game.board[1][0] && game.board[0][0] == game.board[2][0]{
		return game.turn
	}else if game.board[0][1] != "" &&
		game.board[0][1] == game.board[1][1] && game.board[0][1] == game.board[2][1]{
		return game.turn
	}else if game.board[0][2] != "" &&
		game.board[0][2] == game.board[1][2] && game.board[0][2] == game.board[2][2]{
		return game.turn
	}else if game.board[0][0] != "" &&
		game.board[0][0] == game.board[1][1] && game.board[0][0] == game.board[2][2]{
		return game.turn
	}else if game.board[2][2] != "" &&
		game.board[2][2] == game.board[1][1] && game.board[2][2] == game.board[0][2]{
		return game.turn
	}else{
		for i := 0; i < 3; i++{
			for j := 0; j < 3; j++{
				if game.board[i][j] == ""{
					return "go"
				}
			}
		}
		return "tie"
	}
}

func (game *GoPlayBoard) Reset()  {
	for i := 0; i < 3; i++{
		for j := 0; j < 3; j++{
			 game.board[i][j] = ""

		}
	}

}