import (
    "sort"
    "strconv"
    "strings" 
)

func threeSum(nums []int) [][]int {
    negative, zero, positive := []int{}, []int{}, []int{}
    negativeSet, positiveSet := map[int]int{}, map[int]int{}
    
    // divide nums into slices of negatives, positives, and zeros
    // additionally, add negative and positive to sets for O(1) lookups
    for _, v := range nums {
        if v < 0 {
            negative = append(negative, v)
            negativeSet[v]++
        } else if v > 0 {
            positive = append(positive, v)
            positiveSet[v]++
        } else {
            zero = append(zero, v)
        }
    }
    
    solutionSet := NewSet()
    
    if len(zero) > 0 {
        for nkey, _ := range negativeSet {
            comp := nkey * -1
            for pkey, _ := range positiveSet {
                if pkey == comp {
                    solutionSet.Add([]int{nkey, 0, pkey})
                    break
                }
            }
        }
    }
    
    if len(zero) > 2 {
        solutionSet.Add([]int{0, 0, 0})
    }
    
    for i := 0; i < len(negative); i++ {
        for j := i+1; j < len(negative); j++ {
            comp := (negative[j] + negative[i]) * -1
            if _, ok := positiveSet[comp]; ok {
                solutionSet.Add([]int{negative[j], negative[i], comp})
            }
        }
    }
    
    for i := 0; i < len(positive); i++ {
        for j := i+1; j < len(positive); j++ {
            comp := (positive[j] + positive[i]) * -1
            if _, ok := negativeSet[comp]; ok {
                solutionSet.Add([]int{positive[j], positive[i], comp})
            }
        }
    }
    
    return solutionSet.ToSlice()
}

func NewSet() *Set {
    return &Set{
        m: make(map[string]int, 0),
    }
}

type Set struct {
    m map[string]int
}

func (s *Set) Add(elem []int) {
    sort.Ints(elem)
    
    key := ""
    for _, v := range elem {
        key += strconv.Itoa(v) + ","
    }
    key = key[:len(key)-1] // remove last comma
    
    s.m[key]++
}

func (s *Set) ToSlice() [][]int {
    result := make([][]int, 0, len(s.m))
    
    for k, _ := range s.m {
        strSlice := strings.Split(k, ",")
        intSlice := make([]int, 0, 3)
        for _, v := range strSlice {
            num, _ := strconv.Atoi(v)
            intSlice = append(intSlice, num)
        }
        result = append(result, intSlice)
    }
    
    return result
}