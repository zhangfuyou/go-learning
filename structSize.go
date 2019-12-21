package main

//struct内存大小，内存对齐
import (
	"fmt"
	"unsafe"
)

type user1 struct {
	b byte
	i int32
	j int64
}

type user2 struct {
	b byte
	j int64
	i int32
}

type user3 struct {
	i int32
	b byte
	j int64
}

type user4 struct {
	i int32
	j int64
	b byte
}

type user5 struct {
	j int64
	b byte
	i int32
}

type user6 struct {
	j int64
	i int32
	b byte
}

func main() {
	var u1 user1
	var u2 user2
	var u3 user3
	var u4 user4
	var u5 user5
	var u6 user6

	fmt.Println("u1 size is ",unsafe.Sizeof(u1))
	fmt.Println("u2 size is ",unsafe.Sizeof(u2))
	fmt.Println("u3 size is ",unsafe.Sizeof(u3))
	fmt.Println("u4 size is ",unsafe.Sizeof(u4))
	fmt.Println("u5 size is ",unsafe.Sizeof(u5))
	fmt.Println("u6 size is ",unsafe.Sizeof(u6))

	/*result
	u1 size is  16
	u2 size is  24
	u3 size is  16
	u4 size is  24
	u5 size is  16
	u6 size is  16
	*/

	/*
	结果说明：
	1.内存对齐影响struct的大小
	2.struct的字段顺序影响struct的大小
	*/

	/*内存对齐规则：
	1.对于具体类型来说，对齐值=min(编译器默认对齐值，类型大小Sizeof长度）。也就是在默认设置的对齐值和类型的内存占用
	大小之间，取最小值为该类型的对齐值。
	2.struct在每个字段都内存对齐后，其本身也要进行对齐，对齐值=min(默认对齐值，字段最大类型长度)。即struct的所有字段中，
	最大的那个类型的长度及默认对齐值之间，取最小那个。
	*/

	/*user1结果解释：
	byte, int32, int64的对齐值分别为1，4，8，内存占用大小也是1，4，8.那么对于user1，它的字段顺序是byte,int32,int64
	1. 首先使用第一条内存规则进行内存对齐，其内存结构如下：
		bxxx|iiii|jjjj|jjjj
	1）uise1类型，第一个字段是byte，对齐值是1，大小为1，所以放在内存布局中的第一位
	2）第二个字段是int32，对齐值是4，大小是4，所以它的内存偏移值必须是4的倍数，在当前user1中，就不能从第2位开始了，必须从第5
	位开始，也就是偏移量为4.第2、3、4位由编译器进行填充，一般为0值，也称为内存空洞，所以第5为到第8位为第二个字段.
	3）第三个字段，int64，对齐值为8，内存占用大小也是8，因为user1前两个字段已经排到了第8位，所以下一位的偏移量正好是8，是第三个字段对齐值的
	倍数，不用填充，可以直接排列第三个字段，也就是第9为到第16位为第三个字段。
	2. 根据第一条内存规则对齐后，内存长度已经为16个字节，现在开始使用第二条规则进行对齐，根据第二条规则，默认对齐值为8，字段中最大类型长度也是8，所以
	求出结构体的对齐值为8，结构体目前的内存长度为16，是8的倍数，已经实现了对齐。
	所以到此为止，结构体user1的内存占用大小为16字节

	*/

	/*user2结果解释：
	1。 根据第一条内存对齐规则对齐，其内存结构如下：
		bxxx|xxxx|jjjj|jjjj|iiii
	1) 按对齐值和内存占用的大小，第一个字段b偏移量为0，占用1个字节放在第一位
	2） 第二个字段int64，对齐值和大小都是8，所以要从偏移量8开始，也就是第9位到16位位j，第2位到第8位被编译器填充
	3） 目前整个内存布局已经偏移了16位，正好是第三个字段i的对齐值4的倍数，所以不用填充，可以直接排列，第17位到20位为i

	2. 现在所有字段对齐好了，整个内存大小为1+7+8+4=20个字节，开始使用内存对齐的第二条规则，也就是结构体的对齐，通过默认对齐值
	和最大的字段大小，求出结构体的对齐值为8.
		现在整个内存布局大小为20，不是8的倍数，所以需要进行内存填充，补足到8的倍数，最小的就是24，所以对齐后整个内存布局为：
		bxxx|xxxx|jjjj|jjjj|iiii|xxxx
	所以，最终user2的大小为24字节
	*/

	/*user3的内存布局
		iiii|bxxx|jjjj|jjjj  16字节
	*/

	/*user4的内存布局
		iiii|xxxx|jjjj|jjjj|bxxxx|xxxx 24字节
	*/

	/*user5的内存布局
		jjjj|jjjj|bxxx|iiii		16字节
	*/

	/*user6的内存布局
		jjjj|jjjj|iiii|bxxx		16字节
	*/
}
