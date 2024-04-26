<<<<<<< HEAD
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func listFiles(directory string) ([]string, error) {
	var files []string

	fileInfo, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	for _, file := range fileInfo {
		if file.IsDir() {
			continue
		}
		if filepath.Ext(file.Name()) == ".24bes" {
			files = append(files, file.Name())
		}
	}

	return files, nil
}

func decrypt(path, file string) error {
	key := "bhUlIshutrea98liOp"
	keyLen := len(key)
	decryptFolder := filepath.Join(path, "decrypt")

	err := os.MkdirAll(decryptFolder, os.ModePerm)
	if err != nil {
		return err
	}

	decryptBeforePath := filepath.Join(path, file)
	decryptAfterPath := filepath.Join(decryptFolder, file[:len(file)-6])

	data, err := ioutil.ReadFile(decryptBeforePath)
	if err != nil {
		return err
	}

	var content []byte
	for i, v := range data {
		content = append(content, v^byte(key[i%keyLen]))
	}

	err = ioutil.WriteFile(decryptAfterPath, content, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("[+] Decrypted: %s\n", decryptAfterPath)
	return nil
}

func main() {
	path := "/home/kali/htb/sherlock/lockpic/forela-criticaldata"
	files, err := listFiles(path)
	if err != nil {
		fmt.Printf("[-] Error: %s\n", err)
		return
	}

	filesLen := len(files)
	if filesLen == 0 {
		fmt.Println("[-] No files found for decryption.")
		return
	}

	for i, file := range files {
		err := decrypt(path, file)
		if err != nil {
			fmt.Printf("[-] Error decrypting %s: %s\n", file, err)
		}

		fmt.Printf("\r[+] %.0f%% [%s]", float64(i+1)/float64(filesLen)*100, "▓"+string((filesLen-i-1)/filesLen*10))
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("\n[+] Decryption complete.")
}
=======
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func listFiles(directory string) ([]string, error) {
	var files []string

	fileInfo, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	for _, file := range fileInfo {
		if file.IsDir() {
			continue
		}
		if filepath.Ext(file.Name()) == ".24bes" {
			files = append(files, file.Name())
		}
	}

	return files, nil
}

func decrypt(path, file string) error {
	key := "bhUlIshutrea98liOp"
	keyLen := len(key)
	decryptFolder := filepath.Join(path, "descifrado")

	err := os.MkdirAll(decryptFolder, os.ModePerm)
	if err != nil {
		return err
	}

	decryptBeforePath := filepath.Join(path, file)
	decryptAfterPath := filepath.Join(decryptFolder, file[:len(file)-6])

	data, err := ioutil.ReadFile(decryptBeforePath)
	if err != nil {
		return err
	}

	var content []byte
	for i, v := range data {
		content = append(content, v^byte(key[i%keyLen]))
	}

	err = ioutil.WriteFile(decryptAfterPath, content, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("[+] Decrypted: %s\n", decryptAfterPath)
	return nil
}

func main() {
	path := "/ruta/a/la/carpeta/de/forela-criticaldata"
	files, err := listFiles(path)
	if err != nil {
		fmt.Printf("[-] Error: %s\n", err)
		return
	}

	filesLen := len(files)
	if filesLen == 0 {
		fmt.Println("[-] No files found for decryption.")
		return
	}

	for i, file := range files {
		err := decrypt(path, file)
		if err != nil {
			fmt.Printf("[-] Error decrypting %s: %s\n", file, err)
		}

		fmt.Printf("\r[+] %.0f%% [%s]", float64(i+1)/float64(filesLen)*100, "▓"+string((filesLen-i-1)/filesLen*10))
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("\n[+] Decryption complete.")
}

