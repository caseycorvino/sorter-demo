#  Frameworks and drivers:

![alt text](https://github.com/caseycorvino/sorter-demo/blob/master/sorter-architecture/p3.png "Logo Title Text 1")

## AWS
### EC2
Host our website, it's core actions, and our API. API must be constantly live. 

### RDS
Stores our MongoDB database. 

### Redis and Elisticache
Serves as a cache for out MongoDB RDS. Faster data retrieval. Especially important for our API.

### S3
Hosts our static content - most importantly the CSV's. Uploading the CSV triggers the Lambda pipeline. 

### Lambda
Serverless code, only need when CSV is uploaded. Only active on event trigger. This is where the core of our datapipeline lives. 


## CI/CD

### Jenkins
Used for continous deployment when pushing to EC2.

### Travis CI
Used for continous integration when creating PRs.

### Bazel
Used for unit testing. Extremely fast as it only tests changed code.

### Git and Github
Versioning for our repository.

**Need to do more research on CD with lambda**

## Other techonologies
### Github task boards 
Used to facilitate an Agile Scrum environment. Product backlog with Userstories
