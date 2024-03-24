# it-revolution-test backend

## Endpoints

All bodies must be with header `Content-Type: application/json`

### `POST /api/transform` 200 - Short link

Body:

```json
{
  "original_link": "string"
}
```

Result (text/plain):

```
"short link"
```

### `GET /api/original/:id` 200 - Get original link

Result (text/plain):

```
"original link"
```

### `GET /api/statistics` 200 - Get statistics for all shortened links

Result (application/json):

```json
[
    {
        "short_link": "string",
        "count": "int"
    }
]
```