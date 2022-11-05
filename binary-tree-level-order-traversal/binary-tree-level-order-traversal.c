// https://leetcode.com/problems/binary-tree-level-order-traversal/
// 0 ms 7.1 MB
#include<stdio.h>
#include<stdlib.h>

struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
};
/**
 * Return an array of arrays of size *returnSize.
 * The sizes of the arrays are returned as *returnColumnSizes array.
 * Note: Both returned array and *columnSizes array must be malloced, assume caller calls free().
 */

struct queue{
    struct TreeNode** list;
    int len;
    int cap;
};

struct queue* newQueue(int cap){
    struct queue* q = malloc(sizeof * q);
    q->list = malloc(cap * sizeof(struct TreeNode*));
    q->len = 0;
    q->cap = cap;
    return q;
}

void enQueue(struct queue* q, struct TreeNode* val){
    if(q->len == q->cap){
        struct TreeNode** new = malloc(sizeof(struct TreeNode*) * q->cap << 1);
        memcpy(new, q->list, sizeof(struct TreeNode*) * q->cap);
        free(q->list);
        q->list = new;
        q->cap = q->cap << 1;
    }
    q->list[q->len++] = val;
}

int** levelOrder(struct TreeNode* root, int* returnSize, int** returnColumnSizes){
    struct queue* q1 = newQueue(16);
    struct queue* q2 = newQueue(16);
    struct queue* tmp;
    if(root) enQueue(q1, root);

    int retCap = 16;
    int retLen = 0;
    int** ret = malloc(sizeof(int*) * retCap);
    *returnColumnSizes = malloc(sizeof(int) * retCap);

    int i;
    while(q1->len > 0){
        // initializae q2
        q2->len = 0;
        ret[retLen] = malloc(sizeof(int) * q1->len);
        (*returnColumnSizes)[retLen] = q1->len;

        for(i = 0; i < q1->len; i++){
            // get
            struct TreeNode* node = q1->list[i];
            ret[retLen][i] = node->val;
            if(node->left) enQueue(q2, node->left);
            if(node->right) enQueue(q2, node->right);
        }
        retLen++;
        // enlarge
        if(retLen == retCap){
            int** newRet = malloc(sizeof(int*) * retCap << 1);
            int* newReturnColumnSizes = malloc(sizeof(int) * retCap << 1);
            memcpy(newRet, ret, sizeof(int*) * retCap);
            memcpy(newReturnColumnSizes, *returnColumnSizes, sizeof(int) * retCap);
            free(ret);
            free(*returnColumnSizes);
            ret = newRet;
            *returnColumnSizes = newReturnColumnSizes;
            retCap = retCap << 1;
        }

        tmp = q1;
        q1 = q2;
        q2 = tmp;
    }

    *returnSize = retLen;

    return ret;
}