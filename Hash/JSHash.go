package Hash

func JSHash(word string) int {			// 基本返回长度皆为10位以内
	var i = 0
	var hash = 1315423911
	for {
		if i>=len(word) {
			break
		}
		hash ^= ((hash << 5) + int(word[i]) + (hash >> 2))
		i ++
	}
	//fmt.Println("hash:", hash)
	return (hash & 0x7FFFFFFF / 3000);
}

