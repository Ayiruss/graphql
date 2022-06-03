package helper

import "fmt"

func ToFloat(price int64) string {
	fValue := float64(price) / 100
	return fmt.Sprintf("%.2f", fValue)
}

func ToInt(price float64) int64 {
	return int64(price * 100)
}
