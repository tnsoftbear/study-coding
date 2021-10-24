<?php
$cn = "checkme";
$value = 1;
$expires = time() + 60*60*24;
$options = [
//  "path" => '/',
//  "domain" => 'samesite.lv',
  "expires" => $expires,
  "httponly" => false,
  "samesite" => $_GET['ss'],
  "secure" => true,
];
$result = (int)setcookie($cn, $value, $options);
echo "Cookie set: {$result}</br>";
print_r($options);