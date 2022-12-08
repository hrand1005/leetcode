/*
func moveZeroes(nums []int)  {
    i, j := 0, 0

    for {

        for nums[i] != 0 {
            i++
            if i == len(nums) {
                return
            }
        }

        j = i

        for nums[j] == 0 {
            j++
            if j == len(nums) {
                return
            }
        }

        nums[i], nums[j] = nums[j], nums[i]

        i++
    }
}

func moveZeroes(nums []int)  {
    i, j := 0, 0

    for j < len(nums) {
        
        if nums[i] == 0 && nums[j] != 0 {
            nums[i], nums[j] = nums[j], nums[i]
        }

        if nums[i] != 0 {
            i++
        }

        j++
    }
}

*/

func moveZeroes(nums []int)  {
    zero := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[i], nums[zero] = nums[zero], nums[i]
            zero++
        }
    }
}