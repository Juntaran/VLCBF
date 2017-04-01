package Hash

func SDBMHash(word string) int {		// 基本返回长度皆为10位以内
	var i = 0
	var hash int
	for {
		if i>=len(word) {
			break
		}
		//fmt.Printf("word[%d]:%c\n", i, word[i])
		//fmt.Printf("hash:%d\n", hash)
		hash = int(word[i]) + (hash << 6) + (hash << 16) - hash
		i++
		//fmt.Printf("hash:%d\n", hash)
	}
	//fmt.Println("hash:", hash)

	return GetPrefix(hash & 0x7FFFFFFF);
}
