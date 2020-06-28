# 学习笔记

---
## 知识点分析

### 递归： 
**要点：**    
**1.找到重复的子问题，把重复的子问题**  
**2.利用最小的重复子问题架构，写出相关的递归的函数**


### 分治、回溯： 
**要点：**  
**1.找到重复的子问题，把重复的子问题**  
**2.利用最小的重复子问题架构，写出相关的递归的函数**  
**3.把大问题拆分为小问题，然后最后在对每个小问题进行迭代**  
**4.最后把所有小问题的结果作为结果返回**

---
## 模板

### **<font color='blure'>递归模板</font>**

```go
//golang 
func recursion(level,pra1,pra2..){
    //recursion termination
    if level>MaxLevl{
        processResult
        return
    }

    //process logic in current level
    process(level,data..)

    //drill down
    recursion(level+1,pra1,...)

    //reverse the current status if needed 
}
```
### **<font color='blure'>分治模板</font>**

```go
//golang 
func divide_conquer(level,pra1,pra2..){
    //divide_conquer termination
    if level>MaxLevl{
        processResult
        return
    }

    //prepare data 
    data = prepare_data(problem) 
    subproblems = split_problem(problem, data)

    //conquer subproblems 
    subresult1 = divide_conquer(subproblems[0], p1, ...) 
    subresult2 = divide_conquer(subproblems[1], p1, ...) 
    subresult3 = divide_conquer(subproblems[2], p1, ...) 

    // process and generate the final result 
    result = process_result(subresult1, subresult2, subresult3, …)

    //reverse the current status if needed 
}
```



---
## 题目分析

### 46. 全排列

**要点：**  
**1.包含的是唯一不重复的数字**  
**2.需要完全排列**

#### 思路：
>可以分成不同的2个子数组进行排序  
>可以利用分治的思想来不断进行迭代  
>然后把每一次两个子数组的答案不断整合在一起  
><font color=#00ffff>时间：O(n^n),空间：O（n^n）</font>

### 50.pow(x,n)

**要点：**  
**1.不可使用库**  
**2.最小子问题**  
**3.考虑中间变量的使用**  
**4.需要考虑迭代终止的条件**  

#### 思路：
>因为2^2*2^2=2^4   
>次方的计算其实就是使用每次降一半的想法，并且把这个一半的数进行储存和运用，可以大大减少运算的耗时，提高程序的效率  
>需要注意迭代的终止条件，以及n是负数的情况  
><font color=#00ffff>时间:O(logN）,空间:O(logN)</font>


### 77.组合

**要点：**  
**1.重复性问题**  
**2.最小子问题**  
**3.各个子问题的输入输出形式固定**  
**4.应该考虑迭代的方法进行解决**  

#### 思路：
>每次输出数组的大小是固定的，所以可以套用迭代的模板  
>需要注意迭代的终止条件  
>暴力解法解出答案   
><font color=#00ffff>时间:O(n^n）,空间:O(n^n)</font>

---

## <font color=yellow>心得体会</font>

这周的内容难于理解，就像二叉树那周的内容一样。  
不过我发现，在学习过程中，上周可能我不太能理解二叉树。  
但是这周当我回过去再看一遍的话，我又有了新的心得体会。  
**这大概就是超哥所说的，五毒神掌吧。**  
很多时候，脑子说：我会了。手就会说：不，你不会。  
所以学习讲究的是一个持续性的过程。  
这大概就是我这一个月过来的所拥有的感受。 
人是一定会遇到很多困难的，但重要的是，人是可以选择直面还是逃避。  
我之前可能一直在逃避，但是这一次，我想变得更强。  
我不会敷衍我自己，我会努力去学，去努力。  

<font size=4 color=grey>*人生的很多路，是需要靠自己的毅力走出来的。  
多给自己一些自信，只要坚持，把每一步走好，自己的路一定就可以走出来。*</font> 





