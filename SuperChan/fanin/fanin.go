package fanin

import "sync"

func NewFanIn(chansIn []chan string, out chan string) {
    var wg sync.WaitGroup

    wg.Add(len(chansIn))

    fn := func(c chan string) {
        defer wg.Done()
        for v := range c {
            out <- v
        }
    }

    for _, c := range chansIn {
        go fn(c)
    }

    wg.Wait()
}
