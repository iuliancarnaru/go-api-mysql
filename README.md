docker run -d --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password --restart=always mysql

mysql -u root -p
CREATE DATABASE projectmanager;
SHOW DATABASES;

## Api diagram

![alt text](https://github.com/iuliancarnaru/go-api-mysql/blob/main/assets/Screenshot%202024-03-01%20101244.png "schema")

