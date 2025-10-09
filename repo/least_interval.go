package repo

func leastInterval(tasks []byte, n int) int {
	// 贪心
	// 时间取决于两者
	// 1.所有任务按时执行，不存在冷却时间
	// 2.最多任务的最长冷却+执行时间
	maxCounter := 0
	counter := make(map[byte]int)
	count := 0
	for _, t := range tasks {
		counter[t]++
		if counter[t] > maxCounter {
			maxCounter = counter[t]
			count = 0
		} else if counter[t] == maxCounter {
			count++
		}
	}
	return max((n+1)*maxCounter-n+count, len(tasks))

}
