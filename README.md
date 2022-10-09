# Basic Blog App REST API With Go
> Basic Blog App Example to learn golang. Usage packages Gorm, gin-gonic
***

### to run the REST Api follow the following steps;

- run: `$ go run .`

### Project will start port `:8080`

***

## Endpoints

#### list posts

```
curl --location --insecure  --request GET 'http://localhost:8080/post/list'
```

#### create post

```
curl --location --insecure  --request POST 'http://localhost:8080/post/create' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "example post title",
    "content": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry'\''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.",
    "status": 1,
    "user_id": 1,
    "category_id": 1
}'
```

***

### Packages

- Gorm: https://gorm.io/
- Gin-Gonic: https://github.com/gin-gonic/gin
