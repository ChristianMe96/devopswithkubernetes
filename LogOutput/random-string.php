<?php

use Random\RandomException;

/**
 * @throws RandomException
 */
function getRandomString(): void
{
    $random_hex = bin2hex(random_bytes(18));

    error_log($random_hex);

    sleep(5);

    getRandomString();
}

getRandomString();