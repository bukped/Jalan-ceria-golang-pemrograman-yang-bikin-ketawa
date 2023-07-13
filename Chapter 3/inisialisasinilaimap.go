var data map[string]int
data["one"] = 1
// akan muncul error!
data = map[string]int{}
data["one"] = 1
// tidak ada error

// cara horizontal
var chickenl = map[string]int{"januari“: 50, "februari": 40}
// cara vertical
var chickenz = map[string]int{
"januari": 50,
"februari": 40,
}

