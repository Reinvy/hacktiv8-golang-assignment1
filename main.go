package main

import (
	"assignment1/models"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var rawArgs = os.Args

	var args, err = strconv.Atoi(rawArgs[1:][0])

	var i = args - 1

	if err != nil || i < 0 {
		fmt.Println("Argument Tidak Valid!!")
		fmt.Println("Mohon hanya input argument berupa bilangan bulat")
	} else {
		displayData(i)
	}

}

func displayData(i int) {
	person := models.ListPerson
	if i >= len(person) {
		fmt.Println("Data tidak ditemukan")
	} else {
		fmt.Println("ID		:", person[i].ID)
		fmt.Println("Nama		:", person[i].Nama)
		fmt.Println("Alamat		:", person[i].Alamat)
		fmt.Println("Pekerjaan	:", person[i].Pekerjaan)
		fmt.Println("Alasan		:", person[i].Alasan)
	}
}
