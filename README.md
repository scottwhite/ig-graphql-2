### round 2 graphql with graphql-go
heavily used
https://github.com/deltaskelta/graphql-go-pets-example


### Diffs with gqlgen
Nothing is generated, you wire it up, resolvers.go is that 

Tells you want is not implemented based off the schema so it's schema first like gqlgen
Simpiler, cause you make the cake pretty much.


### Files
dao -> data access layer
models -> models.. duh
resolvers.go -> matches up the models with what the schema wants
Resolver Structs (just naming conventions, nothing special with the name)
Resolver -> root that gets hooked up with the daos
CiResolver -> mapper struct to the Ci model in the schema
LocationResolver -> Return wrapper for Location model in the schema for Ci.location



### special types
ID -> graphql.ID
All fields using ID in the schema have to return this
Any field optional needs to return a pointer (similar to gqlgen)

### firing it up
Create config.toml file
ex:

```
[db]
host = "localhost"
port = 5432
database = "postgres"
username = "bacon"
password = "pants"
```


```
go run .
```

if it all works you should see

```
Successfully connected!
INFO[0000] starting server                               fields.time="2020-02-11 19:28:33.022083949 -0500 EST m=+0.003508368"

```

Go to http://localhost:8080


Sample query (db has to exists and be populated)

```
{ getCis{name,isCloud,location{name}}}
```


mutations

create conductor
```json
mutation{
 createConductor(conductor:{
  name:"bacon pants",
  organizationId:"00000000-0000-0000-0000-000000000000",
  fingerprint:"asdfasdfasdfasdfasdf",
  ipAddress:"127.0.0.1",
  port:22
}){
  id,
  ipAddress,
  name
}
}
```


create schedule
```json
mutation{
 createSchedule(schedule:{
  name:"bacon pants sched",
  organizationId:"00000000-0000-0000-0000-000000000000",
  conductorId:"228ecd37-4d21-494f-9062-c23043c51e30",
  interval:1581487323,
}){
  id,
  interval,
  name
}
}
```

