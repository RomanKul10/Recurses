package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func reportPanic() {
	p := recover() //створюємо функцію і викликаємо рековер
	if p == nil {  // якщо все ок(nil) - повертаємо функцію
		return
	}
	err, ok := p.(error) // створємо зміні для помилки
	if ok {
		fmt.Println(err) //друкуємо помилку
	} else {
		panic(p) //якщо якась паніка але не error, то виводимо саму паніку
	}
}

func scanDirectory(path string) { // створюємо рекурсію (path-шлях) і повертаємо помилку
	fmt.Println(path)

	files, err := ioutil.ReadDir(path) //створюємо змінні в які будемо отримувати файли
	if err != nil {
		panic(err)
	}
	for _, file := range files { // створюємо цикл з нашими файлами
		filePath := filepath.Join(path, file.Name())

		if file.IsDir() { //якщо знаходить першу директорію(папку)
			scanDirectory(filePath) // знову запускаємо скан але вже в цій папці що знайшли

		} else {
			fmt.Println(filePath)
		}

	}
}

func main() {
	defer reportPanic() //defer відкладений виклик, тобто він викличиться тоді коли закінчиться робота scanDirectory

	scanDirectory("my_directory")
}

/*func main() {
	err := scanDirectory("my_directory") // зтворюємо зміну "помилку", якщо буде помилка і викликати функцію скан
	// з каталогом в якому проводимо сканування
	if err != nil { //  і обробляємо помилку
		log.Fatal(err)
	}
}*/

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
