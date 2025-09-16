## 总结
(1) 建议使用 freecacheKit，原因:
* 支持TTL
* 协程安全访问  

(2) 用作: 本地缓存.

## LRU淘汰算法
LRU算法是一种缓存淘汰策略，它的全称是最近最少使用（Least Recently Used）。  
它的基本思想: 当缓存空间不足时，优先删除最久未被访问的数据，从而为新的数据腾出空间。LRU算法的核心数据结构是哈希链表，它是双向链表和哈希表的结合体，可以实现快速的查找、插入和删除操作。

## 第三方库
- [coocood/freecache(4.8k Star)](https://github.com/coocood/freecache)
- [hashicorp/golang-lru(3.9k Star)](https://github.com/hashicorp/golang-lru)
- lancet
- GoFrame
