/*
 * Created on Wed Sep 18 2024
 *
 * Copyright (c) 2024 Rizki Hadiaturrasyid
 */

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	inputFiles, err := os.ReadDir("input/")
	if err != nil {
		fmt.Println(err)
	}

	os.Mkdir("output", 0755)

	for _, file := range inputFiles {
		inputName := file.Name()
		outputName := strings.Split(file.Name(), ".")[0] + "_padded" + ".jpg"
		size := "4896x4896"

		cmd2 := exec.Command("convert", "input/"+inputName, "-gravity", "center", "-background", "black", "-extent", size, "output/"+outputName)
		err = cmd2.Run()
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("Done!")
}
