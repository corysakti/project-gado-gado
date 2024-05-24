package standard_library

import (
	"bufio"
	"os"
)

func MainBufioWrite() {
	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("hello world\n")
	_, _ = writer.WriteString("Selamat belajar\n")
}
