package Week_04

func lemonadeChange(bills []int) bool {
    five := 0
    ten := 0
    for _, v := range bills {
        switch v {
            case 5:
            five++
            case 10:
            if five < 1 {
                return false
            }
            five--
            ten++
            case 20:
            if ten > 0 && five > 0 {
                ten--
                five--
            } else if five >= 3 {
                five -= 3
            } else {
                return false
            }
        }
    }
    return true
}