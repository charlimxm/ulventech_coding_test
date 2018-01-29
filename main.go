package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func wordCount(paras []string) map[string]int {
	counts := make(map[string]int)
	for _, word := range paras {
		counts[word]++
	}
	return counts
}

type Format struct {
	Word  string
	Count int
}

func (f Format) String() string {
	return fmt.Sprintf("%s: %d", f.Word, f.Count)
}

type ByCount []Format

func (a ByCount) Len() int           { return len(a) }
func (a ByCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCount) Less(i, j int) bool { return a[i].Count < a[j].Count }

func main() {
	scn := bufio.NewScanner(os.Stdin)
	var lines []string

	fmt.Println("Enter text (<Ctrl + ]> + <Enter> on a new line to terminate input):")

	for scn.Scan() {
		line := strings.ToLower(scn.Text())
		reg := regexp.MustCompile("[[:alnum:]']+")
		processedString := reg.FindAllString(line, -1)

		if len(line) == 1 {
			if line[0] == '\x1D' {
				break
			}
		}
		lines = append(lines, processedString...)
	}

	if len(lines) > 0 {
		fmt.Println("Top ten most used words:")
		results := wordCount(lines)

		bc := make(ByCount, len(results))
		i := 0
		for k, v := range results {
			bc[i] = Format{k, v}
			i++
		}

		sort.Sort(sort.Reverse(bc))

		topTenWords := bc[0:10]
		for i := 0; i < len(topTenWords); i++ {
			fmt.Println(topTenWords[i])
		}
	}
}
