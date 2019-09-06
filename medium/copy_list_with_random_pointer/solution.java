/*
// Definition for a Node.
class Node {
    public int val;
    public Node next;
    public Node random;

    public Node() {}

    public Node(int _val,Node _next,Node _random) {
        val = _val;
        next = _next;
        random = _random;
    }
};
*/

class Solution {
    // map of node value to node copy
    Map<Integer, Node> copies = new HashMap<Integer,Node>();
    
    public Node copyRandomList(Node head) {
        // safety catch
        if (head == null) {
            return null;
        }
        
        Node n = new Node(head.val, null, null);
        copies.put(head.val, n);
        if (head.next != null) {
            n.next = copyRandomList(head.next);
        }
        if (head.random != null) {
            n.random = copies.get(head.random.val);
        }
        return n;
    }
}