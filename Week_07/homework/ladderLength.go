package homework

func ladderLength(beginWord string, endWord string, wordList []string) int {
    wordMap := make(map[string]int, 0)
    for i, v := range wordList {
        wordMap[v] = i
    }
    // if endword not in wordlist, return 0
    if _, ok := wordMap[endWord]; !ok {
        return 0
    }
    startQueue := []string{beginWord}
    endQueue := []string{endWord}
    visited := make([]string, len(wordList))
    step := 1
    for len(startQueue) > 0 && len(endQueue) > 0{
        if len(startQueue) > len(endQueue) {
            startQueue, endQueue = endQueue, startQueue
        }
        temp := make([]string, 0)
        for i := 0; i < len(startQueue); i++ {
            chars := []byte(startQueue[i])
            for j := 0; j < len(chars); j++ {
                o := chars[j]
                for c := 'a'; c <= 'z'; c++ {
                    chars[j] = byte(c)
                    target := string(chars)
                    if containsInArr(endQueue, target) {
                        return step+1
                    }

                    _, ok := wordMap[target]
                    if ok && !containsInArr(visited, target) {
                       temp = append(temp, target)
                       visited = append(visited, target)
                    }
                    chars[j] = o

                }
            }
        }
        startQueue = temp
        step++
    }
    return 0
}

func containsInArr(arr []string, word string) bool {
    for _, v :=  range arr {
        if v == word {
            return true
        }
    }
    return false
}