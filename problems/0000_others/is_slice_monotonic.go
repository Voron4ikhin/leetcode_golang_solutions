package _000_others

//Является ли слайс монотонным?
//Монотонная функция - функция одной переменной, определённая на некотором
//подмножестве действительных чисел, которая либо везде (на области своего
//определения) не убывает, либо везде не возрастает.

func IsMonotonic(in []int) bool {
	isUp, isDown := true, true
	prev := in[0]
	for i := 1; i < len(in); i++ {
		if isUp {
			if prev > in[i] {
				isUp = false
			}
		}
		if isDown {
			if prev < in[i] {
				isDown = false
			}
		}
		prev = in[i]
	}
	return isUp || isDown
}
