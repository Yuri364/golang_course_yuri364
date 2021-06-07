package main

import (
    "encoding/json" // пакет используется для проверки ответа, не удаляйте его
    "fmt"           // пакет используется для проверки ответа, не удаляйте его
    "os"            // пакет используется для проверки ответа, не удаляйте его
)

func main() {
    value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
    // все полученные значения имеют тип пустого интерфейса
    var a, b float64
    var ok bool

    if a, ok = value1.(float64); !ok {
        fmt.Printf("value=%v: %T", value1, value1)
        os.Exit(0)
    }
    if b, ok = value2.(float64); !ok {
        fmt.Printf("value=%v: %T", value2, value2)
        os.Exit(0)
    }
    switch op, _ := operation.(string); {
    case op == "+":
        fmt.Printf("%.4f", a + b)
    case op == "-":
        fmt.Printf("%.4f", a - b)
    case op == "/":
        fmt.Printf("%.4f", a / b)
    case op == "*":
        fmt.Printf("%.4f", a * b)
    default:
        fmt.Printf("неизвестная операция")
    }
}