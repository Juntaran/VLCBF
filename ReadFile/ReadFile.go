/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/2 00:33
  */

package ReadFile

import (
	"os"
	"bufio"
	"strings"
	"io"
	//"fmt"
	"fmt"
)

func ReadFile(fileName string, stringSlice []string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	buf := bufio.NewReader(file)
	for i:=0; ; i++ {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		//fmt.Println(line)
		stringSlice = append(stringSlice, line)
		//fmt.Println(stringSlice)
		if err != nil {
			fmt.Println(err)
			if err == io.EOF {
				return stringSlice, nil
			}
			return nil, err
		}
	}
	return stringSlice, err
}
