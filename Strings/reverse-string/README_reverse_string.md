
# DSA 01 — Reverse a String (Go)

**Category:** Strings  
**Difficulty:** Easy  
**Patterns:** String Traversal, Slices, Two Pointers  
**Language:** Go  
**Status:** Completed ✅

---

## 1. Understanding the Problem

The problem is simple: given a string, return the same string but in reverse order.

For example:

| Input | Output |
|---|---|
| `"hello"` | `"olleh"` |
| `"google"` | `"elgoog"` |
| `"abcd"` | `"dcba"` |
| `"a"` | `"a"` |
| `""` | `""` |

My first thought was to start from the last character and build a new string by adding each character until I reach the beginning.

But before implementing it, there is something important to understand about Go.

Go strings are immutable and UTF-8 encoded. This means we cannot directly modify a character inside a string, and accessing `s[i]` gives us a byte, not necessarily a complete character.

So, we need to be careful about whether we're reversing bytes or actual Unicode characters.

For this problem, I'll try three approaches, starting with the simplest one.

---

## 2. Approach 1 — Simple Approach (Brute Force)

### My Thinking

The first idea that came to mind was to traverse the string backwards.

If the input is `"hello"`, I can start from the last character (`o`) and keep adding characters to a new string until I reach the beginning.

### Code

```go
func reverseBruteForce(s string) string {
    result := ""

    for i := len(s) - 1; i >= 0; i-- {
        result += string(s[i])
    }

    return result
}
```

### Dry Run

Input: `"hello"`

| Step | Character | Result |
|---|---|---|
| 1 | `o` | `"o"` |
| 2 | `l` | `"ol"` |
| 3 | `l` | `"oll"` |
| 4 | `e` | `"olle"` |
| 5 | `h` | `"olleh"` |

Finally, we get `"olleh"`, which is what we wanted.

### Complexity Analysis

Let `n` be the input length in bytes.

- **Time Complexity:** O(n²) in the worst case because of repeated string concatenation.
- **Space Complexity:** O(n) for the resulting string.

### What I Noticed

This approach works, but there is a problem.

Go strings are immutable. Every time we concatenate using `+`, a new string may need to be allocated and the previous content copied.

So, as the result grows, we keep doing unnecessary copying.

I wanted to avoid that, so I tried using a slice instead.

---

## 3. Approach 2 — Using a Byte Slice and Append

### My Thinking

Instead of building the result using string concatenation, I can convert the input string into a byte slice.

A slice is mutable, so I can work with its elements more efficiently.

I'll create an empty result slice, allocate enough capacity in advance, and append bytes while traversing the input backwards.

### Code

```go
func reverseBetter(s string) string {
    bytes := []byte(s)
    result := make([]byte, 0, len(bytes))

    for i := len(bytes) - 1; i >= 0; i-- {
        result = append(result, bytes[i])
    }

    return string(result)
}
```

### Dry Run

Input: `"hello"`

| Step | Byte Added | Result |
|---|---|---|
| 1 | `o` | `"o"` |
| 2 | `l` | `"ol"` |
| 3 | `l` | `"oll"` |
| 4 | `e` | `"olle"` |
| 5 | `h` | `"olleh"` |

The result is the same, but now we're using a slice instead of repeatedly concatenating strings.

### Complexity Analysis

- **Time Complexity:** O(n)
- **Space Complexity:** O(n)

### Why This Is Better

We preallocated the result slice using:

```go
result := make([]byte, 0, len(bytes))
```

This gives the slice enough capacity for the entire result and avoids unnecessary capacity growth.

### One Important Limitation

This approach reverses bytes, not Unicode characters.

For example, emojis and characters from languages like Marathi or Hindi can occupy multiple bytes in UTF-8.

Reversing those bytes individually can produce invalid UTF-8.

So, this approach is suitable for ASCII input or explicitly byte-oriented data.

I wanted a solution that handles Unicode code points too.

---

## 4. Approach 3 — Two Pointers (Final Approach)

### My Thinking

Instead of creating a new result by traversing backwards, what if I simply reverse the elements in place?

This is where the **Two Pointers** pattern comes in.

The idea is simple:

1. Convert the string into a rune slice.
2. Keep one pointer at the beginning.
3. Keep another pointer at the end.
4. Swap the elements at both positions.
5. Move the left pointer forward and the right pointer backward.
6. Continue until both pointers meet or cross.

### Code

```go
func reverseString(s string) string {
    runes := []rune(s)

    for left, right := 0, len(runes)-1; left < right; left, right = left+1, right-1 {
        runes[left], runes[right] = runes[right], runes[left]
    }

    return string(runes)
}
```

### Dry Run

Input: `"hello"`

First, convert the string into a rune slice.

```text
Index:  0   1   2   3   4
Value:  h   e   l   l   o
        L               R
```

**Step 1:** Swap the leftmost and rightmost elements.

```text
Before: h e l l o
After:  o e l l h
          L   R
```

**Step 2:** Move both pointers towards the center and swap again.

```text
Before: o e l l h
After:  o l l e h
            L R
```

Now the pointers have met or crossed, so we stop.

Final output:

```text
olleh
```

That's it! We reversed the string using two pointers.

### Complexity Analysis

Let `n` be the number of Unicode code points.

- **Time Complexity:** O(n)
- **Space Complexity:** O(n)

