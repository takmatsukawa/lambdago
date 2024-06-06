#!/bin/bash

aws --endpoint-url http://localhost:4566 iam create-role --role-name lambda-common --assume-role-policy-document '{ "Version": "2012-10-17", "Statement": [ { "Effect": "Allow", "Principal": { "Service": "lambda.amazonaws.com" }, "Action": "sts:AssumeRole" } ] }'
aws --endpoint-url http://localhost:4566 lambda create-function --function-name lambdago \
                                              --code S3Bucket="hot-reload",S3Key="/go/src/lambda" \
                                              --handler app \
                                              --runtime go1.x  \
                                              --timeout 90  \
                                              --environment Variables={ENV=local} \
                                              --role arn:aws:iam::000000000000:role/lambda-common | cat -