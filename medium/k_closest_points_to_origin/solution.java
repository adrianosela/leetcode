class Point implements Comparable<Point> {
    public int x;
    public int y;
    public double distance;
    
    public Point(final int x, final int y) {
        this.x = x;
        this.y = y;
        this.distance = Math.sqrt((x*x) + (y*y));
    }
    
    public int compareTo(final Point p) {
        return distance == p.distance ? 0 : ((distance < p.distance) ? -1 : 1);
    }
}

class Solution {
    public int[][] kClosest(int[][] points, int K) {
        
        PriorityQueue<Point> queue = new PriorityQueue<>();
        for (int i = 0; i < points.length; i++) {
            Point p = new Point(points[i][0], points[i][1]);
            queue.add(p);
        }
        
        int[][] ans = new int[K][2];
        for (int i = 0; i < K; i++) {
            Point p = queue.poll();
            ans[i][0] = p.x;
            ans[i][1] = p.y;
        }
        return ans;
    }
}