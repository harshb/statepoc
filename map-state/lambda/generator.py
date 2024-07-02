import json


def lambda_handler(event, context):
    # Generate a list of items
    items = [{"id": i, "value": f"Item {i}"} for i in range(1, 11)]

    # Return the list of items
    return {
        'statusCode': 200,
        'body': json.dumps(items)
    }
