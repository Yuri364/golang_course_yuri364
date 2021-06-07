// package main уже объявлен. 
func main() {
    // Все указанные в тексте задачи пакеты уже импортированы (strconv, bufio, os, io)
    // Другие использовать нельзя.

    var sl []string
    var r int64
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        n := scanner.Text()

        sl = append(sl, n)
    }

    for _, v := range sl {
        i, _ := strconv.ParseInt(v, 10, 64)
        r = r+i
    }

    io.WriteString(os.Stdout,strconv.Itoa(int(r)))
}