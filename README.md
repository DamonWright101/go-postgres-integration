users generated using the following script

```
for i in $(seq -w 1 100); do echo "INSERT INTO persons(id, email, username,fname,lname) VALUES ((SELECT COALESCE(MAX(id) + 1, 1) FROM persons),'test$i@email.com','username$i','test$i','test');"; done > postgres/data.sql
```