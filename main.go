package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func scanDirectory(path string) error { // створюємо рекурсію (path-шлях) і повертаємо помилку
	fmt.Println(path)

	files, err := ioutil.ReadDir(path) //створюємо змінні в які будемо отримувати файли
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files { // створюємо цикл з нашими файлами
		filePath := filepath.Join(path, file.Name())

		if file.IsDir() { //якщо знаходить першу директорію(папку)
			err := scanDirectory(filePath) // знову запускаємо скан але вже в цій папці що знайшли
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println(filePath)
		}

	}
	return nil
}

func main() {
	err := scanDirectory("my_directory") // зтворюємо зміну "помилку", якщо буде помилка і викликати функцію скан
	// з каталогом в якому проводимо сканування
	if err != nil { //  і обробляємо помилку
		log.Fatal(err)
	}
}

/*func main() {
	files, err := ioutil.ReadDir("my_directory") //створюємо змінні в які будемо отримувати файли
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Dir: ", file.Name())
		} else {
			fmt.Println("File: ", file.Name())
		}
	}
}*/
