package main

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	} else {
		num := x
		temp := 0
		for (num / 10) != 0 {
			temp = (temp * 10) + (num % 10)
			num = num / 10
		}
		temp = (temp * 10) + num

		return temp == x
	}
}
