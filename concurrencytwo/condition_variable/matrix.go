package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const matrixSize = 250

var matrixA = [matrixSize][matrixSize]int{}

var matrixB = [matrixSize][matrixSize]int{}

var result = [matrixSize][matrixSize]int{}

var rwLock = sync.RWMutex{}
var cond = sync.NewCond(rwLock.RLocker())
var wg = sync.WaitGroup{}

func main() {
	n := time.Now()
	wg.Add(matrixSize)

	for row := 0; row < matrixSize; row++ {
		go workOutRow(row)
	}
	
	for i := 0; i < 100; i++ {
		wg.Wait()
		rwLock.Lock()
		generateRandomMatrix(&matrixA)
		generateRandomMatrix(&matrixB)

		wg.Add(matrixSize)
		rwLock.Unlock()
		cond.Broadcast()

	}

	fmt.Println(time.Since(n))
}

func workOutRow(row int) {
	rwLock.RLock()
	for {
		wg.Done()
		cond.Wait()
		for col := 0; col < matrixSize; col++ {
			for i := 0; i < matrixSize; i++ {
				result[row][col] += matrixA[row][i] * matrixB[i][col]
			}
		}
	}

}

func generateRandomMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			matrix[row][col] += rand.Intn(10) - 5
		}
	}
}
