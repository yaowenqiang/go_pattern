package singleton

import (
    "fmt"
    "errors"
)

func _main() {
    addN := func(m int) func(int) int {
        return func(n int) int {
            return m + n
        }
    }

    addFive := addN(5)
    result := addFive(6)
    fmt.Println(result)
    fmt.Println(divisibleBy(10,3))
    fmt.Println(divisibleBy(10,0))

    /*
    err := doesReturnError()

    if err != nil {
        panic(err)
    }
    */
    fmt.Printf("%d\n", sum())
    fmt.Printf("%d\n", sum(1,2,3))
    fmt.Printf("%d\n", sum(1,2,3,4,5,6))
}

func doesReturnError() error {
    err := errors.New("This function simply returns an error")
    return err
}


func sum(args ...int) (result int) {
    for _, v := range args {
        result += v
    }
    return
}

func divisibleBy(n, divisor int) (bool, error) {
    if divisor == 0 {
        return false, errors.New("A number cannot be divided by zero")
    }
    return (n % divisor == 0), nil
}
