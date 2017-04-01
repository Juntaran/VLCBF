/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/4/1 23:31
 */

package Hash

func GetPrefix(number int) int {
	var count int
	if number < 300000 {
		count = number
	}
	count = (number/(number/300000+1) - 3154*2)
	return count
}
