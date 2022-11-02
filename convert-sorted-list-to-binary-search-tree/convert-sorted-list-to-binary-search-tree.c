// https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/
// 10 ms 10 MB
#include<stdio.h>
#include<stdlib.h>

struct ListNode {
    int val;
    struct ListNode *next;
};

struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
};

struct TreeNode* generate(struct TreeNode* list, int len){
    if(len == 0) return NULL;
    int i = len >> 1;
    list[i].left = generate(list, i);
    list[i].right = generate(list+i+1, len-i-1);
    return list+i;
}

struct TreeNode* sortedListToBST(struct ListNode* head){
    struct ListNode* current = head;
    int len = 0;
    for(current = head; current; current=current->next) len++;
    struct TreeNode* list = malloc(len*sizeof(struct TreeNode));
    for(current = head; current; current=current->next){
        list[0].val = current->val;
        list++;
    }
    return generate(list-len, len);
}