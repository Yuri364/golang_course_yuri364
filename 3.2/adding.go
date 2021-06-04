func adding(a string, b string) int64 {
    var _a,_b string

    for _,v := range []rune(a) {
        if unicode.IsDigit(v) {
        _a += string(v)
        }
    }

    for _,v := range []rune(b) {
        if unicode.IsDigit(v) {
        _b += string(v)
        }
    }

    n1,_ := strconv.ParseInt(_a, 10, 64)
    n2,_ := strconv.ParseInt(_b, 10, 64)


    return n1+n2
}