import (
    "fmt"
)

func fizzBuzz(n int) []string {
    answer := make([]string, 0, n+1)

    for i := 1; i < n+1; i++ {
        v := ""

        if i % 3 == 0 {
            v += "Fizz"
        }

        if i % 5 == 0 {
            v += "Buzz"
        }

        if len(v) == 0 {
            v += fmt.Sprintf("%v", i)
        }

        answer = append(answer, v)
    }

    return answer
}