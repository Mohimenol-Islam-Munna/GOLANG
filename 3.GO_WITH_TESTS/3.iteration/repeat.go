package iteration 

func Repeat(l string) string {
	count := 5
	result := ""

	for i:= 0; i<count; i++{
		result += l
	}

	return result
}

