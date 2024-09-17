package main

func main() {
	var minhaVar interface{} = "Kevin"
	println(minhaVar.(string))
	res, ok := minhaVar.(int)
	println(res, ok)
}