package main

import (
	
	"os"
)

func atoi(str string) (int, bool) {
	sign := 1
	num := 0

	for i, v := range str {
		if i == 0 && v == '+' {
			sign = 1
		} else if i == 0 && v == '-' {
			sign = -1
		} else if v >= 48 && v <= 57 {
			num = num*10 + int(v-48)
		} else {
			return 0, true
		}
	}
	result:=sign*num

	// handling overflows in atoi
	if str[0] == '-' && result > 0 {
		return 0, true // Overflow detected
	} else if str[0] != '-' && result < 0 {
		return 0, true // overflow detected
	} 
	
	return result, false // no overflow

}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func modulo(a, b int) int {
	return a % b
}

func divide(a, b int) int {
	return a / b
}

func subtract(a, b int) int {
	return a - b
}

func itoa(num int) string {

	if num == 0 {
		return "0"
	}
	sign := ""
	if num < 0 {
		sign = "-"
		num = -num
	}
	slais := []byte{}
	for num > 0 {
		digit := num % 10
		slais = append(slais, byte(digit)+48)
		num /= 10
	}

	for i := 0; i < len(slais)-1; i++ {
		for j := i + 1; j < len(slais); j++ {
			slais[i], slais[j] = slais[j], slais[i]
		}
	}
	return sign + string(slais)
}

func validate(str1 string) {
	for i, v := range str1 {
		if i == 0 && v == '-' || v == '+'{
			continue
		}
		if v < 48 || v > 57 {
			return
		}
	}
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	value1 := os.Args[1]
	ops := os.Args[2]
	value3 := os.Args[3]

	// handling invalid values
	validate(value1)
	validate(value3)

	// handling invalid operators
	if ops != "+" && ops != "-" && ops != "/" && ops != "%" && ops != "*" {
		return
	}

	//Converting value1 and value2 to integer
	num1, overflow := atoi(value1)
	if overflow {
		return
	}
	num2, overflow := atoi(value3)
	if overflow {
		return
	}

	// implementing operations
	var result int
	if ops == "+" {
		if num1 > 0 && num2 > 0 && add(num1, num2) < 0 {
			return
		}
		result = add(num1, num2)
	} else if ops == "-" {
		if num1 < 0 && num2 < 0 && subtract(num1, num2) > 0 {
			return
		} 
		result = subtract(num1, num2)
	} else if ops == "/" {
		if num2 == 0{
			os.Stdout.WriteString("No division by 0\n")
			return
		}		
		result = divide(num1, num2)
	} else if ops == "%" {
		if num2 == 0{
			os.Stdout.WriteString("No modulo by 0\n")
			return
		}
		result = modulo(num1, num2)
	} else if ops == "*" {
		if num1 != 0 && multiply(num1, num2) / num1 != num2 {
			return
		}
		result = multiply(num1, num2)
	}
	finalResult := itoa(result)
	// fmt.Println(finalResult)
	os.Stdout.WriteString(finalResult + "\n")
}
