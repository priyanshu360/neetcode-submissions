func isAnagram(s string, t string) bool {
    a:= make([]int, 27)

    for _, c := range s {
        a[int(c - 'a')] += 1
    }

    for _, c := range t {
        a[int(c - 'a')] -= 1
    }

    for _, val := range(a) {
        if val != 0 {
            return false
        }
    }
    return true
}
