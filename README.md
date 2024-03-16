# Go AWS Dynamo with Docker Compose

This repository is a simple example project using DynamoDB with the Go AWS SDK V2 and Docker Compose.

## Prerequisites
- Docker and Docker Compose
- Go

## Steps to Execute
1. Install Go dependencies specified in the `go.mod`: `go mod download`
2. Run `docker-compose up` to start the DynamoDB container.
3. Run `go run ./cmd/main.go` to execute the Go application.

### To stop and remove Docker containers
`docker-compose down`

## Explanation
The setup involves the creation of the database structure in the `SetupDatabase` function. It checks if the `MainSingleTable` exists; if not, it creates the table. The table structure consists of a 'pk' and an 'sk' following the [SingleTable Desing](https://www.alexdebrie.com/posts/dynamodb-single-table/), with the example entity User having the key definition:
- pk: USER#username
- sk: USER#username

After table is created, three users are inserted using 'PutItem'. Finally, the user with the username 'seconduser' is retrieved using 'GetItem'.

## How to Check the Data
- You can install [NoSQL Workbench](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/workbench.html) for DynamoDB for a comprehensive view.
- The Docker Compose runs a Node image with DynamoDB Admin GUI installed. Access it at 'localhost:8001'.

## References
- [AWS SDK for Go V2 Documentation](https://aws.github.io/aws-sdk-go-v2/docs/)
- [AWS SDK for Go V2 GitHub Repository](https://github.com/aws/aws-sdk-go-v2)
- [AWS SDK for Go V2 Examples](https://github.com/awsdocs/aws-doc-sdk-examples/tree/main/gov2/dynamodb)
- [Deploying DynamoDB locally on your computer](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DynamoDBLocal.DownloadingAndRunning.html)
- [DynamoDB Admin GUI](https://github.com/aaronshaf/dynamodb-admin)