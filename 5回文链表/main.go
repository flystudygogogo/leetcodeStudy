/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    temp:=&ListNode{}
    c:=true
    temp=head
    var a []int

    
    if temp==nil&&temp.Next==nil {
        return c
    }

    for head!=nil{
        a=append(a,head.Val)
        head=head.Next
    }

    n:=len(a)

        b:=n/2
        j:=n-1
        for i:=0;i<b;i++{
            if a[i]==a[j]{
                if i==n/2-1{
                    return c
                }
       
            }else {
                c=false
                return c
            }
            j--
        } 
        return c
 }