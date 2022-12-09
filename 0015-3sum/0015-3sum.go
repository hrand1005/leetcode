/*
import (
    "sort"
    "strconv"
    "strings" 
)

func threeSum(nums []int) [][]int {
    negative, zero, positive := []int{}, []int{}, []int{}
    negativeMap, positiveMap := map[int]int{}, map[int]int{}
    
    // divide nums into slices of negatives, positives, and zeros
    // additionally, add negative and positive to maps for O(1) lookups
    for _, v := range nums {
        if v < 0 {
            negative = append(negative, v)
            negativeMap[v]++
        } else if v > 0 {
            positive = append(positive, v)
            positiveMap[v]++
        } else {
            zero = append(zero, v)
        }
    }
    
    solutionSet := NewSet()
    
    // if 0 exists in 'nums', that means there may exist some 3-sum
    // solution (-n, 0, n) if -n is in negativeMap and n is in positiveMap
    if len(zero) > 0 {
        for nkey, _ := range negativeMap {
            comp := nkey * -1
            for pkey, _ := range positiveMap {
                if pkey == comp {
                    solutionSet.Add([]int{nkey, 0, pkey})
                    break
                }
            }
        }
    }
    
    // if there are enough zeroes we can add the special solution set (0, 0, 0)
    if len(zero) > 2 {
        solutionSet.Add([]int{0, 0, 0})
    }
    
    // check whether each negative pair has a complement positive
    // that sums to 0, and add it to the solution set
    for i := 0; i < len(negative); i++ {
        for j := i+1; j < len(negative); j++ {
            comp := (negative[j] + negative[i]) * -1
            if _, ok := positiveMap[comp]; ok {
                solutionSet.Add([]int{negative[j], negative[i], comp})
            }
        }
    }
    
    // check whether each positive pair has a complement negative 
    // that sums to 0, and add it to the solution set
    for i := 0; i < len(positive); i++ {
        for j := i+1; j < len(positive); j++ {
            comp := (positive[j] + positive[i]) * -1
            if _, ok := negativeMap[comp]; ok {
                solutionSet.Add([]int{positive[j], positive[i], comp})
            }
        }
    }
    
    return solutionSet.ToSlice()
}

// NewSet creates a new, empty set intended to store
// triplet-pairs of integers
func NewSet() *Set {
    return &Set{
        m: make(map[string]int, 0),
    }
}

type Set struct {
    m map[string]int
}

// Add sorts our int slice and then converts it to a
// string. This is how we convert our triplet int slice
// into a hashable data type.
func (s *Set) Add(elem []int) {
    sort.Ints(elem)
    
    b := strings.Builder{}
    
    for i := 0; i < len(elem); i++ {
        b.WriteString(strconv.Itoa(elem[i]))
        if i < len(elem) - 1 {
            b.WriteString(",")
        }
    }
    
    key := b.String()
    s.m[key]++
}

// ToSlice creates our int-slice solution set by taking
// each key in our internal map, and converting it back
// into an int-slice.
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
*/

func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    solutions := [][]int{}
    
    for i := 0; i < len(nums); i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        
        l, r := i+1, len(nums) - 1
        for l < r {
            sum := nums[i] + nums[l] + nums[r]
            
            if sum < 0 {
                l++
                continue
            } else if sum > 0 {
                r--
                continue
            } else {
                solutions = append(solutions, []int{nums[i], nums[l], nums[r]})
                // increment left and decrement right until they are no longer
                // at the index of a duplicate value
                for l < r && nums[l] == nums[l+1] {
                    l++
                }
                for l < r && nums[r] == nums[r-1] {
                    r--
                }
                l++
                r--
            }
                
        }
    }
    
    return solutions
}