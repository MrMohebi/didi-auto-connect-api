<?php
namespace Mrmmsi\DidiAutoConnectApi\configs;

include_once "env.php";
require_once __DIR__ . '/../../vendor/autoload.php';

use MongoDB\Client;

class Models
{

    public Client $client;

    public function __construct(){
        $this->client = new Client(
            "mongodb://".$_ENV['DB_USER'].":".$_ENV['DB_PASS']."@".$_ENV['DB_URL']."/?retryWrites=true&w=majority"
        );
    }

    public function Users(){
        return $this->client->{$_ENV['DB_DBNAME']}->users;
    }

    public function Devices(){
        return $this->client->{$_ENV['DB_DBNAME']}->devices;
    }
}