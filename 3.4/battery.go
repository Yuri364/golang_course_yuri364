package main

import (
    "fmt" // пакет используется для проверки ответа, не удаляйте его
    "strings"
)

type Battery string

func (b Battery) String() string {
    return fmt.Sprintf("[%10s]", strings.Repeat("X", strings.Count(string(b), "1")))
}

func main() {
    // batteryForTest - не забывайте об имени
    var batteryForTest Battery
    fmt.Scan(&batteryForTest)
    // } Скобка, закрывающая функцию main() вам не видна, но она есть
