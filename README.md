# coffeeshop-be

.env file

```env
PORT=3000

# DB
DB_USERNAME=(your_db_username)
DB_PASSWORD=(your_db_password)
DB_HOST=(your_db_host)
DB_PORT=(your_db_port)
DB_NAME=(your_db_name)

DB_URL=${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?charset=utf8mb4&parseTime=True&loc=Local
```

DB_URL reference: https://gorm.io/docs/connecting_to_the_database.html