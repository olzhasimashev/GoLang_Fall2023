# Readme

**Author:** Imashev Olzhas

## API Commands and Responses

### Health Check
```bash
$ curl -i localhost:4000/v1/healthcheck
```
**Output:**
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 19 Oct 2023 06:20:05 GMT
Content-Length: 102
{
        "status": "available",
        "system_info": {
                "environment": "development",
                "version": "1.0.0"
        }
}
```

### Get Blender by ID
```bash
$ curl -i http://localhost:4000/v1/blenders/123
```
**Output:**
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 19 Oct 2023 06:21:41 GMT
Content-Length: 215
{
        "blender": {
                "id": 123,
                "name": "Blender Experia 3000",
                "capacity": "3 litres",
                "material": "Alluminium",
                "categories": [
                        "blender",
                        "cooking applience",
                        "electronics"
                ],
                "version": 1
        }
}
```

### Create a New Blender
```bash
$ BODY='{"name":"Bosch Blender 3000","year":2022,"capacity":"3 litres","categories":["blenders","electronics"]}'
$ curl -i -d "$BODY" localhost:4000/v1/blenders
```
**Output:**
```
HTTP/1.1 200 OK
Date: Thu, 19 Oct 2023 06:17:53 GMT
Content-Length: 81
Content-Type: text/plain; charset=utf-8
{Name:Bosch Blender 3000 Year:2022 Capacity:3 Categories:[blenders electronics]}
```

### Database Migrations
```bash
$ migrate -path=./migrations -database=$GREENLIGHT_DB_DSN up
$ migrate -path=./migrations -database=$EXAMPLE_DSN down
```

