# IP Tree

Check if an IP is in the Set. use tree structure to represent a blocked ip list. less memory, faster speed.

### Install

```sh
go get github.com/DeronW/iptree
```

### Usage

```go
package main

import "github.com/DeronW/iptree"

func main(){
    set := ipset.New()
    ipset.Append("10.0.12.1")
    ipset.Append("192.168.0.1/24")
    ipset.Append("1.1.1.1-2.2.2.2")

    ipset.Has("10.0.12.1") // #=> true
    ipset.Has("192.168.0.8") // #=> true
    ipset.Has("1.2.3.4") // #=> true
    ipset.Has("::1") // #=> false

    ipset.Append("::1")
    ipset.Has("::1") // #=> true

    ipset.Values() // =>
    // ::1
    // 1.1.1.1-2.2.2.2
    // 10.0.12.1
    // 192.168.0/24
}
```


### progress

- [x] Parse

- [ ] New
- [ ] Node structure
- [ ] AppendIP
- [ ] RemoveIP
- [ ] AppendRange
- [ ] RemoveRange
- [ ] AppendCIDR
- [ ] RemoveCIDR
- [ ] Append
- [ ] Remove

- [ ] Count, counting how many IP are thre in the tree
- [ ] Has, give the result of if a IP is in the tree
- [ ] Contains, give the result of if a set is contained by the tree
- [ ] Difference, give the difference about a range and the tree

- [ ] testing
- [ ] benchmark