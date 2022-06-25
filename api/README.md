# FindNUS_api

API documentation for FindNUS backend services. Handles the retrieval, processing and management of Lost Items found in NUS.

## Table of Contents

* [Servers](#servers)
* [Paths](#paths)
  - [`GET` /debug/ping](#op-get-debug-ping) 
  - [`GET` /debug/checkAuth](#op-get-debug-checkauth) 
  - [`GET` /debug/getDemoItem](#op-get-debug-getdemoitem) 
  - [`POST` /item](#op-post-item) 
  - [`PATCH` /item](#op-patch-item) 
  - [`GET` /item](#op-get-item) 
  - [`DELETE` /item](#op-delete-item) 
  - [`GET` /item/peek](#op-get-item-peek) 
  - [`GET` /search](#op-get-search) 
* [Schemas](#schemas)
  - Item
  - MiniItem
  - NewItem
  - DeleteItem
  - PatchItem
  - Category
  - ContactMethod


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
      <td>Production cluster that is hosting the backend services for FindNUS</td>
    </tr>
    <tr>
      <td><a href="https://uat-findnus.herokuapp.com" target="_blank">https://uat-findnus.herokuapp.com</a></td>
      <td>User-Acceptance Testing cluster environment for testing</td>
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
"Authorization: my-firebase-idToken"
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
        <td>Id <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>The MongoDB ObjectID associated to this Item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Name <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Name of lost/found item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Date <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Date-time where item is lost/found</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Location <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Where the item was found</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Category <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Type of item</td>
        <td><code>Etc</code>, <code>Cards</code>, <code>Notes</code>, <code>Electronics</code>, <code>Bottles</code></td>
      </tr>
      <tr>
        <td>Contact_method</td>
        <td>
          string
        </td>
        <td>Founder/Lostee Contact Method</td>
        <td><code>nus_security</code>, <code>telegram</code>, <code>whatsapp</code>, <code>wechat</code>, <code>line</code>, <code>phone_number</code></td>
      </tr>
      <tr>
        <td>Contact_details</td>
        <td>
          string
        </td>
        <td>Contact details of Founder/Lostee</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Item_details</td>
        <td>
          string
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Image_url</td>
        <td>
          string
        </td>
        <td>Item's accompanying image link</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>User_id</td>
        <td>
          string
        </td>
        <td>UserID associated to this item. Only applicable for Lookout Items.</td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>


##### Example _(generated)_

```json
{
  "Id": "98721yrr0u14oure",
  "Name": "Water Bottle",
  "Date": "2019-08-24T14:15:22Z",
  "Location": "E4A DSA Lab",
  "Category": "Cards",
  "Contact_method": "Telegram",
  "Contact_details": "FindNUS",
  "Item_details": "Blue, with a sticker and broken handle",
  "Image_url": "https://imgur.com/gallery/RaHyECD",
  "User_id": "string"
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

### `POST` /item
<a id="op-post-item" />

Add new Lost item to be put on Lookout on the database.




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
"Authorization: my-firebase-idToken"
```




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
        <td>User_id</td>
        <td>
          string
        </td>
        <td>Unique User_id generated by firebase to associate a user to a Lookout item.</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Name <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Name of new lost/found item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Date <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Date-time where item is lost/found</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Location <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Where the item was found</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Category <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Type of item</td>
        <td><code>Etc</code>, <code>Cards</code>, <code>Notes</code>, <code>Electronics</code>, <code>Bottles</code></td>
      </tr>
      <tr>
        <td>Contact_method</td>
        <td>
          string
        </td>
        <td>Founder/Lostee Contact Method</td>
        <td><code>nus_security</code>, <code>telegram</code>, <code>whatsapp</code>, <code>wechat</code>, <code>line</code>, <code>phone_number</code></td>
      </tr>
      <tr>
        <td>Contact_details</td>
        <td>
          string
        </td>
        <td>Contact details of Founder/Lostee</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Item_details</td>
        <td>
          string
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Image_base64</td>
        <td>
          string
        </td>
        <td>Accompanying image of new Lost/Found item, if applicable</td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>


##### Example _(generated)_

```json
{
  "User_id": "string",
  "Name": "Water Bottle",
  "Date": "2019-08-24T14:15:22Z",
  "Location": "E4A DSA Lab",
  "Category": "Cards",
  "Contact_method": "Telegram",
  "Contact_details": "FindNUS",
  "Item_details": "Blue, with a sticker and broken handle",
  "Image_base64": "string"
}
```




#### Responses


##### ▶ 200 - Item registered into database

###### Headers
_No headers specified_

##### ▶ 400 - Rejected new item into database

###### Headers
_No headers specified_

##### ▶ 401 - Firebase credentials not invalid

###### Headers
_No headers specified_


</div>

### `PATCH` /item
<a id="op-patch-item" />

Update details of an item on the database.




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
"Authorization: my-firebase-idToken"
```


#### Query parameters

##### &#9655; Id

MongoDB ID of the Item


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
        <td>Id  <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>query</td>
        <td>MongoDB ID of the Item</td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>

##### Example

```json
"Id=629cc52563533a84f60c4c68"
```

##### &#9655; User_id

FindNUS User Id (for lost item lookout requests). Include this to remove from Lost (Lookout) Items collection.



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
        <td>User_id  <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>query</td>
        <td><p>FindNUS User Id (for lost item lookout requests). Include this to remove from Lost (Lookout) Items collection.</p>
      </td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>

##### Example

```json
"User_id=196afas7"
```



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
        <td>Id <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>The MongoDB ObjectID associated to this Item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>User_id</td>
        <td>
          string
        </td>
        <td>Unique User_id generated by firebase to associate a user to a Lookout item.</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Name</td>
        <td>
          string
        </td>
        <td>Name of new lost/found item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Date</td>
        <td>
          string
        </td>
        <td>Date-time where item is lost/found</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Location</td>
        <td>
          string
        </td>
        <td>Where the item was found</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Category</td>
        <td>
          string
        </td>
        <td>Type of item</td>
        <td><code>Etc</code>, <code>Cards</code>, <code>Notes</code>, <code>Electronics</code>, <code>Bottles</code></td>
      </tr>
      <tr>
        <td>Contact_method</td>
        <td>
          string
        </td>
        <td>Founder/Lostee Contact Method</td>
        <td><code>nus_security</code>, <code>telegram</code>, <code>whatsapp</code>, <code>wechat</code>, <code>line</code>, <code>phone_number</code></td>
      </tr>
      <tr>
        <td>Contact_details</td>
        <td>
          string
        </td>
        <td>Contact details of Founder/Lostee</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Item_details</td>
        <td>
          string
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Image_base64</td>
        <td>
          string
        </td>
        <td>Updated image of Lost/Found item, if applicable</td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>


##### Example _(generated)_

```json
{
  "Id": "98721yrr0u14oure",
  "User_id": "string",
  "Name": "Water Bottle",
  "Date": "2019-08-24T14:15:22Z",
  "Location": "E4A DSA Lab",
  "Category": "Cards",
  "Contact_method": "Telegram",
  "Contact_details": "FindNUS",
  "Item_details": "Blue, with a sticker and broken handle",
  "Image_base64": "string"
}
```




#### Responses


##### ▶ 200 - OK

###### Headers
_No headers specified_

##### ▶ 401 - Firebase credentials not invalid

###### Headers
_No headers specified_


</div>

### `GET` /item
<a id="op-get-item" />

Get a particular item's full details





#### Query parameters

##### &#9655; Id

Item Id reference. Case sensitive.


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
        <td>Id  <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>query</td>
        <td>Item Id reference. Case sensitive.</td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>


##### &#9655; User_id

User_id filter to search for this Id in the LOST collection. Case sensitive.


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
        <td>User_id </td>
        <td>
          string
        </td>
        <td>query</td>
        <td>User_id filter to search for this Id in the LOST collection. Case sensitive.</td>
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
        <td>Id <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>The MongoDB ObjectID associated to this Item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Name <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Name of lost/found item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Date <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Date-time where item is lost/found</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Location <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Where the item was found</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Category <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Type of item</td>
        <td><code>Etc</code>, <code>Cards</code>, <code>Notes</code>, <code>Electronics</code>, <code>Bottles</code></td>
      </tr>
      <tr>
        <td>Contact_method</td>
        <td>
          string
        </td>
        <td>Founder/Lostee Contact Method</td>
        <td><code>nus_security</code>, <code>telegram</code>, <code>whatsapp</code>, <code>wechat</code>, <code>line</code>, <code>phone_number</code></td>
      </tr>
      <tr>
        <td>Contact_details</td>
        <td>
          string
        </td>
        <td>Contact details of Founder/Lostee</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Item_details</td>
        <td>
          string
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Image_url</td>
        <td>
          string
        </td>
        <td>Item's accompanying image link</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>User_id</td>
        <td>
          string
        </td>
        <td>UserID associated to this item. Only applicable for Lookout Items.</td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>


##### Example _(generated)_

```json
{
  "Id": "98721yrr0u14oure",
  "Name": "Water Bottle",
  "Date": "2019-08-24T14:15:22Z",
  "Location": "E4A DSA Lab",
  "Category": "Cards",
  "Contact_method": "Telegram",
  "Contact_details": "FindNUS",
  "Item_details": "Blue, with a sticker and broken handle",
  "Image_url": "https://imgur.com/gallery/RaHyECD",
  "User_id": "string"
}
```
##### ▶ 404 - Item not found

###### Headers
_No headers specified_

##### ▶ 500 - Internal server error. Likely to be a message queue fault.

###### Headers
_No headers specified_


</div>

### `DELETE` /item
<a id="op-delete-item" />

Remove an item listing on the database.




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
"Authorization: my-firebase-idToken"
```


#### Query parameters

##### &#9655; Id

MongoDB ID of the Item


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
        <td>Id  <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>query</td>
        <td>MongoDB ID of the Item</td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>

##### Example

```json
"Id=629cc52563533a84f60c4c68"
```

##### &#9655; User_id

FindNUS User_Id (for lost item lookout requests). 
Include this to remove from Lost (Lookout) Items collection. 
Case sensitive.



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
        <td>User_id </td>
        <td>
          string
        </td>
        <td>query</td>
        <td><p>FindNUS User_Id (for lost item lookout requests).
      Include this to remove from Lost (Lookout) Items collection.
      Case sensitive.</p>
      </td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>

##### Example

```json
"User_id=196afas7"
```



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
        <td>Id <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>The MongoDB ObjectID associated to this Item</td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>


##### Example _(generated)_

```json
{
  "Id": "98721yrr0u14oure"
}
```




#### Responses


##### ▶ 200 - Deletion request received and will be processed if the item exists.

###### Headers
_No headers specified_

##### ▶ 401 - Firebase credentials not invalid

###### Headers
_No headers specified_


</div>

### `GET` /item/peek
<a id="op-get-item-peek" />

Get a list of lost items sorted by date.
These items are paginated and filtered by category, if requested.
The default returns the latest 20 items, with no category filter.





#### Query parameters

##### &#9655; offset

Number of items to skip (Case sensitive)


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
        <td>offset </td>
        <td>
          integer
        </td>
        <td>query</td>
        <td>Number of items to skip (Case sensitive)</td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>


##### &#9655; limit

Number of items to return (Case sensitive)


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
        <td>limit </td>
        <td>
          integer
        </td>
        <td>query</td>
        <td>Number of items to return (Case sensitive)</td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>


##### &#9655; category

Types of category to filter by.
Chain multiple category values to filter by the
For example, category=Cards&category=Etc will include results from both Cards and Etc. 



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
        <td>category </td>
        <td>
          string
        </td>
        <td>query</td>
        <td><p>Types of category to filter by.
      Chain multiple category values to filter by the
      For example, category=Cards&amp;category=Etc will include results from both Cards and Etc.</p>
      </td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>






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
        <td>Response.Id <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>The MongoDB ObjectID associated to this Item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Response.Name <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Name of lost/found item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Response.Date <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Date-time where item is lost/found</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Response.Location <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Where the item was found</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Response.Category <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Type of item</td>
        <td><code>Etc</code>, <code>Cards</code>, <code>Notes</code>, <code>Electronics</code>, <code>Bottles</code></td>
      </tr>
      <tr>
        <td>Response.Image_url</td>
        <td>
          string
        </td>
        <td>Item's accompanying image link</td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>


##### Example _(generated)_

```json
[
  {
    "Id": "98721yrr0u14oure",
    "Name": "Water Bottle",
    "Date": "2019-08-24T14:15:22Z",
    "Location": "E4A DSA Lab",
    "Category": "Cards",
    "Image_url": "https://imgur.com/gallery/RaHyECD"
  }
]
```
##### ▶ 500 - Internal server error. Likely to be a message queue fault.

###### Headers
_No headers specified_


</div>

### `GET` /search
<a id="op-get-search" />

Text-based search for an item.





#### Query parameters

##### &#9655; query

Text query to search for lost items. 
Can be any arbitrary string - the ElasticSearch engine will attempt to best-match the query.
The query will be performed over the FOUND collection's Name, Category, Location and Item Detail fields.



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
        <td>query </td>
        <td>
          string
        </td>
        <td>query</td>
        <td><p>Text query to search for lost items.
      Can be any arbitrary string - the ElasticSearch engine will attempt to best-match the query.
      The query will be performed over the FOUND collection's Name, Category, Location and Item Detail fields.</p>
      </td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>






#### Responses


##### ▶ 200 - Returns an array of Found items that were matched to the query string.

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
        <td>Response.Id <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>The MongoDB ObjectID associated to this Item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Response.Name <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Name of lost/found item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Response.Date <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Date-time where item is lost/found</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Response.Location <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Where the item was found</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Response.Category <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Type of item</td>
        <td><code>Etc</code>, <code>Cards</code>, <code>Notes</code>, <code>Electronics</code>, <code>Bottles</code></td>
      </tr>
      <tr>
        <td>Response.Image_url</td>
        <td>
          string
        </td>
        <td>Item's accompanying image link</td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>


##### Example _(generated)_

```json
[
  {
    "Id": "98721yrr0u14oure",
    "Name": "Water Bottle",
    "Date": "2019-08-24T14:15:22Z",
    "Location": "E4A DSA Lab",
    "Category": "Cards",
    "Image_url": "https://imgur.com/gallery/RaHyECD"
  }
]
```
##### ▶ 500 - Internal server error. Likely to be a message queue fault.

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
        <td>Id <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>The MongoDB ObjectID associated to this Item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Name <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Name of lost/found item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Date <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Date-time where item is lost/found</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Location <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Where the item was found</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Category <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Type of item</td>
        <td><code>Etc</code>, <code>Cards</code>, <code>Notes</code>, <code>Electronics</code>, <code>Bottles</code></td>
      </tr>
      <tr>
        <td>Contact_method</td>
        <td>
          string
        </td>
        <td>Founder/Lostee Contact Method</td>
        <td><code>nus_security</code>, <code>telegram</code>, <code>whatsapp</code>, <code>wechat</code>, <code>line</code>, <code>phone_number</code></td>
      </tr>
      <tr>
        <td>Contact_details</td>
        <td>
          string
        </td>
        <td>Contact details of Founder/Lostee</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Item_details</td>
        <td>
          string
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Image_url</td>
        <td>
          string
        </td>
        <td>Item's accompanying image link</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>User_id</td>
        <td>
          string
        </td>
        <td>UserID associated to this item. Only applicable for Lookout Items.</td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>

##### Example _(generated)_

```json
{
  "Id": "98721yrr0u14oure",
  "Name": "Water Bottle",
  "Date": "2019-08-24T14:15:22Z",
  "Location": "E4A DSA Lab",
  "Category": "Cards",
  "Contact_method": "Telegram",
  "Contact_details": "FindNUS",
  "Item_details": "Blue, with a sticker and broken handle",
  "Image_url": "https://imgur.com/gallery/RaHyECD",
  "User_id": "string"
}
```
<a id="" />

#### MiniItem

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
        <td>Id <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>The MongoDB ObjectID associated to this Item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Name <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Name of lost/found item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Date <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Date-time where item is lost/found</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Location <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Where the item was found</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Category <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Type of item</td>
        <td><code>Etc</code>, <code>Cards</code>, <code>Notes</code>, <code>Electronics</code>, <code>Bottles</code></td>
      </tr>
      <tr>
        <td>Image_url</td>
        <td>
          string
        </td>
        <td>Item's accompanying image link</td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>

##### Example _(generated)_

```json
{
  "Id": "98721yrr0u14oure",
  "Name": "Water Bottle",
  "Date": "2019-08-24T14:15:22Z",
  "Location": "E4A DSA Lab",
  "Category": "Cards",
  "Image_url": "https://imgur.com/gallery/RaHyECD"
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
        <td>User_id</td>
        <td>
          string
        </td>
        <td>Unique User_id generated by firebase to associate a user to a Lookout item.</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Name <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Name of new lost/found item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Date <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Date-time where item is lost/found</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Location <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Where the item was found</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Category <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>Type of item</td>
        <td><code>Etc</code>, <code>Cards</code>, <code>Notes</code>, <code>Electronics</code>, <code>Bottles</code></td>
      </tr>
      <tr>
        <td>Contact_method</td>
        <td>
          string
        </td>
        <td>Founder/Lostee Contact Method</td>
        <td><code>nus_security</code>, <code>telegram</code>, <code>whatsapp</code>, <code>wechat</code>, <code>line</code>, <code>phone_number</code></td>
      </tr>
      <tr>
        <td>Contact_details</td>
        <td>
          string
        </td>
        <td>Contact details of Founder/Lostee</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Item_details</td>
        <td>
          string
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Image_base64</td>
        <td>
          string
        </td>
        <td>Accompanying image of new Lost/Found item, if applicable</td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>

##### Example _(generated)_

```json
{
  "User_id": "string",
  "Name": "Water Bottle",
  "Date": "2019-08-24T14:15:22Z",
  "Location": "E4A DSA Lab",
  "Category": "Cards",
  "Contact_method": "Telegram",
  "Contact_details": "FindNUS",
  "Item_details": "Blue, with a sticker and broken handle",
  "Image_base64": "string"
}
```
<a id="" />

#### DeleteItem

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
        <td>Id <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>The MongoDB ObjectID associated to this Item</td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>

##### Example _(generated)_

```json
{
  "Id": "98721yrr0u14oure"
}
```
<a id="" />

#### PatchItem

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
        <td>Id <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td>The MongoDB ObjectID associated to this Item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>User_id</td>
        <td>
          string
        </td>
        <td>Unique User_id generated by firebase to associate a user to a Lookout item.</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Name</td>
        <td>
          string
        </td>
        <td>Name of new lost/found item</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Date</td>
        <td>
          string
        </td>
        <td>Date-time where item is lost/found</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Location</td>
        <td>
          string
        </td>
        <td>Where the item was found</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Category</td>
        <td>
          string
        </td>
        <td>Type of item</td>
        <td><code>Etc</code>, <code>Cards</code>, <code>Notes</code>, <code>Electronics</code>, <code>Bottles</code></td>
      </tr>
      <tr>
        <td>Contact_method</td>
        <td>
          string
        </td>
        <td>Founder/Lostee Contact Method</td>
        <td><code>nus_security</code>, <code>telegram</code>, <code>whatsapp</code>, <code>wechat</code>, <code>line</code>, <code>phone_number</code></td>
      </tr>
      <tr>
        <td>Contact_details</td>
        <td>
          string
        </td>
        <td>Contact details of Founder/Lostee</td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Item_details</td>
        <td>
          string
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>Image_base64</td>
        <td>
          string
        </td>
        <td>Updated image of Lost/Found item, if applicable</td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>

##### Example _(generated)_

```json
{
  "Id": "98721yrr0u14oure",
  "User_id": "string",
  "Name": "Water Bottle",
  "Date": "2019-08-24T14:15:22Z",
  "Location": "E4A DSA Lab",
  "Category": "Cards",
  "Contact_method": "Telegram",
  "Contact_details": "FindNUS",
  "Item_details": "Blue, with a sticker and broken handle",
  "Image_base64": "string"
}
```
<a id="" />

#### Category

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
        <td>Category</td>
        <td>
          string
        </td>
        <td>Non-case sensitive category name</td>
        <td><code>Etc</code>, <code>Cards</code>, <code>Notes</code>, <code>Electronics</code>, <code>Bottles</code></td>
      </tr>
  </tbody>
</table>

##### Example _(generated)_

```json
"Etc"
```
<a id="" />

#### ContactMethod

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
        <td>ContactMethod</td>
        <td>
          string
        </td>
        <td>Non-case sensitive contact method</td>
        <td><code>nus_security</code>, <code>telegram</code>, <code>whatsapp</code>, <code>wechat</code>, <code>line</code>, <code>phone_number</code></td>
      </tr>
  </tbody>
</table>

##### Example _(generated)_

```json
"nus_security"
```
