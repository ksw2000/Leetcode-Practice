<?php
// https://leetcode.com/problems/add-two-numbers/
// 12 ms 15 MB

/**
 * Definition for a singly-linked list->
 * class ListNode {
 *     public $val = 0;
 *     public $next = null;
 *     function __construct($val = 0, $next = null) {
 *         $this->val = $val;
 *         $this->next = $next;
 *     }
 * }
 */

 class ListNode {
    public $val = 0;
    public $next = null;
    function __construct($val = 0, $next = null) {
        $this->val = $val;
        $this->next = $next;
    }
}

class Solution {
    function addTwoNumbers($l1, $l2) {
        $carry = 0;
        $head = new ListNode();
        $tail = $head;

        for(; $l1 != null && $l2 != null; $l1 = $l1->next, $l2 = $l2->next){
            $sum        = $l1->val + $l2->val + $carry;
            $carry      = (int)($sum / 10);
            $tail->next = new ListNode((int)($sum % 10));
            $tail       = $tail->next;
        }

        for(; $l1 != null; $l1 = $l1->next){
            $sum       = $l1->val + $carry;
            $carry     = (int)($sum / 10);
            $tail->next = new ListNode((int)($sum % 10));
            $tail      = $tail->next;
        }

        for(; $l2 != null; $l2 = $l2->next){
            $sum       = $l2->val + $carry;
            $carry     = (int)($sum / 10);
            $tail->next = new ListNode((int)($sum % 10));
            $tail      = $tail->next;
        }

        if($carry > 0){
            $tail->next = new ListNode($carry);
            $tail      = $tail->next;
        }

        return $head->next;
    }
}

function makeLL($arr){
    $len = count($arr);
    $root = null;
    for($i = $len-1; $i>=0; $i--){
        $new = new ListNode($arr[$i]);
        $new->next = $root;
        $root = $new;
    }
    return $root;
}

function printLL($list){
    $str = "";
    for($c = $list; $list!=null; $list=$list->next){
        $str .= $list->val." -> ";
    }
    $str.= "null\n";
    echo $str;
}

$l1 = makeLL([2,4,3]);
printLL($l1);
$l2 = makeLL([5,6,4]);
printLL($l2);

$l3 = (new Solution())->addTwoNumbers($l1, $l2);
printLL($l3);
?>
