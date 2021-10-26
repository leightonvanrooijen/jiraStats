package utils

import "fmt"

func CheckError(ok bool, message string) {
	if !ok {
		fmt.Println(message)
	}

}

func ConvertString(data interface{}) string {
	if data != nil {
		return data.(string)
	}
	return ""
}
