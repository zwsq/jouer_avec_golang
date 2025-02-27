package main

func main() {
	numero := 1
	for numero < 101 {
		if numero%2 == 0 {
			println(numero)
		}
		numero++
	}
}
