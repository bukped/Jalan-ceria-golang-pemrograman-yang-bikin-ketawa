fruits := []string{"apple", "banana", "cherry", "durian"}
slice := fruits[0:2:2]
fmt.Println(slice)      // ["apple", "banana"]
fmt.Println(len(slice)) // 2
fmt.Println(cap(slice)) // 2
