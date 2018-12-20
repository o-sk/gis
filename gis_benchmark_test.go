package gis

import (
	"os"
	"testing"
)

func BenchmarkDownload(b *testing.B) {
	dir := "tmp/download"
	_ = os.MkdirAll(dir, 0777)
	filename := "cat"
	images, _ := Search("猫")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Download(dir, filename, images)
	}
}

func BenchmarkSerialDownload(b *testing.B) {
	dir := "tmp/serial_download"
	_ = os.MkdirAll(dir, 0777)
	filename := "cat"
	images, _ := Search("猫")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SerialDownload(dir, filename, images)
	}
}
