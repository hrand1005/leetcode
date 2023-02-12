public class Solution {
    public int[] TwoSum(int[] nums, int target) {
        Dictionary<int, int> valIdx = new Dictionary<int, int>();
        for (int i = 0; i < nums.Length; i++) 
        {
            int compIdx;
            int comp = target - nums[i];
            if (valIdx.TryGetValue(comp, out compIdx)) {
                return new int[] { compIdx, i };
            }
            valIdx.TryAdd(nums[i], i);
        }
        return null;
    }
}