---
title: Delete GDrive File
weight: 1
---

# Delete GDrive File
This activity allows you to delete a file from Gdrive account. This activity is built by Team AllStars

## Installation

### Flogo CLI
```bash
flogo install github.com/DipeshTest/gdrivedelete
```

### Third-party libraries used
- #### package drive - "google.golang.org/api/drive/v3":
	Package drive provides access to the Drive API. For more details, check https://developers.google.com/drive/
- #### package googleapi - "google.golang.org/api/googleapi":
	Package googleapi contains the common code shared by all Google API libraries.

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
Please refer activity_test.go

## Response Codes
### Google Drive Create
| ResponseCode     | Type | Description |
|:------------|:---------|:------------|
|200 |OK| The request was successful. |
|106 |INVALID INPUT| File ID is not specified.|
|401 |AUTHENTICATION ERROR| Invalid Access Token.|
|404 |SERVER ERROR| File Not Found.|