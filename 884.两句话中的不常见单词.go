func uncommonFromSentences(A string, B string) []string {
    AA := strings.Fields(strings.TrimSpace(A))
    BB := strings.Fields(strings.TrimSpace(B))
    
    CC := make(map[string]int)

    for _, a := range AA {
    	if _, ok := CC[a]; ok {
    		CC[a]++
    	} else {
    		CC[a] = 1
    	}
    }

    for _, b := range BB {
    	if _, ok := CC[b]; ok {
    		CC[b]++
    	} else {
    		CC[b] = 1
    	}    	
    }
    
    var AB []string
    for ab, c := range CC {
    	if c == 1 {
    		AB = append(AB, ab)
    	}
    }

    return AB
}
