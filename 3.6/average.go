package main

import (
    "encoding/json"
    "io/ioutil"
    "os"
)

type Student struct {
    Rating []int
}

type Base struct {
    Students []Student
}

type Average struct {
    Average float64
}

func main() {
    data, _ := ioutil.ReadAll(os.Stdin)
    var r []int

    var s Base
    var _ = json.Unmarshal(data, &s)

    for _, v := range s.Students {
        r = append(r, len(v.Rating))
    }

    n := len(r)
    sum := 0
    for i := 0; i < n; i++ {
        sum += r[i]
    }

    avg := (float64(sum)) / (float64(n))

    out, _ := json.MarshalIndent(Average{avg}, "", "    ")

    _, _ = os.Stdout.Write(out)

}