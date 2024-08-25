# Microservices with Golang

This repository contains examples of various microservice architectures using Golang. The aim is to provide a practical guide for building, deploying, and managing microservices-based systems.

## Request

- **URL**: `localhost:8080/orders`
- **Method**: `POST`
- **Possible message receivers**: nats,kafka,rest,rabbitmq
- **Body**:

```json
    {
        "order_id": "1",
        "send_by": "rabbitmq"
    }
```

- **Response**:

```json
{
    "message": "order created successfully"
}
```

- **CURL request**

```bash
curl --location 'localhost:8080/orders' \
--header 'Content-Type: application/json' \
--data '{
    "order_id": "1",
    "send_by": "rabbitmq"
}'
```
