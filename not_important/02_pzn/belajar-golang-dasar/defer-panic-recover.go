package main

import "fmt"

// digunakan untuk memastikan bahwa func logging akan dijalankan meskipun terjadi error
// mirip seperti finally pada try catch
func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging()
	fmt.Println("Run Application")
}

func endApp() {
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}
}

// defer panic, defer awal panic akhir, melanggar peraturan akan membuat defer tidak tereksekusi
func main() {
	runApplication()
	fmt.Println("EOL")
}

func recoverFailApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}

	message := recover()
	fmt.Println("Terjadi Error", message)
}

func recoverSuccessApp(error bool) {
	defer endAppRecover()
	if error {
		panic("ERROR")
	}

}

func endAppRecover() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Terjadi Error", message)
}
