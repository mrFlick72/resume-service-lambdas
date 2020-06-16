import boto3
import json

def lambda_handler(event, context):
    s3 = boto3.resource('s3')
    obj = s3.Object(event['bucket'], event['bucket-key'])
    content = obj.get()['Body'].read().decode('utf-8')

    return {
        'statusCode': 200,
        'body': json.loads(content)
    }
