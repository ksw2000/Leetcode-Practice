// URL : https://leetcode.com/problems/swap-nodes-in-pairs/
// 0 ms	7.3 MB
#include<stdio.h>
#include<stdlib.h>
struct ListNode {
    int val;
    struct ListNode *next;
};
typedef struct ListNode   link;
typedef struct ListNode* _link;

//----------------------------------- Answer Begin -----------------------------------//
struct ListNode* swapPairs(struct ListNode* head){
    struct ListNode* current;
    struct ListNode* first;
    struct ListNode* second;
    struct ListNode* newHead;
    struct ListNode* nextNode;

    if(head){
        if(head->next){
            newHead=head->next;
        }else{
            newHead=head;
        }

        for(current=head; current; current=nextNode){
            if(current->next){
                first=current;
                second=current->next;
                nextNode=current->next->next;
                first->next=(second->next)? ((second->next->next)? second->next->next : second->next) : NULL;
                second->next=first;
            }else{
                nextNode=current->next;
            }
        }
        return newHead;
    }
    return head;
}
//----------------------------------- Answer END -----------------------------------//

void append(struct ListNode* head, int val){
    struct ListNode* newNode=malloc(sizeof(struct ListNode));
    newNode->val=val;
    newNode->next=NULL;

    struct ListNode* current;
    for(current=head; current->next; current=current->next);
    current->next=newNode;
}

int main(){
    struct ListNode* head=malloc(sizeof(struct ListNode));
    struct ListNode* current;
    head->val=1;
    head->next=NULL;
    append(head,2);
    append(head,3);
    //append(head,4);
    for(current=head; current ;current=current->next){
        printf("%d->",current->val);
    }
    printf("\n");
    for(current=swapPairs(head); current; current=current->next){
        printf("%d->",current->val);
    }
    printf("\n");
    system("pause");
    return 0;
}
