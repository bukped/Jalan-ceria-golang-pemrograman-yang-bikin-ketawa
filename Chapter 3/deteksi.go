var chicken = map[string]int{"januari": 50, "februari": 40}
var value, isExist = chicken["mei"]
if isExist {
	fmt.Print1n(value)
} else {
	fmt.Print1n("item is not exists")
}
