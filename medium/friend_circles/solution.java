class FriendCircles {
    public int findCircleNum(int[][] M) {
        if (M == null || M.length == 0) {
            return 0;
        }
        
        int nodes = M.length;
        
        // initialize disjoint sets where every node
        // is its own parent
        int parents[] = new int[nodes];
        for (int node = 0; node < nodes; node++) {
            parents[node] = node;
        }

        for (int node_a = 0; node_a < nodes; node_a++) {
            for (int node_b = node_a + 1; node_b < nodes; node_b++) {
                if (M[node_a][node_b] == 1) {
                    union(parents, node_a, node_b);
                }
            }
        }
        
        Set<Integer> s = new HashSet();
        for (int node = 0; node < nodes; node++) {
            s.add(find(parents, node));
        }
        
        return s.size();
    }
    
    private void union(int[] p, int node_a, int node_b) {
        p[find(p, node_b)] = find(p, node_a);
    }
    
    private int find(int[] p, int node) {
        while (p[node] != node) {
            node = p[node];
        }
        return p[node];
    }
}