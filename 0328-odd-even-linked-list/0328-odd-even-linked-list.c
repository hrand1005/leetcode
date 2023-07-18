/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode* oddEvenList(struct ListNode* head){
    struct ListNode *odd, *even, *evenh, *n;
    
    if (head == NULL || head->next == NULL)
        return head;
    
    odd = head;
    evenh = even = head->next;
    while (even != NULL && even->next != NULL) {
        odd->next = even->next;
        odd = odd->next;
        even->next = even->next->next;
        even = even->next;
    }
    odd->next = evenh;
    
    return head;
}