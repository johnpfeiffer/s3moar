# s3moar
more things you always wanted to do with s3

go mod tidy
go run main.go BUCKETNAME

## Usage

To run the program, use the following command:

```
go run main.go BUCKETNAME
```

Replace `BUCKETNAME` with the name of your S3 bucket.

## Pagination

The program retrieves objects from the specified S3 bucket using pagination. By default, it fetches up to 1000 keys per page. The results are fetched concurrently to improve performance.
