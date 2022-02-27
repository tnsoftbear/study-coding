<?php
/**
 * Replace arguments order with help of auto-refactoring tool, you will see broken code in caller scope
 */

class Foo {
    public static $bar = [1, 2];
}

$bar2 = [1, 2];

function baz($arg1, $arg2) {
}

baz(...Foo::$bar);
baz(...$bar2);
