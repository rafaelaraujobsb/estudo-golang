package main

func main() {
	funcsESalaraios := map[string]float64{
		"Zé":    12362.98,
		"Maria": 12370.,
	}

	funcsESalaraios["João"] = 14523.78

	delete(funcsESalaraios, "Não existe")
}
