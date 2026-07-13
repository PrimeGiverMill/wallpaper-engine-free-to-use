package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "2.3" )

func E4SPf4UiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XoQfDJIcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jswNZiY2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9VgSKSExWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s2hQ62CzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WMdjV05tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lce8UobRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func THRTzLGhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qxVRxv9VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iJPyixpNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6F4sCDN5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VEpm6KPiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8V88ieb5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0G3lYDjwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EOpG0byyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DjpA0T3tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func glvaxvzdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z8kBsrviWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pFA310W7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aXuRnZoUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ahgPpjntWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AcbmyJaaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xEXOecGhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NaCfEFnnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gUm6oFmVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rdW1VIwnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 866cqsRyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LDkqAw2uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YvlRkbUsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nTInvaYpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AiH2ixuIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CiQouZgyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mZ0P94eMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nld1FUB5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tRvH2Up7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5z4zlLoyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KdYOhsQ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QfXVdgZ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jwyTwG5XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iLFZt3DtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h6oAHAOoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3xldO5dqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OjqaU6bLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NnKXTBnSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WxS7qrIpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W38e4qzFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0OtoaztcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Bd7LLyzyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EDkD2r1YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OA1iOk8sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DKGoZcJPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PzgIz2vHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V7R24t4AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UQzrbZLiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QIY3TbSnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T7Qzu3SJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 214VsDJ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PDWZiyR6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func afbiZ1rrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ouHEgNx2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qPSOzdwfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hNHlxjIWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aD6bGSHkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FOPdQ4dZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lLkm1XqmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1WCtXk4OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m6FBYW9AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OfOTRzEjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qrF0eHlsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8BEh7dZRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0CTl9sojWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PxMTgaPNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DUybIP45Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gNuh9JtQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dEP5WLjEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QjBqouqBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KsjMe9gMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 15RBrz2RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wd3lA8NDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X8ZwZURnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oXs3QXnyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9PdaXN2BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rhNzOerBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kmsODBWrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wdeXFRb0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PpPoUg99Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 61wxZ4jPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bh0iLH9SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5qlRc42bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rGVP80TEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rxXiPZgNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0u49pAJKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jkLy4WFoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5aVAKo01Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UjI8sKAZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IXfCrJq2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i52Y0TK8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XisYNrRcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lyeVJJZmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XbPDCL8SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BjIEUZbDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PxEPqkFQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oZOGsuOQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iIIRBDLPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9Z4H2NEFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tGNp0sccWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IGTusY2FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5yGkhsFzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aiMkEsGqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hozMIR55Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UEqINzaoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I2Qnvnf9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9WSzix0qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FdXvMaFQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ps32zIqVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ErKCiRgvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GOIZ3Z24Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D295tqK2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YzbJCGzGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EozquNynWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XoRlPFyDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sGdUutMSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func arnpwkZpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BocB3XA4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AojblPZuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ft7BvBgxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K1rn0CwxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xQ92O4YCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T1jkI1UQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tSKOhmzEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6vl3YcSeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q19MN6G8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oxbrn4ZlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xSH7CXeSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zkz3x3GYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oYRkNgAMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iOKoFvFtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2gMVx1GmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CgPQLTEzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wZcM1hQGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dhEezXs1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func edH4rHi7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hb8hZnINWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q8nklef7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SIKYQg8rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F0H9S70nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e3oRcEwsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FXKHJsFFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e677GV1CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rmoy3pHLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M6brzYgSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iM4hCaaQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ypXdMU8UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eBHEFucDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vvtft9apWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i83r5TrqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m1bt1nQKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lf9nlg93Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vFASU8lmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0WoI0ZQ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aXp0HLffWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cB7ee1yiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M2btXTSCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8q2emISlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8j08BFFlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ViTvPG6hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fF76L0qIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9LB9KUQzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QJtC2a1AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z38c9hiWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lJcuQdboWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZvLMr0YzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rMyqxWPnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bY9uXQXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gNOHGjesWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EJOU5qepWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tbNUpRDOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WBckXZHfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d6Ds9sHPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vDRsa8siWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LjF1juHlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H8LtlKipWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hIppgdMgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MdwZSUteWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pbDN4mrIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0lAuFWnqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hB6FofF5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lzne56tdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B5MTfrV5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vOucAIgkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u89DojVQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sqseAwfSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HfxlZw26Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0XviqodmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7DPtw4qBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VC2QC0eFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3H1gtbjfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t79WLsHtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GdJP0gFCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PmH9LRcsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MyYTmPc6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P8AApEnFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DRKY457pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GBCSHj2KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6qFDyXtpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rGXNdX2jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cDjAou58Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T1eygLxYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HXgQxFPlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func acDPwCXWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fPh3PfxhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WEb4oVhTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u4txjvESWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DpheZI7uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LusEpPzBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nvCbZOvHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2cELIi2bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TpizVFS4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Frb6EQekWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bALCo94HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func liXrv0U0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k7KvHDdJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9AS8Qu7AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UxywXwz4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p5D7vuzbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2GK0ZzGaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oh4cMvCpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y0XC5iyJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oK4PmkSyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LWwF7D0hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EewbaUyZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yLWNWyvfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4vD3We99Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jAOEFi22Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lR772yTRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OquN6c16Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bLbbkY6DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zk5bU8ctWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JpauQgdDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9uDQRi8VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i8EG2KUAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sjvdpImxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xqmguoYKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xu5e7ZhbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dyiUyEuvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i7W30mizWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r3evUyNiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7o9GYyLBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CmLNzVccWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func whPVissQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ROlUDacoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5hO6OntaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0sRYBvGKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OyB3G0GoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B1N1gg4WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XVT0twggWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AyOt5jKvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G258ki3bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gdQjpHttWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uAJ2DEX7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QMVR6NKtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ThK30vtXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6NRfaRXLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P0iJpkXwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ce8ok1CMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CNKKFp3JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cAQf2unAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XY6L7H3mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OsklaQ2nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RmdY7sN7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2GcEePWAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UEoSMFiuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7td22CNBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ijCd1W79Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s9FwyUcvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RVOJDgYyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GJaYtZy2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j3aBkQfbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ezYxK85rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func siNLyBdoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GP1bWnI4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pcv8VoHzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X30YCDBJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6xvAlGzXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wfj5L1rKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eWQAraTnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pHjhkfC9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qz31CC7OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HUvQcxreWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3AjvahygWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OOhtRvHvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9D6b5BUIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aJGARQvQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3mlB8vj7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t9cMy6W9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VzwI4MqZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func szwGx6IwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ELQggYKCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func taxzFFknWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n2HkGMbKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PoGdM2ZbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yhSobwMYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yhwgllvNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jxteGsuDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f3A9vGGgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9PxdCkIPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ebQyIoDBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UFO3jx2IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i33zJnfgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RjRZysMyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8p0FkNYZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6XYESy4NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fGMwGhOKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HSzNa9y1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kgvUEPTBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3VJCALGWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lBd5MNuJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L659o6l4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fsvIxipVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9pvMadtTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WFqMMrVIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cPl4h1POWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lFRx8EyJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7WSCpDicWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ygVuvRlvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8v2mspTdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U1odYVKUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1PgQ7FYiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yCJ2ACfYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v6kxSe7lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 93nS4I2QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UYaoFDVcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m7hmAmdUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EPzxp57KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hWblQxApWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fWX1RrCqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mkwPUsvSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gU1sjMD5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AK6NidOtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iZdJrj8pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WQFwxbUqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rqoFbtR8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rwsQGA5zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hYMYLLlqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MSAxBhQsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8LVs7AcIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bjaVWRJ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AqSvzU0sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 49dfpAktWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8OuTC1BRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ogbOZIVHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L0OmNuIOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4kJkVA0mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q2F8hBVMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O2PFT1XxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xQIpMPrJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WDVulHOEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jPMdaVgHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func krBeif2NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uvZyECIvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B1yKK8dKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0G6fxqbqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GdLGwo57Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QxhJ4wOFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J4HTv1P7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cDsnf1wFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GxYZlDYRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Heep5jbGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EEdIWFXsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GbU6QcMlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 420qEm8AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oBNA4uCPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NDKijm1tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EopPSCIkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jPFt7vvGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sViKu5gwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fdQ0pDyJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OhRxb74vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6TLKO7FKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zbxxB3LZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KIJjrO3dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9n1ED8v5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j2VJB7mZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G6AiJasNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yHcg85okWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VFoLW8x6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func npidTSQIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7xRmGjf8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KFUJ9U0QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qkq11ikTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iVnoUJNGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func toqlMMYyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NzRHLevhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a3J7T6clWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HY4fPB9tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZhdBCxneWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZscJ53jeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OdZGI8EpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3vXA7rl5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 45bfzpPeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v2oDAii4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3KkbAZeUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ALLAIzVtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YO01hns9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ScqNk6TuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r6idF4idWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FHzgyiu9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G8Uc6GesWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xrVHkpgXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UWjfTU5gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vOGDUPGnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v9uhcO4VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xZHZuhJzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0xaz1xvjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PB3nT2DkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N96myfrbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c1HdxeGMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iy4A7tqpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WBERS3eWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ipzU4LGNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dLOC430cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fxhFy0wlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qIYWN8CMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s0PfXkvsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 31WseYtoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ltUplPM9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nC51rzjfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func If60Bx6sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IjVmarNIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gG76IuCLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xOlwBsvxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I5lgKffXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y8iSd5BUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kZMVkoYQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yLqxNRAOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XtsCCpBXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EN8Pnc6hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dx9VIEeLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JwUGcLgQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cBaeiHn2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2IxtFlBoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OpBFFWqCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yCSn13V5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rp2oqZ5hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qCXbNbrvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RJiTs2TqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3YQErErIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RSTF2iqwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func upyJ6YlGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O7ElwNZSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4ZheDaNqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NlIRgLO5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4M3GaKuUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EVqAOcKIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kLjRVkXYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tPHb4kw5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cFSpHhZZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I1hic7UsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 478fXtysWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZECwqCtDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Ri7ssKSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JpF6m8L5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r6wd64CCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sYrci3JiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wnzBJWQOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9FtJs2uHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func huMSyMdkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jMGMVzI2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h4HEnKvUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vZ1kin20Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tyHWvJulWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dqWxPFoFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UDIL34EFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n4m1O92EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uFWRxRsTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6khSCok9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func htQcUzhQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nMCF71InWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sQOPTwBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FBznTrLNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ijB7Ex3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CGxJwXbDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PsSH0MOuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CLBaomAlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2oTRIMzmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ef7SARMWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hRyOneWaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hDC6yt0hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rXG2xYVvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9C48xMEMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZiOt5zkRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vgTAJktEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IFLULKVhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CvY4qqvmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pAppuB47Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GUflfJsOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yOOzjAJ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j88VsK8uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J6hhGfAoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mizyItiiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gPTZcMtSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6M8bEezcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JoPDUi01Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mTWs89cHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qeXIyg77Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b1D4BYEUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YURJkdhNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yJgu9L4CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NDAWiy3KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e8dwim9wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fHklxpcRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mKBcyz0OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sUlbidq4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8XgdG7JTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SwaiDTwRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SP6MQumSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2obnNWL2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Eerj0ipQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yx9i9LzsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5xLBUGNdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TBmS4yhcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rufY0FVZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bkwHPwUOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uXXwYj0iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a0KTOXfmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vBGTSng6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Td6PUO2TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E9o6lKwFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wOtx9gWtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1kzdMeLwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qs50JzZWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xGJhd7sZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wPppM1JAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vU7pcMy9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OEEYMiGWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9d2T939pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Uia73yOiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ox7L5PjeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H3vPQFFzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vo8CwAreWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V5FVfpE6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4XzMWNnWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Azi2oHz5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g57uRI8pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZHYCf3JvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HFGN9KAsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2NyJjCuGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4LIMiHFGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ufsLTI7mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mnHFGF8rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IuNBDbZKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HVpeBKxkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eWnOA6FAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WaiKkL8kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tqbknu7JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gAlpH2oNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FAgGJEu9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XYaltsyXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WavO14gJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HJYH9aPrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UoXIe00SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OafwmWn9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GqsVXrgfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CdUGeLNAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E5lcI2x0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BJXxXyv5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HuEhl8VgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mi1t1LSAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CYghbl3fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HWyXxbKDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GjxBSyOHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8GKaoG6xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p3yQQbDAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func odGLMXwiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r23n17O8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bpE3crjjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WzCLJ78LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v0dDz5OcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rfP3wGQPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cSVzVUnuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o3hbrFJ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xitf3BgBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MlmXUR5YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NKbepXQJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dBVjfUMJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SngVidb8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6MoAqKddWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2D98ShECWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sMB1LelBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EK9eNPcdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r6q33jyXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vIB7cLdOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zFcHztdAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T4TYwM8RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AMLTNOAaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aeOYMmEcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JAVzlPEdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0bn3tJr1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9N0xhCbGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oHEhN7qgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func grdbzZclWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1pAhCyWBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vzU12m49Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tPCsKoacWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 24dfVihoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3fRdt3qiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xIx62RVPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iGlEXlDDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OZG98KNuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZLyWCMP4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nriTPgrKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JVFloao4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QdrJRB4ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jK1raATfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qoyVin5BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5ifhNM5jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qsPbQ1J7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZxjAWhnoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iZm0lkD0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c8l8VoQIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gsdw9BnUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KEuPgIGyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UCgjTUV7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TG6qAKPcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4qsnKLKwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lHF43N9xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lx5audfvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SVwbsabLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KTeUHAObWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JtXBWie8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0rYyQ6D6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x6UAC5rrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aqEi4hOzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pvfEqRKiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tNNax4CPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O7MSuZExWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9rTSYiYbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9zBwcs9PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t6mDiD9oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o76zJ3esWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TGxz9ZPlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nLFMZHc0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TL7EKWBXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uy59GiSEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ofAKvI6tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 05XlcSCJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rKrfpt5RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KEN87H23Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e6NBshdjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J4So4PNaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yIorQI5aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P0NFRPn7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fHwfVsffWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fTAi7iWHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yrzvDbulWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T1VUNyGxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func khXz9QRBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yUeYYBx0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ST4Bt3VuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UZ6zQGOwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dl7OaI5RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hpBBmAX9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qNBl0nG0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kFLh4uI0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9JPsuLEcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kf4CpZVvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rxwaumCAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kiLSUqVNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kHI1Kw5vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IczzZ6SxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JTkIN5mzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LLUPtBSPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9kVBxs2nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QBaVPeEVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rMXPAfhyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PdfxsiuWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d3cjOsMiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tNcQjgjiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xIjxmhUSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m34GS8ubWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func idhHbAuSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zDVAGfxlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JBi3GMUMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3qUrivsOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GMSRUKb5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UlfJcd8ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9950tWlJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Je0BM0wIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wo4w66xCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rZmvqnsGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OOkAyoY1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v5JuSuBlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func etzGPW7kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eHcKilu4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MJVqQBomWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AnS1Gy9RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CY86GzOuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sxAlzOwuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7X2zXIO3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A4IIAp4GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UtRTAyQVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eAnRGYvaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p8m7Fnx5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OE930tFXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8Q8Fa8FMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sUl1rtG4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2HWQUctcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1DgSpUkGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yS74XLMIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jiqYYYNyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uiRH7q1wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wdkkHEXiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I6hjHu7nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P5Lo1mvtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VKnNq3ocWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dmhgsVLBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BT2bvEX8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func acDmcuyLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YkZgVbcbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jphcDzd5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fsh1IWY6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WmHKLGUGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BtAR3pyPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ERmHFOLBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e8ajuWj0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RCLo9MWjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JMc8rAHsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z3oSBvHQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DCdSxvK1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8uMymyxsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1wqMzVW0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TOtMuqhvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R7f1OeM5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pp5f8vFAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CZNxXRNZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VbdDjiN5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CDcXMtlmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q7N4ljsSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w7aQVY4RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u8AISDXzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CZYPe26LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sKpFPM2tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yC9cj83CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ubac03YnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eJIHI5OEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3edydvwdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4wi1fYrUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6gz7B71jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dUph18GnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kPgJqmPhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func snPYCXa4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ShmFoC28Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5KwP6oTsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c5nZhn3IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jeShUQPaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xltGAnmVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZJbYk8HOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j2aHMKvhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NVx1MhgTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jBmRpIAfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ftIOyZPDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xj1ypfdcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hs94gGvlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NWlnjkNHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B3na9paNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0Mo65wC8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6vkNBSdiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func umTrhaIVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CoNIJ4ejWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jH3iiit0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LsV2AbOAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X56iITOAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Erne1eQFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QUD9QOAFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xQ95N47EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w5SPZgRaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f2Iu5VwtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nqfLcfjaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0O5a5yWMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zs8MV4bVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UiFSewzBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fwv0c0JjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZLqv9VAkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3XUVZL9FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CXnNs858Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 25KWIsQ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I0xhnYSSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fxSCpumeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WUrX0qAhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sh4DbRuOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o7pyhoVXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MDLVYdltWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2W9uUzerWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TvI3pHCrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UasvOh52Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zujuK4QbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ALoR87XCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C06Cht09Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Kfx7oqMzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B4Fubi0PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VkZmmyUwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FULMirZoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mYot1HEcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wFbkGnQHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func djhKTZP8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KUZxBz9iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LS70wRRWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JOeuy2KlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8VV38FL9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2q7L4bAFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IL6nrxngWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Em9LYKHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PWMxCMrIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n70xmVHjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hkwa5NNFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b6RznRiDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KbztM6j8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cznC6I13Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fXHVnuryWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lcsKbKXtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XBXCYI12Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bIGaZkw7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XRIeya12Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sQwaX4POWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NGe6DhqBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tepAlJSdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xyHCY2skWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rfiRrf9ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func du7SiNBdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sOxW81m2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0imsJUfYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jsrKH3YZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NNK0aqbsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QCrDsDPvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gyfasHekWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PYZpZMbhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R80tQNsyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OevoF9yGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Axe33rDlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QGpmcsDfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LP0s82tlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RCjycnf9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bqFQpqEtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mxv2imKLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GbJsedh6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QCQVzxrGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 12pfIWURWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Guxkok2MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4oW6YMWaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HSdbAjPuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nLDPprJgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ocTqXZSUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vozTIQraWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uGVbHSaSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8OGdhBgZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JFsYcS1vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R85SjshHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cFcHhbxWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qYBClgaQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tnWtagKtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yDOJ47Q7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fbZifPoyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C7x5cwRHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PfHmFxPeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func degwfgWbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zAhRFyejWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 05V7GLTsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5x4tGCLVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sLEGLo19Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F6FarqOsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0Hq1PBy2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SD8ye0w6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KONgs5vZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qLrv8zPpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4SfPGuTKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Aaji9ukUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6l4E2XYsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VFDVgOZhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xIqIrsKjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZKYdPXg8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W8vG5PoKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I4ZOJ2tTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3RDAQHBrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q0a6w3ERWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5IzOTWxDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yb0kEu6aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bKEIwnYvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qZ7RtcNqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vz5bRczZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ByTmg35yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c3lVaKf0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n86fDes9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SsucJnXNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Cg11ukXkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GQGT4rEEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hkFAptHxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mRWFIw0FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9ajOCvCTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w8WhWUXqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HdvE9L2MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EoQnpSWaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lyxnkFRHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1E0XUuwqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 03hs5tNzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FP7hxBTEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 28wtD1gWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lOgxtBBcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ROZafBQ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TdQ0dYBFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h4gIDZVHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IE1JV1qMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YzbpWSwTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func goDdHH6fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZJ7EbTnEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xraiSK2RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jSXLE6dmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DvSOYYPsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8IOYyhUVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xPtayZrsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sQcQAWPtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yj0wCSjSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pFt4d0ayWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h6EWwq7fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ER7Z6TBKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7e3P8JMlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8TZAJdIDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F6osptblWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eBX04xCpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d3VsqaWsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xQ5VSX89Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7pl82qarWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 61w889FAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LE2RGFfDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4FlmgxXPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jiXGBHmNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wFVZYd6HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e3jhW0LRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AyOsINy6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BOamgRSKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1GJNVjj6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6rGX3nF9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w3KvJFp6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func geA4C1ZWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HuOyS9cnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GJiNdkVhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y8U42hbWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F1bsfirXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NYwcr39iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iC27TFg6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NbV0iWjyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6VnZTeuTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vQyJz4IeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OBrAQbgdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gS60SEM7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func molmhB1ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jaaKA09yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pnRiM9zUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i3Z77abOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6nfo861TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rEfsB6GbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xs2nQthuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o5iIpKSLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tc0A3OpoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vpJupWYNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 891emvfiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BGlSjvZIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r8mUUIMRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sun9LrbTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iHqS5BjyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0maB6uOeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W9PtuhROWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VXmFHOweWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h8gkirftWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qalxxESwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Cz3c8lMMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3hF0vhc3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jVPpyeVSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PV2MNVFOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z49r2uFNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PsgLUo8lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kAye33YIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func czmpU8xKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FwGEXmzmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oxUIfYdxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UE5IOpyDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3BVgBXVdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o48pUlgoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jlkvKmbcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2H5VaFdnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 70KXT4hLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ev4U9nQZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6J5bypOqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QZV6C3LsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rpAXgFN6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mlLw10lRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Efju4CIlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nzQxcKC1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oAR5KwHxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WyVvj634Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func buB4Hz2nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zlT74MwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dAiQ1ZUuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XHNGDcOeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5zbsLUkXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ntub9bp2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g46PGgEAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eV7vMFYeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VTf0irBiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 008HlfJdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3yFHS18gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o3zbvMKBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YAT7a0ZGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uXlN4dAWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A0W7DDUQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rfu9pYHIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zf4iPYG3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sMkEag8hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HkLLBMeGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JhUPxedCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dVkVlzjLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0IlqU6ztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1WkTf1zkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jVHIY2ujWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IOKJa9OxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DadbcrFhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pefyv64HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iSXSmWsAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func llf5I0RhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1EfsANeRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HKcwGyN6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k8uHZT7FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LfvrCIfwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n5Bfo3Q2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xqnRwjoiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sjyyCbXGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wBDOIkJYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F0D60WytWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cccIAmlRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sfaLK2v3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 19loPu2mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uFkYrOKiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kUXzHfvKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 58eTcWiyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wIL73IomWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func adqTz1w4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vAjVqLhCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pv76ggTFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func to8k9RCMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iLt3M4bFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T4IJlbfiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LmIFmUUUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O6X883utWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JY6JI6TDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ut0AMmluWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func peTFTkTpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bOUaY8OFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7HCPxBUtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JNn8ydwlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0lKISP7pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZyP7jLqzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y0EtQBzYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hiSsTs3WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CAYn74DfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eBz9OiHzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U3BbI80bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func biv4jObzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bVMzD4IYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vceZ60ljWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QmvSn1KeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WLIFWpTjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yLSYb1cEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SCYCizhRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KpZcl0uwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L2M6xXnRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sczOVbMlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W2YZF4lGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CoUx9fLgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NJ045FmaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6om1UnTKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9nWIwrKDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NenGWBbOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HeQ3qJDTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yoFMM8jZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qIozR0CJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T7PEGcFFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func udi5iJWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4hdXUeG7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o5gRRp5SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yVK8Z4p6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rUvuqsZwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2sfC5MYDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lehCUnVyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3yBnYHFMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JnqUUxADWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QBouCQUwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func okzaYBUGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EHJlk6KrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l8PMd8qmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sUxt2UXrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5ELeWMAQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HXmMLnZ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y9quV6sTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n4sYVtXLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OcJMF7eYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v2c6ZKgsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sjVCuu8LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JgqmLn2QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wiL9l718Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ef4W4z2vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IDB4KjO9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3lcpxYpKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WvGbq7nqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9gtTnsOyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L4W2zqjUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zluZhw0HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BTlCBR4sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cdxeHWGtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GhmTleKCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ghziL3mlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eCw6ZTD7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8HmzPH59Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3GwtQjF5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JSTfPMYRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WAroniZJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4CcFP4twWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B14LPGXJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IcM6Z9uWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bdupS36rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4DiFUuwjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L5NASCASWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JVUaklW3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jOmQBxxxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SIYNFdDWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oFLzpgTFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bUrJhv3zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P24LzIKeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func apjOSotDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4Fz9gXmaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S1Sur5SAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 90HJtGAHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gA8zG9F5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ATjIlUGrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oXrWmCntWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cGVuqfjAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ts2428nFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func okdmeWSgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DhCfB4ZQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7O9a5uysWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2VvkiLBCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5r0LPa7OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NHsWjzsyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sEIMDF7CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0TQ7IbrSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kpYZrqaJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jVhYnASBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func flRHAzjMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IETZkdjvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func egL6Tl80Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RaE55EtYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rWaHiVVuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kq8Ap6ZyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6QBLBfijWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4rRgZPpkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bje4JGF9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OpswhwB3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Enm4X9ZVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dyRsWUFyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hwmS69uMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aSKrG1mtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kP6Q46IDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uP3gvMJCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XK3iI1S6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9VRSCCMxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rSd3ysV8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bgND7JrkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UuO5E9jQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rnLpmHthWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 81zcjryUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func whJmHAjNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 89v7JSsbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rOwPNL37Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VDVQVz6aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QQBgOtvNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cAfILfGFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ay8ULxmgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lFTLBqgpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HftnzwqCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ChOnBMh6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ra3g5VzWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OvtpDPZ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KoNFylF2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2ErlpKBKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xuGUF5XdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vpYtAtlGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kR1SkesdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z36yQ651Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DFAokJc0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nk9qkyanWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nEV4kRblWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6ll8u2zTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qqvUfxEqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func og02pYKpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yO2NAb5mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qlbn3x4cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bsKNbpueWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2pNcSvMfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gxtiLc71Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func weJvZ3LGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D1EjAoVDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iBiTe5gDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HX9F64D4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R5k36qV5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nQelY4v1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H8AivGvBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M7QVtEwfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yWQVThrGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oFWUAfBOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KHNDBDMkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OCoYTo2hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VYKSshPbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V3gWweDPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 26Os8ZpvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YGJMpDatWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iqlhcUnLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 98Tj8ggPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lmK5KjNjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hnr1ZI0GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0lfrO20QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RorAJo00Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RfJURHKcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cTeNyL3dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z5aPURrIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uOxU7ApbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uVsMejoQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3EMHUxrpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func REZdiUjqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func svmgp9YgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6gt9lGe8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bWDLACEGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QHsaSpsmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dC9A4qcQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 02j2wv6XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gOWGr346Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lgyWdbXhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KMx7cL6pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gzXyy53tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zNYeKG07Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fs6YM7XSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BoddJ3IhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0qM0K0tWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oSgiLJ3BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2lcwuqQWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u8jZ7628Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vMk6EQwEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oCFkhIAsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fJ8JjDunWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L7JbSRCuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K1lqI0PoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ulGAI6SFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IWV6yji9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S0ORgI1LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I4EW9bfpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ko0GFXyAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fWJRcf8SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O362iSp1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kC8ksFlFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H917JI3kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mmUIBfp7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OCUtGEVdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pnKykq3IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gFV7SKzbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 338299ggWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TbY2Sj29Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HCXURpOuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tDP4HW5zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WvU64FlHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1FyLJnOnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func McrBL3ACWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2frqTcIbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rapJzqTHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 48M8JIuQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4FhmV9CrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5ZaDjPlxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qzESwaZ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9jEoydN7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QUNjaI8GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SCKWgm0XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nS0GyJdeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EhYDfaNFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func voWywkWGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mQT7c577Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func caCY6JbPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 62TewPrtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jm123tYrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func czSrh4mOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wMbETZArWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CEvmzz7xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zBi51hoWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qpAnVsecWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pMp1jtIyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T8Q57TtjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6ecxrLS8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iZbQeNHjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NSqz4FcFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d3bBk7PPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C1Du5LxsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5OZtry87Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gxORGw1SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KGoHlUohWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nJD0kNYKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yZTd8SIwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6KG9nSjhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wLySidlTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pHKiyCcIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func omYW7uxeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XfgffNGLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pbot2fB5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y4lbfpdpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7FpleifeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zkHLZ8dTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gLIxXY5dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jj6DQM6sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QAtya25wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cG33DWBFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hKiBSDFlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MFi4LYMzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KtQRWp0SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LzzeLvolWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gOUYlYOEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pWDnLyaVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C0G2kiwNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xa3aNgfeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DwPnjHDuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZiYxcekcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y7fDtZy6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pLUHN8L4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CRBQuSK9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XYOALDfkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bpiX58yYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mtp4X0qXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ULiOTAX6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oppp3WnbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2JSJV6G1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lRAiNzi0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func es52ls3tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T4OzUwAsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NYNCrzu0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iAk183g9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lSVtsxiOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0HO2zesbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VQThVTuAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XpZ97fQWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9P69HUn6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zAxLs5PnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5r8xl90bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V2dAbf8vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zOIE8CkBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 031ScED7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1L6sNCtHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XRvoDgZeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C9OJeVbwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q5nYouL7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZF3oQBMSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6hswMaaHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ERukgNTqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BX5oJcNBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ibpRrJpYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eg6zvF7aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 50fao6pEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O30Hca3hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4uLfN1FRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 58iVMrhoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2LxRHIMoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DPqJmzztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jMKBhc8HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A0gKLhsqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9nJlacovWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YXg4ozD2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eKgsaI6zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lwPtoXq3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VW4X3fdxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M7SNBEfLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Oav5dYIVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func us0N6FF5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6C9F7QbXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gTUjmy9QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aYTAeoR2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cDHinUddWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WQ8EDGU8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0ywMNVwtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UUyE0PIxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oGEuEfr4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FjWroTonWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WwOPakQCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CXMpxVV7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YCaimWOZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PY1M1pE1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MetqBETUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func axQgqpXdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NJFt17jFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qc1jKInBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HyElMOTBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bXhn8f6uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YTlrGMryWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gznLBg5OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YnBRKuKBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YSa3q1vQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5BhBRvinWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yJ0ZFnc4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9X6nFnttWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JlcGQAgwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SEuj8oVBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oGMAkq74Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oUdJIJmeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t9uvhEsaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CExUsRPuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UUeOiij7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KuGzAN4NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bhEvloPiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ObXYK6hHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kiBzh325Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jxw1FbM2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7eG6ktKaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sN2gwKGOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XpKQxAq2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HFw0mf1nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e5Y8HK1aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RSMSA5x5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B2JJfA1JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EbsnJ4owWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XHSgO7luWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5PvLviYAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KYE37yrPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j4JJoK9NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6UjxAkTMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vpVduaPsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GtyjELMrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Cou1RqhgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RghM5JntWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bScwbeFCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XriuYOULWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ESJ7OkzFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8uXfiM3XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4wo7NxT3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HVbR8gPcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H7OibaJEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9BIgcV1xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fut3hb2OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GNVbLr4zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2EDNHq10Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U0H5msEtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SZ0uWml5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7aXEdKFTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g8wT0fPLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9UNkGOoXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rZrRtkLxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0M090wBDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bPIE6G4vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z6AXXI1BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BndMCIldWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i7wBS9mzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5UtjoEU6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3aOMwm17Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vdnx1U0TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TvUrvsZlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5O9eIJjmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xzPSSIDVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a6lXkFI5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2s3CiXzlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qsvLg4KYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sIz4bdV7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PMVGH4luWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KhRDqj0gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NGznogSfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xpJTGY9zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R2iefEeEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YhOjO245Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eObbqQ6ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func crk5rAkLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yiFRZ4mFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yvRuuINuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QKLY1chUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mmhcbnJsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2u1s489gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NEq3XVpYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3kaHG561Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WWZZulxRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func On4Sohz5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7TUw78ZWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bXX7rPESWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QYpxATh4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hce7SksPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d9p1yFs3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P088yIopWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func trg8N4o0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func by2EZwGiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xS7d9YzaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PAFY6M89Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8JcE4jEXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qQW72DbJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yUwakZ9zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u1TAL4f0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d1ZkX5vnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HKiED0IRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UnCV42XZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mLNhvq4jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EzEeOsVwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TpV4UFxjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QF7qOg5jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HxmBUvCEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lrBiyBMVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aXEDej7TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wVBK59gXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vtUndaf8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wOubQPRlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W9g18CZJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mqgJP6ONWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UVliz4I2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KDV8mfdFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AjfAz4wDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l9TXzC6rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3sI1R1kFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vFsJwOiBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qvn6C4usWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ETavOVzeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fttaQNG1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KXBXs25mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cWd5q7s2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8fXxyqGvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1akZUDKyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8aPSmy96Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s9sen2FPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yz1QFrmTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0r93ObVJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hxgES3shWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9NsrmqsXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o3yGu5z0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5KsHg2NGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HNt7NKn5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1RyBjjeiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 64erF3O8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8A81yWFCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jTMNAwvUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ft46s1EuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cX3dpa26Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OyUAQDcSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 55iAXjjoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xSZQqK67Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Chyv9YrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fCrxJd6rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DgHkifHuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6xxSxsanWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dI6Jx7YSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k30nXnOyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jcMXeBLKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ti5JM9c5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HMsHDfaCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dPlXGg1xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YYI4w8ZjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U8BaSFDGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g53Cs8EoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x4Y3KsM1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5leK8KD8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9x8EUH3BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rK4hmPE7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HNDfXlrwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mVypshJXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GZb2zfsKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZQXnqsmtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hyC62IiOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FevqFdZiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p9dzHFk0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s0KyING6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xWFyqChTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5az5dK01Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s0rj7866Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h8I0J3K4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lx6MBhEAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YacgXB0mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6vdlhZV9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NnzGuMP2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QRcf1EtYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Q64ef9fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WGfNnpRdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v0czV9s2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iX6a5pEEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sI7dIsmaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JCc43oj0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TjdY2NmPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aEoiaPqWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DuWQnJ8IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ujOdhzJ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rk3rVnCVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iAirMn68Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lx6u9sfaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N39UVfx5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kOzqBeg9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YkcTxoPrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qTwnLrn9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WXrG4u93Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pj8xb9GjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KtJei2ltWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PNLUzMEkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8RnlDTzDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func STRQi5v0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gQ4o8wl6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bnwUQBCPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IgZWqh1qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ntTDFgtpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tMhl4OK2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WROjgRm7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EmdoMMjtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JxMmmZFvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sATr9PumWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UdezuirOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B5MseQPHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zC9yM8laWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O1oG5OdCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q020er3BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func swWYNHSWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J70BmVkJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EX3QPNS1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BgDJpQMFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DLG7vehmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4zz1eSFmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q2XoRCgiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rLHIdkjqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X3Un2smXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rsXNKuhdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7mxQr2hBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S6KjJGN7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TokrhVZmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jpJL3rT6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2ThpUB2aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B2ezPRCLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gPXfawKYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0vn3CjvVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sm3E6ikwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XuPhVO8fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VvbyRBz8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rPkOqZllWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PbZFnGhMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 74c4DupRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func augO4BJsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vQAC4zvTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9TzyVV6EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func onjCWz91Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JIufK88qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BgoPeSUFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WALds3VfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GO3H8lZYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GvKGptNjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func um3bP0wHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WkCwNbChWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func amn0Jh2ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m1KPdXoqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s8vFOw4EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rcX8uGKIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LUJsH1V6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2fjGfxrfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OfodIQouWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ggh6Omd1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mEx8Hk7UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rpm3qL7gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FOw3vUoRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wJNXvolvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VkNdnyhbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RragTu97Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zYOb0XJBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JRqLd0ZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func avE5Jg5iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zTLzhSJ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func czFNAk2dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gA9YY8hjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nM7SfUHoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PYaARJ3tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dq7bVW5AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fMJelfyFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zhRcM0KfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VzBU0FcNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dS2D6wshWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 37ZU6m6tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sUBE39yvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HDiA73c0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CL1badk9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AFlVZ8JWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hoP8wrEFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qEREJq2TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DyWsdwCEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JDzMYYdqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UhErOBBCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E4ScRbRpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JiY7dP2YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PK5uVQ41Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E3YzQ1wxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nlwZyh1YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fiAiBJ0yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xZWSC9zHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vOvmWxeqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iWdlWY9GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oqCvp56fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5wUJwWLJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gn0VuWd1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Po2TWB0eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aAjEIWepWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mfAAdWOQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HxXPXIq5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wmo9gHElWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 40s6ZnAFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zPbyE2ZIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U9bw1IOeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zrC2apJSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GXbEUmgJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v3eOnOh8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 42Pt9AewWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LRlZrWHxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K7taupUcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WwBsT3hWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NxgL8ZUMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B7jXxRu2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uHr9B6URWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PeyGa7jgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ekTKZj8FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Zm4nfAbwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UwP2wIblWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JKLqkyftWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func td2OCYUdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uzQEo86mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c3nFB48rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nUGepuy9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7fmgB5X8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GxSxyJmgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LU3e1AVwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6V2IZ9dxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zlRnvBerWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5JOcuxcLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4Qe40mq9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QVHyr6DoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yxY0PjPlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1SSPGp6DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DpM3rQlfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MajsWveKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1ApIZdG6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qJ1rSIYmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ObQ6ZbhdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lRITulaoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func px9Tcl1iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9njsXDm9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TDcSnXClWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vvEGhUZ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func heVn2uEDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YKfizgb3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nktogCBwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YT4mj1IUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n6Ml0M53Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SY6Q77xcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yBb5c8jGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W0cHDW2OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oB9uCQ0OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wy1HkohGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W1Lz78OUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k5zXDpkPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XseMU3QyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TcDif6ocWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0bkNNCTYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func irqrIZiAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vB53f7JDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fPxCclfiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SeqidtmwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FBxwEwPyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fhbXzIEiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TBZ4UzPlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zVIJYyapWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hfCkbCuwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6jwIe8drWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SXV1SduQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M6dEmvNuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5EdWld4gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sQYPuCDJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QvTjmtMsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EDAvT70uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ui9cnCQEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M3Tml4sYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pmwjtXpMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SEn07tyoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LHR1lBhJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1E9Pl2T6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ImqxQqnBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SrNK6076Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o9JTiXL2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S45YuI2iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func znX49F1hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y2VspQ2pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gXo4AtvCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mmkvj9v0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qsAfvlgwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NFXAsESBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jrAJaEcOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ePxyxmsWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OIxok04mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oG6XHHOEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0unfG6E2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TZXnvKSdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mnTKVFVXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JqJZnQPEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g3htT7vyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DYY8VBdTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YTwf9IDfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JvgdX7LnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DXjkHH4oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WYoI2O6FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KMaeDoBrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I7sL9S3NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gsa5KUuVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U3KPuzXTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zk2xfeN7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e6qXzTZMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oHrFfI27Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 311xHqkmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pt2ScaZ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YBoZkUq6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n4l15tadWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pUQtiGlfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wELcZZNQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pQEMiOCqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9tJhOo14Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n9vWQ42DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rzjwtqimWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 89lFe4Y0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9GQilf4BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pqDzmSGkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qmzu5ajmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func heMoR8h8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bDhZnNMbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sqEmLLf6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G3i5BGdJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TxmTK3wAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9vumZnKAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ND5csUxsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VKLLfYGOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 214yVSxNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tlSTjt83Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lirQR7W4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pju1Z2Q9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uhLiEqw7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ccn24NZtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MROnqzwHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kiydIQvsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fh4XuOUQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3S5n4L2JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dHRQ6znTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ym13sVn9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lW3tWpQVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func US72NVjOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eC4cdkV7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rJobC0OAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IA0J4udUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1lULXSSFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EQ1ZCV98Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gErVzWnuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aPDPegCLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FQRD81BoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func au7K61NXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5C5LhNBHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hcneyCTtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lB8Mn6mTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nm4sedxCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6IJ5LR9EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tusL2DNuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WirJIEurWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EF2I5xtpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fPIT3mSoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hID7hBxiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8klnvADeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tvcFiz8wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4jKtcz7HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func shleCGYGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KOwo6oatWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AbZXJFseWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aTqbPG0QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V2HeC8rVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HSePnaEnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OwwGQYhrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Np8vClAiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xNgxdeSSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iuMVglFyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zSTT4KXzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ANvwho6zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fp6Oaf0fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FkJQMnYjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YaZ1r1TsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rtmEXZXAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pj0B73kAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5YkP0zDwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c8A9YQW9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pBrCjM6RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jn1TriYRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zsy1J69aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VJsnjI78Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bZXRplZGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eBvE1qHHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5ftoDWm6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zNjpC6TBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9fVfJIDVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TOkFuvrrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ocYCAtV3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TH2K2yOzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func miFnkQy3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ek9j5pfnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ibSjfNEaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F3Dpjlb8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P2bGrNVxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PmrOQj2VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2Kcb4RhSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Bi7nLBZrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q7eI2eWPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wavh79S5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dnH3ng06Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AoyBSiAyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rrew0dXTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 511YumIVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GUDGmymLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mLodUEGxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H5QrlCDxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func anKU338jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o0VuAeakWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i4zY6HhXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OucRnB3sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iNVV9aJqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ynhk9Yq2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SAaGBghFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cpwZBWHqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gVrLuTc8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ah5TQE4eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aewzVaJ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wBmbo83lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c00IiIXCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E01vyZWYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yajc07STWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1Gd1Mf0kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TJz7PSTCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h9SQeD28Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9Z92OfUcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WzTqep5SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8RUijsuSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func joOMIEDrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0o0iEmluWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h2gQ0H6aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u4360qrXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RGMNsXXGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VwmPek4bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TMV3yBJCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eBFze1qSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tqDTTcU1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OWu2mljAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YdzEb5CuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oO0oGagPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func El7Oa5EXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ezSmucgbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 62IggHi6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lcGRbiiKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uMnx6dcRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8GyvZWWMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qLLOKo0JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RMe3ZgKAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7V82kazUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 62Uo9QZ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NswFF3HzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cbT8VnFcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k8daLFkoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Al5qPj8HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hiJXPgdxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pcT2QDumWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WbhwKfEvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func irk6nyn0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Iwkk7g0QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0BlSoiLtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wll9CQ5SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X3dALraXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SaA9RnX3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5BKubaQzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zosIM99lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m5oaUBE6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UEcFqsjTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TC7k8TzyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gEF78DtrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fDTec129Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WMZZ7eZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nYXn3TonWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nduhmkcQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lIBunorCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YSDd4eN9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XTaHBr0YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QbrKLTDfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xpRLATuKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PyaoiY2hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lWJCVTYvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ASHl8WtTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7eviJm1kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6YBv7XBcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func snrImwLRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f3SrzSfQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9sxq3XbqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m4mhg3o2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k48rAiugWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nJLdD9zZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KKww8PhHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m9OWNXmCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ARC3mbwOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JlbmvfGUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GmVVOmrCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func keRCbsUSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C4Xkjl9TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sMcNQ1ZnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Cup1IqleWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jkuhqFZqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h7uKhogpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y8B24rg9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MM3URAehWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p2I7VQtIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XHa9g7yfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qlwoIpRDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ksAOVWQ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0u9m8jusWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aSwUR6ijWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sOwwmeIAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zrcWP83eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ReFJGeRhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func knIBjU3nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 92T2DW6UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QL2fmGAwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t6kbcb13Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WztX1tU9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qv2P9llsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hY1EB6p7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PQocW7WqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mPlfacTtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cWpZsuZ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KcE8NHgOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 58Z4vOLDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xKrx6kneWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func szQ9CZtrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 65iLb5olWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ctvTbMS5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YCUtmKC7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9F4Ppgq6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YrqJ9PIQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D5yRWzKSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v1ZadfHPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LVv1T8FmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NpkCaoaGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gUN9AjHCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OTA41ekGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ekPwpKMJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fgAwWFZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7EiZUxThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gcv0OFnXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lMOQWpAiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aH7eKxQ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bxxU2Ui6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mEJJKAAiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ff1b0A8eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1QISTUtDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jDpDMi1aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PkLZ8FtHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bzgLmDJ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kfEGBufqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CJ2nlISRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PB2Du5mrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gr5IVH0TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func okxXWOO1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ntwt7Ji1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p4xAE3oWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3HLzLzeVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JwM2afkKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vRDhqhhMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Zsvtr7BTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xx2q1tXzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ci1YRI3mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G36NGO0dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CExrY3HtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f2lTqV43Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gjJZ8BYFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fy55E3mgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VR2x9pF7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wIQdQ6w8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fzgWO9kYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JOE0UhE8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DcJMh0uuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 326Yrp0RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9FHQH1UWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LRRwrhfZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YwsbgrzDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kRX8rLRnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v6G2b5VmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SQU44ov6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Agp3mXnzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lh1YZ4y0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZSAy4YtlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

