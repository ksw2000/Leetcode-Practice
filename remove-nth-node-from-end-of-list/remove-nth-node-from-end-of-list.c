// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
// 0 ms 5.4 MB

#include <stdio.h>

struct ListNode {
    int val;
    struct ListNode* next;
};

struct ListNode* removeNthFromEnd(struct ListNode* head, int n) {
    struct ListNode *current, *pioneer;
    current = pioneer = NULL;
    int count = 0;
    for (pioneer = head; pioneer && count <= n;
         pioneer = pioneer->next, count++);
    if (count <= n) {
        return head->next;
    }

    for (current = head; pioneer;
         current = current->next, pioneer = pioneer->next);
    current->next = current->next->next;

    return head;
}