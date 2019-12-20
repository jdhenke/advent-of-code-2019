package input

import (
	"bufio"
	"os"
	"strconv"
)

func ForEachNumInFile(path string, fn func(i int)) error {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = f.Close()
	}()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		fn(i)
	}
	return scanner.Err()
}
