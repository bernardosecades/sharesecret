# ShareSecret

[![Test](https://github.com/bernardosecades/sharesecret/workflows/Test/badge.svg)](https://github.com/bernardosecades/sharesecret/actions)
[![Super-Linter](https://github.com/bernardosecades/sharesecret/workflows/Super-Linter/badge.svg)](https://github.com/bernardosecades/sharesecret/actions)

ShareSecret is a service to share sensitive information that's both simple and secure.

If you share some text will be display it once and then delete it. After that it's gone forever.

We keep secrets for up to 5 days.

## Why should I trust you?

General we can't do anything with your information even if we wanted to (which we don't). If it's a password for example, we don't know the username or even the application that the credentials are for.

If you include a password, we use it to encrypt the secret. We don't store the password (only a crypted hash) so we can never know what the secret is because we can't decrypt it.


## Docker
 
Build image with tag version:

`sudo docker build --tag bernardosecades/api-share-secret:v1`

Build container from image:

`docker run -p 8080:8080 -e SECRET_KEY=11111111111111111111111111111111 -e DEFAULT_PASSWORD=@myPassword -e MONGODB_URI=mongodb://root:example@192.168.1.132:27017 bernardosecades/api-share-secret:v1`

Push to docker hub:

`sudo docker login -u bernardosecades -p <YOUR_PASSWORD>`

`sudo docker push bernardosecades/api-share-secret:v1`