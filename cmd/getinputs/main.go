package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	var day int
	var sessionFile string
	var baseDir string
	var all bool
	var stub bool

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	date := time.Now().Local()

	flag.IntVar(&day, "day", date.Day(), "The day to download input for. By default, the current day is used.")
	flag.StringVar(&sessionFile, "session", "session", "Path to a file containing the session cookie.")
	flag.StringVar(&baseDir, "dir", wd, "Base directory to save input files to. By default, the current directory is used.")
	flag.BoolVar(&all, "all", false, "Download input for all days up to the chosen day.")
	flag.BoolVar(&stub, "stub", false, "Create empty files instead of downloading input.")
	flag.Parse()

	absBaseDir, err := filepath.Abs(baseDir)
	if err != nil {
		panic(fmt.Sprintf("Failed to get absolute path for base directory: %v", err))
	}

	var session string
	if !stub {
		sessionBuf, err := os.ReadFile(filepath.Join(absBaseDir, sessionFile))
		if err != nil {
			panic(fmt.Sprintf("Failed to read session file: %v", err))
		}
		session = strings.TrimSpace(string(sessionBuf))
	}

	if all {
		fmt.Printf("Downloading input for all days up to day %d to %s\n", day, absBaseDir)
		for i := 1; i <= day; i++ {
			if err := downloadOrStubInput(stub, i, session, absBaseDir); err != nil {
				panic(fmt.Sprintf("Failed to get input for day %d: %v", i, err))
			}
		}
	} else {
		fmt.Printf("Downloading input for day %d to %s\n", day, absBaseDir)
		if err := downloadOrStubInput(stub, day, session, absBaseDir); err != nil {
			panic(fmt.Sprintf("Failed to get input for day %d: %v", day, err))
		}
	}
}

func downloadOrStubInput(stub bool, day int, session, baseDir string) error {
	err := os.MkdirAll(fmt.Sprintf("%s/day%d", baseDir, day), 0755)
	if err != nil {
		return err
	}
	if stub {
		return stubInput(day, baseDir)
	}
	return downloadInput(day, session, baseDir)
}

func stubInput(day int, baseDir string) error {
	file, err := os.Create(fmt.Sprintf("%s/day%d/input.txt", baseDir, day))
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func downloadInput(day int, session, baseDir string) error {
	url := fmt.Sprintf("https://adventofcode.com/2023/day/%d/input", day)
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "github.com/johningve/adventofcode2023")

	cookie := new(http.Cookie)
	cookie.Name, cookie.Value = "session", session
	req.AddCookie(cookie)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to get input for day %d: %s", day, resp.Status)
	}

	file, err := os.Create(fmt.Sprintf("%s/day%d/input.txt", baseDir, day))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
