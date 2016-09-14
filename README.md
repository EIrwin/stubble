#stubby - Mock JSON API Generator

# What is stubby?
Stubby is a mock JSON API generator that uses a YAML specification to define mock API endpoints and responses.

# Why stubby?
Current API response mocking solutions bloat client and/or server side code. Stubby can be ran 100% from your client and server leaving it clean and free of unecessary bloat.

# Example
  
Stubby uses a simple `YAML` specification to generate a mock JSON API.
 
 ``` 
host: "localhost"

port: "8282"

endpoints:
  - "GET /api/v1/users responses/users_get.json"
  - "POST /api/v1/users responses/users_post.json"
  - "PUT /api/v1/users responses/users_put.json"
  - "GET /api/v1/groups responses/groups_get.json"
  - "POST /api/v1/groups responses/groups_post.json"
  ```

