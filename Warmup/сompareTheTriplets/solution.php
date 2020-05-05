<?php

// Complete the compareTriplets function below.
function compareTriplets($a, $b)
{
    $alice = 0;
    $bob = 0;

    foreach ($a as $key => $value) {
        if ($value > $b[$key]) {
            $alice += 1;
        } else if ($value < $b[$key]) {
            $bob += 1;
        }
    }

    return [$alice, $bob];
}
