<?php

require_once "../configs/Models.php";
use Mrmmsi\DidiAutoConnectApi\configs\Models;

if(isset($_POST['username'])&&isset($_POST["password"])&&isset($_POST['deviceHash'])){
    $username = $_POST['username'];
    $deviceHash = $_POST['deviceHash'];


    $models = new Models();
    $users = $models->Users();
    $devices = $models->Devices();

    $isLimit = false;
    $token = bin2hex(openssl_random_pseudo_bytes(16));

    if($user = $users->findOne(['username' => new MongoDB\BSON\Regex("^$username$", 'i')])){
        if(!password_verify($_POST["password"], $user->password)){
            exit("401");
        }

        $users->updateOne(
            ['_id' => $user->_id],
            ['$set' => ['token' => $token, 'lastLogin'=>time()]]
        );

        $device = $devices->findOne(['hash' => $deviceHash]);
        //create new device
        if(!$device){
            $device = $devices->insertOne([
                'userID' => $user->_id,
                'hash' => $deviceHash,
                "isActive"=>false,
                "lastLogin"=>null,
                "createdAt"=>time()
            ]);
            $device = $devices->findOne(['_id' => $device->getInsertedId()]);
        }

        // check if it has limitation
        if(!$device->isActive){
            $activeDevice = $devices->findOne(['userID' => $user->_id, 'isActive'=>true]);
            $LIMITATION_TIME = 5 * 60 * 60;
            if($activeDevice && $activeDevice->lastLogin > (time()-$LIMITATION_TIME)){
                $isLimit = true;
            }
        }else{
            $device->updateOne(
                ['userID' => $user->_id, 'isActive'=>true],
                ['$set' => ["isActive"=>false]]
            );
            $device->updateOne(
                ['_id' => $device->_id],
                ['$set' => ["isActive"=>true,'lastLogin'=>time()]]
            );
        }
    }else{
        $password = password_hash($_POST["password"], PASSWORD_DEFAULT);

        $user = $users->insertOne([
            'username' => $username,
            'password' => $password,
            'token'=>$token,
            "lastLogin"=>time(),
            "createdAt"=>time()
        ]);
        $device = $devices->insertOne([
            'userID' => $user->getInsertedId(),
            'hash' => $deviceHash,
            "isActive"=>true,
            "lastLogin"=>time(),
            "createdAt"=>time()
        ]);
    }


    exit(json_encode([
        "token"=>$token,
        "hasAccess"=>!$isLimit,
        "isLimit"=>$isLimit,
    ]));
}else{
    exit('400');
}