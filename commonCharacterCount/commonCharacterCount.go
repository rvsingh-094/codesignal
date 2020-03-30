func commonCharacterCount(s1 string, s2 string) int {
    strMap := make(map[byte]int,0)
    for i:=0;i<len(s1);i++{
        if  _,found := strMap[s1[i]]; found {
            strMap[s1[i]] += 1
        }else {
            strMap[s1[i]] = 1
        }
    }
    repeatedCount := 0
    for i:=0;i<len(s2);i++{
        if _, found := strMap[s2[i]]; found {
            if strMap[s2[i]] > 0 {
                repeatedCount += 1
                strMap[s2[i]] -= 1
            }
        }
    }
    return repeatedCount 
}

