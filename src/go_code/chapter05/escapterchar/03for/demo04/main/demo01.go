package main

import "fmt"

func main() {
	var classNum int
	var stuNum int
	fmt.Printf("请输入班级数量")
	fmt.Scanln(&classNum)
	fmt.Printf("请输入班级学生数量")
	fmt.Scanln(&stuNum)
	totalSum := 0.0
	for j := 1; j <= classNum; j++ {
		sum := 0.0
		for i := 1; i <= stuNum; i++ {
			var score float64
			fmt.Printf("请输入第%d班 第%d个学生的成绩\n", j, i)
			fmt.Scanln(&score)
			sum += score
		}
		fmt.Printf("第%d班级的平均分是%v\n", j, sum/float64(stuNum))
		totalSum += sum
	}
	fmt.Printf("所有班级的总成绩%v 所有班级的平均分是%v\n", totalSum, totalSum/float64(stuNum*classNum))

}
