<?php

require_once "../configs/Models.php";

use Mrmmsi\DidiAutoConnectApi\configs\Models;

if(isset($_POST['username'])&&isset($_POST["password"])&&isset($_POST['deviceHash'])){
    $password = password_hash($_POST["password"], PASSWORD_DEFAULT);
    $username = $_POST['username'];
    $device_hash = $_POST['deviceHash'];
    $token = bin2hex(openssl_random_pseudo_bytes(16));

    $models = new Models();
    $users = $models->Users();
    $devices = $models->Devices();

    if($users->findOne(['username' => $username])){
        exit("409");
    }

    $user = $users->insertOne([
        'username' => $username,
        'password' => $password,
        'token'=>$token,
        "lastLogin"=>time(),
        "createdAt"=>time()
    ]);

    $device = $devices->insertOne([
        'userID' => $user->getInsertedId(),
        'hash' => $device_hash,
        "lastLogin"=>time(),
        "createdAt"=>time()
    ]);

    $user = $users->updateOne(
        ['_id' => $user->getInsertedId()],
        ['$set' => ['activeDeviceID' => $device->getInsertedId()]]
    );

    exit(json_encode([
        "token"=>$token,
        "hasAccess"=>true,
    ]));
}else{
    exit('400');
}