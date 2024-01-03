package main

//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"time"
//)
//
//func main() {
//	printInitialInfo()
//	playGame()
//}
//
//func printInitialInfo() {
//	fmt.Println("Escolha a posição onde deseja jogar, sendo elas:")
//	fmt.Println()
//	fmt.Println("   |SE|SM|SD|")
//	fmt.Println("   |ME|MM|MD|")
//	fmt.Println("   |IE|IM|ID|")
//	fmt.Println()
//
//	currentPlayer, nextPlayer := setFirstPlayer()
//	fmt.Println("Primeiro jogador:", currentPlayer)
//	fmt.Println("Segundo jogador:", nextPlayer)
//	fmt.Println()
//
//	board := createBoard()
//	printBoard(board)
//	fmt.Println()
//}
//
//func createBoard() [3][7]string {
//	var board [3][7]string
//
//	for i := range board {
//		for j := range board[i] {
//			if j%2 != 0 {
//				board[i][j] = "-"
//			} else {
//				board[i][j] = "|"
//			}
//		}
//	}
//
//	return board
//}
//
//func printBoard(board [3][7]string) {
//	fmt.Println()
//
//	for i := range board {
//		//para o tabuleiro descola da parede
//		fmt.Print("    ")
//
//		for j := range board[i] {
//			//assim eu printo os elementos/strings [j] da linha [i] um a um (sem os colchetes do array)...
//			fmt.Print(board[i][j])
//		}
//		//...e pulo uma linha ao terminar para ir pra próxima
//		fmt.Println()
//	}
//}
//
//func setFirstPlayer() (string, string) {
//	t := time.Now()
//
//	//para usar .Second() tem que ser numa variável do tipo time.Now()
//	if t.Second()%2 == 0 {
//		return "O", "X"
//	} else {
//		return "X", "O"
//	}
//}
//
//func setPlayId(board [3][7]string) map[string][2]int {
//	arrayPositionId := [9]string{"SE", "SM", "SD", "ME", "MM", "MD", "IE", "IM", "ID"}
//	playIdMap := make(map[string][2]int)
//	counterId := 0
//
//	for i := range board {
//		for j := range board[i] {
//			//para setar as posições corretamente já que as colunas tem borda lateral (tamanho=7)
//			if j%2 != 0 {
//				playIdMap[arrayPositionId[counterId]] = [2]int{i, j}
//				counterId++
//			}
//		}
//	}
//
//	return playIdMap
//}
//
//func isBoardComplete(board [3][7]string) bool {
//	emptyMap := make(map[string]int)
//	var isComplete bool
//
//	for i := range board {
//		for j := range board[i] {
//			if board[i][j] == "-" {
//				emptyMap["-"]++
//			}
//		}
//	}
//
//	if emptyMap["-"] == 0 {
//		isComplete = true
//	} else {
//		isComplete = false
//	}
//
//	return isComplete
//}
//
//func playGame() {
//	currentPlayer, nextPlayer := setFirstPlayer()
//	usedIdMap := make(map[string]int)
//
//	board := createBoard()
//	playIdMap := setPlayId(board)
//
//	var haveRowWinner, haveColumnWinner, haveDiagonalWinner = false, false, false
//	var rowWinner, columnWinner, diagonalWinner string
//
//	for !isBoardComplete(board) && !haveRowWinner && !haveColumnWinner && !haveDiagonalWinner {
//		scanner := bufio.NewScanner(os.Stdin)
//		var positionId string
//
//		fmt.Printf("Player %s: ", currentPlayer)
//		scanner.Scan()
//		positionId = scanner.Text()
//
//		for !verifyPositionId(positionId) || !isPositionIdUsed(positionId, usedIdMap) {
//			fmt.Printf("Quadrante %s inválido, Player %s: ", positionId, currentPlayer)
//			scanner.Scan()
//			positionId = scanner.Text()
//		}
//
//		usedIdMap[positionId]++
//		r := playIdMap[positionId][0]
//		c := playIdMap[positionId][1]
//
//		board[r][c] = currentPlayer
//
//		currentPlayer, nextPlayer = nextPlayer, currentPlayer
//
//		printBoard(board)
//		haveDiagonalWinner, _ = verifyDiagonals(board)
//		haveColumnWinner, _ = verifyColumns(board)
//		haveRowWinner, _ = verifyRows(board)
//	}
//
//	haveRowWinner, rowWinner = verifyRows(board)
//	haveColumnWinner, columnWinner = verifyColumns(board)
//	haveDiagonalWinner, diagonalWinner = verifyDiagonals(board)
//
//	switch {
//	case haveRowWinner:
//		fmt.Printf("Resultado: %s ganhou!", rowWinner)
//	case haveColumnWinner:
//		fmt.Printf("Resultado: %s ganhou!", columnWinner)
//	case haveDiagonalWinner:
//		fmt.Printf("Resultado: %s ganhou!", diagonalWinner)
//	default:
//		fmt.Print("Resultado: Velha!")
//	}
//}
//
//func verifyPositionId(positionId string) bool {
//	arrayPositionId := [9]string{"SE", "SM", "SD", "ME", "MM", "MD", "IE", "IM", "ID"}
//
//	for _, id := range arrayPositionId {
//		if positionId == id {
//			return true
//		}
//	}
//
//	return false
//}
//
//func isPositionIdUsed(positionId string, usedIdMap map[string]int) bool {
//	// if usedIdMap[positionId] == 0 {return true}; return false
//	//sugestão do VSCode, retorna true se for 0 ou falso se não for
//	return usedIdMap[positionId] == 0
//}
//
//func verifyRows(board [3][7]string) (bool, string) {
//	var iCountX, iCountO int
//
//	for i := range board {
//		iCountX = 0
//		iCountO = 0
//
//		for j := range board[i] {
//			switch {
//			case board[i][j] == "X":
//				iCountX++
//			case board[i][j] == "O":
//				iCountO++
//			}
//		}
//
//		if iCountX == 3 {
//			return true, "X"
//		} else if iCountO == 3 {
//			return true, "O"
//		}
//	}
//
//	return false, ""
//}
//
//func verifyColumns(board [3][7]string) (bool, string) {
//	var jCountX, jCountO int
//	//preciso verificar a mesma coluna (j) para cada linha (i)
//	//i vai de 0 a 2 (tamanho 3)
//	//j vai de 0 a 6 (tamanho 7)
//	//entao preciso fazer que isso aconteça:
//	//[i][j]
//	//[0][0] -> [i][j]
//	//[1][0] -> [i+1][j]
//	//[2][0] -> [i+2][j]
//
//	for j := 0; j < 7; j++ {
//		jCountX = 0
//		jCountO = 0
//
//		for i := 0; i < 3; i++ {
//			if j%2 != 0 {
//				switch {
//				case board[i][j] == "X":
//					jCountX++
//				case board[i][j] == "O":
//					jCountO++
//				}
//			}
//		}
//
//		if jCountX == 3 {
//			return true, "X"
//		} else if jCountO == 3 {
//			return true, "O"
//		}
//	}
//
//	return false, ""
//}
//
//func verifyDiagonals(board [3][7]string) (bool, string) {
//	if board[0][1] == "X" && board[1][3] == "X" && board[2][5] == "X" || board[2][1] == "X" && (board[1][3] == "X" && board[0][5] == "X") {
//		return true, "X"
//	}
//
//	if board[0][1] == "O" && board[1][3] == "O" && board[2][5] == "O" || board[2][1] == "O" && (board[1][3] == "O" && board[0][5] == "O") {
//		return true, "O"
//	}
//
//	return false, ""
//}
