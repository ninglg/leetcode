func subdomainVisits(cpdomains []string) []string {
    info := make(map[string]int)
    var result []string
    
    for _, cpdomain := range cpdomains {
        cd := strings.Fields(cpdomain)
        c, _ := strconv.Atoi(cd[0])
        d := cd[1]
        ds := strings.Split(d, ".")
        info[d] += c
        
        if len(ds) == 3 {
            info[ds[1]+"."+ds[2]] += c
            info[ds[2]] += c
        } else if len(ds) == 2 {
            info[ds[1]] += c
        }
    }
    
    for x, y := range info {
        result = append(result, strconv.Itoa(y)+" "+x)
    }
    
    return result
}
