Send emails with golang in ~50 lines of code. This repo shows you how to send emails with golang using the net/smtp package. For the full tutorial check out this [medium article](https://medium.com/@toshvelaga/sending-emails-with-golang-and-aws-lambda-functions-ff93329dac43). To run the code, you need to clone the repo, and create an .env file with the variables from .env.example. Then use the following command to actually run the code:

```
go run .
```

NOTE: if you want the AWS lambda function code checkout the AWS branch in this repo. For that you will need to use the build.sh script to compile and zip your code for aws. You can just run the script with the following command:

```
sh build.sh
```
