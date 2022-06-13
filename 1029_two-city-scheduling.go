func twoCitySchedCost(costs [][]int) int {
    var res int

    sort.Slice(costs, func(i, j int)bool{
        return (costs[i][0] - costs[i][1]) < (costs[j][0] - costs[j][1])
    })

    for i := 0; i < len(costs) / 2; i++{
        res += costs[i][0]
    }

    for i := len(costs) / 2; i < len(costs); i++{
        res += costs[i][1]
    }

    return res
}

