class Solution {
    public int[] nextGreaterElements(int[] nums) {
        
        int len = nums.length;
        int ret[] = new int[len];
        
        int cur_num;
        for (int i = 0; i < len; i++) {
            cur_num = nums[i];
            ret[i] = -1;
            
            for (int j = (i+1)%len; j != i; j=(j+1)%len) {
                int replacement = nums[j];
                if (replacement > cur_num) {
                    ret[i] = replacement;
                    break;
                }
            }
            
        }

        return ret;
    }
}
