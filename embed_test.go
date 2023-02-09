package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

/*
! Embed Package

* Package embed adalah fitur baru untuk mempermudah membaca isi file pada saat
* compile time secara otomatis dimasukkan isi file nya dalam variable


! Cara Embed File
* Untuk melakukan embed file ke variable, kita bisa mengimport package embed terlebih dahulu
* Selajutnya kita bisa tambahkan komentar //go:embed diikuti dengan nama file nya, diatas variable yang kita tuju
* Variable yang dituju tersebut nanti secara otomatis akan berisi konten file yang kita inginkan secara otomatis ketika kode golang di compile
* Variable yang dituju tidak bisa disimpan di dalam function
*/

/*
? Embed File ke String
* Embed file bisa kita lakukan ke variable dengan tipe data string
* Secara otomatis isi file akan dibaca sebagai text dan masukkan ke variable tersebut

*/

//go:embed version.txt
var version1 string

//go:embed version.txt
var version2 string

func TestString(t *testing.T) {
	fmt.Println(version1)
	fmt.Println(version2)
}

/*
? Embed File ke []byte
* Selain ke tipe data String, embed file juga bisa dilakukan ke variable tipe data []byte
* Ini cocok sekali jika kita ingin melakukan embed file dalam bentuk binary, seperti gambar dan lain-lain
*/

//go:embed logo-udemy.svg
var logo1 []byte

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("logo_new.jpg ", logo1, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

/*
? Embed Multiple Files
* Kadang ada kebutuhan kita ingin melakukan embed beberapa file sekaligus
* Hal ini juga bisa dilakukan menggunakan embed package
* Kita bisa menambahkan komentar //go:embed lebih dari satu baris
* Selain itu variable nya bisa kita gunakan tipe data embed.FS
*/

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	b, _ := files.ReadFile("files/b.txt")
	c, _ := files.ReadFile("files/c.txt")

	fmt.Println(string(a))
	fmt.Println(string(b))
	fmt.Println(string(c))
}

/*
? Path Mathcer
* mengguakan patch matcher untuk membaca multiple file yang kita inginkan
* Ini sangat cocok ketika misal kita punya pola jenis file yang kita inginkan untuk kita baca
* Caranya, kita perlu menggunakan path matcher seperti pada package function path.Match

*/

//go:embed files/*.txt
var path1 embed.FS

func TestPathMathcer(t *testing.T) {
	dirEntries, _ := path1.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path1.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}
