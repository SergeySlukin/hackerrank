<?php
function staircase($n) {
    for ($i = 1; $i <= $n; $i++) {
        $j = $n - $i;
        echo str_repeat(" ", $j) . str_repeat("#", $n - $j) . "\n";
    }
}