---
title: Delete GDrive File
weight: 1
---

# Counter
This activity allows you to delete a file from Gdrive account.

## Installation
### Flogo Web
This activity comes out of the box with the Flogo Web UI
### Flogo CLI
```bash
flogo add activity github.com/DipeshTest/gdrivedelete
```

## Schema
Inputs and Outputs:

```json
"inputs":[
  {
"name": "accessToken",
"type": "string",
  "required": true
},
{
"name": "fileId",
"type": "string",
"required": true

},{
"name": "timeout",
"type": "string",
"value":"120"
} ],
"outputs": [
  {
    "name": "statusCode",
    "type": "string"
  },
  {
    "name": "message",
    "type": "any"
  }
]
}
```
## Settings
| Setting     | Required | Description |
|:------------|:---------|:------------|
| accessToken | True     | The access token for your account |         
| fileId   | False    | File Id of the file to delete |
| timeout       | False    | Timeout value for the delete call, default value is 120 seconds|

## Examples
### Increment
The below example for a sample delete:

```json

```