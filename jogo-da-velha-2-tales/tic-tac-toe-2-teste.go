// /*	REQUISITOS:
//
//  1. TABULEIRO(S) PEQUENO(S);
//
//  2. TABULEIRO GRANDE;
//     2.1) A JOGADA NO PEQUENO DEFINE A POSIÇÃO A SER JOGADA NO GRANDE;
//     2.2) IDENTIFICAÇÕES PARA CADA QUADRANTE;
//
//  3. TABULEIROS PEQUENOS TEM QUE SER INDEPENDENTES UM DO OUTRO;
//
//  4. VENCEDOR OU EMPATE DO(S) TABULEIRO(S) PEQUENO(S);
//     4.1) EMPATE NO PEQUENO VALE PRA QUALQUER UM NO GRANDE;
//     4.2) CONDIÇÕES DE VITÓRIA;
//
//  5. VENCEDOR OU EMPATE DO TABULEIRO GRANDE;
//     5.1) CONDIÇÕES DE VITÓRIA;
//
//  6. JOGADOR INICIAL ESCOLHIDO ALEATORIAMENTE;
//     6.1) APENAS O PRIMEIRO JOGADOR PODE ESCOLHER EM QUAL TABULEIRO JOGAR;
//
//  7. VERIFICADORES PARA VER SE ONDE O JOGADOR QUER JOGAR É VÁLIDO;
//
//  8. REGRAS NO INÍCIO;	*/
package main

//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"time"
//)
//
////NOTES:
//
////i => linhas smallBoard 0-2
////j => colunas smallBoard 0-2
////b => cada smallBoard 0-8
//
//func main() {
//	bigBoard := createBigBoard()
//	printBigBoard(bigBoard)
//
//	playGame()
//}
//
//func createBigBoard() [9][3][3]string {
//	var bigBoard [9][3][3]string
//
//	for b := range bigBoard {
//		for i := range bigBoard[b] {
//			for j := range bigBoard[b][i] {
//				bigBoard[b][i][j] = "-"
//			}
//		}
//	}
//
//	return bigBoard
//}
//
//func printBigBoard(bigBoard [9][3][3]string) {
//	for b := range bigBoard {
//		for i := range bigBoard[b] {
//			fmt.Print("| ")
//
//			for j := range bigBoard[b][i] {
//				fmt.Print(bigBoard[b][i][j])
//				fmt.Print(" ")
//
//				if j == 2 {
//					fmt.Print("| ")
//				}
//			}
//
//			if (b == 2 || b == 5) && i == 2 {
//				fmt.Println()
//				fmt.Print("-----------------------------")
//			}
//		}
//
//		fmt.Println()
//	}
//}
//
//func setPlayers() (string, string) {
//	t := time.Now()
//
//	if t.Second()%2 == 0 {
//		return "X", "O"
//	}
//
//	return "O", "X"
//}
//
//// a identificacao dos boards vai de 0 a 8
//func setBoardId(bigBoard [9][3][3]string) map[string]int {
//	arrBoardId := [9]string{"BSE", "BSM", "BSD", "BME", "BMM", "BMD", "BIE", "BIM", "BID"}
//	boardIdMap := make(map[string]int)
//
//	for b := range bigBoard {
//		boardIdMap[arrBoardId[b]] = b
//	}
//
//	return boardIdMap
//}
//
//func setPlayId(bigBoard [9][3][3]string, currentBoardId int) map[string][2]int {
//	arrPositionId := [9]string{"SE", "SM", "SD", "ME", "MM", "MD", "IE", "IM", "ID"}
//	playIdMap := make(map[string][2]int)
//	b := currentBoardId
//	counterId := 0
//
//	for i := range bigBoard[b] {
//		for j := range bigBoard[b][i] {
//			playIdMap[arrPositionId[counterId]] = [2]int{i, j}
//		}
//	}
//
//	return playIdMap
//}
//
//func getFirstBoardId(bigBoard [9][3][3]string) int {
//	var boardId string
//
//	scanner := bufio.NewScanner(os.Stdin)
//	boardIdMap := setBoardId(bigBoard)
//	currentPlayer, _ := setPlayers()
//
//	fmt.Printf("Player %s, escolha o primeiro tabuleiro: ", currentPlayer)
//	scanner.Scan()
//	boardId = scanner.Text()
//
//	for !verifyFirstBoardId(boardId, boardIdMap) {
//		fmt.Printf("Player %s, tabuleiro inválido. Escolha outra: ", currentPlayer)
//		scanner.Scan()
//		boardId = scanner.Text()
//	}
//
//	currentBoardId := boardIdMap[boardId]
//
//	return currentBoardId
//}
//
//func verifyFirstBoardId(boardId string, boardIdMap map[string]int) bool {
//	_, ok := boardIdMap[boardId]
//
//	return ok
//}
//
////func getPlayId(currentBoardId int) string {
////	currentPlayer, nextPlayer := setPlayers()
////
////	bigBoard := createBigBoard()
////	playIdMap := setPlayId(bigBoard, currentBoardId)
////
////	scanner := bufio.NewScanner(os.Stdin)
////	var positionId string
////
////	fmt.Printf("Player %s, escolha a posição: ", currentPlayer)
////	scanner.Scan()
////	positionId = scanner.Text()
////
////	if !verifyPositionId(positionId) {
////		fmt.Printf("Player %s, posição inválida. Escolha outra: ", currentPlayer)
////		scanner.Scan()
////		positionId = scanner.Text()
////	}
////
////	currentPlayer, nextPlayer = nextPlayer, currentPlayer
////}
//
//func verifyPositionId(positionId string) bool {
//	arrPositionId := [9]string{"SE", "SM", "SD", "ME", "MM", "MD", "IE", "IM", "ID"}
//
//	for _, id := range arrPositionId {
//		if positionId == id {
//			return true
//		}
//	}
//
//	return false
//}
//
//func playGame() {
//	bigBoard := createBigBoard()
//
//	currentPlayer, nextPlayer := setPlayers()
//	currentBoardId := getFirstBoardId(bigBoard)
//
//	boardIdMap := setBoardId(bigBoard)
//	//playIdMap := setPlayId(bigBoard, currentBoardId) --  PODE SER DENTRO DO FOR QUE O JOGO NAO ACABOU
//	fmt.Printf("crrPly = %s\nnxtPly = %s\ncrrBoardId = %d\nboardIdMap = %v", currentPlayer, nextPlayer, currentBoardId, boardIdMap)
//}
