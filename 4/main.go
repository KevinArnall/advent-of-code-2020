package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	f, err := os.Open("./4/input.txt")
	if err != nil {
		fmt.Print("Error opening file: ", err.Error())
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var validPassportsP1 int
	var validPassportsP2 int

	passport := make(map[string]string)
	for scanner.Scan() {

		line := scanner.Text()
		if line == "" {
			if checkValidityP1(passport) {
				validPassportsP1++
				if checkValidityP2(passport) {
					validPassportsP2++
				}
			}
			passport = make(map[string]string)
			continue
		}

		entries := strings.Split(line, " ")

		for _, entry := range entries {
			part := strings.Split(entry, ":")

			key := part[0]
			value := part[1]

			passport[key] = value
		}
	}

	// Need to get the last one because the file doesn't end in an empty line
	if checkValidityP1(passport) {
		validPassportsP1++
		if checkValidityP2(passport) {
			validPassportsP2++
		}
	}

	fmt.Printf("Found %d valid passports for part 1\n", validPassportsP1)
	fmt.Printf("Found %d valid passports for part 2", validPassportsP2)

}

func checkValidityP2(p map[string]string) bool {

	byr, _ := strconv.Atoi(p["byr"])
	if !(byr >= 1920 && byr <= 2002) {
		return false
	}

	iyr, _ := strconv.Atoi(p["iyr"])
	if !(iyr >= 2010 && iyr <= 2020) {
		return false
	}

	eyr, _ := strconv.Atoi(p["eyr"])
	if !(eyr >= 2020 && eyr <= 2030) {
		return false
	}

	if strings.HasSuffix(p["hgt"], "cm") {
		p["hgt"] = strings.Replace(p["hgt"], "cm", "", 1)

		hgt, _ := strconv.Atoi(p["hgt"])
		if !(hgt >= 150 && hgt <= 193) {
			return false
		}
	} else if strings.HasSuffix(p["hgt"], "in") {
		p["hgt"] = strings.Replace(p["hgt"], "in", "", 1)

		hgt, _ := strconv.Atoi(p["hgt"])
		if !(hgt >= 59 && hgt <= 76) {
			return false
		}
	} else {
		return false
	}

	if len(p["hcl"]) != 7 || p["hcl"][0] != '#' {
		return false
	}

	validEyeColors := map[string]bool{"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true}

	if !validEyeColors[p["ecl"]] {
		return false
	}

	if len(p["pid"]) != 9 {
		return false
	}

	return true
}

func checkValidityP1(p map[string]string) bool {
	if len(p) == 8 || (len(p) == 7 && p["cid"] == "") {
		return true
	} else {
		return false
	}
}
