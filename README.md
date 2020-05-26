# Skeleton Go Web Server

A small repo useful to fork and start your new web service in Go.
It is opinionated and intended to act as an example of best practices.

## What's implemented?

A web server which:

- Exposes one method (/api/answer), which allows GET and POST and responds with JSON

```
GET  /api/answer               // 'Hello!'
POST /api/answer 'Stefano'     // 'Hello, Stefano!'
```

- Has tests

## Design choices at a glance

- Server shuts down gracefully
- Routes go in their own file

## Contributing

Submit a pull request with:

1. New functionality or edits.
2. Tests covering the above.
3. Reasoning as to why it should be included.
