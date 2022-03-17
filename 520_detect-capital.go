func detectCapitalUse(word string) bool {
t := []rune(word)

	   if t[0] >= 'A' && t[0] <= 'Z' {
		   if len(t) > 1 {
			   if t[1] >= 'A' && t[1] <= 'Z' {
				   for i := 2; i < len(t); i++ {
					   if t[i] >= 'a' && t[i] <= 'z' {
						   return false
					   }
				   }
			   } else {
				   for i := 2; i < len(t); i++ {
					   if t[i] >= 'A' && t[i] <= 'Z' {
						   return false
					   }
				   }
			   }
		   }
	   } else {
		   for i := 1; i < len(t); i++ {
			   if t[i] >= 'A' && t[i] <= 'Z' {
				   return false
			   }
		   }
	   }

   return true
}
