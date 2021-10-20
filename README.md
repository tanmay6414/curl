## Update made in application
1. Application is not running properly, giving panic error so updated some code for sending mail
2. Also application does not contain any code for starting webserver, added this code to start webserver while starting 
## Working with Dockerfile
### Completing below task
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

### Build and Run Dockerfile
1. Move to the root directory of this repository.
2. Open terminal and apply below command
```
docker build -t <REPO>/<IMAGE-NAME>:<TAG> .
```

3. For running this newly created image 
```
docker run  -p 80:80  --env SENDER_USERNAME="valid email address" -e SENDER_PASSWORD="valid password for this email" -e RECEIVER_MAIL="valid sender email address" <REPO>/<IMAGE-NAME>:<TAG>
``` 
4. Additionally you can pull image push by me on my dockerhub, this is public repo do not required any authentication.
```
docker run  -p 80:80  --env SENDER_USERNAME="valid email address" -e SENDER_PASSWORD="valid password for this email" -e RECEIVER_MAIL="valid sender email address" tanmay8898/curl:1.0
```
> :warning: **Tou need to enable less secure app if using gmail**: https://support.google.com/accounts/answer/6010255?hl=en
