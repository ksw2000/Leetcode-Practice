// 0 ms	2.7 MB
// URL : https://leetcode.com/problems/valid-sudoku/
package main
import "fmt"

func isValidSudoku(board [][]byte) bool {
    var marked [10]int

    //sub-boxes
    for i:=0; i<3; i++{
        for j:=0; j<3; j++{
            for k:=0; k<10; k++{
                marked[k]=0
            }
            for k:=0; k<3; k++{
                for m:=0; m<3; m++{
                    if board[3*i+k][3*j+m]!='.'{
                        if marked[board[3*i+k][3*j+m]-'0']==0{
                            marked[board[3*i+k][3*j+m]-'0']=1
                        }else{
                            return false
                        }
                    }
                }
            }
        }
    }

    //row
    for i:=0; i<9; i++{
        for j:=0; j<10; j++{
            marked[j]=0
        }
        for j:=0; j<9; j++{
            if board[i][j]!='.'{
                if marked[board[i][j]-'0']==0{
                    marked[board[i][j]-'0']=1
                }else{
                    return false
                }
            }
        }
    }

    //col
    for i:=0; i<9; i++{
        for j:=0; j<10; j++{
            marked[j]=0
        }
        for j:=0; j<9; j++{
            if board[j][i]!='.'{
                if marked[board[j][i]-'0']==0{
                    marked[board[j][i]-'0']=1
                }else{
                    return false
                }
            }
        }
    }


    return true
}

func main(){
    var sudoku = [][]byte{
        {'8','3','.','.','7','.','.','.','.'},
        {'6','.','.','1','9','5','.','.','.'},
        {'.','9','8','.','.','.','.','6','.'},
        {'8','.','.','.','6','.','.','.','3'},
        {'4','.','.','8','.','3','.','.','1'},
        {'7','.','.','.','2','.','.','.','6'},
        {'.','6','.','.','.','.','2','8','.'},
        {'.','.','.','4','1','9','.','.','5'},
        {'.','.','.','.','8','.','.','7','9'},
    }
    fmt.Println(isValidSudoku(sudoku))

    sudoku2 := [][]byte{
      {'5','3','.','.','7','.','.','.','.'},
      {'6','.','.','1','9','5','.','.','.'},
      {'.','9','8','.','.','.','.','6','.'},
      {'8','.','.','.','6','.','.','.','3'},
      {'4','.','.','8','.','3','.','.','1'},
      {'7','.','.','.','2','.','.','.','6'},
      {'.','6','.','.','.','.','2','8','.'},
      {'.','.','.','4','1','9','.','.','5'},
      {'.','.','.','.','8','.','.','7','9'},
    }
    fmt.Println(isValidSudoku(sudoku2))
}
