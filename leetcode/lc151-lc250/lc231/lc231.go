//lc231-2的幂
//[位运算]

//给你一个整数n，请你判断该整数是否是2的幂次方。
//如果是，返回true；否则，返回false。
//如果存在一个整数x使得n==2^x，则认为n是2的幂次方。
//-2^31<=n<=2^31-1

//n和n-1按位与运算，如果结果为0，则n是2的幂次方。
//例如31，31=011111，30=011110,31&30=000001
//例如32，32=100000，31=011111,32&31=000000
//例如33，33=100001，32=100000,33&32=000001

package main

import "fmt"

func main() {
  fmt.Println(isPowerOfTwo((1024)))

}

func isPowerOfTwo(n int) bool {
  return n > 0 && n&(n-1) == 0
}
