package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	FirstName string
	LastName  string
	Phone     string
	Address   string
}

var adresDefteri []Person

func getInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func findPersonIndex(firstName, lastName string) int {
	for index, person := range adresDefteri {
		if person.FirstName == firstName && person.LastName == lastName {
			return index
		}
	}
	return -1
}

func main() {
	for {
		fmt.Println("\nAdres Defteri Uygulaması")
		fmt.Println("1. Yeni Kişi Ekle")
		fmt.Println("2. Kişi Bilgilerini Düzenle")
		fmt.Println("3. Kişi Sil")
		fmt.Println("4. Kişi Listesi Görüntüle")
		fmt.Println("5. Çıkış")
		choice := getInput("Bir seçenek girin: ")

		switch choice {
		case "1":
			addPerson()
		case "2":
			editPerson()
		case "3":
			deletePerson()
		case "4":
			listPersons()
		case "5":
			fmt.Println("Çıkılıyor...")
			return
		default:
			fmt.Println("Geçersiz seçenek. Tekrar deneyin.")
		}
	}
}

func addPerson() {
	firstName := getInput("Ad: ")
	lastName := getInput("Soyad: ")
	phone := getInput("Telefon: ")
	address := getInput("Adres: ")

	if firstName == "" || lastName == "" || phone == "" || address == "" {
		fmt.Println("Tüm alanları doldurmanız gerekmektedir.")
		return
	}

	adresDefteri = append(adresDefteri, Person{firstName, lastName, phone, address})
	fmt.Println("Kişi başarıyla eklendi.")
}

func editPerson() {
	firstName := getInput("Düzenlemek istediğiniz kişinin adı: ")
	lastName := getInput("Düzenlemek istediğiniz kişinin soyadı: ")

	index := findPersonIndex(firstName, lastName)
	if index == -1 {
		fmt.Println("Kişi bulunamadı.")
		return
	}

	fmt.Println("Mevcut bilgiler: ", adresDefteri[index])
	newFirstName := getInput("Yeni Ad (Boş bırakılırsa eski değer korunur): ")
	newLastName := getInput("Yeni Soyad (Boş bırakılırsa eski değer korunur): ")
	newPhone := getInput("Yeni Telefon (Boş bırakılırsa eski değer korunur): ")
	newAddress := getInput("Yeni Adres (Boş bırakılırsa eski değer korunur): ")

	if newFirstName != "" {
		adresDefteri[index].FirstName = newFirstName
	}
	if newLastName != "" {
		adresDefteri[index].LastName = newLastName
	}
	if newPhone != "" {
		adresDefteri[index].Phone = newPhone
	}
	if newAddress != "" {
		adresDefteri[index].Address = newAddress
	}

	fmt.Println("Kişi bilgileri başarıyla güncellendi.")
}

func deletePerson() {
	firstName := getInput("Silmek istediğiniz kişinin adı: ")
	lastName := getInput("Silmek istediğiniz kişinin soyadı: ")

	index := findPersonIndex(firstName, lastName)
	if index == -1 {
		fmt.Println("Kişi bulunamadı.")
		return
	}

	confirmation := getInput("Bu kişiyi silmek istediğinize emin misiniz? (E/H): ")
	if strings.ToLower(confirmation) == "e" {
		adresDefteri = append(adresDefteri[:index], adresDefteri[index+1:]...)
		fmt.Println("Kişi başarıyla silindi.")
	} else {
		fmt.Println("Silme işlemi iptal edildi.")
	}
}

func listPersons() {
	if len(adresDefteri) == 0 {
		fmt.Println("Adres defteri boş.")
		return
	}

	fmt.Println("Adres Defterindeki Kişiler:")
	for _, person := range adresDefteri {
		fmt.Printf("Ad: %s\n Soyad: %s\n Telefon: %s\n Adres: %s\n", person.FirstName, person.LastName, person.Phone, person.Address)
	}
}
