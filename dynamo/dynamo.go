package dynamo

import (
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func Dbn() string {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// Create table VTResults
	tableName := "VTResults"

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Scan_id"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("positives"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Scan_id"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("positives"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}

	_, err := svc.CreateTable(input)
	if err != nil {
		if strings.Contains(err.Error(), "Table already exists") {
			// The TB is the way we filter the string in main
			// TBE mean Table Exist
			log.Println("[Main-Package] WARN:  When Calling Dynamo Package! Please Check OutPut =>", err)
			return "[Dynamo-Package] WARN: TBE Table already exists!"
		} else {
			// The OE is the way we filter the string in main
			// OE mean Other Error
			log.Fatalf("[Dynamo-Package] FATAL: OE calling CreateTable: %s", err)
			return "[Dynamo-Package] Error: OE Other Error!"
		}
	}
	return "Created the table VTResults"
}

// ################################################ Data Format That Will Be Saved To DynamoDb
// 	   "scan_id": "927f903c8658b68c4fcc2f21fd5df21d64871345b6715a79fe6eda52371d7-1673295888",
//     "sha1": "0f72b87630947d376fb7457d7a0553f88d669aab",
//     "resource": "927f903c8658b68c4fcc25d4d6db276ec9d64871345b6715a79fsdfsfdsfd",
//     "response_code": 1,
//     "sha256": "927f903c8658b68c4fcc25d4d6db276ec9d64871345b6715a79fsdfsfdsfd",
//     "permalink": "https://www.virustotal.com/gui/file/927f903c8658b68c4fcc25d4d6db276ec9d64871345b6715a79fsdfsfdsfd/detection/f-927f903c8658b68c4fcc25d4d6db276ec9d64871345b6715a79fsdfsfdsfd-1673295888",
//     "md5": "66d0ae2b23cdd90crtf7a923d3044fc3",
//     "verbose_msg": "Scan request successfully queued, come back later for the report",
//     "filename": "EpicPhoto.png"
