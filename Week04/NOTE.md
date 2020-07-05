学习笔记

---
## 知识点分析

### 深度优先： 
**要点：**    
**1.深度优先就想成先遍历一棵树的最深的那一边**  
**2.重要的就是遍历完一个子树才遍历另一个子树**  
**3.而且这个算法在遍历的过程中不会等到这个函数调用完才去**   


### 广度优先： 
**要点：**  
**1.用队列实现**  
**2.广度优先是按层次遍历，可以记录当前层的层数**  
**3.等每一层的队列都清空了才会到下一层去（与dfs不同）**  
**4.类似水滴波纹散开**

---

**栈实现**

### 模板
1.递归写法：
```go
//matain a map
vistied=make(map[int]bool)
func dfs(node binarayTree,*visited map){
    //first check the node is valid or not
    if visited[node.Val]{
        return
    }
    //else will update the visited map
    visited[node.Val]=true
    //process the current node
    ...
    //next node 
    for _,nodeNext :=range(nodeChildren){
        if !visited[nodeNext.Val]{
            dfs(nodeNext,&visited)
        }
    }
}
```

2.非递归写法：
```go
vistied=make(map[int]bool)
func dfs(tree treeNode){
    //first check the node is valid or not
    if tree.root ==nil{
        return 
    }
    visited,stack:=[]int{},[]int{tree.root}
    for stack{
        node:=stack.pop()
        visited.add(node)
        process(node)
        nodes := generate_related_nodes(node)
        stack.push(nodes)
    }
    //process
    ...
}

```


## 广度优先

广度优先的话就是逐层遍历。
遍历完一层才会往下遍历新的一层。
**队列实现**
### 模板


```go
//matain a map
func bfs(graph tree,start,end int){
    vistied:=make(map[int]bool)
    queue:=append([]int{},start)

    //start loop
    for queue{
        curNode:=queue.pop()
        visited[curNode]=true

        nodesNextLvl:=generate_related_nodes(node)
        queue=append(queue,nodesNextLvl...)
    }
    //process
    ...
}
```


## 二分查找

二分查找需要满足特定的条件才可以使用
或者可以变化为可以二分查找的题目

### 需要满足的条件
1.目标函数单调性(单调递增或者递减)
2.存在上下界(bound)
3.能够索引（index accessible）

### 代码模板


```go
//matain a map
func binarydivide(target int, array []int)bool{
    left,right:=0,len(array)-1
    for left<=right{
        mid=(left+right)/2
        if array[mid]==target{
            return true// or break the loop
        }else if array[mid]<target{ 
            //target should on the right of the mid array
            left=mid+1
        }else{
            //target should on the left of the mid array
            right=mid-1
        }
    }
    return false
}
```



---
## 题目分析

### 69.平方根

**要点：**  
**1.y=x^2,在y轴的右侧单调递增**  
**2.有界**
**3.index，mid可以同时作为下标和值**

#### 思路：
>x*x就是一个值的平方，所以可以用x得到平方根  
>单调有界，适合使用二分查找  
><font color=#00ffff>时间：O(logN),空间：O(N)</font>

### 122.最大利润

**要点：**  
**1.贪心算法**  
**2.局部最优解等于全局最优解**

#### 思路：
>后一天比前一天多就是适合于卖出  
>单调有界，适合使用二分查找  
><font color=#00ffff>时间：O(N),空间：O(1)</font>


### 127.

**要点：**  
**1.贪心算法**  
**2.局部最优解等于全局最优解**

#### 思路：单词接龙
>单词的每一次迭代，会产生出一个新的分支，从开始的词到结尾的词
>典型的层序遍历，每一层就是一个替换词的深度  
>最后当遇到与结尾相同的词，就返回当前的深度  
><font color=#00ffff>时间：O(N^N),空间：O(2N)</font>

---
## <font color=yellow>心得体会</font>
没有人随随便便成功，只有不断进行磨炼。
我中途被很多东西打扰，我贪婪我懒惰我嗔怪。
但是我明白我很想要一件东西的时候，我会迸发出很大的潜力。
**做事的时候要一心一意，但是努力的话需要多线程努力。**
人生不是只为了一件事而活，尽量去体会不同的人生。
