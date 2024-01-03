package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Board struct {
	Board [3][3]string
}

// global variables
var arrPlayId = [9]string{"SE", "SM", "SD", "ME", "MM", "MD", "IE", "IM", "ID"}
var scanner = bufio.NewScanner(os.Stdin)

var notUsedBoardIdMap = createNotUsedBoardIdMap()
var arrNotUsedIdMap = createArrUsedIdMap()
var arrBoards = setGameBoards()
var rBoard = setResultBoard()

func setBoard(id string, arrBoards [9]Board) Board {
	var b Board
	var i int

	for i, b = range arrBoards {
		if id == arrPlayId[i] {
			break
		}
	}

	return b
}

func setPlayers() (string, string) {
	t := time.Now()

	if t.Second()%2 == 0 {
		return "O", "X"
	} else {
		return "X", "O"
	}
}

// only need one function to create both boards
func createBoard() Board {
	var b Board

	for i := range b.Board {
		for j := range b.Board[i] {
			b.Board[i][j] = "-"
		}
	}

	return b
}

func setGameBoards() [9]Board {
	gBoard1 := createBoard()
	gBoard2 := createBoard()
	gBoard3 := createBoard()
	gBoard4 := createBoard()
	gBoard5 := createBoard()
	gBoard6 := createBoard()
	gBoard7 := createBoard()
	gBoard8 := createBoard()
	gBoard9 := createBoard()

	arrayBoards := [9]Board{gBoard1, gBoard2, gBoard3, gBoard4, gBoard5, gBoard6, gBoard7, gBoard8, gBoard9}

	return arrayBoards
}

func createArrUsedIdMap() [9]map[string]int {
	notUsedIdMap1 := make(map[string]int)
	notUsedIdMap2 := make(map[string]int)
	notUsedIdMap3 := make(map[string]int)
	notUsedIdMap4 := make(map[string]int)
	notUsedIdMap5 := make(map[string]int)
	notUsedIdMap6 := make(map[string]int)
	notUsedIdMap7 := make(map[string]int)
	notUsedIdMap8 := make(map[string]int)
	notUsedIdMap9 := make(map[string]int)

	arrayNotUsedIdMaps := [9]map[string]int{notUsedIdMap1, notUsedIdMap2, notUsedIdMap3, notUsedIdMap4, notUsedIdMap5,
		notUsedIdMap6, notUsedIdMap7, notUsedIdMap8, notUsedIdMap9}

	for _, notUsedIdMap := range arrayNotUsedIdMaps {
		for i := range arrPlayId {
			notUsedIdMap[arrPlayId[i]]++
		}
	}

	return arrayNotUsedIdMaps
}

func createNotUsedBoardIdMap() map[string]int {
	nUsedBoardIdMap := make(map[string]int)

	for i := range arrPlayId {
		nUsedBoardIdMap[arrPlayId[i]]++
	}

	return nUsedBoardIdMap
}

func setResultBoard() Board {
	resultBoard := createBoard()

	return resultBoard
}

func printGameBoard(b Board, boardId string) {
	fmt.Printf("\nBoard: %s\n", boardId)

	for i := range b.Board {
		fmt.Print("     | ")

		for j := range b.Board[i] {
			fmt.Print(b.Board[i][j], " | ")
		}

		fmt.Println()
	}
}

//func printResultBoard(r Board) {
//	fmt.Printf("Result board: \n")
//
//	for i := range r.Board {
//		fmt.Print("     | ")
//
//		for j := range r.Board[i] {
//			fmt.Print(r.Board[i][j], " | ")
//		}
//
//		fmt.Println()
//	}
//}

func setPlayIdMap(b Board) map[string][2]int {
	playIdMap := make(map[string][2]int)
	counter := 0

	for i := range b.Board {
		for j := range b.Board[i] {
			playIdMap[arrPlayId[counter]] = [2]int{i, j}
			counter++
		}
	}

	return playIdMap
}

func getBoardId(crrPlayer string) string {
	var boardId string

	fmt.Printf("\nPlayer %s, choose the board: ", crrPlayer)
	scanner.Scan()
	boardId = scanner.Text()

	for !verifyBoardId(boardId) {
		fmt.Printf("Invalid board. Player %s, choose another: ", crrPlayer)
		scanner.Scan()
		boardId = scanner.Text()
	}

	return boardId
}

func setNextBoardId(playId string) string {
	nextBoardId := playId
	return nextBoardId
}

