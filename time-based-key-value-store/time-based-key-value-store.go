// https://leetcode.com/problems/time-based-key-value-store/
// 401 ms 69.2 MB
// Your runtime beats 100.00 % of golang submissions.
// Your memory usage beats 25.91 % of golang submissions.

package main

type TimeMap struct {
    bucket    []*Node
    log2Size    uint
}

type Val struct{
    Value     string
    Timestamp int
}

type Node struct{
    Key       string
    Next      *Node
    Values    []Val
}

func Constructor() TimeMap {
    return TimeMap{
        bucket: make([]*Node, 128),
        log2Size: 7, // log2(len(bucket))
    }
}

func (this *TimeMap) Set(key string, value string, timestamp int)  {
    h := hash(key) & uint32(this.log2Size-1)
    for current := this.bucket[h]; current != nil; current = current.Next{
        if current.Key == key{
            current.Values = append(current.Values, Val{
                Value: value,
                Timestamp: timestamp,
            })
            return
        }
    }
    this.bucket[h] = &Node{
        Key:   key,
        Next: this.bucket[h],
        Values: []Val{
            Val{
                Value: value,
                Timestamp: timestamp,
            },
        },
    }
}

func (this *TimeMap) Get(key string, timestamp int) string {
    h := hash(key) & uint32(this.log2Size-1)
    for current := this.bucket[h]; current != nil; current = current.Next{
        if current.Key == key{
            if current.Values[0].Timestamp > timestamp{
                return ""
            }
            index := search(current.Values, timestamp)
            return current.Values[index].Value
        }
    }
    return ""
}

func hash(s string) uint32{
    h := fnv.New32a()
    h.Write([]byte(s))
    return h.Sum32()
}

// return the index i, where list[i].timestamp <= n
func search(list []Val, n int) int{
    l := 0
    r := len(list)
    var mid int
    for l+1 < r{
        mid = (l+r)>>1
        if list[mid].Timestamp > n{
            r = mid
        }else if list[mid].Timestamp < n{
            l = mid
        }else{
            return mid
        }
    }
    return l
}

// Another way
// 452 ms 68.3 MB
// type TimeMap map[string][]Val

// type Val struct{
//     Value     string
//     Timestamp int
// }

// func Constructor() TimeMap {
//     return TimeMap{}
// }

// func (this *TimeMap) Set(key string, value string, timestamp int)  {
//     (*this)[key] = append((*this)[key], Val{
//         Value: value,
//         Timestamp: timestamp,
//     })
// }

// func (this *TimeMap) Get(key string, timestamp int) string {
//     if list, ok := (*this)[key]; ok{
//         if list[0].Timestamp > timestamp{
//             return ""
//         }
//         return list[search(list, timestamp)].Value
//     }
    
//     return ""
// }

// // return the index i, where list[i].timestamp <= n
// func search(list []Val, n int) int{
//     l := 0
//     r := len(list)
//     var mid int
//     for l+1 < r{
//         mid = (l+r)>>1
//         if list[mid].Timestamp > n{
//             r = mid
//         }else if list[mid].Timestamp < n{
//             l = mid
//         }else{
//             return mid
//         }
//     }
//     return l
// }