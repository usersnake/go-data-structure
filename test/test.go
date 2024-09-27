package main

import(
	"fmt"
	"dataStructure/linkedList"
)

func main() {
    // 创建一个Buckets大小为5的哈希表
    hashTable := linkedList.NewHashTable(5)

    // 向哈希表中插入键值对
    hashTable.Insert(1, 100)
    hashTable.Insert(6, 200)
    hashTable.Insert(11, 300)
    hashTable.Insert(16, 100)
    hashTable.Insert(17, "hash")
    hashTable.Print()
    fmt.Println("\n")

	value, found := hashTable.Find(6)
    if found {
        fmt.Printf("找到的值为: %d\n", value)
    } else {
        fmt.Println("不存在")
    }
    value, found = hashTable.Find(10)
    if found {
        fmt.Printf("找到的值为: %d\n", value)
    } else {
        fmt.Println("不存在")
    }

    hashTable.Change(3, 300)
    hashTable.Print()

	if hashTable.Delete(6) {
        fmt.Println("删除成功")
    } else {
        fmt.Println("不存在")
    }
    hashTable.Print()

    if hashTable.Delete(10) {
        fmt.Println("删除成功")
    } else {
        fmt.Println("不存在")
    }
    hashTable.Print()
}