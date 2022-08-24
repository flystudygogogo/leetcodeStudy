func hasCycle(head *ListNode)bool{

//如果只有0-1个节点就直接返回false
    if head == nil || head.Next == nil{
        return false
    }

//定义两个快慢指针。如果定义在一个节点上，那么第一次循环时就相等了。
    slow,fast := head,head.Next

//只有快慢指针不等时才进入循环，相等的话直接返回true
    for slow != fast{

//如果fast走完全程快慢指针都不相同，那么说明链表没有循环，直接返回false
        if fast == nil || fast.Next == nil{
            return false
        }

//快指针移动两步，慢指针移动一步
        fast = fast.Next.Next
        slow = slow.Next
    }

    return true
}
