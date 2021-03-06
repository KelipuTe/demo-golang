//冒泡排序，BubbleSort
package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  var arr1numsLen = 20
  var arr1nums []int = make([]int, arr1numsLen)

  rand.Seed(time.Now().UnixNano())
  for i := 0; i < arr1numsLen; i++ {
    arr1nums[i] = rand.Intn(999)
  }

  mao4pao4(arr1nums, true)

  rand.Seed(time.Now().UnixNano())
  for i := 0; i < arr1numsLen; i++ {
    arr1nums[i] = rand.Intn(999)
  }

  mao4pao4(arr1nums, false)
}

func mao4pao4(arr1nums []int, isASC bool) {
  fmt.Println(arr1nums)

  var arr1numsLen int = len(arr1nums)
  var indexLast int = arr1numsLen - 1 //最后交换的元素的下标
  var isExchange bool = false         //一轮排序是否有元素被交换

  //无论是升序排序还是降序排序，每轮比较结束之后，最后一个元素一定是最大或者最小元素，后续的循环就不需要进行比较了
  //在此基础上，如果一轮排序中，位于后面的一部分元素没有交换，则说明这部分已经有序，后续的循环就不需要进行比较了

  for indexi := 0; indexi < arr1numsLen; indexi++ { //第n次循环找到第n小（大）的元素
    indexLastTemp := indexLast
    if isASC { //升序排序
      for indexj := 0; indexj < indexLastTemp; indexj++ {
        if arr1nums[indexj] > arr1nums[indexj+1] { //如果前面比后面大就交换
          arr1nums[indexj], arr1nums[indexj+1] = arr1nums[indexj+1], arr1nums[indexj]
          indexLast = indexj
          isExchange = true
        }
      }
    } else { //降序排序
      for indexj := 0; indexj < indexLastTemp; indexj++ {
        if arr1nums[indexj] < arr1nums[indexj+1] { //如果前面比后面小就交换
          arr1nums[indexj], arr1nums[indexj+1] = arr1nums[indexj+1], arr1nums[indexj]
          indexLast = indexj
          isExchange = true
        }
      }
    }
    if !isExchange { //如果一轮排序没有元素被交换，则序列已经有序
      break
    }
  }

  fmt.Println(arr1nums)
}
