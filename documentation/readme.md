# Table of Contents

- [faq-service-backend-v1-migration](#faq-service-backend-v1-migration)
- [Setting up DATABASE](#setting-up-database)
  - [Start Postgres on Local Machine](#postgres-start-on-local-machine)
  - [Postgres Status Check](#postgres-status-check)
  - [Accessing Postgres in Ubuntu](#accessing-postgres-in-ubuntu)
  - [Create User](#create-user)
  - [Create Database](#create-database)
  - [Change User Password](#change-user-password)
  - [Grant User Permissions](#grant-user-permissions)
  - [Other Useful Commands](#other-useful-commands)
- [Installing and Connecting to Database using DBeaver](#installing-and-connecting-to-database-using-dbeaver)
  - [Installation](#installation)
  - [Connecting to a Database](#connecting-to-a-database)
- [Run](#running-this-repo)
- [Environments](#Environments)


# file-service-backend-v1-migration

This repository contains the migration scripts and seeds for the `file-service` backend.


# Setting up DATABASE

## Postgres Start on Local Machine

To start the Postgres service on your local machine, run the following command:

```shell
sudo service postgresql start
```

## Postgres Status Check

To check the status of the Postgres service, use the following command:

```shell
sudo service postgresql status
```

## Accessing Postgres in Ubuntu

To access the Postgres command line in Ubuntu, use the following command:

```shell
sudo -u postgres psql
```

## Create User

To create a new user in Postgres, run the following command in the Postgres command line:

```shell
create user studymitr;
```

## Create Database

To create a new database in Postgres, run the following command in the Postgres command line:

```shell
create database faq_service_local;
```

## Change User Password

To change the password for a user in Postgres, run the following command in the Postgres command line:

```shell
alter user testDay password 'studymitr';
```

## Grant User Permissions

To grant permissions to a user in Postgres, run the following command in the Postgres command line:

```shell
GRANT admin TO studymitr;
```

## Other Useful Commands

To list all databases, use the following command in the Postgres command line:

```shell
\l
```

To list all users and their roles, use the following command in the Postgres command line:

```shell
\du
```

# Installing and Connecting to Database using DBeaver

This guide will walk you through the process of installing DBeaver and connecting to a database.

## Installation

1. Go to the [DBeaver website](https://dbeaver.io/) and download the appropriate version for your operating system.
2. Run the installer and follow the on-screen instructions to complete the installation.

## Connecting to a Database

1. Launch DBeaver.
2. Click on the "New Connection" button in the toolbar, or go to `Database` -> `New Connection`.

3. In the "New Connection" dialog, select the database type you want to connect to (e.g., PostgreSQL, MySQL, Oracle, etc.) and click "Next".


4. Fill in the connection details:

   - **Host:** The hostname or IP address of the database server.
   - **Port:** The port number on which the database server is running.
   - **Database:** The name of the database you want to connect to.
   - **Username:** Your database username.
   - **Password:** Your database password.


5. Click "Test Connection" to verify that the connection settings are correct and DBeaver can connect to the database successfully. If the test is successful, click "Finish" to save the connection.


6. The new connection will appear in the DBeaver workspace. Double-click on it to establish a connection to the database.


7. You can now explore the database structure, run queries, and perform various database operations using DBeaver.


# Running this repo

this repo can  `drop`, `migrate` and `seed` database.

to drop database run : 

```shell
go run main.go droptables
```

to migrate run : 

```shell
go run main.go migrate
```

to seed run : 
```shell
go run main.go seed
```

Or you can run following command to run all three at once : 

```shell
make run
```


# Environments
	LOCAL-"local"
	DEVELOPMENT- "development"
	UAT-"uat"
	PRODUCTION -"production"