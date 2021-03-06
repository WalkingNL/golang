## go language memory model

### 分配方法
1. 线性分配器（Sequential Allocator，Bump Allocator）
2. 空闲链表分配器（Free-List Allocator）
3. 分级分配
  · 线程缓存分配（Thread-Caching Malloc）.
  

Go语言的内存分配器会根据申请分配的内存大小选择不同的处理逻辑，运行时根据对象的大小将对象分成微对象、小对象和大对象三种：
  
类别   | 大小
----- | -----
微对象 | (0, 16B) 
小对象 | (16B, 32KB)
大对象 | (32KB, +∞)

多级缓存。go语言的内存分配器会将内存分配成不同的级别进行分别管理。go运行时分配器都会引入线程缓存（Thread Cache）、中心缓存（Central Cache）和页堆（Page Heap）三个组件分级管理内存，这三个组件之间的关系如下图所示

![]()

### 虚拟内存布局
#### 线性内存

#### 稀疏内存

#### 地址空间
  状态      | 解释
---------- ｜ ---------------------------------------------------------------------------
`None`       | 内存没有被保留或者映射
`Reserved`   | 运行时持有改地址空间，但是访问该内存，会导致错误
`Prepared`   | 内存被保留，一般没有对应的物理内存访问该片内存的行为是未定义的可以快速转换到`Ready`状态
`Ready`      | 可以被安全访问

