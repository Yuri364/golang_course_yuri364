/*
    Внутри функции main (объявлять ее не нужно) вы должны объявить функцию вида func(uint) uint, которую внутри
    функции main в дальнейшем можно будет вызвать по имени fn.

    Функция на вход получает целое положительное число (uint), возвращает число того же типа в котором отсутствуют нечетные
    цифры и цифра 0. Если же получившееся число равно 0, то возвращается число 100. Цифры в возвращаемом числе имеют тот же
    порядок, что и в исходном числе.
 */

fn := func(n uint) uint {
    var r uint
    var _n string

    for n!=0 {
        r = n%10
        if r%2==0 && r!=0 {
            _n = strconv.FormatUint(uint64(r), 10)+_n
        }
        n /= 10
    }

    num,_ := strconv.ParseUint(_n, 10, 32)

    if num == 0 {
        num = 100
    }

    return uint(num)
}