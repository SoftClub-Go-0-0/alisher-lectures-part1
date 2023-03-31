package main

func removeLongStrings(data map[string]string) int {
	var cmt int=0
	for k,v := range data{
		if len(v) > 30 {
			delete(data,k)
		}
	
	}
	for k,_ := range data{
		k+=""
		cmt++
	}
	return cmt
	
	// удалите все элементы, длина значения которых превышает 10 символов
    // вернуть количество элементов после удаления
}

func main() {
	
}