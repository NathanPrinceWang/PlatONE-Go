# Ethereum 代码风格指南
## 1. 头文件
* `#pragma once` 保护：所有头文件都应该使用 `#pragma once` 来防止头文件被多重包含。
* 前置声明：尽可能地避免使用前置声明。使用 #include 包含需要的头文件即可。
* 内联函数：只有当函数只有 10 行甚至更少时才将其定义为内联函数。
* `#include` 的路径及顺序：使用标准的头文件包含顺序可增强可读性，避免隐藏依赖: 相关头文件，C 库，C++ 库，其他库的.h，本项目内的.h。
  
## 2. 作用域
* 命名空间：命名空间将全局作用域细分为独立的，具名的作用域，可有效防止全局作用域的命名冲突。
  * 在命名空间的最后注释出命名空间的名字
* 局部变量：将函数变量尽可能置于最小作用域内，并在变量声明时进行初始化。
  
## 3. 类
* 构造函数：不要在构造函数中调用虚函数，也不要在无法报出错误时进行可能失败的初始化。
* 隐式类型转换：不要定义隐式类型转换。 对于转换运算符和单参数构造函数，请使用 `explicit` 关键字。
* 可拷贝类型和可移动类型：如果你的类型需要，就让它们支持拷贝 / 移动。 否则，就把隐式产生的拷贝和移动函数禁用。
* 结构体与类：仅当只有数据成员时使用 `struct`，其它一概使用 `class`。
* 继承：使用组合常常比使用继承更合理。 如果使用继承的话，定义为 public 继承。
  * 所有继承必须是 public 的。 如果你想使用私有继承，你应该替换成把基类的实例作为成员对象的方式
  * 如果你的类有虚函数，则析构函数也应该为虚函数
  * 对于重载的虚函数或虚析构函数，使用 override，或 (较不常用的) final 关键字显式地进行标记
* 多重继承：只有当所有父类除第一个外都是 纯接口类 时，才允许使用多重继承。 为确保它们是纯接口，这些类必须以 Interface 为后缀。
* 运算符重载：除少数特定环境外，不要重载运算符。 也不要创建用户定义字面量。只有在意义明显，不会出现奇怪的行为并且与对应的内建运算符的行为一致时才定义重载运算符。
* 存取控制：将所有数据成员声明为 private，除非是 static const 类型成员。
* 声明顺序：一般应以`public:`开始，后跟`protected:`，最后是`private:`。 省略空部分。

## 4. 函数
* 参数顺序：输入参数在先，后跟输出参数。
* 编写简短函数：我们倾向于编写简短，凝练的函数。如果函数超过 40 行，可以思索一下能不能在不影响程序结构的前提下对其进行分割。
* 引用参数：所有按引用传递的参数必须加上 const。
* 函数重载：若要使用函数重载，则必须能让读者一看调用点就胸有成竹，而不用花心思猜测调用的重载函数到底是哪一种。 这一规则也适用于构造函数。
* 缺省参数：只允许在非虚函数中使用缺省参数，且必须保证缺省参数的值始终一致。
* 函数返回类型后置语法：只有在常规写法 (返回类型前置) 不便于书写或不便于阅读时使用返回类型后置语法。如`template <class T，class U> auto add(T t，U u) -> decltype(t + u);`。
* 函数参数变量：入参均已`_`开头，以区分函数体里面声明的局部变量。如`bool Block::sync(BlockChain const& _bc，h256 const& _block，BlockHeader const& _bi)`。

## 5. 命名约定
* 文件命名：类或模块定义时文件名一般一一对应，如TransactionQueue.cpp和TransactionQueue.h
* 类命名：类型名称的每个单词首字母均大写，不包含下划线: MyExcitingClass，MyExcitingEnum。
* 变量命名：变量以驼峰式命名。类的成员变量m_开头，但结构体的就不用。
  * 普通变量命名：如`string blockData`。
  * 类数据成员：类的成员变量以前缀`m_`开头，以区分局部变量。
  * 结构体变量：不管是静态的还是非静态的，结构体数据成员都可以和普通变量一样，不用像类那样加前缀`m_`。
  * 常量命名：声明为 constexpr 或 const 的变量，或在程序运行期间其值始终保持不变的，命名时以`c_`开头，大小写混合。如`static const int64_t c_maxGasEstimate = 50000000;`。
  * 全局变量：命名时以`g_`开头，大小写混合。除非不得已，否则应该尽可能少使用全局变量！除非不得已，否则应该尽可能少使用全局变量！除非不得已，否则应该尽可能少使用全局变量！
* 函数命名：常规函数使用大小写混合，首字母小写，参考普通变量命名规则。
* 命名空间命名：命名空间以小写字母命名。 最高级命名空间的名字取决于项目名称。 要注意避免嵌套命名空间的名字之间和常见的顶级命名空间的名字之间发生冲突。
* 枚举命名：以驼峰式命名，第一个字母大写。建议以C++11语法枚举。如:
  ```
  enum class Verification {
      Skip，
      Normal
   };
  ```