The conversion to `[]rune` takes O(n), and the reversal also takes O(n).

The rune slice requires O(n) space.

### Why I Chose This Approach

This is the approach I prefer when the requirement is to reverse Unicode code points.

- No repeated string concatenation.
- Easy to understand and implement.
- Uses the two-pointer pattern.
- Handles multi-byte UTF-8 code points correctly.

One thing to remember: `[]rune` reverses Unicode code points, not necessarily complete user-perceived characters. Some emojis and combining characters require grapheme-cluster handling.

---

## 5. Comparing All Three Approaches

After trying all three approaches, here's what I learned:

| Approach | Technique | Time | Space |
|---|---|---|---|
| Brute Force | String concatenation | O(n²) | O(n) |
| Better | Byte slice + append | O(n) | O(n) |
| Final | Rune slice + Two Pointers | O(n) | O(n) |

The important difference isn't just the complexity.

It's understanding which representation to use depending on the input:

- `string` — Immutable UTF-8 byte sequence.
- `[]byte` — Mutable bytes, suitable for ASCII and byte-oriented operations.
- `[]rune` — Mutable Unicode code points.

For ASCII-only input, we could also reverse a byte slice in place using two pointers.

---

## 6. Edge Cases I Checked

Before considering the solution complete, I wanted to make sure it handles more than just `"hello"`.

| Test Case | Input | Expected Output |
|---|---|---|
| Empty string | `""` | `""` |
| Single character | `"a"` | `"a"` |
| Two characters | `"ab"` | `"ba"` |
| Odd length | `"hello"` | `"olleh"` |
| Even length | `"abcd"` | `"dcba"` |
| Repeated characters | `"aaaa"` | `"aaaa"` |
| Spaces | `"hello world"` | `"dlrow olleh"` |
| Unicode | `"A😊B"` | `"B😊A"` |

I also considered Marathi text (`"नमस्ते"`) to understand how Unicode code-point reversal differs from reversing visible characters.

---

## 7. Go Unit Tests

I wrote a few test cases to verify that my final implementation works as expected.

**reverse_test.go**

```go
package main

import "testing"

func TestReverseString(t *testing.T) {
    tests := []struct {
        name  string
        input string
        want  string
    }{
        {"normal string", "hello", "olleh"},
        {"empty string", "", ""},
        {"single character", "a", "a"},
        {"even length", "abcd", "dcba"},
        {"repeated characters", "aaaa", "aaaa"},
        {"unicode code points", "A😊B", "B😊A"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := reverseString(tt.input)

            if got != tt.want {
                t.Errorf(
                    "reverseString(%q) = %q; want %q",
                    tt.input,
                    got,
                    tt.want,
                )
            }
        })
    }
}
```

Run the tests:

```bash
go test -v
```

All six test cases should pass.

---

## 8. What I Learned From This Problem

This problem looks easy, but it helped me understand a few important things about Go.

1. **Strings in Go:** A string is immutable, and indexing it gives a byte.
2. **Slices:** When I need to modify elements, I can convert the string into a mutable slice.
3. **Time Complexity:** Repeated string concatenation can be expensive.
4. **Two Pointers:** Instead of building a reversed result, I can swap elements from both ends.
5. **Unicode:** Bytes, runes, and visible characters are not always the same thing.

Most importantly, I learned not to jump directly to the optimized solution.

Starting with a simple approach, understanding its limitations, and then improving it made the problem much easier to understand.

---

## 9. Interview Explanation

If I were asked to solve this problem in an interview, this is how I would explain my approach:

> First, I would clarify whether the input contains only ASCII characters or arbitrary Unicode text, because Go strings are UTF-8 encoded and indexing a string returns a byte.
>
> My initial approach would be to traverse the string from right to left and build a new result. However, repeated string concatenation can lead to O(n²) time complexity.
>
> To improve it, I can use a mutable slice. For Unicode code-point reversal, I would convert the string into a rune slice and use two pointers.
>
> One pointer starts at the beginning and the other at the end. I swap their elements and move both pointers towards the center until they meet or cross.
>
> This gives us O(n) time complexity and O(n) auxiliary space.

### Questions I Might Get Asked

**Q1. Why can't we reverse a string directly in Go?**

Because strings are immutable in Go. We need to convert them into a mutable representation like `[]byte` or `[]rune`.

**Q2. What's the difference between `len(s)` and `len([]rune(s))`?**

`len(s)` returns the number of bytes, whereas `len([]rune(s))` returns the number of Unicode code points.

**Q3. Can we solve this with O(1) extra space?**

The swapping operation itself needs O(1) extra space if we already have a mutable slice. However, converting a Go string into a slice and creating the returned string require additional memory.

**Q4. Does `[]rune` handle every Unicode character correctly?**

It handles Unicode code points, but not necessarily grapheme clusters. Some visible characters consist of multiple code points.

**Q5. Why use two pointers?**

Because we can swap elements from both ends and move towards the center, avoiding the need to build a second reversed slice.

---

## 10. Final Takeaway

My final choice for this problem is the **Rune Slice + Two Pointers** approach when Unicode code-point reversal is required.

The biggest takeaway isn't just reversing a string. It's understanding how Go handles strings, when to use slices, and how the two-pointer pattern can simplify a problem.

**Pattern to remember:** Whenever you need to reverse or swap elements from opposite ends of a sequence, think about the Two Pointers approach.
