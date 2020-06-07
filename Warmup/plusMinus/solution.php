<?php

function plusMinus($arr) {
    $denominator = count($arr);
    $positive = 0;
    $negative = 0;
    $zero = 0;

    foreach ($arr as $a) {
        switch (true) {
            case $a > 0:
                $positive++;
                break;
            case $a < 0:
                $negative++;
                break;
            default:
                $zero++;
                break;
        }
    }

    printResult($positive, $denominator);
    printResult($negative, $denominator);
    printResult($zero, $denominator);
}

function printResult($number, $denominator)
{
    echo $number / $denominator . "\n";
}