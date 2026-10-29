# DSA Problem 1: Reverse a String (Go)

**Group:** Strings  
**Difficulty:** Easy  
**Patterns:** String traversal, slices, two pointers  
**Status:** Completed

---

## 1. Problem Statement

Given a string `s`, return its characters in reverse order.

### Examples

| Input | Output |
|---|---|
| `"hello"` | `"olleh"` |
| `"google"` | `"elgoog"` |
| `""` | `""` |
| `"a"` | `"a"` |
| `"abcd"` | `"dcba"` |

### Clarify the requirement

In Go, a string is a read-only sequence of bytes, usually encoded as UTF-8. `len(s)` returns the number of bytes, and `s[i]` returns a byte.

Before implementing a general solution, clarify what “character” means:

- **ASCII/bytes:** reverse individual bytes; appropriate for ASCII input.
- **Unicode code points:** convert to `[]rune` and reverse the runes.
- **User-perceived characters (grapheme clusters):** use grapheme-cluster segmentation when combining marks or joined emoji must remain together.

The optimized solution below reverses Unicode code points, not necessarily grapheme clusters.

---

## 2. How to Think About the Problem

For input `"hello"`, the required output is `"olleh"`.

Two initial ideas:

1. Start at the last position and build a new result from right to left.
2. Put one pointer at the beginning and another at the end, swap the elements, and move both pointers toward the middle.

The first idea is the simplest baseline. The second leads to the **two-pointer pattern**.

---

## 3. Approach 1 — Brute Force: Repeated String Concatenation

### Idea

Traverse the input from the last byte to the first and append each byte to a result string.

### Go implementation

```go
package main

import "fmt"

func reverseBruteForce(s string) string {
	result := ""

	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}

	return result
}

func main() {
	fmt.Println(reverseBruteForce("hello"))  // olleh
	fmt.Println(reverseBruteForce("google")) // elgoog
}
```

### Dry run: `"hello"`

| Iteration | Byte read | Result |
|---:|---|---|
| 1 | `o` | `"o"` |
| 2 | `l` | `"ol"` |
| 3 | `l` | `"oll"` |
| 4 | `e` | `"olle"` |
| 5 | `h` | `"olleh"` |

### Complexity

Let `n` be the input length in bytes.

- **Time:** O(n²) in the worst case due to repeated copying as the result grows.
- **Auxiliary space:** O(n) for the resulting string.

### Limitations

- Repeated concatenation can be inefficient for large inputs.
- This byte-indexing implementation is only suitable for ASCII input; it does not correctly reverse arbitrary UTF-8 text.

---

## 4. Approach 2 — Better: Byte Slice and `append`

### Idea

For ASCII input, convert the string to a mutable byte slice. Traverse it backward and append bytes to a result slice. Preallocate the result capacity to avoid unnecessary slice growth.

### Go implementation

```go
package main

import "fmt"

func reverseBetter(s string) string {
	bytes := []byte(s)
	result := make([]byte, 0, len(bytes))

	for i := len(bytes) - 1; i >= 0; i-- {
		result = append(result, bytes[i])
	}

	return string(result)
}

func main() {
	fmt.Println(reverseBetter("hello"))  // olleh
	fmt.Println(reverseBetter("google")) // elgoog
}
```

### Dry run: `"hello"`

| Byte read | Result slice |
|---|---|
| `o` | `"o"` |
| `l` | `"ol"` |
| `l` | `"oll"` |
| `e` | `"olle"` |
| `h` | `"olleh"` |

### Complexity

- **Time:** O(n).
- **Auxiliary space:** O(n) for the input byte copy and result slice.

### Why it is better

Appending to a preallocated byte slice avoids repeatedly rebuilding an immutable string.

### Limitation

A UTF-8 code point may occupy multiple bytes. Reversing individual bytes can produce invalid UTF-8, so this implementation is intended for ASCII input or explicitly byte-oriented data.

---

## 5. Approach 3 — Optimized for Unicode Code Points: Two Pointers

### Idea

1. Convert the string to a rune slice.
2. Initialize `left` to the first index and `right` to the last index.
3. Swap `runes[left]` and `runes[right]`.
4. Move `left` forward and `right` backward.
5. Stop when `left >= right`.
6. Convert the rune slice back to a string.

### Go implementation

```go
package main

import "fmt"

func reverseString(s string) string {
	runes := []rune(s)

	for left, right := 0, len(runes)-1; left < right; left, right = left+1, right-1 {
		runes[left], runes[right] = runes[right], runes[left]
	}

	return string(runes)
}

func main() {
	fmt.Println(reverseString("hello"))  // olleh
	fmt.Println(reverseString("google")) // elgoog
	fmt.Println(reverseString(""))       // ""
	fmt.Println(reverseString("a"))      // a
	fmt.Println(reverseString("नमस्ते")) // ेस्तमन (code-point reversal)
	fmt.Println(reverseString("A😊B"))   // B😊A
}
```

### Dry run: `"hello"`

Initial rune slice:

```text
Index:  0   1   2   3   4
Value:  h   e   l   l   o
        L               R
```

**First swap:** swap `h` and `o`.

