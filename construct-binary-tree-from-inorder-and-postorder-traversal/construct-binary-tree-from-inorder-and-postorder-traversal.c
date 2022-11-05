// https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
// 10 ms 8.3 MB
#include <stdio.h>
#include <stdlib.h>

struct TreeNode {
    int val;
    struct TreeNode* left;
    struct TreeNode* right;
};

struct hashEl {
    int key;
    int val;
    struct hashEl* next;
};

struct hash {
    struct hashEl** list;
    int buckets;
};

struct hash* Hash(int buckets) {
    struct hash* h = malloc(sizeof(struct hash));
    h->list = calloc(buckets, sizeof(struct hashEl*));
    h->buckets = buckets;
    return h;
}

__inline__ int hashCode(struct hash* h, int key) {
    return key & (h->buckets - 1);
}

__inline__ void put(struct hash* h, int key, int val) {
    int code = hashCode(h, key);
    struct hashEl* e = malloc(sizeof(struct hashEl));
    e->key = key;
    e->val = val;
    e->next = h->list[code];
    h->list[code] = e;
}

__inline__ int get(struct hash* h, int key) {
    int code = hashCode(h, key);
    struct hashEl* current = h->list[code];
    for (; current; current = current->next) {
        if (current->key == key)
            return current->val;
    }
    fprintf(stderr, "ERROR");
    return 0;
}

struct el {
    int inorderSize;
    int postorderSize;
    int* inorder;
    int* postorder;
    struct TreeNode* me;
};

struct stack {
    struct el** list;
    int len;
    int cap;
};

struct stack* Stack(int cap) {
    struct stack* s = malloc(sizeof(struct stack));
    s->list = malloc(sizeof(struct el*) * cap);
    s->cap = cap;
    s->len = 0;
    return s;
}

__inline__ void push(struct stack* s, struct el* status) {
    if (s->len == s->cap) {
        struct el** new = malloc(sizeof(struct el*) * s->cap << 1);
        memcpy(new, s->list, sizeof(struct el*) * s->cap);
        free(s->list);
        s->list = new;
        s->cap = s->cap << 1;
    }
    s->list[s->len++] = status;
}

__inline__ struct el* pop(struct stack* s) {
    return s->list[--s->len];
}

struct TreeNode* buildTree(int* inorder,
                           int inorderSize,
                           int* postorder,
                           int postorderSize) {
    if (postorderSize == 0)
        return NULL;
    struct stack* s = Stack(32);

    struct el* root = malloc(sizeof(struct el));
    struct el* status;
    struct el* tmp;
    root->inorder = inorder;
    root->inorderSize = inorderSize;
    root->postorder = postorder;
    root->postorderSize = postorderSize;
    root->me = malloc(sizeof(struct TreeNode));
    push(s, root);

    // create hashmap
    struct hash* map = Hash(256);
    int k;
    for (k = 0; k < inorderSize; k++) {
        put(map, inorder[k], k);
    }

    while (s->len > 0) {
        status = pop(s);
        if (status->postorderSize > 0) {
            int val = status->postorder[status->postorderSize - 1];
            status->me->val = val;
            // find by hashmap
            int i = get(map, val);

            // [ ] [ ] [ ] [ i ] [ ] ... [ ]
            //  ^ inorder
            //         [ ] [ i'] [ ]
            //          ^ status->inorder
            //  |_______|
            //    offset

            i = i - (status->inorder - inorder);

            // enqueue for left
            if (i > 0) {
                tmp = malloc(sizeof(struct el));
                tmp->inorder = status->inorder;
                tmp->inorderSize = i;
                tmp->postorder = status->postorder;
                tmp->postorderSize = i;
                tmp->me = malloc(sizeof(struct TreeNode));
                status->me->left = tmp->me;
                push(s, tmp);
            } else {
                status->me->left = NULL;
            }

            // enqueue for right
            if (status->inorderSize - i - 1 > 0) {
                tmp = malloc(sizeof(struct el));
                tmp->inorder = (status->inorder) + i + 1;
                tmp->inorderSize = status->inorderSize - i - 1;
                tmp->postorder = (status->postorder) + i;
                tmp->postorderSize = status->inorderSize - i - 1;
                tmp->me = malloc(sizeof(struct TreeNode));
                status->me->right = tmp->me;
                push(s, tmp);
            } else {
                status->me->right = NULL;
            }
        }
    }

    return root->me;
}