<?php

$file = fopen('/Users/ryuya/workspace/resize-api/tmp/original/gopher.png', 'rb');

$curl = curl_init('https://3d8r7a230b.execute-api.ap-northeast-1.amazonaws.com/default/resize-api');
curl_setopt_array($curl, [
  CURLOPT_UPLOAD => true,
  CURLOPT_CUSTOMREQUEST => "POST",
  CURLOPT_RETURNTRANSFER => true,
  CURLOPT_BINARYTRANSFER => true,
  CURLOPT_INFILE => $file,
]);

$response = curl_exec($curl);
curl_close($curl);

$result = json_decode($response, true);
print_r($result);
?>