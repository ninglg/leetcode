/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    start := 1
    end := n

    for {
        m := (start+end)/2
        t := guess(m)
        switch t {
            case -1:
                end = m
            case 1:
                if m == start {
                    start = m+1
                } else {
                    start = m
                }
            case 0:
                return m
        }
    }
}
