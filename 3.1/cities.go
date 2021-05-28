/*
 * Группировка городов по численности населения в тысячах человек
 * от 10 до 100, от 100 до 1000 и более 1000:
 * groupCity map[int][]string{
 *	 10: []string{...},
 *	 100: []string{...},
 *	 1000: []string{...},
 * }
 *
 * Население городов в тысячах человек:
 * cityPopulation map[string]int{...}
 */

cities := append(groupCity[10], append(groupCity[1000])...)

for _, city := range cities {
	if _, ok := cityPopulation[city]; ok {
		delete(cityPopulation, city)
	}
}
