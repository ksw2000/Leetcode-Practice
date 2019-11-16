// URL : https://leetcode.com/problems/palindrome-linked-list/
#include<stdio.h>
#include<stdlib.h>
typedef enum{
    false,true
}bool;

struct ListNode{
    int val;
    struct ListNode *next;
};
typedef struct ListNode list;

//----------------------------------- Answer Begin -----------------------------------//
typedef struct s{
    int val;
    struct s* ptr;
}stack;

void push(stack** s, int value){
    stack* temp=malloc(sizeof(stack));
    temp->val=value;
    temp->ptr=*s;
    *s=temp;
}

int pop(stack** s){
    stack* temp;
    int re=(*s)->val;
    temp=(*s)->ptr;
    free(*s);
    (*s)=temp;
    return re;
}

bool isPalindrome(struct ListNode* head){
    struct ListNode* current;
    stack* s=NULL;
    for(current=head; current; current=current->next){
        push(&s,current->val);
    }

    for(current=head; current; current=current->next){
        if(pop(&s) != current->val){
            return false;
        }
    }
    return true;
}
//----------------------------------- Answer END -----------------------------------//

int main(){
    int arr[]={9,4,8,7};
    int i;
    list* head=NULL;
    for(i=0; i<4; i++){
        list* tmp=malloc(sizeof(list));
        tmp->next=head;
        tmp->val=arr[i];
        head=tmp;
    }
    printf("%d\n",isPalindrome(head));
    system("pause");
    return 0;
}
