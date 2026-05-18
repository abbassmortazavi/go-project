# Create users Table

```sql
drop table if exists users;
```


````sql
create table users (
    id serial primary key,
    age int,
    first_name text,
    last_name text,
    email varchar(255) unique not null
);
```