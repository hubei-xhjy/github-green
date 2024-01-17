package main

import (
	"fmt"
	"math/rand"
	"os/exec"
	"time"
)

func main() {
	// 设置开始时间
	startTime := time.Date(2023, 5, 16, 2, 3, 3, 0, time.Local)
	// 结束时间为当前日期
	endTime := time.Now().Local()

	rand.Seed(time.Now().UnixNano()) // 初始化随机数种子

	// end < start 返回 true
	for endTime.After(startTime) {
		numCommits := rand.Intn(6) // 生成 0 - 5 之间的随机数
		for i := 0; i < numCommits; i++ {
			// 生成随机时分秒
			randomTime := rand.Intn(24 * 60 * 60) // 一天的秒数

			commitDate := startTime.Add(time.Duration(randomTime) * time.Second).Format("Mon Jan 02 15:04:05 2006 -0700")
			// --allow-empty 允许空提交 --date 设置提交时间
			cmd := exec.Command("git", "commit", "--allow-empty", "--date", commitDate, "-m", "Green everyday!")

			if err := cmd.Run(); err != nil {
				fmt.Println(err)
			}
		}

		// 增加一天
		startTime = startTime.AddDate(0, 0, 1)
	}
	fmt.Println("Done!")
}
