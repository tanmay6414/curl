## Update made in application
1. Application is not running properly, giving panic error so updated some code for sending mail
2. Also application does not contain any code for starting webserver, added this code to start webserver while starting 
## Working with Dockerfile
####SS Completing below task
1. Use latest stable ```golang:1.17.2-alpine``` image as base image which is of much smaller in size.
```
FROM golang:1.17.2-alpine as builder
```
2. Build multi build Dockerfile
3. Also make use of ```scratch```  image.
```
FROM scratch
```
4. Also set default env variable for sender username, receiver username and sender mail password. you can override it during runtime.
5. Use entrypoint to ```RUN``` executable.