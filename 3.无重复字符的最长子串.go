func lengthOfLongestSubstring(s string) int {
    if len(s) == 0 {
        return 0
    }
    
    ss := []rune(s)

    l, m := 1, 1
    var p map[rune]struct{}
    for i := 0; i < len(ss); i++ {
    	l = 1
    	p = map[rune]struct{}{ss[i]:struct{}{}}
        for j := i+1; j < len(ss); j++ {
    		if _, ok := p[ss[j]]; ok {
    			break
    		}
    		
    		p[ss[j]] = struct{}{}
    		l++
    	}

    	if l > m {
    		m = l
    	}
    }

    return m
}
