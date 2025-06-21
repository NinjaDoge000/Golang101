# Golang 101

func (receiver) methodName[T Type](args) returnType


| Type          | Reference-like? | Pass pointer needed? | Notes                                                |
| ------------- | --------------- | -------------------- | ---------------------------------------------------- |
| `interface{}` | ✅ Yes           | ❌                    | Already reference                                    |
| `map`         | ✅ Yes           | ❌                    | Shareable and mutable                                |
| `slice`       | ✅ Kinda         | ❌ or ✅               | Depends if you need to mutate slice vs backing array |
| `chan`        | ✅ Yes           | ❌                    | Already reference                                    |
| `func`        | ✅ Yes           | ❌                    | Already reference                                    |
| `string`      | ✅ Internally    | ❌ (usually)          | Immutable, pass by value fine                        |
| `struct`      | ❌ No            | ✅                    | Pointer for mutation or performance                  |
