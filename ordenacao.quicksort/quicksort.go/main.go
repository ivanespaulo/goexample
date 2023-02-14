// Quick Sort in Golang
package main

import (
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "time"
)

func main() {

    //debug.SetGCPercent(-1)
    qtd, _ := strconv.Atoi(os.Args[1])
    if qtd <= 1 {
        println("O valor tem que ser maior que 1!")
        return
    }

    start := time.Now()
    slice := generateSlice(qtd)
    fmt.Println("slice load:", time.Now().Sub(start))

    t1 := time.Now()
    //fmt.Println("\n--- Unsorted --- \n\n", slice)
    //quicksort(slice, 0, len(slice)-1)
    quicksort(slice)
    t2 := time.Now()
    println("Time:", t2.Sub(t1).String())
    //fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

func generateSlice2(qtd int) []int {
    rand.Seed(time.Now().UnixNano())
    slice := make([]int, 0, qtd)
    times := qtd
    for i := 0; i < times; i++ {
        val := rand.Intn(20000000)
        slice = append(slice, val)
    }
    return slice
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {
    slice := make([]int, size, size)
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < size; i++ {
        slice[i] = rand.Intn(999) - rand.Intn(999)
    }
    return slice
}

func quicksort(a []int) []int {
    if len(a) < 2 {
        return a
    }

    left, right := 0, len(a)-1
    pivot := right / 2
    a[pivot], a[right] = a[right], a[pivot]

    for i, _ := range a {
        if a[i] < a[right] {
            a[left], a[i] = a[i], a[left]
            left++
        }
    }

    a[left], a[right] = a[right], a[left]

    quicksort(a[:left])
    quicksort(a[left+1:])

    return a
}
