package tool

import (
	"regexp"
	"strings"
)

const IdCardregular = `^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`

/*
身份证号码各位数字的含义如下：
0）前1、2位数字表示：所在省份的代码；
1）第3、4位数字表示：所在城市的代码；
2）第5、6位数字表示：所在区县的代码；
3）第7~14位数字表示：出生年、月、日；
4）第15、16位数字表示：所在地的派出所的代码；
5）第17位数字表示性别：奇数表示男性，偶数表示女性；
6）第18位数字是校检码
*/
// 验证18位身份证
func CheckIdCardNumber18(number string) (bool, error) {
	// 00.首先用正则表达式验证身份证号码格式是否正确
	if IsIdCardNumberRegular(number) {
		// 01.遍历1～17位，分别乘以7-9-10-5-8-4-2-1-6-3-7-9-10-5-8-4-2，并相加
		factor := [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
		var sum = 0
		for index, num := range number[:17] {
			sum = sum + int(num-48)*factor[index]
		}
		// 02.将上一步相加的和除以11，得出余数
		remainder := sum % 11
		// 03.由于数字的特殊性，这些余数只可能是0-10这11个数字，身份证最后一位的对应数字为1-0-X-9-8-7-6-5-4-3-2。例上面的余数结果为3那么对应身份证号码的最后一位就是9，如果是10，身份证最后一位便是2。
		checkStr := "10X98765432"
		return checkStr[remainder] == strings.ToUpper(number)[17], nil
	}
	return false, nil
}

// 正则表达式，验证身份证号是否符合格式，18位身份证号码
func IsIdCardNumberRegular(number string) bool {
	if len(number) != 18 {
		return false
	}
	reg := regexp.MustCompile(IdCardregular)
	return reg.MatchString(number)
}
