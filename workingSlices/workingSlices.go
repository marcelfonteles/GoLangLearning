package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func init() {
	fmt.Println("Surprise Motherfucker!")
}

func main() {
	dt := time.Now()
	dtString := dt.Format("02-Jan-2006 15:04:05")
	texto := [][]string{{"Surprise", "Motherfucker!"}, {"Data", dtString}}
	fmt.Println(texto)

	fileCSV, _ := os.Create("/home/marcelvieira/surprise.csv")
	fileTXT, _ := os.Create("/home/marcelvieira/surprise.txt")
	defer fileCSV.Close()
	defer fileTXT.Close()
	writerCSV := csv.NewWriter(fileCSV)
	defer writerCSV.Flush()

	for _, value := range texto {
		for _, innerValue := range value {
			fmt.Fprint(fileTXT, innerValue)
			fmt.Fprint(fileTXT, " ")
			//fileTXT.WriteString(innerValue)
			if value[len(value)-1] == innerValue {
				fmt.Fprintln(fileTXT, "")
			}
		}
		writerCSV.Write(value)
	}
	forLoopSlice()


}

func forLoopSlice() {
	slice := []int{1, 2, 4, 3, 5, 6, 7}
	fmt.Println(slice)
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}
