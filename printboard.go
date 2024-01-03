package main

import (
	"fmt"
)

func printCompleteBoard(arrayBoards [9]Board) {
	printLine1(arrayBoards)
	printLine2(arrayBoards)
	printLine3(arrayBoards)
	fmt.Printf("------+-------+------\n")
	printLine4(arrayBoards)
	printLine5(arrayBoards)
	printLine6(arrayBoards)
	fmt.Printf("------+-------+------\n")
	printLine7(arrayBoards)
	printLine8(arrayBoards)
	printLine9(arrayBoards)
}

func printInfo() {
	printCompleteBoard(arrBoards)
	fmt.Printf("\nThis is is TicTacToe 2!\nPress R to see how it works and C to continue... ")
	scanner.Scan()
	key := scanner.Text()

	for key != "R" && key != "C" {
		fmt.Printf("Invalid key (R/C), press another: ")
		scanner.Scan()
		key = scanner.Text()
	}

	if key == "R" {
		fmt.Printf("\nTo play you need to inform the board or the position, their ids are:\n" +
			"\n     SE|SM|SD\n" +
			"     ME|MM|MD\n" +
			"     IE|IM|ID\n" +
			"\nThey select both, board or position.\nLet's start!\n")
	}
}

func printLine1(arrayBoards [9]Board) {
	//fmt.Print("     ")
	for j := 0; j < 3; j++ {
		fmt.Print(arrayBoards[0].Board[0][j], " ")
	}
	fmt.Print("| ")

	for j := 0; j < 3; j++ {
		fmt.Print(arrayBoards[1].Board[0][j], " ")
	}
	fmt.Print("| ")

	for j := 0; j < 3; j++ {
		fmt.Print(arrayBoards[2].Board[0][j], " ")
	}
	fmt.Println()
}

func printLine2(arrayBoards [9]Board) {
	//fmt.Print("     ")
	for j := 0; j < 3; j++ {
		fmt.Print(arrayBoards[0].Board[1][j], " ")
	}
	fmt.Print("| ")

	for j := 0; j < 3; j++ {
		fmt.Print(arrayBoards[1].Board[1][j], " ")
	}
	fmt.Print("| ")

	for j := 0; j < 3; j++ {
		fmt.Print(arrayBoards[2].Board[1][j], " ")
	}
	fmt.Println()
}

func printLine3(arrayBoards [9]Board) {
	//fmt.Print("     ")
	for j := 0; j < 3; j++ {
		fmt.Print(arrayBoards[0].Board[2][j], " ")
	}
	fmt.Print("| ")

	for j := 0; j < 3; j++ {
		fmt.Print(arrayBoards[1].Board[2][j], " ")
	}
	fmt.Print("| ")

	for j := 0; j < 3; j++ {
		fmt.Print(arrayBoards[2].Board[2][j], " ")
	}
	fmt.Println()
}

func printLine4(arrayBoards [9]Board) {
	//fmt.Print("     ")
	for j := 0; j < 3; j++ {
		fmt.Print(arrayBoards[3].Board[0][j], " ")
	}
	fmt.Print("| ")

	for j := 0; j < 3; j++ {
		fmt.Print(arrayBoards[4].Board[0][j], " ")
	}
	fmt.Print("| ")

	for j := 0; j < 3; j++ {
		fmt.Print(arrayBoards[5].Board[0][j], " ")
	}
	fmt.Println()
}

func printLine5(arrayBoards [9]Board) {
	//fmt.Print("     ")
	for j := 0; j < 3; j++ {
		fmt.Print(arrayBoards[3].Board[1][j], " ")
	}
	fmt.Print("| ")

	for j := 0; j < 3; j++ {
		fmt.Print(arrayBoards[4].Board[1][j], " ")
	}
	fmt.Print("| ")

	for j := 0; j < 3; j++ {
		fmt.Print(arrayBoards[5].Board[1][j], " ")
	}
	fmt.Println()
}

func printLine6(arrayBoards [9]Board) {
	//fmt.Print("     ")
	for j := 0; j < 3; j++ {
		fmt.Print(arrayBoards[3].Board[2][j], " ")
	}
	fmt.Print("| ")

	for j := 0; j < 3; j++ {
		fmt.Print(arrayBoards[4].Board[2][j], " ")
	}
	fmt.Print("| ")

	for j := 0; j < 3; j++ {
		fmt.Print(arrayBoards[5].Board[2][j], " ")
	}
	fmt.Println()
}

func printLine7(arrayBoards [9]Board) {
	//fmt.Print("     ")
	for j := 0; j < 3; j++ {
		fmt.Print(arrayBoards[6].Board[0][j], " ")
	}
	fmt.Print("| ")

	for j := 0; j < 3; j++ {
		fmt.Print(arrayBoards[7].Board[0][j], " ")
	}
	fmt.Print("| ")

	for j := 0; j < 3; j++ {
		fmt.Print(arrayBoards[8].Board[0][j], " ")
	}
	fmt.Println()
}

func printLine8(arrayBoards [9]Board) {
	//fmt.Print("     ")
	for j := 0; j < 3; j++ {
		fmt.Print(arrayBoards[6].Board[1][j], " ")
	}
	fmt.Print("| ")

	for j := 0; j < 3; j++ {
		fmt.Print(arrayBoards[7].Board[1][j], " ")
	}
	fmt.Print("| ")

	for j := 0; j < 3; j++ {
		fmt.Print(arrayBoards[8].Board[1][j], " ")
	}
	fmt.Println()
}

func printLine9(arrayBoards [9]Board) {
	//fmt.Print("     ")
	for j := 0; j < 3; j++ {
		fmt.Print(arrayBoards[6].Board[2][j], " ")
	}
	fmt.Print("| ")

	for j := 0; j < 3; j++ {
		fmt.Print(arrayBoards[7].Board[2][j], " ")
	}
	fmt.Print("| ")

	for j := 0; j < 3; j++ {
		fmt.Print(arrayBoards[8].Board[2][j], " ")
	}
	fmt.Println()
}
