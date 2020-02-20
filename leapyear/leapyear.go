package leapyear

// IsLeapYear  ...function checks if the entered year is a leap year.
func IsLeapYear(year uint16) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}
