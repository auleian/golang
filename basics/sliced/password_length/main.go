package main

func isValidPassword(password string) bool {
	if len(password) < 5 || len(password) > 12  {
		return false
	}
	
	cap := false
	num := false

	for _, char := range []byte(password) {
		if char >= 'A' && char <= 'Z' {
			cap = true
		}
		if char >= '0' && char <= '9' {
			num = true
		}
		if cap && num {
			break
		}
	}

	return cap && num
	
}
