package main

import (
	"bytes"
	"log"
	"os"
)

func main() {
	f, err := createFile("./file.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer removeFile(f)
	err = writeToFile(f)
	if err != nil {
		log.Println(err)
		return
	}
	err = readFromFile(f)
	if err != nil {
		log.Println(err)
	}
}

func createFile(p string) (*os.File, error) {
	log.Println("creating")
	file, err := os.Create(p)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func writeToFile(f *os.File) error {
	log.Println("writing")
	content := "some content"
	_, err := f.Write([]byte(content))

	return err
}

func readFromFile(f *os.File) error {
	log.Println("reading")
	reader := &bytes.Buffer{}
	f.Seek(0, 0)
	_, err := reader.ReadFrom(f)
	if err != nil {
		return err
	}

	log.Printf("file content: \"%s\"\n", reader.String())
	return nil
}

func removeFile(f *os.File) error {
	log.Println("closing")
	err := f.Close()
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("removing")
	os.Remove(f.Name())

	return nil
}
