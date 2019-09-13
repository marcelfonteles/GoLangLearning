package main

import "fmt"

func init() {

}

func main() {
	array := [5]int{1,2,3,4}
	fmt.Println(array)

	slice := []int{1,2,3,4}
	fmt.Println(slice)
	slice = append(slice, 5)
	fmt.Println(slice)

	slice = append(slice[:1], slice[3:]...)
	fmt.Println(slice)

	newSlice := []int{42,43,44,45,46,47,48,49,50,51}

	// [42 43 44 45 46]
	slice1 := newSlice[0:5]
	fmt.Println("slice1:", slice1)
	// [47 48 49 50 51]
	slice2 := newSlice[5:]
	fmt.Println("Slice2:", slice2)
	// [44 45 46 47 48]
	slice3 := newSlice[2:7]
	fmt.Println("Slice3:", slice3)
	// [43 44 45 46 47]
	slice4 := newSlice[1:6]
	fmt.Println("Slice4:", slice4)
	// [43 44 47 48 49 50 51]
	sliceExtra := append(newSlice[1:3], newSlice[5:]...)
	fmt.Println("SliceExtra:", sliceExtra)
	// Multiple Operations
	sliceExtra2 := append(newSlice, 52)
	fmt.Println("SliceExtra2:", sliceExtra2)

	sliceExtra2 = append(sliceExtra2, []int{53,54,55}...)
	fmt.Println("SliceExtra2:", sliceExtra2)
	y := []int{56,57,58,59,60}
	sliceExtra2 = append(sliceExtra2, y...)
	fmt.Println("SliceExtra2:", sliceExtra2)


	// Delete from Slice: Using APPEND with SLICING

	sliceMake := make([]int, 5)
	fmt.Println("Using make:", sliceMake)
	fmt.Println("Using make: (len)", len(sliceMake))
	fmt.Println("Using make: (cap)", cap(sliceMake))

	sliceMake2 := make([]int, 3, 5)
	fmt.Println("Using make:", sliceMake2)
	fmt.Println("Using make: (len)", len(sliceMake2))
	fmt.Println("Using make: (cap)", cap(sliceMake2))

	// Slice And Maps
	mapEstados := map[string]string{
		"AC": "Acre",
		"AL": "Alagoas",
		"AP": "Amapá",
		"AM": "Amazonas",
		"BA": "Bahia",
		"CE": "Ceará",
		"DF": "Distrito Federal",
		"ES": "Espírito Santo",
		"GO": "Goiás",
		"MA": "Maranhão",
		"MT": "Mato Grosso",
		"MS": "Mato Grosso do Sul",
		"MG": "Minas Gerais",
		"PA": "Pará",
		"PB": "Paraíba",
		"PR": "Paraná",
		"PE": "Pernambuco",
		"PI": "Piauí",
		"RJ": "Rio de Janeiro",
		"RN": "Rio Grande do Norte",
		"RS": "Rio Grande do Sul",
		"RO": "Rondônia",
		"RR": "Roraima",
		"SC": "Santa Catarina",
		"SP": "São Paulo",
		"SE": "Sergipe",
		"TO": "Tocantins",
	}
	fmt.Println(mapEstados)
	fmt.Println(len(mapEstados))

	sliceEstados := make([]string, 0, 27)
	fmt.Println(sliceEstados)
	for _, value := range mapEstados {
		sliceEstados = append(sliceEstados, value)
	}
	fmt.Println(sliceEstados)

	// Slice of slice of string

	sliceSliceString := [][]string{}
	/*
	[
		["Marcel", "Rocha", "Fonteles", "Vieira"]
		["Marcel", "Rocha", "Fonteles", "Vieira"]
		["Marcel", "Rocha", "Fonteles", "Vieira"]
	]
	 */
	sliceSliceString = append(sliceSliceString, []string{"Marcel", "Rocha", "Fonteles", "Vieira"})
	sliceSliceString = append(sliceSliceString, []string{"Larissa", "Rocha", "Fonteles", "Vieira"})
	sliceSliceString = append(sliceSliceString, []string{"Silvana", "Rocha", "Fonteles", "Albuquerque"})
	sliceSliceString = append(sliceSliceString, []string{"Ricardo", "Damasceno", "de", "Albuquerque"})

	fmt.Println(sliceSliceString)

	fmt.Println("---------------")
	for _, value := range sliceSliceString {
		fmt.Println(value)
	}
	fmt.Println("---------------")
	for _, value := range sliceSliceString {
		for _, value2 := range value {
			fmt.Println(value2)
		}
	}

	// Map: key -> string / values -> slice of string

	mapWhatever := map[string][]string{
		"Marcel": []string{"Cachaça", "Cerveja", "Alcool"},
		"James": []string{"Shaken not stirred", "Martinis", "Women"},
	}

	fmt.Println(mapWhatever)

	for key, value := range mapWhatever {
		fmt.Println("Pessoa:",key,"/ Preferência:", value)
	}
	fmt.Println("--------------")
	printMap(mapWhatever)

	mapWhatever["Cairo"] = []string{"Cerveja", "Fortaleza", "Maconha"}

	printMap(mapWhatever)

	delete(mapWhatever, "Cairo")
	printMap(mapWhatever)


}

func printMap(mapWhatever map[string][]string) {
	for key, value := range mapWhatever {
		fmt.Printf("Pessoa: %v / Preferência: ", key)
		for _, value2 := range value {
			fmt.Printf("%v", value2)
			if value2 != value[len(value)-1] {
				fmt.Print(", ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Println("------------")
}
