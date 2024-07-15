import json


def lambda_handler(event, context):
    # Define a list of items for ECS template
    ecs_template_items = [
        {"id": 1, "name": "ecs"},
        {"id": 2, "name": "sqs"},
        {"id": 3, "name": "elasticache"},
    ]

    # Define a list of general items
    items = [
        {"id": 1, "name": "eks"},
        {"id": 2, "name": "dynamodb"},
        {"id": 3, "name": "s3"}
    ]

    # Check if the 'template' parameter is provided and equals 'ecs'
    if 'template' in event and event['template'] == 'ecs':
        response_items = ecs_template_items
    else:
        response_items = items

    # Return the appropriate items based on the template parameter
    return {
        'statusCode': 200,
        'body': {"items": response_items}
    }
