package config

import (
	"bufio"
	"os"
)

func GenerateBanner() {

	readFile, err := os.Open("banner.txt")

	if err != nil {
		return
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	println("------------------------------------------------------------------------------")
	for fileScanner.Scan() {
		println(fileScanner.Text())
	}
	println("------------------------------------------------------------------------------")

	err = readFile.Close()

	if err != nil {
		return
	}

}
