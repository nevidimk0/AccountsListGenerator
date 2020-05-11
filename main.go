//
//	This tool generate account names according to original people names and  naming conventions
//
//	Account name generation rules:
//	1. Complete first name plus last name
//	2. Initial of first name and complete last name
//	3. First three characters of the first name and first three of last name
//	4. First three characters of the first name and complete last name
//
//	Usage: go run ./main.go users.txt
//
//	users.txt:
//	Bill Murray
//	Mike Myers
//
//
//

package main
 
import (
	"bufio"
	"fmt"
	"log"
	"strings"
	"regexp"
	"os"
)

func unique(srtingSlice []string) []string {
    keys := make(map[string]bool)
    list := []string{} 
    for _, entry := range srtingSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }    
    return list
}

func get_all_combinations(parts []string) ([]string) {
	var combinations []string
	
	first_part := parts[0]
	second_part := parts[1]
	combinations = append(combinations,first_part+second_part) // Type 1
	combinations = append(combinations,first_part+"."+second_part) // Type 1 + dot
	combinations = append(combinations,first_part+"-"+second_part) // Type 1 + hyphen
	
	combinations = append(combinations,second_part+first_part) // Type 1 alt
	combinations = append(combinations,second_part+"."+first_part) // Type 1 alt + dot
	combinations = append(combinations,second_part+"-"+first_part) // Type 1 alt + hyphen
	
	
	combinations = append(combinations,string(first_part[0])+second_part) // Type 2
	combinations = append(combinations,string(first_part[0])+"."+second_part) // Type 2 + dot
	combinations = append(combinations,string(first_part[0])+"-"+second_part) // Type 2 + hyphen
	
	combinations = append(combinations,string(second_part[0])+first_part) // Type 2 alt
	combinations = append(combinations,string(second_part[0])+"."+first_part) // Type 2 alt + dot
	combinations = append(combinations,string(second_part[0])+"-"+first_part) // Type 2 alt + hyphen
	
	if len(first_part) >= 3 && len(second_part) >= 3 {
		combinations = append(combinations,string(first_part[0:3])+second_part[0:3]) // Type 3
		combinations = append(combinations,string(first_part[0:3])+"."+second_part[0:3]) // Type 3 + dot
		combinations = append(combinations,string(first_part[0:3])+"-"+second_part[0:3]) // Type 3 + hyphen
		
		combinations = append(combinations,string(second_part[0:3])+first_part[0:3]) // Type 3 alt
		combinations = append(combinations,string(second_part[0:3])+"."+first_part[0:3]) // Type 3 alt + dot
		combinations = append(combinations,string(second_part[0:3])+"-"+first_part[0:3]) // Type 3 alt + hyphen
	}
	
	
	if len(first_part) >= 3 {
		combinations = append(combinations,string(first_part[0:3])+second_part) // Type 4
		combinations = append(combinations,string(first_part[0:3])+"."+second_part) // Type 4 + dot
		combinations = append(combinations,string(first_part[0:3])+"-"+second_part) // Type 4 + hyphen
	}
	if len(second_part) >= 3 {
		combinations = append(combinations,string(second_part[0:3])+first_part) // Type 4 alt
		combinations = append(combinations,string(second_part[0:3])+"."+first_part) // Type 4 alt + dot
		combinations = append(combinations,string(second_part[0:3])+"-"+first_part) // Type 4 alt + hyphen
	}
	
	return combinations
}

func main() {
	file_to_process := os.Args[1]
	file, err := os.Open(file_to_process)
 
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
 
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var total_combinations []string
	for scanner.Scan() {
		line_to_process := scanner.Text()
		if strings.Contains(line_to_process, " ") {
			space := regexp.MustCompile(`\s+`)
			line_to_process = space.ReplaceAllString(line_to_process, " ")
			username_slice := strings.Split(line_to_process," ")
			total_combinations = append(total_combinations,get_all_combinations(username_slice)...)
		} else {
			total_combinations = append(total_combinations,line_to_process)
		}
	}
	unique_combinations := unique(total_combinations)
	for _,final_item := range unique_combinations {
		fmt.Println(final_item)
	}
}