// https://leetcode.com/problems/swap-nodes-in-pairs/
// 0 ms 5.9MB
#include<stdio.h>

struct ListNode {
    int val;
    struct ListNode* next;
};

struct ListNode* swapPairs(struct ListNode* head) {
    struct ListNode pseudo = {.next = head};
    head = &pseudo;
    struct ListNode* p1 = head;
    struct ListNode* p2;
    struct ListNode* p3;

    while (p1 && p1->next && p1->next->next) {
        p2 = p1->next->next;
        p3 = p1->next->next->next;
        p1->next->next->next = p1->next;
        p1->next->next = p3;
        p1->next = p2;
        p1 = p1->next->next;
    }
    return head->next;
}