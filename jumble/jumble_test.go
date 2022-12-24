package jumble

import (
    "testing"
)

func TestSolve(t *testing.T) {
    wordlist := createWordlist()
    examples := map[string][]string{
        "wakte": []string{"tweak"},
        "yoppp": []string{"poppy"},
        "lautes": []string{"elatus", "salute", "setula"},
        "sneewt": []string{"newest"},
        "SNEEWT": []string{"newest"},
        "SnEeWt \n\n": []string{"newest"},
        "": []string{""},
    }
    for input, output := range examples {
        testOutput := Solve(input, wordlist)
        for i, _ := range testOutput {
            if testOutput[i] != output[i] {
                t.Errorf("Solve(%s) = %s; want %s", input, testOutput, output)
            }
        }
    }
}
