package main

import (
	"log"
	"os"
	"os/exec"
	"runtime"
)

var CMakeList = []string{
	"cmake_minimum_required(VERSION 3.22.1)",
	"",
	"project(project_name)",
	"set(CMAKE_CXX_STANDARD 20)",
	"set(CMAKE_CXX_STANDARD_REQUIRED TRUE)",
	"set(CMAKE_EXPORT_COMPILE_COMMANDS ON)",
	"set(SRC_DIR \"${CMAKE_CURRENT_SOURCE_DIR}/src\")",
	"set(INCLUDE_DIR \"${CMAKE_CURRENT_SOURCE_DIR}/include\")",
	"",
	"aux_source_directory(${SRC_DIR} SRC_FILES)",
	"include_directories(${INCLUDE_DIR})",
	"",
	"add_executable(${PROJECT_NAME} main.cpp)",
}

var MainCpp = []string{
	"#include <iostream>",
	"",
	"int main() {",
	"\tstd::cout << \"Sample\" << std::endl;",
	"\treturn 0;",
	"}",
}

func main() {
	err := os.Mkdir("build", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Mkdir("include", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Mkdir("src", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	createCMakeList()
	createMainCPP()

	if runtime.GOOS == "windows" {
		cmakeExec := exec.Command("cmake", "-S", ".", "-B", "build/", "-G", "MinGW Makefiles")
		err := cmakeExec.Run()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		cmakeExec := exec.Command("cmake", "-S", ".", "-B", "build/")
		err := cmakeExec.Run()
		if err != nil {
			log.Fatal(err)
		}
	}

}

func createCMakeList() {
	cmakeFile, err := os.Create("CMakeLists.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer cmakeFile.Close()

	for _, line := range CMakeList {
		_, err := cmakeFile.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

}

func createMainCPP() {
	mainFile, err := os.Create("main.cpp")
	if err != nil {
		log.Fatal(err)
	}
	defer mainFile.Close()

	for _, line := range MainCpp {
		_, err := mainFile.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}
