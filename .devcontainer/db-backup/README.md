# DATABASE BACKUP RESTORE UTILITY TOOL
This utility tool provides the user to make backup and restore of supported databases i.e. MySql, Postgres, Mariadb, MongoDb

### Prerequisites
 Lists : 
 - Needs docker installed
 - After installing docker run the docker compose up command
 - To use the tool run docker exec -it db-backup sh, which will launch interactive shell and allow you to use the tool without downloading any prerequisite dependencies


### Installing
<pre>
db-backup comes pre installed in the docker container, you will be able to use db-backup command as soon as you run docker exec -it db-backup sh
## Usage
  db-backup [command]  

Available Commands:
  backup      backup is a command to backup databases
  help        Help about any command
  list        list is a command to list backups for a specific DBMS
  restore     restore is a command to restore databases

 db-backup backup [flags] is a command to backup databases.
It supports various DBMS like Postgres, MongoDB, MySQL, and MariaDB.
You can specify the DBMS using the --dbms flag.
        Example: db-backup backup --dbms postgres


Flags:
  -d, --dbms          Database Management System (DBMS) name (e.g., postgres, mongodb, mysql, mariadb)
  -h, --help          help for backup

db-backup restore [flags] is a command to restore databases.
It supports various DBMS like Postgres, MongoDB, MySQL, and MariaDB.
You can specify the DBMS using the --dbms flag.
Example: db-backup restore --dbms postgres

Usage:
  db-backup restore [flags]

Flags:
  -d, --dbms string   DBMS name
  -h, --help          help for restore
</pre>
## License
This project is licensed under the GNU GENERAL PUBLIC LICENSE Version 2, see the LICENSE file for details.
