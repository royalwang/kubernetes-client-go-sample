package main

import (
	"fmt"
	"os"
	"testing"
)

func setup() {
	fmt.Println("Before all tests")
}

func teardown() {
	fmt.Println("After all tests")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

// Test starting ...

func TestHomeDir(t *testing.T) {
	dir := homeDir()
	println(dir)
}

func TestKubeConfig(t *testing.T) {
	config := kubeConfig()
	fmt.Printf("%+v\n", config)
}