func getPlayId(crrPlayer string, mIndex int) (string, [9]map[string]int) {
	var playId string

	fmt.Printf("\nPlayer %s, choose your position: ", crrPlayer)
	scanner.Scan()
	playId = scanner.Text()

	for !verifyPlayId(playId, mIndex) {
		fmt.Printf("Invalid position. Player %s, choose another: ", crrPlayer)
		scanner.Scan()
		playId = scanner.Text()
	}

	arrNotUsedIdMap = deleteUsedId(playId, mIndex)

	return playId, arrNotUsedIdMap
}

func setPlayId(b Board, bIndex int, playId, crrPlayer string) ([9]Board, [3][3]string) {
	playIdMap := setPlayIdMap(b)
	i := playIdMap[playId][0]
	j := playIdMap[playId][1]
	b.Board[i][j] = crrPlayer

	arrBoards[bIndex] = b

	return arrBoards, b.Board
}

func setArraysIndex(boardId string) int {
	var i int

	for i = range arrPlayId {
		if boardId == arrPlayId[i] {
			break
		}
	}

	return i
}

func verifyBoardId(boardId string) bool {
	for id := range notUsedBoardIdMap {
		if boardId == id {
			return true
		}
	}

	return false
}

func verifyPlayId(playId string, bIndex int) bool {
	for i, notUsedMap := range arrNotUsedIdMap {
		if i == bIndex {
			_, ok := notUsedMap[playId]
			if ok {
				return true
			}
		}
	}

	return false
}

func deleteUsedId(playId string, bIndex int) [9]map[string]int {
	for i, notUsedMap := range arrNotUsedIdMap {
		if i == bIndex {
			delete(notUsedMap, playId)
		}
	}

	return arrNotUsedIdMap
}

func checkBoardWin(b Board, bIndex int, boardId, crrPlayer string) (bool, Board, map[string]int) {
	boardIdMap := setPlayIdMap(b)
	i := boardIdMap[boardId][0]
	j := boardIdMap[boardId][1]

	if checkRows(b, crrPlayer) || checkColumns(b, crrPlayer) || checkDiagonals(b, crrPlayer) {
		fmt.Printf("\nPlayer %s wins on the board %s!\n", crrPlayer, boardId)

		rBoard.Board[i][j] = crrPlayer
		arrBoards = uptadeResultBoard(b, bIndex, crrPlayer)
		delete(notUsedBoardIdMap, boardId)

		return true, rBoard, notUsedBoardIdMap
	} else if len(arrNotUsedIdMap[bIndex]) == 0 {
		fmt.Printf("\nBoard %s result: Ø!\n", boardId)

		rBoard.Board[i][j] = "Ø"

		arrBoards = uptadeResultBoard(b, bIndex, "Ø")
		delete(notUsedBoardIdMap, boardId)

		return true, rBoard, notUsedBoardIdMap
	}

	return false, rBoard, notUsedBoardIdMap
}

func checkFinalWin(r Board, ntxPlayer string) bool {
	if checkFinalRows(r, ntxPlayer) || checkFinalColumns(r, ntxPlayer) || checkFinalDiagonals(r, ntxPlayer) {
		fmt.Printf("\nCongratulations Player %s, you won!\n", ntxPlayer)
		return true
	} else if checkFinalDraw(r) {
		fmt.Printf("\nGame result: Draw!\n")
		return true
	}

	return false
}

func checkFinalDraw(r Board) bool {
	cellsMap := make(map[string]int)

	for i := range r.Board {
		for j := range r.Board[i] {
			switch r.Board[i][j] {
			case "X":
				cellsMap["X"]++
			case "O":
				cellsMap["O"]++
			case "Ø":
				cellsMap["Ø"]++
			}
		}
	}

	if cellsMap["X"]+cellsMap["O"]+cellsMap["Ø"] == 9 {
		return true
	}

	return false
}

func checkFinalRows(b Board, crrPlayer string) bool {
	var iCount, jokerCount int

	for i := range b.Board {
		jokerCount = 0
		iCount = 0

		for j := range b.Board[i] {
			switch b.Board[i][j] {
			case crrPlayer:
				iCount++
			case "Ø":
				jokerCount++
			}
		}

		if iCount+jokerCount == 3 {
			return true
		}
	}

	return false
}

func checkFinalColumns(b Board, crrPlayer string) bool {
	var jCount, jokerCount int

	for i := range b.Board {
		jokerCount = 0
		jCount = 0

		for j := range b.Board[i] {
			switch b.Board[j][i] {
			case crrPlayer:
				jCount++
			case "Ø":
				jokerCount++
			}
		}

		if jCount+jokerCount == 3 {
			return true
		}
	}

	return false
}

