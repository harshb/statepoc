import boto3
import json
import logging

# Initialize a DynamoDB client
dynamodb = boto3.client('dynamodb')

# Set up logging
logger = logging.getLogger()
logger.setLevel(logging.INFO)


def lambda_handler(event, context):
    # Extract the task token and correlation_id without default values
    task_token = event.get('task_token')
    correlation_id = event.get('correlation_id')

    # Extract other fields with a default value of space
    name = event.get('name', " ")
    component_type = event.get('component_type', " ")
    component_group = event.get('component_group', " ")

    # Log the name or any meaningful identifier
    logger.info(f"Processing item: {name}")

    # Define the DynamoDB table name
    table_name = 'TokenTable'

    # Prepare the item to be inserted into DynamoDB
    item = {
        'TaskToken': {'S': task_token},
        'Name': {'S': name},
        'ComponentType': {'S': component_type},
        'ComponentGroup': {'S': component_group},
        'CorrelationId': {'S': correlation_id}
    }

    # Insert the item into the DynamoDB table
    response = dynamodb.put_item(
        TableName=table_name,
        Item=item
    )

    # Log the response from DynamoDB
    logger.info(f"DynamoDB response: {json.dumps(response)}")

    # Return the response from DynamoDB
    return {
        'statusCode': 200,
        'body': json.dumps('Item saved successfully')
    }