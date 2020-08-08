<?php

// Complete the miniMaxSum function below.
function miniMaxSum($arr) {

    $b = array_flip($arr);

    $min = null;
    $max = null;

    if (count($b) === 1) {
        $sum = current($arr) * 4;
        $min = $sum;
        $max = $sum;

        echo $min . ' ' . $max;
        return;
    }

    foreach ($b as $k => $v) {
        $c = $b;
        unset($c[$k]);
        $sum = array_sum(array_flip($c));
        if (\is_null($min)) {
            $min = $sum;
        }
        if (\is_null($max)) {
            $max = $sum;
        }

        if ($sum < $min) {
            $min = $sum;
            continue;
        }
        if ($sum > $max) {
            $max = $sum;
        }
    }

    echo $min . ' ' . $max;
}

$stdin = fopen("php://stdin", "r");

fscanf($stdin, "%[^\n]", $arr_temp);

$arr = array_map('intval', preg_split('/ /', $arr_temp, -1, PREG_SPLIT_NO_EMPTY));

miniMaxSum($arr);

fclose($stdin);
