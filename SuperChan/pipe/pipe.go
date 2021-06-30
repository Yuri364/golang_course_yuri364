package pipe

func NewPipe(ch1 chan string, ch2 chan string, fn func(string) string) {
    for v := range ch1 {
        ch2 <- fn(v)
    }
}