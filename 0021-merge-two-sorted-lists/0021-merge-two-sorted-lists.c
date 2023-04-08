/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode* mergeTwoLists(struct ListNode* list1, struct ListNode* list2){
    if (!list1 || !list2) {
        return (!list1) ? list2 : list1;
    }
    
    struct ListNode* head;
    if (list1->val <= list2->val) {
        head = list1;
        list1 = list1->next;
    } else {
        head = list2;
        list2 = list2->next;
    }
    head->next = mergeTwoLists(list1, list2);
    return head;
}