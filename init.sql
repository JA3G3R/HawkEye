-- Create the "bounty" database
CREATE DATABASE IF NOT EXISTS bounty;

-- Switch to the "bounty" database
USE bounty;

-- Create the "targets" table
CREATE TABLE IF NOT EXISTS targets (
	TargetId INT AUTO_INCREMENT PRIMARY KEY,
    NAME varchar(100) NOT NULL,
    PLATFORM varchar(100) DEFAULT 'Independent',
    SCOPE varchar(5000),
    OUTSCOPE varchar(5000),
    PORTS varchar(1000),
    WEBSERVER varchar(100),
    CLOUD varchar(100),
    KNOWNDBS varchar(100),
    WILDCARD_SCOPE bit(1)
);
