## 压缩级别level

      #   gzip.HuffmanOnly			  -2	仅使用 Huffman 编码，不做 LZ77 匹配，速度极快但压缩率低
      #   gzip.DefaultCompression	  -1	默认压缩级别（推荐，平衡速度与压缩率）（对应的实际级别是 6）
      #   gzip.NoCompression		  0		不压缩（仅封装 gzip 格式，无实际压缩）
      #   gzip.BestSpeed			  1		最快速度，压缩率最低
      #   2~8						  2~8   中间级别，数字越大压缩率越高、速度越慢
      #   gzip.BestCompression		  9		最高压缩率，速度最慢

## pgzip - 并发压缩、解压

- [github](https://github.com/klauspost/pgzip)
- [「GoCN酷Go推荐」高性能 gzip 压缩工具 pgzip](https://mp.weixin.qq.com/s/C95QcppVZ108CMTEhYRzzw)
