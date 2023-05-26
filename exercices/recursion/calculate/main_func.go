package calculate

var arr [1001]int

const m int = 10e9 + 7

func Calculating(a, b int) int {
	arr[0] = 1
	if b <= 0 {
		return arr[0]

	} else {
		if b%2 == 0 {
			x := Calculating(a, b/2) % m
			arr[b] = (x * x) % m
			return arr[b]
		} else {
			x := (Calculating(a, b-1) % m)
			arr[b] = (x % m) * (a % m) % m
			return arr[b]
		}
	}
}
