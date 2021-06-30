package fanout

import "sync"

func NewFanOut(in chan string, chansOut []chan string) {
    var wg sync.WaitGroup

    wg.Add(len(chansOut))

    fn := func(out chan string) {
        defer wg.Done()
        for v := range in {
            out <- v
        }
    }

    for _, c := range chansOut {
        go fn(c)
    }

    wg.Wait()
}
