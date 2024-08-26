package lasagna

func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

func Quantities(layers []string) (noodles int, sauce float64) {
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles += 50
		}
		if layers[i] == "sauce" {
			sauce += 0.2
		}
	}
	return
}

func AddSecretIngredient(friendsList []string, myList []string) []string {
	return append(myList, friendsList[len(friendsList)-1])
}

func ScaleRecipe(quantities []float64, amount int) []float64 {
	var newQuantities = make([]float64, len(quantities))
	for i, v := range quantities {
		newQuantities[i] = (v / 2) * float64(amount)
	}
	return newQuantities
}
