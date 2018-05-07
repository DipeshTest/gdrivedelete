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
| fileId   | True    | File Id of the file to delete |
| timeout       | False    | Timeout value for the delete call, default value is 120 seconds|

## Examples
### Increment
The below example for a sample delete:

```json
{
	"accessToken": "ya29.GlurBW7n5A2Fk_rstX9KMVeXLEOT4k0OhmSnF_w7626K9kgKmempF_xTDJ6uQVMkdWWWIMiNcb-ht6Rv9cnhsUb2VhtF9h7nltFw0iniwp10dmDQsFT49giOqFR8",
	"fileId": "1M1mRDaQKzl6N_V6_6WqzlIoe4mTVKA38",
	"timeout": "120"
}
```


## Response Codes
### Google Drive Create
| ResponseCode     | Type | Description |
|:------------|:---------|:------------|
|200 |OK| The request was successful. |
|106 |INVALID INPUT| File ID is not specified.|
|401 |AUTHENTICATION ERROR| Invalid Access Token.|
|404 |SERVER ERROR| File Not Found.|