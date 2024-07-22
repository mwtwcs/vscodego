package main

func  IntsCopy(src []int, maxLen int) []int {
    if maxLen <= 0 {
        return []int{} 
    }
    copyLen := maxLen
    if copyLen > len(src) {
        copyLen = len(src)
    }

    
    copySlice := make([]int, copyLen)

 
    copy(copySlice, src)

    return copySlice
}
