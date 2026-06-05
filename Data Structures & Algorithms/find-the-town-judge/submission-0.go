func findJudge(n int, trust [][]int) int {
    trustGraph := make(map[int]map[int]struct{}, n)
    for i:=1; i<=n; i++ {
        trustGraph[i] = map[int]struct{}{}
    }

    for _, t := range trust {
        if _, found := trustGraph[t[0]]; found {
            trustGraph[t[0]][t[1]] = struct{}{}
        }
    }

    judge := 0
    for i:=1; i<=n; i++ {
        if len(trustGraph[i]) != 0 {
            continue
        }
        if judge != 0 {
            return -1
        }
        judge = i
    }

    for i:=1; i<=n; i++ {
        if i == judge {
            continue
        }
        if _, found := trustGraph[i][judge]; !found{
            return -1
        }
    }

    return judge
}
