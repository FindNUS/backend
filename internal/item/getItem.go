package main

// Handler for specific get item calls.
// Determines the collection the GET is querying,
// then gets the item from MongoDB and returns it
func DoGetItem(msg ItemMsgJSON) Item {
	var item Item
	id, ok := msg.Params["Id"]
	if !ok {
		return item // no Id exists -- should no happen
	}
	userid, ok := msg.Params["User_id"]
	if ok {
		// User id exists, get request should point to LOST pool
		item = MongoGetItem(COLL_LOST, id[0], userid[0])
		return item
	}
	// User id does not exist, get request should point to FOUND pool
	item = MongoGetItem(COLL_FOUND, id[0], "")
	return item
}

// Hanlder for /peek with filters
// Parses the filters
func DoGetManyItems(msg ItemMsgJSON) []Item {
	// TODO specify filters to include:
	// Limit, Offset, SortBy, FilterBy
	params := msg.Params
	var items []Item
	if _, ok := params["User_id"]; ok {
		items = MongoGetManyItems(COLL_LOST, params)
	} else {
		items = MongoGetManyItems(COLL_FOUND, params)
	}
	return items
}