* 宏命名：全部使用大写，以下划线分割各个单词。如`#define ETH_TIMED_IMPORTS 1`。建议尽量使用常量替代宏的声明。除非不得不用。
  
## 6. 注释
* 注释风格：小段注释用`//`，大段注释用`/* */`。小段注释`//`后面请加空格隔开注释。
* 文件注释：文件注释描述了该文件的内容。必须至少包含文件名，作者，文件创建日期。如：
   ```
  /** @file BlockChain.h
    * @author Gav Wood <i@gavwood。com>
    * @date 2014/08/08
    */
  ```
* 类注释：每个类的定义都要附带一份注释，描述类的功能和用法，除非它的功能相当明显。
* 函数注释：函数声明处的注释描述函数功能; 定义处的注释描述函数实现。
* 变量注释：通常变量名本身应该足以很好说明变量用途。 某些情况下，也需要额外的注释说明。
* 函数参数注释：如果函数参数的意义不明显，考虑用下面的方式进行弥补。
  * 如果参数是一个字面常量，你应当用一个常量名让这一约定变得更明显。
  * 如果某个函数有多个配置选项，你可以考虑定义一个类或结构体以保存所有的选项。
  * 用具名变量代替大段而复杂的嵌套表达式。
* TODO 注释：对那些临时的，短期的解决方案，或已经够好但仍不完美的代码使用 TODO 注释。请在TODO 注释处写上自己的名字。如`// TODO(Zeke，bug1234) change this to use relations。`

## 7. 格式
* 行长度：每一行代码字符数不超过 120。
* 非 ASCII 字符：尽量不使用非 ASCII 字符，使用时必须使用 UTF-8 编码。
* 空格还是制表位：只使用空格，每次缩进 4 个空格。
* 函数声明与定义：返回类型和函数名在同一行，参数也尽量放在同一行，如果放不下就对形参分行，分行方式与函数调用一致。注意以下几点。
  * 使用好的参数名。
  * 只有在参数未被使用或者其用途非常明显时，才能省略参数名
  * 如果返回类型和函数名在一行放不下，分行
  * 如果返回类型与函数声明或定义分行了，不要缩进
  * 左圆括号总是和函数名在同一行
  * 函数名和左圆括号间永远没有空格
  * 圆括号与参数间没有空格
  * 左大括号另起新行
  * 右大括号总是单独位于函数最后一行
  * 所有形参应尽可能对齐
  * 换行后的参数保持 4 个空格的缩进
* Lambda 表达式：Lambda 表达式对形参和函数体的格式化和其他函数一致; 捕获列表同理，表项用逗号隔开。
* 函数调用：要么一行写完函数调用，要么在圆括号里对参数分行，要么参数另起一行且缩进四格。 如果没有其它顾虑的话，尽可能精简行数，比如把多个参数适当地放在同一行里。
* 条件语句：倾向于不在圆括号内使用空格。 关键字 if 和 else 另起一行。
* 循环和开关选择语句：switch 语句可以使用大括号分段，以表明 cases 之间不是连在一起的。 在单语句循环里，括号可用可不用。 空循环体应使用 {} 或 continue
* 指针和引用表达式：句点或箭头前后不要有空格。 指针或地址操作符 (*，&) 之后不能有空格
* 布尔表达式：如果一个布尔表达式超过 标准行宽，断行方式要统一一下。
* 函数返回值：不要在 return 表达式里加上非必须的圆括号
* 变量及数组初始化：用 =，() 和 {} 均可。
* 预处理指令：预处理指令不要缩进，从行首开始。
* 类格式：访问控制块的声明依次序是`public:`，`protected:`，`private:`，不需要进行缩进。
* 构造函数初始值列表：构造函数初始化列表放在同一行或按4个空格缩进并排多行。
* 命名空间格式化：命名空间内容不缩进。
* 水平留白：水平留白的使用根据在代码中的位置决定。 永远不要在行尾添加没意义的留白。

## 8. 其他
* 文件编码：统一用UTF-8编码。
* 引用参数：所有按引用传递的参数必须加上 const。
* 右值引用：只在定义移动构造函数与移动赋值操作时使用右值引用。 不要使用 std::forward。
* 函数重载：若要用好函数重载，最好能让读者一看调用点就胸有成竹，不用花心思猜测调用的重载函数到底是哪一种。
* 类型转换：使用 C++ 的类型转换，如 static_cast<>()。 不要使用 int y = (int)x 或 int y = int(x) 等转换方式。
* 预处理宏：使用宏时要非常谨慎，尽量以内联函数，枚举和常量代替之。
* 0，nullptr 和 NULL：指针均用nullptr替代NULL与0。
* sizeof：尽可能用 sizeof(varname) 代替 sizeof(type)。
* auto：用 auto 绕过烦琐的类型名，只要可读性好就继续用，别用在局部变量之外的地方。
* 列表初始化：可以用C++11列表初始化。
* Lambda 表达式：适当使用 lambda 表达式。别用默认 lambda 捕获，所有捕获都要显式写出来。
* 模板编程：不要使用复杂的模板编程。
