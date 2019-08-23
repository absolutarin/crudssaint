# crudssaint
A highly concurrent CRUD system with REST APIs
 - create channels to process reads
 - create channels to process concurrent writes
 - attach a mongodb to store session IDs
 - ambitious: place a redis cache between requests and responses
 - ambitious: make this serverless
