# findnusapi
API documentation for FindNUS backend services. Handles the retrieval, processing and management of Lost Items found in NUS.

## Version: 0.1-210522

### /debug/ping

#### GET
##### Description:

Returns a 'Hello World' equivalent message. Shows that the backend connection works.


##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | A hello world message. |

### /debug/checkAuth

#### GET
##### Description:

Check with backend if the Firebase token is valid.


##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | AUTH OK |
| 401 | AUTH NOT OK |

### /item/new

#### POST
##### Description:

Submit a new item to be stored into the database.


##### Responses

| Code | Description |
| ---- | ----------- |
| 201 | Item registered into database |
| 400 | Rejected new item |

##### Security

| Security Schema | Scopes |
| --- | --- |
| firebaseAuth | |

### /item/get/{itemId}

#### GET
##### Description:

Get a particular item's full details


##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| itemId | path | Item Id reference | Yes | long |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | A Lost Item's details |
| 404 | Item not found |

##### Security

| Security Schema | Scopes |
| --- | --- |
| firebaseAuth | |

### /item/peek

#### GET
##### Description:

Get a list of lost items that can be sorted.
Peek at the database's latest finds.
(Sorting and filtering to be implemented in the future)


##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | Returns an array of lost items that may be filtered |

### /item/request

#### POST
##### Description:

Add a lost item request to the server. 


##### Responses

| Code | Description |
| ---- | ----------- |
| 201 | Item lookout request added to server. |

##### Security

| Security Schema | Scopes |
| --- | --- |
| firebaseAuth | |

### /search

#### GET
##### Description:

Elastisearch for an item.


##### Responses

| Code | Description |
| ---- | ----------- |
| 501 | Function not added yet. |

### Models


#### Item

A Lost Item's full details

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | integer |  | Yes |
| name | string |  | Yes |
| date | dateTime |  | Yes |
| location | string |  | Yes |
| category | integer |  | Yes |
| contact_method | integer |  | Yes |
| contact_details | string |  | Yes |
| item_details | string |  | Yes |
| image_url | byte |  | Yes |

#### NewItem

A New Lost Item with accompanying Image

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| name | string |  | Yes |
| date | dateTime |  | Yes |
| location | string |  | Yes |
| category | integer |  | Yes |
| contact_method | integer |  | Yes |
| contact_details | string |  | Yes |
| item_details | string |  | Yes |
| image_base64 | byte |  | No |

#### SearchItem

Truncated details for a Lost Item

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | integer |  | Yes |
| name | string |  | Yes |
| date | dateTime |  | Yes |
| location | string |  | Yes |
| category | integer |  | Yes |