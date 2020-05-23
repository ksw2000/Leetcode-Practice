// https://leetcode.com/problems/add-two-numbers/
// 108 ms 42.8 MB

/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */

function ListNode(val, next) {
    this.val = (val===undefined ? 0 : val);
    this.next = (next===undefined ? null : next);
}

function printLink(link){
    let str = '';
    for(let cur = link; cur!=null ; cur=cur.next){
        str += cur.val + ' -> ';
    }
    console.log(str);
}

function makeLL(nums){
    let root = null;
    for(let i = nums.length-1; i>=0; i--){
        root = new ListNode(nums[i], root);
    }
    return root;
}

//----------------------------------- Answer Begin -----------------------------------//
var addTwoNumbers = function(l1, l2) {
    let carry = 0;
    let sum;
    let cur1, cur2;
    let head, tail;
    head = new ListNode();
    tail = head;
    for(cur1 = l1, cur2 = l2; cur1 != null && cur2 != null; cur1 = cur1.next, cur2 = cur2.next){
        sum       = cur1.val + cur2.val + carry;
        carry     = Math.floor(sum / 10);
        tail.next = new ListNode(Math.floor(sum % 10));
        tail      = tail.next;
    }

    for(; cur1 != null; cur1 = cur1.next){
        sum       = cur1.val + carry;
        carry     = Math.floor(sum / 10);
        tail.next = new ListNode(Math.floor(sum % 10));
        tail      = tail.next;
    }

    for(; cur2 != null; cur2 = cur2.next){
        sum       = cur2.val + carry;
        carry     = Math.floor(sum / 10);
        tail.next = new ListNode(Math.floor(sum % 10));
        tail      = tail.next;
    }

    if(carry > 0){
        tail.next = new ListNode(carry);
        tail      = tail.next;
    }

    return head.next;
};

//----------------------------------- Answer END -----------------------------------//

var l1 = makeLL([2,4,3]);
var l2 = makeLL([5,6,4]);
printLink(l1);
printLink(l2);
printLink(addTwoNumbers(l1,l2));
