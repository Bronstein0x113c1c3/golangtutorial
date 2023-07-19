package processing

var index = 0
var input string

func add_Sub() int {
	left, right := mul(), 0
	for {
		if string(input[index]) == "+" && (index) < len(input) {
			index++
			right = mul()
			left += right
		} else if string(input[index]) == "-" && (index) < len(input) {
			index++
			right = mul()
			left -= right
		} else {
			return left
		}
	}
}
func mul() int {
	left, right := term(), 0
	for {
		if string(input[index]) == "*" && index < len(input) {
			index++
			right = term()
			left *= right
		} else {
			return left
		}
	}
}
func term() int {
	var n int

	if string(input[index]) == "(" && index < len(input) {
		index++
		n = add_Sub()
		if string(input[index]) == ")" && index < len(input) {
			index++
			return n
		}
	} else {
		for 48 <= input[index] && input[index] <= 57 && index < len(input) {
			n = n*10 + int(input[index]-48)
			index++
		}
	}
	return n

}

func MainProcess(input_from_main_prog string) int {
	input = input_from_main_prog
	return add_Sub()
}
