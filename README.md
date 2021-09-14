# Storj-REST-API

<img src="https://github.com/mohammedfajer/Storj-REST-API/blob/main/2021-09-14%2017-08-49.gif"/>


© 2020/2021 The University of Manchester and Mohammed Akram Fajer


Presenting Master Project conducted at the University of Manchester. Reference this project when you choose to use any of the ideas, source code, images or other material found in this repository.

#### LaTeX Reference:

```
@misc{fajer21master,
author         = {Mohammed Akram Fajer},
title          = {{Blockchain Technology Based Patient-Centric Electronic Health Record}},
submissiondate = {17/09/2021},
year           = {2021},
url            = {{https://github.com/mohammedfajer/EHRethChain}},
note           = {School of Computing, The University of Manchester}
}
```



## Install Instruction

### Dependencies
- Go Language

### Tech Stack
- PostgreSQL - database (storing data)
- GORM - golang library for communicating with database
- Env vars - storing sensitive information
- Gorrilla/Mux - golang library to serve APIs
- Postmen - test API
- Insomina - test API as shown in the gif video above

### Run API
MacOS installation
Prerequisites:
- Install homebrew


1. Install postgresql

```bash
$ brew tap homebrew/services
$ brew install postgresql
$ initdb /usr/local/var/postgres
```

2. Make new password for using the database

```bash
$ sudo passwd postgres
New password: AEHRethChain
Retype new passowrd: AEHRethChain
```
3. Run database

```bash
$ sudo brew services start postgresql
```

4. Shell into database

```bash
$ sudo -u postgres psql
```

5. Final setup steps

```
$ postgres=# ALTER USER postgres PASSWORD 'AEHRethChain'
$ postgres=# CREATE DATABASE dapp_users;
$ postgres=# \q
```

6. Run golang 

```
$ go run main.go 
```



