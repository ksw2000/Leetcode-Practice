// https://leetcode.com/problems/add-two-numbers/
// 8 ms	7.5 MB
#include<stdio.h>
#include<stdlib.h>

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */

struct ListNode{
    int val;
    struct ListNode *next;
};

//----------------------------------- Answer Begin -----------------------------------//
struct ListNode* addTwoNumbers(struct ListNode* l1, struct ListNode* l2){
    struct ListNode *cur1, *cur2;
    struct ListNode *head, *tail;
    head = calloc(1, sizeof(*head));
    tail = head;
    int carry = 0;
    int sum;
    for(cur1 = l1, cur2 = l2; cur1 && cur2; cur1 = cur1->next, cur2 = cur2->next){
        sum       = cur1->val + cur2->val + carry;
        carry     = sum / 10;

        struct ListNode* new = malloc(sizeof(*new));
        new->val   = sum % 10;
        new->next  = NULL;
        tail->next = new;
        tail       = new;
    }

    for(; cur1; cur1 = cur1->next){
        sum       = cur1->val + carry;
        carry     = sum / 10;

        struct ListNode* new = malloc(sizeof(*new));
        new->val   = sum % 10;
        new->next  = NULL;
        tail->next = new;
        tail       = new;
    }

    for(; cur2; cur2 = cur2->next){
        sum       = cur2->val + carry;
        carry     = sum / 10;

        struct ListNode* new = malloc(sizeof(*new));
        new->val   = sum % 10;
        new->next  = NULL;
        tail->next = new;
        tail       = new;
    }

    if(carry > 0){
        struct ListNode* new = malloc(sizeof(*new));
        new->val  = carry;
        new->next = NULL;
        tail->next = new;
        tail       = new;
    }

    tail->next = NULL;
    return head->next;
}
//----------------------------------- Answer END -----------------------------------//

struct ListNode* makeLL(int* array, int size){
    struct ListNode* root = NULL;
    for(size--; size>=0; size--){
        struct ListNode* new = malloc(sizeof(*new));
        new -> val = array[size];
        new -> next = root;
        root = new;
    }
    return root;
}

void printLL(struct ListNode* root){
    struct ListNode* cur;
    for(cur = root; cur; cur=cur->next){
        printf("%d -> ", cur->val);
    }
    printf("NULL \n");
}

int main(){
    int arr1[] = {2,4,3,7};
    int arr2[] = {5,6,4,8};
    struct ListNode *l1 = makeLL(arr1, 4);
    struct ListNode *l2 = makeLL(arr2, 4);
    printLL(l1);
    printLL(l2);

    l1 = addTwoNumbers(l1, l2);
    printLL(l1);

    return 0;
}
