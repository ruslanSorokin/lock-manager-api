<h1><a id="summary" class="anchor" aria-hidden="true"></a>Interprocess Communication Lock Manager API</h1>

[`Application`](https://github.com/ruslanSorokin/lock-manager)

## RPC

### Lock resource

#### Request

```protobuf
// LockReq is a request for Lock call.
message LockReq{
  // Required.
  optional string resource_id = 1;
}
```

#### Response

```protobuf
// LockRes is a response for LockReq.
message LockRes{
  // Error is a enum error. Sent in status error message if any.
  enum Error{
    ERR_UNSPECIFIED = 0;
    // ERR_INVALID_RESOURCE_ID means that resource_id violates one of its constraint
    ERR_INVALID_RESOURCE_ID = 1;
    // ERR_LOCK_ALREADY_EXISTS means that resource_id is already locked
    ERR_LOCK_ALREADY_EXISTS = 2;
  }
  // Required.
  optional string token = 1;
}
```

---

### Unlock resource

#### Request

```protobuf
// UnlockReq is a request for Unlock call.
message UnlockReq{
  // Required.
  optional string resource_id = 1;
  // Required.
  optional string token = 2;
}

```

#### Response

```protobuf
// UnlockRes is a response for UnlockReq.
message UnlockRes{
  // Error is a enum error. Sent in status error message if any.
  enum Error{
    ERR_UNSPECIFIED = 0;
    // ERR_INVALID_TOKEN means that token violates one of its constraint
    ERR_INVALID_TOKEN = 1;
    // ERR_WRONG_TOKEN means that token doesn't fit
    ERR_WRONG_TOKEN = 2;
    // ERR_INVALID_RESOURCE_ID means that resource_id violates one of its constraint
    ERR_INVALID_RESOURCE_ID = 3;
    // ERR_LOCK_NOT_FOUND means that lock with given resource_id was not found
    ERR_LOCK_NOT_FOUND = 4;
  }
  google.protobuf.Empty response = 1;
}
```

---

## REST

### Lock resource

Swagger server is on

#### Request

```yaml
paths:
  /locks:
    post:
      tags:
        - Lock
      summary: >
        Lock a resource.
      description: >
        Creates a new lock with given `resource_id` and returns a token to
        delete the lock.
      parameters:
        - name: resource_id
          in: query
          required: true
          schema:
            allOf:
              - $ref: "#/components/schemas/resource_id"
```

#### Response

```yaml
responses:
  "201":
    description: >
      Lock on the resource was successfully created. You'll get a token
      that can be used to delete the resource and HATEOAS format links in the response body.
    content:
      application/json:
        schema:
          type: object
          properties:
            token:
              allOf:
                - $ref: "#/components/schemas/token"
            _links:
              type: array
              items:
                type: object
                properties:
                  unlock:
                    type: object
                    properties:
                      href:
                        type: string
                        format: url
                        default: http://localhost:9090/api/v1/unlock/{resource_id}
                      type:
                        type: string
                        enum:
                          - get
                          - post
                          - put
                          - delete
                        default: delete

  "400":
    description: >
      Lock on the resource was NOT successfully created, because of the invalid
      parameters of the request. You'll get error in format of predefined enum.
    content:
      application/json:
        schema:
          type: object
          properties:
            error:
              description: >
                Enum format error.
              type: string
              format: enum
              enum:
                - INVALID_RESOURCE_ID # resource_id violates one of its constraint

  "409":
    description: >
      Lock on the resource was NOT successfully created, because lock on
      the given resource already exists. You'll get error in format of predefined enum.
    content:
      application/json:
        schema:
          type: object
          properties:
            error:
              description: >
                Enum format error.
              type: string
              format: enum
              enum:
                - RESOURCE_ID_ALREADY_EXISTS # resource_id is already locked
```

---

### Unlock resource

#### Request

```yaml
paths:
  /locks/{resource_id}:
    delete:
      tags:
        - Unlock
      summary: >
        Unlock a resource.
      description: >
        Removes lock with given `resource_id` only if the token fits.
      parameters:
        - name: resource_id
          in: path
          required: true
          schema:
            allOf:
              - $ref: "#/components/schemas/resource_id"
        - name: token
          in: query
          required: true
          schema:
            allOf:
              - $ref: "#/components/schemas/token"
```

#### Response

```yaml
responses:
  "200":
    description: >
      Lock was successfully deleted.

  "400":
    description: >
      Lock on the resource was NOT successfully deleted, because of the invalid
      parameters of the request. You'll get error in format of predefined enum.
    content:
      application/json:
        schema:
          type: object
          properties:
            error:
              description: >
                Enum format error.
              type: string
              format: enum
              enum:
                - INVALID_TOKEN # token violates one of its constraint
                - WRONG_TOKEN # token doesn't fit
                - INVALID_RESOURCE_ID # resource_id violates one of its constraint

  "404":
    description: >
      Lock on the resource doesn't exist. You'll get error in format of predefined enum.
    content:
      application/json:
        schema:
          type: object
          properties:
            error:
              description: >
                Enum format error.
              type: string
              format: enum
              enum:
                - RESOURCE_ID_NOT_FOUND # resource_id was not found
```
