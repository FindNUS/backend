# FindNUS_api

API documentation for FindNUS backend services. Handles the retrieval, processing and management of Lost Items found in NUS.

## Table of Contents

* [Servers](#servers)
* [Paths](#paths)
  - [`GET` /debug/ping](#op-get-debug-ping) 
  - [`GET` /debug/checkAuth](#op-get-debug-checkauth) 
  - [`GET` /debug/getDemoItem](#op-get-debug-getdemoitem) 
  - [`POST` /item/new](#op-post-item-new) 
  - [`GET` /item/get/{itemId}](#op-get-item-get-itemid) 
  - [`GET` /item/peek](#op-get-item-peek) 
  - [`POST` /item/request](#op-post-item-request) 
  - [`GET` /search](#op-get-search) 
* [Schemas](#schemas)
  - Item
  - NewItem
  - SearchItem


<a id="servers" />
## Servers

<table>
  <thead>
    <tr>
      <th>URL</th>
      <th>Description</th>
    <tr>
  </thead>
  <tbody>
    <tr>
      <td><a href="https://findnus.herokuapp.com" target="_blank">https://findnus.herokuapp.com</a></td>
      <td>Heroku domain that hosts the backend services for FindNUS</td>
    </tr>
    <tr>
      <td><a href="https://uat-findnus.herokuapp.com" target="_blank">https://uat-findnus.herokuapp.com</a></td>
      <td>Integration environment for User acceptance testing.</td>
    </tr>
  </tbody>
</table>

<a name="security"></a>
## Security

<table class="table">
  <thead class="table__head">
    <tr class="table__head__row">
      <th class="table__head__cell">Type</th>
      <th class="table__head__cell">In</th>
      <th class="table__head__cell">Name</th>
      <th class="table__head__cell">Scheme</th>
      <th class="table__head__cell">Format</th>
      <th class="table__head__cell">Description</th>
    </tr>
  </thead>
  <tbody class="table__body">
    <tr class="table__body__row">
      <td class="table__body__cell">http</td>
      <td class="table__body__cell"></td>
      <td class="table__body__cell"></td>
      <td class="table__body__cell">bearer</td>
      <td class="table__body__cell"></td>
      <td class="table__body__cell"></td>
    </tr>

  </tbody>
</table>

## Paths


### `GET` /debug/ping
<a id="op-get-debug-ping" />

Returns a Hello World equivalent message. Shows that the backend connection works.









#### Responses


##### ▶ 200 - A hello world message.

###### Headers
_No headers specified_

###### text/plain



<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Type</th>
      <th>Description</th>
      <th>Accepted values</th>
    </tr>
  </thead>
  <tbody>
      <tr>
        <td>Response</td>
        <td>
          string
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>


##### Example
```
message: Hi there, you have reached FindNUS!

```

</div>

### `GET` /debug/checkAuth
<a id="op-get-debug-checkauth" />

Check with backend if the Firebase token is valid.




#### Headers

##### &#9655; Authorization

Firebase ID token of user


<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Type</th>
      <th>In</th>
      <th>Description</th>
      <th>Accepted values</th>
    </tr>
  </thead>
  <tbody>
      <tr>
        <td>Authorization  <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>header</td>
        <td>Firebase ID token of user</td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>

##### Example

```json
"Authorization: my-token\n"
```






#### Responses


##### ▶ 200 - Id token is valid

###### Headers
_No headers specified_

##### ▶ 401 - Id token is invalid

###### Headers
_No headers specified_


</div>

### `GET` /debug/getDemoItem
<a id="op-get-debug-getdemoitem" />

Get a demo item for Milestone 1. 





#### Query parameters

##### &#9655; name

Name of the demoset item to be retrieved


<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Type</th>
      <th>In</th>
      <th>Description</th>
      <th>Accepted values</th>
    </tr>
  </thead>
  <tbody>
      <tr>
        <td>name  <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>query</td>
        <td>Name of the demoset item to be retrieved</td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>






#### Responses


##### ▶ 200 - Get request is valid, item is found

###### Headers
_No headers specified_

###### application/json



<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Type</th>
      <th>Description</th>
      <th>Accepted values</th>
    </tr>
  </thead>
  <tbody>
      <tr>
        <td>id <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>The MongoDB ObjectID associated to this Item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>name <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Name of lost item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>date <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Date-time where item is lost/found</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>location <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Where the item was found</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>category <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Type of item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>contact_method</td>
        <td>
          string
        </td>
        <td>Founder/Lostee Contact Method</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>contact_details</td>
        <td>
          string
        </td>
        <td>Contact details of Founder/Lostee</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>item_details</td>
        <td>
          string
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>image_url</td>
        <td>
          string
        </td>
        <td>Item's accompanying image</td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>


##### Example _(generated)_

```json
{
  "id": "98721yrr0u14oure",
  "name": "Water Bottle",
  "date": "2019-08-24T14:15:22Z",
  "location": "E4A DSA Lab",
  "category": "Etc",
  "contact_method": "Telegram",
  "contact_details": "FindNUS",
  "item_details": "Blue, with a sticker and broken handle",
  "image_url": "https://imgur.com/gallery/RaHyECD"
}
```
##### ▶ 404 - Get request is valid, item not found

###### Headers
_No headers specified_

###### text/plain



<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Type</th>
      <th>Description</th>
      <th>Accepted values</th>
    </tr>
  </thead>
  <tbody>
      <tr>
        <td>Response</td>
        <td>
          string
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>


##### Example
```
Nothing Found!

```

</div>

### `POST` /item/new
<a id="op-post-item-new" />

Submit a new item to be stored into the database.







#### Request body
###### application/json



<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Type</th>
      <th>Description</th>
      <th>Accepted values</th>
    </tr>
  </thead>
  <tbody>
      <tr>
        <td>name <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Name of new lost/found item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>date <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Date-time where item is lost/found</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>location <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Where the item was found</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>category <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Type of item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>contact_method</td>
        <td>
          string
        </td>
        <td>Founder/Lostee Contact Method</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>contact_details</td>
        <td>
          string
        </td>
        <td>Contact details of Founder/Lostee</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>item_details</td>
        <td>
          string
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>image_base64</td>
        <td>
          string
        </td>
        <td>Accompanying image of new Lost/Found item</td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>


##### Example _(generated)_

```json
{
  "name": "Water Bottle",
  "date": "2019-08-24T14:15:22Z",
  "location": "E4A DSA Lab",
  "category": "Etc",
  "contact_method": "Telegram",
  "contact_details": "FindNUS",
  "item_details": "Blue, with a sticker and broken handle",
  "image_base64": "string"
}
```




#### Responses


##### ▶ 201 - Item registered into database

###### Headers
_No headers specified_

##### ▶ 400 - Rejected new item into database

###### Headers
_No headers specified_


</div>

### `GET` /item/get/{itemId}
<a id="op-get-item-get-itemid" />

Get a particular item's full details



#### Path parameters

##### &#9655; itemId

Item Id reference


<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Type</th>
      <th>In</th>
      <th>Description</th>
      <th>Accepted values</th>
    </tr>
  </thead>
  <tbody>
      <tr>
        <td>itemId  <strong>(required)</strong></td>
        <td>
          integer
        </td>
        <td>path</td>
        <td>Item Id reference</td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>








#### Responses


##### ▶ 200 - A Lost Item's details

###### Headers
_No headers specified_

###### application/json



<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Type</th>
      <th>Description</th>
      <th>Accepted values</th>
    </tr>
  </thead>
  <tbody>
      <tr>
        <td>id <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>The MongoDB ObjectID associated to this Item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>name <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Name of lost item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>date <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Date-time where item is lost/found</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>location <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Where the item was found</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>category <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Type of item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>contact_method</td>
        <td>
          string
        </td>
        <td>Founder/Lostee Contact Method</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>contact_details</td>
        <td>
          string
        </td>
        <td>Contact details of Founder/Lostee</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>item_details</td>
        <td>
          string
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>image_url</td>
        <td>
          string
        </td>
        <td>Item's accompanying image</td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>


##### Example _(generated)_

```json
{
  "id": "98721yrr0u14oure",
  "name": "Water Bottle",
  "date": "2019-08-24T14:15:22Z",
  "location": "E4A DSA Lab",
  "category": "Etc",
  "contact_method": "Telegram",
  "contact_details": "FindNUS",
  "item_details": "Blue, with a sticker and broken handle",
  "image_url": "https://imgur.com/gallery/RaHyECD"
}
```
##### ▶ 404 - Item not found

###### Headers
_No headers specified_


</div>

### `GET` /item/peek
<a id="op-get-item-peek" />

Get a list of lost items that can be sorted.
Peek at the database's latest finds.
(Sorting and filtering to be implemented in the future)









#### Responses


##### ▶ 200 - Returns an array of lost items that may be filtered

###### Headers
_No headers specified_

###### application/json



<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Type</th>
      <th>Description</th>
      <th>Accepted values</th>
    </tr>
  </thead>
  <tbody>
      <tr>
        <td>Response</td>
        <td>
          array(object)
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Response.id <strong>(required)</strong></td>
        <td>
          integer
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Response.name <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Response.date <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Response.location <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Response.category <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>


##### Example _(generated)_

```json
[
  {
    "id": 0,
    "name": "string",
    "date": "2019-08-24T14:15:22Z",
    "location": "string",
    "category": "string"
  }
]
```

</div>

### `POST` /item/request
<a id="op-post-item-request" />

Add a lost item request to the server. 







#### Request body
###### application/json



<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Type</th>
      <th>Description</th>
      <th>Accepted values</th>
    </tr>
  </thead>
  <tbody>
      <tr>
        <td>name <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Name of new lost/found item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>date <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Date-time where item is lost/found</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>location <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Where the item was found</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>category <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Type of item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>contact_method</td>
        <td>
          string
        </td>
        <td>Founder/Lostee Contact Method</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>contact_details</td>
        <td>
          string
        </td>
        <td>Contact details of Founder/Lostee</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>item_details</td>
        <td>
          string
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>image_base64</td>
        <td>
          string
        </td>
        <td>Accompanying image of new Lost/Found item</td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>


##### Example _(generated)_

```json
{
  "name": "Water Bottle",
  "date": "2019-08-24T14:15:22Z",
  "location": "E4A DSA Lab",
  "category": "Etc",
  "contact_method": "Telegram",
  "contact_details": "FindNUS",
  "item_details": "Blue, with a sticker and broken handle",
  "image_base64": "string"
}
```




#### Responses


##### ▶ 201 - Item lookout request added to server.

###### Headers
_No headers specified_


</div>

### `GET` /search
<a id="op-get-search" />

Elastisearch for an item.









#### Responses


##### ▶ 501 - Function not added yet.

###### Headers
_No headers specified_


</div>

## Schemas

<a id="" />

#### Item

<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Type</th>
      <th>Description</th>
      <th>Accepted values</th>
    </tr>
  </thead>
  <tbody>
      <tr>
        <td>id <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>The MongoDB ObjectID associated to this Item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>name <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Name of lost item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>date <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Date-time where item is lost/found</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>location <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Where the item was found</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>category <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Type of item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>contact_method</td>
        <td>
          string
        </td>
        <td>Founder/Lostee Contact Method</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>contact_details</td>
        <td>
          string
        </td>
        <td>Contact details of Founder/Lostee</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>item_details</td>
        <td>
          string
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>image_url</td>
        <td>
          string
        </td>
        <td>Item's accompanying image</td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>

##### Example _(generated)_

```json
{
  "id": "98721yrr0u14oure",
  "name": "Water Bottle",
  "date": "2019-08-24T14:15:22Z",
  "location": "E4A DSA Lab",
  "category": "Etc",
  "contact_method": "Telegram",
  "contact_details": "FindNUS",
  "item_details": "Blue, with a sticker and broken handle",
  "image_url": "https://imgur.com/gallery/RaHyECD"
}
```
<a id="" />

#### NewItem

<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Type</th>
      <th>Description</th>
      <th>Accepted values</th>
    </tr>
  </thead>
  <tbody>
      <tr>
        <td>name <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Name of new lost/found item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>date <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Date-time where item is lost/found</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>location <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Where the item was found</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>category <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Type of item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>contact_method</td>
        <td>
          string
        </td>
        <td>Founder/Lostee Contact Method</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>contact_details</td>
        <td>
          string
        </td>
        <td>Contact details of Founder/Lostee</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>item_details</td>
        <td>
          string
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>image_base64</td>
        <td>
          string
        </td>
        <td>Accompanying image of new Lost/Found item</td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>

##### Example _(generated)_

```json
{
  "name": "Water Bottle",
  "date": "2019-08-24T14:15:22Z",
  "location": "E4A DSA Lab",
  "category": "Etc",
  "contact_method": "Telegram",
  "contact_details": "FindNUS",
  "item_details": "Blue, with a sticker and broken handle",
  "image_base64": "string"
}
```
<a id="" />

#### SearchItem

<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Type</th>
      <th>Description</th>
      <th>Accepted values</th>
    </tr>
  </thead>
  <tbody>
      <tr>
        <td>id <strong>(required)</strong></td>
        <td>
          integer
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>name <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>date <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>location <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>category <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>

##### Example _(generated)_

```json
{
  "id": 0,
  "name": "string",
  "date": "2019-08-24T14:15:22Z",
  "location": "string",
  "category": "string"
}
```
