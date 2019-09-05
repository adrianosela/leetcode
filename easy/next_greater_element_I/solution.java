class Solution {
    Map<Integer, Integer> indices = new HashMap<Integer, Integer>();
    
    public int[] nextGreaterElement(int[] nums1, int[] nums2) {
        int ret[] = new int[nums1.length];
        
        for (int i = 0; i < nums2.length; i++) {
            indices.put(nums2[i], i);
        }

        int cur_num2;
        for (int i = 0; i < nums1.length; i++) {
            ret[i] = -1;
            for (int j = indices.get(nums1[i]); j < nums2.length; j++) {
                cur_num2 = nums2[j];
                if (cur_num2 > nums1[i]) {
                    ret[i] = cur_num2;
                    break;
                }
            }
        }
        
        return ret;
    }
}
