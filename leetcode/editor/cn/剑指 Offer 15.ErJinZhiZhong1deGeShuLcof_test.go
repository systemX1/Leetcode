package Leetcode

import (
    "log"
    "math"
    "testing"
    "unsafe"
)

func TestErJinZhiZhong1deGeShuLcof(t *testing.T) {
    fn := hammingWeight
    log.Println(fn(11), fn(128), fn(4294967293))
    log.Println(fn(math.MaxUint32), fn(127), fn(4294967292))
    fn2 := hammingWeightS
    log.Println(fn2(-11), fn2(-128), fn2(0xFFFFFFF))
    log.Println(fn2(0x8000000), fn2(math.MinInt32), fn2(1))
    log.Println(fn2(0x7FFFFFFF), fn2(math.MinInt32), fn2(-1))

}

//leetcode submit region begin(Prohibit modification and deletion)
func hammingWeight(num uint32) int {
    ans := 0
    for num != 0 {
        num = (num - 1) & num
        ans++
    }
    return ans
}
//leetcode submit region end(Prohibit modification and deletion)

func hammingWeightS(num int32) int {
    ans := 0
    for num != 0 {
        num = (num - 1) & num
        ans++
    }
    return ans
}

// 书相关题目
// 相关题目1：用一条语句判断一个整数是不是2的整数次方
func f3(num uint32) bool {
    return num & (num - 1) == 0
}
// 相关题目2：输入两个整数m和n，计算需要改变m的二进制表示中的多少位才能得到n。
func f4(m, n uint32) int {
    num := m ^ n
    ans := 0
    for num != 0 {
        num = (num - 1) & num
        ans++
    }
    return ans
}
// 相关题目3：二进制中0的个数
func f5(num uint32) int {
    ans := 0
    for num != 0 {
        num = (num - 1) & num
        ans++
    }
    return int(unsafe.Sizeof(num)) - ans
}

