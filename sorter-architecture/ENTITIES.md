# Entities: 
*Written in mgo MongoDB driver for Go format*

I often use "..." because this should be easily and continuously updated. For example, adding fields like lastUpdated to reduce api calls or changing First and Last name to omitempty. 

All entities will use bcrypt and salting for as many fields as possible, especially the fields associated with Profile. 

## Domain Model:

![domain_model](https://github.com/caseycorvino/sorter-demo/blob/master/sorter-architecture/p1.png "domain model" )

### ApiKey
```Golang
type ApiKey struct {
Company         Company      `json:"company" bson:"company,omitempty"`    
PublicKey       string       `json:"public_key" bson:"public_key,omitempty"`    
SecretKey       string       `json:"secret_key" bson:"secret_key,omitempty"`    
...
}
```

### Company
```Golang
type Company struct {
ID              entity.ID    `json:"id" bson:"_id,omitempty"`
Name            string       `json:"name" bson:"name"`
Users           []*User      `json:"users bson:"users"`
CreatedAt       time.Time    `json:"created_at" bson:"created_at"`
...
}
```

### Profile
```Golang
type Profile struct {
ID              entity.ID    `json:"id" bson:"_id,omitempty"`
FirstName       string       `json:"first_name bson:"first_name"`
LastName        string       `json:"last_name bson:"last_name`
Address         string       `json:"address bson:"address"`
Emails          []*string    `json:"emails" bson:"emails,omitempty"`
DeviceIds       []*string    `json:"deviceIds" bson:"deviceIds,omitempty"`
Profiles        []*string    `json:"profiles" bson:"profiles,omitempty"`
Watson          string       `json:"watson  bson:"profiles,omitempty"`
PYCO            string       `json:"pyco  bson:"profiles,omitempty"`
Salt            string       `json:"salt" bson:"salt,omitempty"`
CreatedAt       time.Time    `json:"created_at" bson:"created_at"`
...
}
```

### Records
```Golang
type Record struct {
ID              entity.ID   `json:"id" bson:"_id,omitempty"`
s3Path          string      `json:"s3_path bson:"s3_path"`
CreatedAt       time.Time   `json:"created_at" bson:"created_at"`
...
}
```

### User
```Golang
type User struct {
ID              entity.ID    `json:"id" bson:"_id,omitempty"`
Email           string       `json:"email" bson:"email"`
Password        string       `json:"password" bson:"password,omitempty"`
Salt            string       `json:"salt" bson:"salt,omitempty"`
Company         Company      `json:"company" bson:"company,omitempty"`
ApiKeys         []*ApiKey    `json:"api_keys" bson:"api_keys,omitempty"`
CreatedAt       time.Time    `json:"created_at" bson:"created_at"`
...
}
```
