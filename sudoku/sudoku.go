package sudoku

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func ImportFile(path string) ([][]int, error) {
	matrix := make([][]int, 0)

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			return nil, err
		}

		nums := strings.Fields(line)
		row := make([]int, 0)
		for _, v := range nums {
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			row = append(row, n)
		}

		if len(row) != 0 {
			matrix = append(matrix, row)
		}

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		if len(row) != len(matrix[0]) {
			return nil, errors.New("wrong number of elements")
		}
	}

	return matrix, nil
}

func ValidMatrix(matrix [][]int) bool {
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			iEl := i * 3
			jEl := j * 3
			uniqueNumbers := make([]int, 9)
			for k := 0; k <= 2; k++ {
				for m := 0; m <= 2; m++ {
					bigI := iEl + k
					bigJ := jEl + m
					val := matrix[bigI][bigJ]
					if val > 0 && val < 10 {
						if uniqueNumbers[val-1] == 1 {
							fmt.Println("Appeared two times: ", val)
							return false
						} else {
							uniqueNumbers[val-1] = 1
						}
					} else {
						fmt.Println("Invalid value: ", val)
						return false
					}
				}
			}
		}
	}

	for i := 0; i < 9; i++ {
		sum := 0
		for j := 0; j < 9; j++ {
			sum += matrix[i][j]
		}
		if sum != 45 {
			return false
		}
		sum = 0
	}

	for i := 0; i < 9; i++ {
		sum := 0
		for j := 0; j < 9; j++ {
			sum += matrix[j][i]
		}
		if sum != 45 {
			return false
		}
		sum = 0
	}

	return true
}

func Sudoku(argumets []string) {
	if len(argumets) != 2 {
		fmt.Println("usage: loadFile textFile size")
		return
	}

	file := argumets[1]
	matrix, err := ImportFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	if ValidMatrix(matrix) {
		fmt.Println("Correct sudoku matrix!")
	} else {
		fmt.Println("Incorrect sudoku matrix!")
	}
}
