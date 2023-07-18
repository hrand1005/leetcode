/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode* oddEvenList(struct ListNode* head){
    struct ListNode *odd, *oddh, *even, *evenh, *n;
    
    if (head == NULL || head->next == NULL)
        return head;
    
    odd = head;
    evenh = even = head->next;
    n = head->next->next;
    
    for (int i = 1; n != NULL; i++) {
        printf("n: %d\n", n->val);
        if (i % 2 == 0) {
            even->next = n;
            even = even->next;
        } else {
            odd->next = n;
            odd = odd->next;
        }
        n = n->next;
    }
    odd->next = evenh;
    even->next = NULL;
    return head;
}