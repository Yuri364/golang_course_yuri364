package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/**
	На стандартный ввод подаются данные о продолжительности периода (формат приведен в примере).
	Кроме того, вам дана дата в формате Unix-Time: 1589570165 в виде константы типа int64
	(наносекунды для целей преобразования равны 0).

	Требуется считать данные о продолжительности периода, преобразовать их в тип Duration, а затем вывести
	(в формате UnixDate) дату и время, получившиеся при добавлении периода к стандартной дате.

	Небольшая подсказка: базовую дату необходимо явно перенести в зону UTC с помощью одноименного метода.
 */

// вам это понадобится
const now = 1589570165
var re = regexp.MustCompile(`(\d+) мин\. (\d+) сек\.`)

func main() {
	startDate := time.Unix(now, 0)

	timeString, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	timeString = strings.TrimSpace(timeString)

	times := re.FindStringSubmatch(timeString)

	min, _ := strconv.Atoi(times[1])
	sec, _ := strconv.Atoi(times[2])

	d := time.Minute*time.Duration(min) + time.Second*time.Duration(sec)
	r := startDate.Add(d).UTC().Format(time.UnixDate)

	fmt.Println(r)
}





