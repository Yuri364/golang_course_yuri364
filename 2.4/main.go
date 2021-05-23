type testStruct struct {
    On bool
    Ammo int
    Power int
}

func (s *testStruct) Shoot() bool {
    if s.On == false {
        return false
    }

    if s.Ammo > 0 {
        s.Ammo--
        return true
    }

    return false
}

func (s *testStruct) RideBike() bool {
    if s.On == false {
        return false
    }

    if s.Power > 0 {
        s.Power--
        return true
    }

    return false
}

func main() {

testStruct := new (testStruct)
/*
 * Экземпляр созданной вами структуры необходимо передать в качестве
 * аргумента функции testStruct, которая выполнит проверку соблюдения
 * всех условий задания/
 */

// }