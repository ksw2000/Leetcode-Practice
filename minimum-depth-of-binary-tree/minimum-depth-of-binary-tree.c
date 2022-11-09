// https://leetcode.com/problems/minimum-depth-of-binary-tree/
// 162 ms 75.8 MB
#include<stdio.h>

struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
};

__inline__ int min(int a, int b) {
    return a < b ? a : b;
}

int minDepth(struct TreeNode* root){
    if(!root) return 0;
    if(!root->left) return 1 + minDepth(root->right);
    if(!root->right) return 1 + minDepth(root->left);
    return 1 + min(minDepth(root->left), minDepth(root->right));
}