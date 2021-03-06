学习笔记

---
## 知识点分析

### 动态规划： 
**要点：**    
**1.可以拆分为重复子问题**  
**2.找到最优子结构**  
**3.适当进行剪枝**   

---

### **感触**  
- 找到最近最简方法，拆成可重复的子问题
- 抵制人肉递归
- 简化复杂问题，通过敲碎成小的子问题
- 把复杂度降低很多
---
## 题目分析

### 53.最大子序和

**要点：**  
**1.动态规划就行递归**  
**2.找到合适的条件进行剪枝**  
**3.找到最简子问题进行迭代**

#### 思路：
> 当前值小于0的话就不会得到最大值，可以进行剪枝    
> 对原数组进行拷贝，然后在新的组数上进行迭代修改，不影响原数组的内容
><font color=#00ffff>时间：O(N),空间：O(N)</font>

### 120.三角形最小路径和

**要点：**  
**1.动态规划**  
**2.类似于机器人走到棋盘问题**  
**3.可以使用一维来储存数据进行降为**  

#### 思路：
>自底向上进行思考，反向思考  
>下一层中的两个与上一层中的一个相关联  
>取出最每一层的最小值再往上一层迭代
><font color=#00ffff>时间：O(N),空间：O(N)</font>