// URL : https://leetcode.com/problems/remove-nth-node-from-end-of-list/
// 0 ms	5.4 MB
#include<stdio.h>
#include<stdlib.h>

struct ListNode {
    int val;
    struct ListNode *next;
};

//----------------------------------- Answer Begin -----------------------------------//
struct ListNode* removeNthFromEnd(struct ListNode* head, int n){
    struct ListNode *current, *pioneer;
    current = pioneer = NULL;
    int count=0;
    for(pioneer=head; pioneer!=NULL && count<=n; pioneer=pioneer->next, count++);
    if(count<=n){
        return head->next;
    }

    for(current=head; pioneer!=NULL; current=current->next, pioneer=pioneer->next);
    current->next = current->next->next;

    return head;
}
//----------------------------------- Answer END -----------------------------------//

void append(struct ListNode** head, struct ListNode** tail, int val){
    struct ListNode* new=malloc(sizeof(*new));
    new->val=val;
    new->next=NULL;
    if(*head==NULL){
        *head=*tail=new;
    }else{
        (*tail)->next=new;
        *tail=(*tail)->next;
    }
}

int main(){
    int arr[]={1,2,3,4,5,6,7};
    int i;
    struct ListNode *head,*tail,*current;
    head=tail=NULL;
    for(i=0; i<sizeof(arr)/sizeof(int); i++){
        append(&head,&tail,arr[i]);
    }
    printf("original:\n");
    for(current=head; current; current=current->next){
        printf("%d->",current->val);
    }
    printf("\nnew:\n");
    struct ListNode* newll=removeNthFromEnd(head,7);
    for(current=newll; current; current=current->next){
        printf("%d->",current->val);
    }
    printf("\n");

    system("pause");
    return 0;
}
