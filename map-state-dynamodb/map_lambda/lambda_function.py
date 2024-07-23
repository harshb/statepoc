import json


def lambda_handler(event, context):
    correlation_id = event.get('correlation_id',
                               'default_id')  # Get correlation_id from event, use 'default_id' if not found

    items = {
        "component_1": [{"name": "lambdarole", "component_type": "Lambda Role", "component_group": "Group 1",
                         "correlation_id": correlation_id}],
        "component_2": [
            {"name": "DynamoDb", "component_type": "DynamoDb", "component_group": "Group 2", "correlation_id": correlation_id}],
        "component_3": [
            {"name": "Lambda", "component_type": "lambda", "component_group": "Group 3", "correlation_id": correlation_id},
            {"name": "NATGateway", "component_type": "skip", "component_group": "Group 3", "correlation_id": correlation_id},
            {"name": "PrivateVPCEndpoints", "component_type": "skip", "component_group": "Group 3",
             "correlation_id": correlation_id},
            {"name": "LambdaBase", "component_type": "skip", "component_group": "Group 3", "correlation_id": correlation_id}
        ],
        "component_4": [
            {"name": "empty", "component_type": "skip", "component_group": "Group 4", "correlation_id": correlation_id}],

        "component_5": [
            {"name": "NATGateway", "component_type": "skip", "component_group": "Group 5", "correlation_id": correlation_id},
            {"name": "PrivateVPCEndpoints", "component_type": "skip", "component_group": "Group 5",
             "correlation_id": correlation_id},
            {"name": "LambdaBase", "component_type": "skip", "component_group": "Group 5", "correlation_id": correlation_id},
        ]
    }

    return {
        'statusCode': 200,
        'body': {"items": items}
    }
