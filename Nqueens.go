package main

import (
	"fmt"
)

var N = 8
//var board = make([][]int, N)

var board = [8][8]int {{0,0,0,0,0,0,0,0}, {0,0,0,0,0,0,0,0}, {0,0,0,0,0,0,0,0},
{0,0,0,0,0,0,0,0}, {0,0,0,0,0,0,0,0}, {0,0,0,0,0,0,0,0}, {0,0,0,0,0,0,0,0}, {0,0,0,0,0,0,0,0}}


func print (board [8][8]int){
	for _,x := range board {
		for _, y := range x {
			fmt.Print(y)
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
}

func isSafe (row int, col int) bool{
	for i:=0; i<N; i++ {
		if (board[row][i] == 1 || board[i][col] == 1) {
			return false
		}
	}
	for j:=0; j<N; j++{
		for k:=0; k<N; k++ {
			if ((j+k == row+col)|| (j-k == row-col)){
				if board[j][k]== 1 {
					return false
				}
			}
		}
	}
	return true
}

func nqueens(n int) bool{
	if (n==0){
		return true
	}
	for i:=0; i<N; i++ {
		for j:=0; j<N; j++ {
			if isSafe(i,j){
				board[i][j] = 1
				if nqueens(n-1){
					return true
				}
				board[i][j] = 0
			}
		}
	}
	return false
}

func main(){
	nqueens(N)
	print(board)
}