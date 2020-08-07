/*
*by wayyoung
题目：二维数组中的查找
在一个二维数组中（每个一维数组的长度相同），
每一行都按照从左到右递增的顺序排序，每一列都
按照从上到下递增的顺序排序。请完成一个函数，输
入这样的一个二维数组和一个整数，判断数组中是否
含有该整数。

*解题思路：
整体结构为：
小1.。。。。大1
小2.。。。。大二
但是每行独立来看都是独立的，因为大1不一定小于小2
所以直接判断头尾就行，
若小x>目标数字时已经决定了之后的数组肯定不会再存
在目标数字，可以直接跳出
*/
package so

import (

	)

func so0(target int,array [][]int) bool {
	if len(array[0])==0 ||len(array)==0{//筛选[ ]的情况
	return false
	}
	for i:=0;i<len(array) ;i++  {//行遍历
		if target>=array[i][0]&&target<=array[i][len(array[0])] {//大于第一个，小于第二个
			for j:=0;j< len(array[0]);i++{
				if array[i][j]==target {
					return true
				}
			}
			if target<array[i][0] {//若小x>目标数字时已经决定了之后的数组肯定不会再存在目标数字，可以直接跳出

				return false
			}
		}

	}
	return false
}
