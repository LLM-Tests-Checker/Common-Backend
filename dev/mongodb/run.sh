docker build -t common-backend-mongo .
docker run -d -p 27017:27017 --name common-backend-mongo common-backend-mongo