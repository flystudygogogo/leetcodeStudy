/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode{

        //init two pointers
        rabbit := &ListNode{}
        turtle := &ListNode{}

        //both set at the head at the begining
        rabbit = head
        turtle = head 

        // if empty, return empty
        if head == nil {
            return nil 
        }

        //if only one element, return empty (as you can only delete this one)
        if head.Next == nil {
            return nil
        }

        // moves rabbit ahead n steps, depends on n, while turtle stays 
        for i :=0; i < n-1; i +=1 {
            if(rabbit.Next != nil){
                rabbit = rabbit.Next
            }
        }

        //defines a helper variable to store the previous node behinds turtle
        //so that the node of turtle can be removed later
        turtle_prev := &ListNode{}
        
        //move both pointers ahead until rabbit reaches the end 
        for rabbit.Next != nil  {
            
            //update the previous node of turtle over the time 
            turtle_prev = turtle

            //move turtle and rabbit
            turtle = turtle.Next
            rabbit = rabbit.Next 
        }

        //edge cases : when the node to remove is the head 
        if(turtle == head){
            return turtle.Next
        }

        //connect the previous node of turtle directly to the one after turtle to 
        //delete the node of turtle 
        turtle_prev.Next = turtle.Next  
        
        //return head 
        return head
}
