// Lambda in Go
// @jeffotoni
// 2019-09-17

package main

import (
    "context"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "log"
)

var SIZE = int64(70000) // 70kb

type MsgResponse struct {
    Message string `json:"Answer:"`
}

func handlerS3(ctx context.Context, s3Event events.S3Event) (MsgResponse, error) {

    var msg string = "Events S3."
    for _, record := range s3Event.Records {

        rs3 := record.S3
        log.Printf("[%s - %s] Bucket = %s, Key = %s Size = %d\n", record.EventSource, record.EventTime, rs3.Bucket.Name, rs3.Object.Key, rs3.Object.Size)

        if rs3.Object.Size > SIZE {
            svc := s3.New(session.New())
            _, err := svc.DeleteObject(&s3.DeleteObjectInput{
                Bucket: aws.String(rs3.Bucket.Name),
                Key:    aws.String(rs3.Object.Key)})
            if err != nil {
                log.Printf("\nNão é possível excluir a chave %q do bucket %q, %v", rs3.Object.Key, rs3.Bucket.Name, err)
                continue
            }

            log.Println("Arquivo maior que o permitido: key: ", rs3.Object.Key, "excluido com [sucesso]")
        }
    }

    return MsgResponse{Message: msg}, nil
}

func main() {
    lambda.Start(handlerS3)
}
