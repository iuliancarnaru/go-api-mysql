docker run -d --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password --restart=always mysql

mysql -u root -p
CREATE DATABASE projectmanager;
SHOW DATABASES;
