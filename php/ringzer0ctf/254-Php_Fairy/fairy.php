<?php
$pass = md5("admin1674227342");
//echo('strlen($pass): ' . strlen($pass) . '; strlen($_GET[pass]): ' . strlen($_GET['pass']) . "; gettype(pass): " . gettype($pass) . "; gettype($_GET[pass]): " . gettype($_GET['pass']) . "\n");
if (
    (($_GET['pass'] == $pass) && ($pass !== $_GET['pass'])) || 
    (($pass == $_GET['pass']) && ($_GET['pass'] !== $pass))) {

    if (strlen($pass) == strlen($_GET['pass'])) {
        $output = "<div class='alert alert-success'>FLAG-XXXXXXXXXXXXXXXXXXXXXXX</div>";
    } else {
        $output = "<div class='alert alert-danger'>Wrong password</div>";
    }
} else {
    $output = "<div class='alert alert-danger'>Wrong password!!</div>";
}

echo $output;