func checkFinalDiagonals(r Board, crrPlayer string) bool {
	switch r.Board[1][1] {
	case crrPlayer:
		if (r.Board[0][0] == crrPlayer || r.Board[0][0] == "Ø") && (r.Board[2][2] == crrPlayer || r.Board[2][2] == "Ø") ||
			(r.Board[2][0] == crrPlayer || r.Board[2][0] == "Ø") && (r.Board[0][2] == crrPlayer || r.Board[0][2] == "Ø") {
			return true
		}
	}

	return false
}

func checkRows(b Board, crrPlayer string) bool {
	var iCount int

	for i := range b.Board {
		iCount = 0

		for j := range b.Board[i] {
			if b.Board[i][j] == crrPlayer {
				iCount++
			}
		}

		if iCount == 3 {
			return true
		}
	}

	return false
}

func checkColumns(b Board, crrPlayer string) bool {
	var jCount int

	for i := range b.Board {
		jCount = 0

		for j := range b.Board[i] {
			if b.Board[j][i] == crrPlayer {
				jCount++
			}
		}

		if jCount == 3 {
			return true
		}
	}

	return false
}

func checkDiagonals(b Board, crrPlayer string) bool {
	switch b.Board[1][1] {
	case crrPlayer:
		if b.Board[0][0] == crrPlayer && b.Board[2][2] == crrPlayer || b.Board[2][0] == crrPlayer && b.Board[0][2] == crrPlayer {
			return true
		}
	}

	return false
}

func uptadeResultBoard(b Board, bIndex int, crrPlayer string) [9]Board {
	for i := range b.Board {
		for j := range b.Board[i] {
			b.Board[i][j] = crrPlayer
		}
	}

	arrBoards[bIndex] = b

	return arrBoards
}

func checkWin(b Board, bIndex int, boardId, crrPlayer string) (bool, map[string]int) {
	if checkRows(b, crrPlayer) || checkColumns(b, crrPlayer) || checkDiagonals(b, crrPlayer) {
		delete(notUsedBoardIdMap, boardId)

		return true, notUsedBoardIdMap
	} else if len(arrNotUsedIdMap[bIndex]) == 0 {
		delete(notUsedBoardIdMap, boardId)

		return true, notUsedBoardIdMap
	}

	return false, notUsedBoardIdMap
}

func getNotUsedBoardId(b Board, boardId, crrPlayer string, bIndex int, arrayBoards [9]Board) (Board, string, map[string]int) {
	_, notUsedBoardIdMap = checkWin(b, bIndex, boardId, crrPlayer)

	//_, ok := notUsedBoardIdMap[boardId]

	for !verifyBoardId(boardId) {
		fmt.Printf("Invalid board. Player %s, choose another: ", crrPlayer)
		scanner.Scan()
		boardId = scanner.Text()
	}

	//delete(notUsedBoardIdMap, boardId)

	crrBoard := setBoard(boardId, arrayBoards)

	return crrBoard, boardId, notUsedBoardIdMap
}

func isBoardPlayable(b Board, bIndex int, boardId, crrPlayer string) bool {
	_, notUsedBoardIdMap = checkWin(b, bIndex, boardId, crrPlayer)
	_, ok := notUsedBoardIdMap[boardId]

	if !ok {
		return false
	}

	return true
}

func playGame() {
	var playId string

	crrPlayer, nxtPlayer := setPlayers()
	boardId := getBoardId(crrPlayer)
	crrBoard := setBoard(boardId, arrBoards)

	//with (crrPlayer), checkFinalWin finish on the next turn, so I used nxtPlayer
	for !checkFinalWin(rBoard, nxtPlayer) {
		bIndex := setArraysIndex(boardId) //por algum motivo o bIndex nao pode ser declarado fora do for

		if !isBoardPlayable(crrBoard, bIndex, boardId, crrPlayer) {
			printCompleteBoard(arrBoards)
			printGameBoard(crrBoard, boardId)
			crrBoard, boardId, notUsedBoardIdMap = getNotUsedBoardId(crrBoard, boardId, crrPlayer, bIndex, arrBoards)
			bIndex = setArraysIndex(boardId)
		}

		printCompleteBoard(arrBoards)
		printGameBoard(crrBoard, boardId)

		playId, arrNotUsedIdMap = getPlayId(crrPlayer, bIndex)
		arrBoards, crrBoard.Board = setPlayId(crrBoard, bIndex, playId, crrPlayer)

		_, rBoard, notUsedBoardIdMap = checkBoardWin(crrBoard, bIndex, boardId, crrPlayer)

		if checkFinalWin(rBoard, nxtPlayer) {
			break
		}

		boardId = setNextBoardId(playId)
		nxtBoard := setBoard(playId, arrBoards)

		crrPlayer, nxtPlayer = nxtPlayer, crrPlayer
		crrBoard, nxtBoard = nxtBoard, crrBoard
	}

	printCompleteBoard(arrBoards)
}
