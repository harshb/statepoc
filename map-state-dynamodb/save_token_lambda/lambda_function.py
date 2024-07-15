import boto3
import json
import logging

# Initialize a DynamoDB client
dynamodb = boto3.client('dynamodb')

# Set up logging
logger = logging.getLogger()
logger.setLevel(logging.INFO)


def lambda_handler(event, context):
    # Extract the task token and other relevant information from the event
    task_token = event.get('taskToken')
    name = event.get('name')
    # Assuming other data from the event might be stored as well
    item_data = event.get('data', {})

    # Log the name or any meaningful identifier
    logger.info(f"Processing item: {name}")

    # Define the DynamoDB table name
    table_name = 'TokenTable'

    # Prepare the item to be inserted into DynamoDB
    item = {
        'TaskToken': {'S': task_token},
        'Name': {'S': name}
        # Add other attributes from item_data if necessary
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
        'body': json.dumps(response)
    }
