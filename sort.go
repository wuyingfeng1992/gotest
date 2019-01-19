package main

import "fmt"

func main(){
	// 插入算法
	arr:= [...]int{1,3,5,2,6,4,6,9,8,7}
	length := len(arr)
	for i:=1;i<length;i++{
		for j:=i;j>0;j--{
			if arr[j-1]<arr[j]{
				break;
			}else{
				arr[j-1],arr[j]=arr[j],arr[j-1]
			}
		}

	}
	fmt.Println(arr)
	count:=0

	b:= [...]int{2,1,4,5,3}
	l:=len(b)
	//选择算法，每次选到最小的交换位置，第一次其实就是找len－1的最小值了
	 for z:=0;z<l-1;z++{
	 	for k:=z+1;k<l;k++{
	 		count++
	 		if b[k]<b[z]{
				b[z],b[k]=b[k],b[z]
			}
		}

	 }

	//arr[1]=29
	fmt.Println(b)





	//冒泡算法
	a3:=[...]int{5,6,4,3,2,1}

	l3:=len(a3)

	for i:=0;i<l3;i++{
		for j:=1;j<l3-i;j++{
			if a3[j]<a3[j-1]{
				a3[j],a3[j-1]=a3[j-1],a3[j]
			}
		}
	}
	fmt.Println(a3)
	a4:=[...]int{5,6,4,3,2,1}
	qsort(a4[:],0,len(a4)-1)
	fmt.Println(a4)

	//



}

func qsort(a []int,left,right int){

	if(left>=right){
		return
	}

	k:=left;
	val:=a[left]
	//确定val正确的位置
	for i:=left+1;i<=right;i++{
		if a[i]<val{
			a[k]=a[i]
			a[i]=a[k+1]
			k++
		}
	}

	a[k]=val


	qsort(a,left,k)
	qsort(a,k+1,right)

}
