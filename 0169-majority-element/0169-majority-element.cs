public class Solution {
    public int MajorityElement(int[] nums) {
        int majorityElement = nums[0];
        int majorityElementCount = 0;
        for (int i = 0; i < nums.Length; i++)
        {
            if (nums[i] == majorityElement)
            {
                majorityElementCount++;
            }
            else
            {
                majorityElementCount--;
            }
            
            if (majorityElementCount == 0)
            {
                majorityElement = nums[i];
                majorityElementCount = 1;
            }
        }
        return majorityElement;
    }
}