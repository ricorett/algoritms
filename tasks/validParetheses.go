func isValid(s string) bool {
    stack := []rune{}

    for _, c := range s {
        switch c {

        case '(':
            stack = append(stack, ')')

        case '{':
            stack = append(stack, '}')

        case '[':
            stack = append(stack, ']')

        default:
            if len(stack) == 0 || stack[len(stack)-1] != c {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }

    return len(stack) == 0
}
