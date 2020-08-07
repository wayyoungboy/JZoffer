/*
*by wayyoung
*
题目描述
写出一个程序，接受一个由字母和数字组成的字符串，和一个字符，然后输出输入字符串中含有该字符的个数。不区分大小写。

输入描述:
第一行输入一个有字母和数字以及空格组成的字符串，第二行输入一个字符。

输出描述:
输出输入字符串中含有该字符的个数。

示例1
	输入:
		ABCDEF
		A
	输出:
		1
 */
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	for {
		var (
			s string
			b string
			err error
		)
		/*由于字符串中存在空格，因此无法直接使用此注释的代码
		_,err=fmt.Scan(&s,&b)
		if err!=nil{
			return
		}
		 */
		//每次读一行
		inputReader := bufio.NewReader(os.Stdin)
		s, err = inputReader.ReadString('\n')
		if err!=nil{
			return
		}
		b,err= inputReader.ReadString('\n')
		if err!=nil{
			return
		}

		re:=so0(s,b[0])
		fmt.Println(re)
	}
}
func so0(s string,b byte) int {
	if len(s)==0 {
		return 0
	}
	re:=0
	for _,x:=range s{
		if byte(x)==b {//不加强制类型转换会报错，不知为啥，索性就加了
			re++
		}
	}
return re
}