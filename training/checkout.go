package main

func checkoutTotal(subtotal int, hasCoupon bool, couponType string, isVIP bool, isFirstOrder bool, country string) int {
	// your code here
	if hasCoupon == true {
		if couponType == "PERCENT10" {
			subtotal *= 0.9
		}else if couponType == "FLAT500" {
			subtotal -= 500
		}else{
			
		}
	}
	return 0
}
