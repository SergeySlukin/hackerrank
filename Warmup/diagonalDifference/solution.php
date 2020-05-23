<?php

function diagonalDifference($arr) {
    $a = 0;
    $b = 0;
    $j = 0;
    $rowElementsCount = count($arr[0]);
    $jj = $rowElementsCount - 1;
    for ($i = 0; $i < $rowElementsCount; $i++) {
        $a += $arr[$i][$j];
        $b += $arr[$i][$jj];
        $j++;
        $jj--;
    }

    return abs($a - $b);
}
