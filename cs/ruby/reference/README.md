# Ruby Reference Model — Explanation

This document explains **how Ruby treats variables, objects, and method arguments**, based on the examples in `reference_study.rb`.

Ruby is often said to use “pass-by-reference,” but the correct explanation is:

> **Ruby is pass-by-value — but the value being passed is a reference to an object.**

This means Ruby behaves *like* pass-by-reference in some cases, while still being fundamentally different from true reference semantics.

---

## 1. Variables store references, not objects

When you write:

```ruby
a = "hello"
b = a
```

Both `a` and `b` point to the **same object**.

```ruby
a.object_id == b.object_id  # => true
```

There is **one object**, referenced by multiple variables.

---

## 2. Destructive methods mutate the object itself

Ruby provides methods ending with `!`, such as:

```ruby
str = "ruby"
str.upcase!   # mutates the underlying string object
```

Since the object changes, any variable referencing it will see the mutation.

---

## 3. Mutations are shared across all references

Arrays and hashes behave the same way.

```ruby
x = [1, 2, 3]
y = x
y << 4
```

Both `x` and `y` now see `[1, 2, 3, 4]` because they reference the same array.

Mutation is shared.
Reassignment is *not*.

---

## 4. Method arguments are “copies of references”

Ruby passes a **copy of the reference**, not the object itself.

```ruby
def modify(arr)
  arr << 999   # mutates the original object
end
```

Calling:

```ruby
list = [1, 2, 3]
modify(list)
```

Both `list` and `arr` refer to the same array, so mutation is visible outside the method.

---

## 5. Reassignment inside methods does NOT affect the caller

```ruby
def reassign(a)
  a = "changed"  # only changes local variable 'a'
end
```

This only changes what the *local* variable `a` refers to.
The caller’s variable remains unchanged.

```ruby
text = "original"
reassign(text)
text  # => "original"
```

This is the key difference from true pass-by-reference.

---

## 6. Hash mutations behave the same way as arrays

```ruby
h1 = { a: 1 }
h2 = h1
h2[:b] = 2
```

Both variables reflect the change because they share the same underlying object.

---

## 7. Arrays and elements are separate objects

```ruby
arr = [1, 2, 3]
arr.object_id        # ID of the array
arr[0].object_id     # ID of the element
```

Ruby distinguishes the collection object and its contained objects.

---

# Summary: What “reference passing” means in Ruby

Ruby behaves as follows:

| Operation                                                      | Affects caller? | Explanation                             |
| -------------------------------------------------------------- | --------------- | --------------------------------------- |
| **Mutating the object (e.g., `<<`, `push`, `upcase!`, `[]=`)** | ✔ Yes           | Caller and callee share the same object |
| **Reassigning a variable (e.g., `a = ...`)**                   | ✘ No            | Only changes the local reference        |

Therefore:

> **Ruby is pass-by-value, but the value being passed is a reference to the object.**
> Mutations affect the caller, but reassignment inside the method does not.

This hybrid model often looks like pass-by-reference, but it is not true reference semantics.
