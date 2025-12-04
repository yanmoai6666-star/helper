package calculator

// CelsiusToFahrenheit 摄氏度转换为华氏度
func CelsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9.0 / 5.0) + 32.0
}

// FahrenheitToCelsius 华氏度转换为摄氏度
func FahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32.0) * 5.0 / 9.0
}

// KilogramsToPounds 千克转换为磅
func KilogramsToPounds(kg float64) float64 {
	return kg * 2.20462
}

// PoundsToKilograms 磅转换为千克
func PoundsToKilograms(lb float64) float64 {
	return lb / 2.20462
}

// KilometersToMiles 公里转换为英里
func KilometersToMiles(km float64) float64 {
	return km * 0.621371
}

// MilesToKilometers 英里转换为公里
func MilesToKilometers(miles float64) float64 {
	return miles / 0.621371
}
