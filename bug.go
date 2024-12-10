func myFunc(x int) (int, error) {
    if x == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return 10 / x, nil
}