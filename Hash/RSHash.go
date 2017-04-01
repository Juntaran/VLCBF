package Hash


func RSHash(word string) int {			// 基本返回长度皆为10位以内
	var i = 0
	var hash int
	var a = 63689
	var b = 378661
	for {
		if i>=len(word) {
			break
		}
		hash = hash * a + int(word[i])
		i ++
		a *= b
		//fmt.Printf("hash:%d\n", hash)
	}
	//fmt.Println("hash:", hash)
	return GetPrefix(hash & 0x7FFFFFFF);
}
