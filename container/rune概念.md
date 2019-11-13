rune相当于go的char

1. 遍历字符串 - UTF-8字符
   []rune(s)
   
2. 使用utf8.RuneCountInString 获得正确的字符数量

3. len 获得字节长度

4. 遍历字符串 - 字节
   []byte(s) 
   