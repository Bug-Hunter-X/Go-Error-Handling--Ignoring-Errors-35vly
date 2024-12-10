# Go Error Handling: Ignoring Errors

This example demonstrates a common error in Go: ignoring errors returned from functions.  Ignoring errors can lead to silent failures and unexpected behavior in your application.  The `bug.go` file shows the problematic code, and `bugSolution.go` provides a corrected version.

## Problem

The function `myFunc` performs division.  If the input `x` is 0, it returns an error. However, the calling code doesn't check for this error, potentially leading to a runtime panic due to division by zero.

## Solution

The `bugSolution.go` file shows the corrected version which explicitly checks for and handles the error, preventing a crash and providing better error handling.