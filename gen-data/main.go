package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	dirName, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	// getThumbCmd := `ffmpeg -i %s -vf "select=eq(n\,0)" -q:v 3 %s`

	for i := 0; i < 20; i++ {
		fileName := dirName + "/" + fmt.Sprint(i+1) + ".mp4"
		fileOutut := dirName + "/" + fmt.Sprint(i+1) + ".png"
		if _, err := os.Stat(fileName); os.IsNotExist(err) {
			fmt.Println("not exist file", fileName)
			break
		}
		// cmd := fmt.Sprintf(getThumbCmd, fileName, fileOutut)

		cmd := exec.Command("ffmpeg", "-i", fileName, "-vframes", "1", "-q:v", "2", fileOutut)
		err := cmd.Run()
		if err != nil {
			log.Fatalln("run command err", err)
		}
	}
}
