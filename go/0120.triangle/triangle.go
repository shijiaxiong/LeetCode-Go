package main

import "fmt"

//class Solution {
//    Integer[][] memo;
//    public int minimumTotal(List<List<Integer>> triangle) {
//        memo = new Integer[triangle.size()][triangle.size()];
//        return  dfs(triangle, 0, 0);
//    }
//
//    private int dfs(List<List<Integer>> triangle, int i, int j) {
//        if (i == triangle.size()) {
//            return 0;
//        }
//        if (memo[i][j] != null) {
//            return memo[i][j];
//        }
//        return memo[i][j] = Math.min(dfs(triangle, i + 1, j), dfs(triangle, i + 1, j + 1)) + triangle.get(i).get(j);
//    }
//}
//

func minimumTotal(triangle [][]int) int {

	return dfs(triangle, 0, 0)
}

func dfs(triangle [][]int, i, j int) int {
	if i == len(triangle) {
		return 0
	}

	return min(dfs(triangle, i+1, j), dfs(triangle, i+1, j+1)+triangle[i][j])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
}
