package main

func romanToInt(s string) int {
	strLen := len(s)
	val := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'I':
			{
				if i < strLen-1 {
					switch s[i+1] {
					case 'V':
						{
							val += 4
							i++
							// break
						}
					case 'X':
						{
							val += 9
							i++
							// break
						}
					default:
						{
							val += 1
							// break
						}
					}
				} else {
					val += 1
				}
				// break
			}
		case 'V':
			{
				val += 5
				// break
			}
		case 'X':
			{
				if i < strLen-1 {
					switch s[i+1] {
					case 'L':
						{
							val += 40
							i++
							// break
						}
					case 'C':
						{
							val += 90
							i++
							// break
						}
					default:
						{
							val += 10
							// break
						}
					}
				} else {
					val += 10
				}
				// break
			}
		case 'L':
			{
				val += 50
				// break
			}
		case 'C':
			{
				if i < strLen-1 {
					switch s[i+1] {
					case 'D':
						{
							val += 400
							i++
							// break
						}
					case 'M':
						{
							val += 900
							i++
							// break
						}
					default:
						{
							val += 100
							// break
						}
					}
				} else {
					val += 100
				}
				// break
			}
		case 'D':
			{
				val += 500
				// break
			}
		case 'M':
			{
				val += 1000
				// break
			}
		}
	}

	return val
}
