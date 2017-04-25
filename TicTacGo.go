package main

import(
	"os"
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/core"
)

type MocLabel struct {
	widgets.QLabel
	_ func(string) `signal:"updateLabel"`
}
type App struct {
	widgets.QApplication

}
type MWidget struct{
	widgets.QWidget

}
type MWindow struct{
	widgets.QMainWindow
}

var application App
var buttonArrays []interface{}
var myBoard GoPlayBoard


func main()  {

	//this creates the game window

	application = *NewApp(len(os.Args), os.Args)
	myBoard = *NewGoPlayBoard()
	NewGrid()


	widgets.QApplication_Exec()
}

func NewGrid() {
	buttonArrays = make([]interface{}, 9)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Tic Tac Go")

	//this creates the layout
	layout := widgets.NewQVBoxLayout()
	//this creates the widget adn sets it layout
	mwidget := widgets.NewQWidget(nil, 0)

	mwidget.SetLayout(layout)

	label := NewMocLabel(nil, 0)
	label.ConnectUpdateLabel(func(s string) {
		switch s {
		case "x", "X":
			label.SetText("Player X's turn")
		case "o", "O":
			label.SetText("Player O's turn")
		}

	})

	label.SetText("Player X's turn")
	layout.AddWidget(label, 0, 0)



	layoutC0 := widgets.NewQHBoxLayout()

	button00 := widgets.NewQPushButton2("", nil)
	button00.SetEnabled(true)
	button00.ConnectClicked(func(checked bool) {
		if button00.IsEnabled() {
			button00.SetEnabled(false)
			button00.SetText(myBoard.WhoTurn())
			setHelper(0, 0)
			label.UpdateLabel(myBoard.WhoTurn())
		}
	})
	layoutC0.AddWidget(button00, 0, core.Qt__AlignTrailing)
	button01 := widgets.NewQPushButton2("", nil)
	button01.SetEnabled(true)
	button01.ConnectClicked(func(checked bool) {
		if button01.IsEnabled() {
			button01.SetEnabled(false)
			button01.SetText(myBoard.WhoTurn())
			setHelper(0, 1)
			label.UpdateLabel(myBoard.WhoTurn())
			//popup("Yo!")
		}
	})

	layoutC0.AddWidget(button01, 0, core.Qt__AlignTrailing)


	button02 := widgets.NewQPushButton2("", nil)
	button02.SetEnabled(true)
	button02.ConnectClicked(func(checked bool) {
		if button02.IsEnabled() {
			button02.SetEnabled(false)
			button02.SetText(myBoard.WhoTurn())
			setHelper(0, 2)
			label.UpdateLabel(myBoard.WhoTurn())
		}
	})
	layoutC0.AddWidget(button02, 0, core.Qt__AlignTrailing)

	layout.AddLayout(layoutC0, 0)
	layoutC1 := widgets.NewQHBoxLayout()
	button10 := widgets.NewQPushButton2("", nil)
	button10.SetEnabled(true)
	button10.ConnectClicked(func(checked bool) {
		if button10.IsEnabled() {
			button10.SetEnabled(false)
			button10.SetText(myBoard.WhoTurn())
			setHelper(1, 0)
			label.UpdateLabel(myBoard.WhoTurn())
		}
	})

	layoutC1.AddWidget(button10, 0, core.Qt__AlignTrailing)
	button11 := widgets.NewQPushButton2("", nil)
	button11.SetEnabled(true)
	button11.ConnectClicked(func(checked bool) {
		if button11.IsEnabled() {

			button11.SetEnabled(false)
			button11.SetText(myBoard.WhoTurn())
			setHelper(1, 1)
			label.UpdateLabel(myBoard.WhoTurn())
		}
	})
	layoutC1.AddWidget(button11, 0, core.Qt__AlignTrailing)

	button12 := widgets.NewQPushButton2("", nil)
	button12.ConnectClicked(func(checked bool) {
		if button12.IsEnabled() {
			button12.SetEnabled(false)
			button12.SetText(myBoard.WhoTurn())
			setHelper(1, 2)
			label.UpdateLabel(myBoard.WhoTurn())
		}
	})
	layoutC1.AddWidget(button12, 0, core.Qt__AlignTrailing)



	layout.AddLayout(layoutC1, 1)

	layoutC2 := widgets.NewQHBoxLayout()


	button20 := widgets.NewQPushButton2("", nil)
	button20.SetEnabled(true)
	button20.ConnectClicked(func(checked bool) {
		if button20.IsEnabled() {
			button20.SetEnabled(false)
			button20.SetText(myBoard.WhoTurn())
			setHelper( 2, 0)
			label.UpdateLabel(myBoard.WhoTurn())
		}

	})

	layoutC2.AddWidget(button20, 0, core.Qt__AlignTrailing)
	button21 := widgets.NewQPushButton2("", nil)
	button21.SetEnabled(true)
	button21.ConnectClicked(func(checked bool) {
		if button21.IsEnabled() {
			button21.SetEnabled(false)
			button21.SetText(myBoard.WhoTurn())
			setHelper( 2, 1)
			label.UpdateLabel(myBoard.WhoTurn())

		}
	})
	layoutC2.AddWidget(button21, 0, core.Qt__AlignTrailing)

	button22 := widgets.NewQPushButton2("", nil)
	button22.SetEnabled(true)
	button22.ConnectClicked(func(checked bool) {
		if button22.IsEnabled() {
			button22.SetEnabled(false)

			button22.SetText(myBoard.WhoTurn())
			setHelper( 2, 2)
			label.UpdateLabel(myBoard.WhoTurn())
			//popup("Yo!")
		}
	})

	layoutC2.AddWidget(button22, 0, core.Qt__AlignTrailing)


	buttonArrays[0] = *button00
	buttonArrays[1] =  *button01
	buttonArrays[2] =  *button02
	buttonArrays[3] = *button10
	buttonArrays[4] =  *button11
	buttonArrays[5] = *button12
	buttonArrays[6] =  *button20
	buttonArrays[7] =  *button21
	buttonArrays[8] =  *button22

	layout.AddLayout(layoutC2, 1)

	window.SetCentralWidget(mwidget)
	window.Show()


}
func reset(){
	for i := 0; i < len(buttonArrays); i++{
		value := buttonArrays[i].(widgets.QPushButton)
		value.SetText("")
		value.SetEnabled(true)
	}

}

func setHelper(i int, j int) {
	myBoard.SetSquare(i, j)
	checkIfGameOver()
	myBoard.SwitchTurn()

}
func checkIfGameOver(){

	//call popup
	winner := myBoard.GameOver()
	if winner != "go"{
		popup(winner)
	}

	//else switch turns
}
func popup(winner string) {
	var myWinner string
	switch winner {
		default:
			myWinner = "There was an error."
	case "x", "X":
		myWinner = "Player X won!!!!!"
	case "o", "O":
		myWinner = "Player O won!!!!!"
	case "tie":
		myWinner = "No one won!!!"

	}
	message := widgets.NewQMessageBox2(widgets.QMessageBox__NoIcon, "Game Over", myWinner, widgets.QMessageBox__Ok, nil, 0)
	message.AddButton3(widgets.QMessageBox__Cancel  )
	rtn := message.Exec()
	switch  rtn {
	case 1024:
		//repopulate here
		reset()
		myBoard.Reset()
	case 4194304:
		//exit here
		//application.QuitDefault()
	}

}




