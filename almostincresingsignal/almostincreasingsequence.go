import (
    "sort"
)

func noRepeating(slice []int) bool {
    intMap := make(map[int]bool,0)
    for _, val := range slice{
        if found,_ := intMap[val]; found{
            return false
        }
        intMap[val] = true
    }
    return true
}


func remove(slice []int, s int) []int {
    //fmt.Println(slice, s)
    tempStr := make([]int,0)
    for i:=0;i<len(slice);i++ {
    	if i == s{
		tempStr = append(tempStr,slice[i+1:]...)
		//fmt.Println(tempStr, slice[i+1:])
		break
	}
	tempStr = append(tempStr,slice[i])
    }
    return tempStr
}
func almostIncreasingSequence(sequence []int) bool {
    if len(sequence) != 2 && sort.IntsAreSorted(sequence){
               return false
    }

    for i:=0;i<len(sequence); i++ {
        tempSequence := sequence
        tempSequence = remove(tempSequence, i)
	//fmt.Println("testHere :", tempSequence)
        if sort.IntsAreSorted(tempSequence[:len(tempSequence)]){
             if i>0 && !noRepeating(tempSequence[i-1:]){
               continue
            }
            return true
        }
    }
    return false
}

