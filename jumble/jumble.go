package jumble

import (
    "bufio"
    "fmt"
    "os"
    "path"
    "sort"
    "strings"
)

func genPermutations(runes []rune, left, right int, permutations []string) []string {
    if left == right {
        return append(permutations, string(runes))
    } else {
        for i := left; i <= right; i++ {
            runes[left], runes[i] = runes[i], runes[left]
            permutations = genPermutations(runes, left+1, right, permutations)
            runes[left], runes[i] = runes[i], runes[left]
        }
    }
    return permutations
}

func createWordlist() map[string]bool {
    wordlistPath := path.Join("/tmp", "words.txt")
    file, err := os.Open(wordlistPath)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    wordlist := make(map[string]bool)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        wordlist[strings.ToLower(scanner.Text())] = true
    }
    return wordlist
}

func Solve(puzzle string, wordlist map[string]bool) (solutions []string) {
    validSolutions := make(map[string]bool)
    var permutations []string
    permutations = genPermutations([]rune(puzzle), 0, len(puzzle)-1, permutations)
    for _, potentialSolution := range permutations {
        if wordlist[potentialSolution] {
            validSolutions[potentialSolution] = true
        }
    }
    for solution, _ := range validSolutions {
        solutions = append(solutions, solution)
    }
    sort.Strings(solutions)
    return
}

func Run() {
    wordlist := createWordlist()
    reader := bufio.NewReader(os.Stdin)
    confirm := false
    for {
        fmt.Print("Jumble: ")
        text, _ := reader.ReadString('\n')
        text = strings.ToLower(strings.TrimSpace(text))
        if text == "" {
            if confirm {
                break
            }
            confirm = true
            fmt.Println("Press Enter again to quit.")
        }
        solutions := Solve(text, wordlist)
        for _, solution := range solutions {
            fmt.Println(solution)
        }
    }
}
