import (
    "sort"
)
type byLength []string
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}


func allLongestStrings(inputArray []string) []string {
    tempDetails := make([]string,0)
    for _, val := range inputArray{
        tempDetails = append(tempDetails, val)
    }
    sort.Sort(byLength(tempDetails))

    fmt.Println(tempDetails, inputArray)
    largestString := len(tempDetails[len(tempDetails)-1])
    tempString := make([]string,0)
    for _,str := range inputArray{
        if len(str) == largestString {
            tempString = append(tempString, str)
        }
    }
    return tempString
}