```text
o e l l h
  L       R
```

**Second swap:** swap `e` and `l`.

```text
o l l e h
    L R
```

The pointers have met or crossed, so the loop stops. Result: `"olleh"`.

### Complexity

Let `n` be the number of Unicode code points.

- **Time:** O(n), including conversion and reversal.
- **Auxiliary space:** O(n) for the rune slice and output conversion.

This is asymptotically optimal for producing a reversed result while supporting Unicode code points.

### Unicode caveat

`[]rune` handles Unicode code points, not necessarily user-perceived characters. For example, a base character and its combining mark may be separate code points. Reversing them individually can change the visual result. If the requirement is grapheme-aware reversal, segment the input into grapheme clusters first.

---

## 6. Comparison of All Approaches

| Approach | Main technique | Time | Auxiliary space | Input assumptions |
|---|---|---:|---:|---|
| Brute force | Repeated string concatenation | O(n²) worst case | O(n) | ASCII |
| Better | Byte slice + `append` | O(n) | O(n) | ASCII / byte-oriented |
| Optimized | Rune slice + two pointers | O(n) | O(n) | Unicode code points |

The better and optimized approaches have the same asymptotic complexity. The rune-based approach is more appropriate when the problem requires Unicode code-point reversal. For ASCII-only input, reversing a byte slice in place with two pointers is also an O(n)-time solution.

---

## 7. Edge Cases and Tests

| Case | Input | Expected output / behavior |
|---|---|---|
| Empty string | `""` | `""` |
| Single character | `"a"` | `"a"` |
| Two characters | `"ab"` | `"ba"` |
| Odd length | `"hello"` | `"olleh"` |
| Even length | `"abcd"` | `"dcba"` |
| Repeated characters | `"aaaa"` | `"aaaa"` |
| Spaces | `"hello world"` | `"dlrow olleh"` |
| Unicode code points | `"A😊B"` | `"B😊A"` |
| Marathi text | `"नमस्ते"` | Reverse code-point sequence; grapheme behavior may differ |

### Go unit test

Save the optimized implementation in `reverse.go` and use this test in `reverse_test.go`:

```go
package main

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "normal string", input: "hello", want: "olleh"},
		{name: "empty string", input: "", want: ""},
		{name: "single character", input: "a", want: "a"},
		{name: "even length", input: "abcd", want: "dcba"},
		{name: "repeated characters", input: "aaaa", want: "aaaa"},
		{name: "unicode code points", input: "A😊B", want: "B😊A"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseString(tt.input)
			if got != tt.want {
				t.Errorf("reverseString(%q) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}
```

Run:

```powershell
go test -v
```

Expected: all six test cases pass.

---

## 8. Correctness Reasoning

For the two-pointer algorithm, each iteration swaps the two outermost elements that have not yet been placed in their final positions. The left pointer moves right and the right pointer moves left. Therefore, after each swap, those two elements are in their reversed positions. When the pointers meet or cross, every pair has been swapped exactly once, and the sequence is reversed.

---

## 9. Interview-Ready Explanation

> I would first clarify whether the input is ASCII or arbitrary Unicode text, because Go strings are UTF-8 byte sequences and string indexing returns a byte.
>
> A straightforward solution traverses the string from right to left and appends each character to a result string. Repeated concatenation can lead to O(n²) time complexity in the worst case.
>
> For ASCII input, I can improve this by using a byte slice and appending into a preallocated result slice, achieving O(n) time.
>
> For Unicode code-point reversal, I convert the string to a rune slice and use two pointers. One starts at the beginning and the other at the end. I swap their elements and move both pointers toward the center until they meet or cross.
>
> This solution takes O(n) time and O(n) auxiliary space. I would test empty strings, single characters, odd and even lengths, repeated characters, and Unicode input. If the requirement is to preserve user-perceived characters, I would use grapheme-cluster segmentation rather than assuming each rune is one visible character.

### Likely follow-up questions

**1. Why not reverse a Go string directly?**  
Go strings are immutable, so we create a mutable representation such as `[]byte` or `[]rune`.

**2. What is the difference between `len(s)` and `len([]rune(s))`?**  
`len(s)` counts bytes. `len([]rune(s))` counts Unicode code points after conversion.

**3. Is O(n) time optimal?**  
Yes, asymptotically: producing the reversed output requires work proportional to the input size.

**4. Does `[]rune` correctly reverse every visible character?**  
No. It reverses Unicode code points, not necessarily grapheme clusters.

**5. Can the ASCII solution use O(1) auxiliary space?**  
The reversal itself can use O(1) extra space by swapping bytes in a mutable byte slice in place. Converting the input string to a byte slice still requires O(n) space, and returning a new string requires output storage.

---

## 10. Final Takeaways

- Start with the simplest correct solution, then identify its bottleneck.
- Repeated string concatenation can cause O(n²) work.
- Use byte slices for byte-oriented or ASCII problems.
- Use rune slices when the requirement is Unicode code-point processing.
- Recognize two pointers as a reusable pattern.
- Always clarify character semantics when Unicode matters.
- State time and space complexity precisely.
