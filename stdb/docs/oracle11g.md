# oracle 11g 操作手册

## 部署

```shell
docker-compose up -d oracle11g

docker-compose exec oracle11g bash

create user springboot identified by 123456;
GRANT ALL PRIVILEGES TO springboot;
COMMIT;
```