# kmp算法个人理解

> 这个算法研究了有三天吧，确实长见识了。翻阅了好多网上的文章，感觉还是有必要记录一下自己的理解。kmp算法的优势理解起来很容易，即利用了已经比较过的信息，i指针不需要回溯，
仅仅将子串向后滑动到一个合适位置，且该位置只和子串有关，与主串无关。
其算法时间复杂度O(m+n),比普通的模式匹配O(m*n)快得多，当主串与子串有很多部分匹配时，优势会更明显。
算法难点在于高效求解部分匹配表next数组。但是仔细观察代码，会发现求解部分匹配表和匹配主体算法很像，
其实生成部分匹配表就是依据上一份表去生成新表，所以算法其实是递归的。理解这一点，
 
## 名词解释

* 主串: 就是较长的一个字符串
* 子串: 子串可能是主串的真子集
* i: 匹配过程中主串当前指针
* j: 匹配过程中子串当前指针
* next: 基于模式串的部分匹配表数组
* 前缀: 除最后一个字符外，所有以字符串头部起始的子串，如ABCAB，前缀有A,AB,ABC,ABCA
* 后缀: 除第一个字符外，所有以字符串尾部结尾的子串，如ABCAB，后缀有BCAB,CAB,AB,B
* 部分匹配值: 字符串前缀后缀相同的最大子串长度值。如ABCAB，前后缀相同最大长度为2, AB


> 笔记本没电了，下次写0.0