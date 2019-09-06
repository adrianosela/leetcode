/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */

class Solution {
    public boolean isSubtree(TreeNode s, TreeNode t) {   
        if(s == null) {
            return false;
        }
    
        if (s.val != t.val) {
            return isSubtree(s.left, t) || isSubtree(s.right, t);
        }
        
        return check(s, t)
                || isSubtree(s.left, t) 
                || isSubtree(s.right, t);   
    }
    
    private boolean check(TreeNode s, TreeNode t) {
        if(s == null && t == null) {
            return true;
        }
        if(s == null && t != null) {
            return false;
        }
        if(s != null && t == null) {
            return false;
        }
        if(s.val != t.val) {
            return false;
        }
        return check(s.left, t.left) && check(s.right, t.right);
    }
}