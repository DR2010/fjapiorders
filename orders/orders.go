package orders

import (
	"database/sql"
	helper "fjapiorders/helper"
	"fmt"
	"log"
	"strconv"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Order is what the client wants
type Order struct {
	SystemID   bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	ID         string        // random ID for order, yet to define algorithm
	ClientName string        // Client Name
	ClientID   string        // Client ID in case they logon
	Atendente  string        // Pessoa atendendo
	Date       string        // Order Date
	Time       string        // Order Time
	Status     string        // Placed, Serving, Completed, Removed
	EventID    string        // Every order belongs to an Event (Events are linked to Organisation)
	EatMode    string        // EatIn, TakeAway, Delivery
	PickUpTime string        // EatIn, TakeAway, Delivery
	TotalGeral string        // Delivery phone number
	Items      []Item
}

// Item represents a single item of an order
type Item struct {
	ID         string // Sequential number of the item
	PratoName  string // Dish ID or unique name from "Dishes"
	Quantidade string // Individual price
	Price      string // Individual price
	Total      string // Total Price
	Tax        string // GST
}

// SearchCriteria is what the client wants
type SearchCriteria struct {
	ID                   string // random ID for order, yet to define algorithm
	ClientName           string // Client Name
	ClientID             string // Client ID in case they logon
	Date                 string // Order Date
	Time                 string // Order Time
	Status               string // Open, Completed, Cancelled
	EatMode              string // EatIn, TakeAway, Delivery
	DeliveryMode         string // Internal, UberEats,
	DeliveryFee          string // Delivery Fee
	DeliveryLocation     string // Address
	DeliveryContactPhone string // Delivery phone number
}

// Add is for export
func Add(objtoinsert Order) helper.Resultado {

	database := helper.GetDBParmFromCache("CollectionOrders")

	session, err := mgo.Dial(database.Location)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	collection := session.DB(database.Database).C(database.Collection)

	err = collection.Insert(objtoinsert)

	if err != nil {
		log.Fatal(err)
	}

	var res helper.Resultado
	res.ErrorCode = "0001"
	res.ErrorDescription = "Order added"
	res.IsSuccessful = "Y"

	return res
}

// Find is to find stuff
func Find(objtofind string) (Order, string) {

	database := helper.GetDBParmFromCache("CollectionOrders")

	objkey := objtofind
	objnull := Order{}

	session, err := mgo.Dial(database.Location)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(database.Database).C(database.Collection)

	result := []Order{}
	err1 := c.Find(bson.M{"id": objkey}).All(&result)
	if err1 != nil {
		log.Fatal(err1)
	}

	var numrecsel = len(result)

	if numrecsel <= 0 {
		return objnull, "404 Not found"
	}

	return result[0], "200 OK"
}

// Getall works
func Getall() []Order {

	database := helper.GetDBParmFromCache("CollectionOrders")

	fmt.Println("database.Location")
	fmt.Println(database.Location)

	session, err := mgo.Dial(database.Location)

	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(database.Database).C(database.Collection)

	var results []Order

	err = c.Find(nil).All(&results)
	if err != nil {
		// TODO: Do something about the error
	} else {
		return results
	}

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// GetallbyUser works
func GetallbyUser(userid string) []Order {

	database := helper.GetDBParmFromCache("CollectionOrders")

	fmt.Println("database.Location")
	fmt.Println(database.Location)

	session, err := mgo.Dial(database.Location)

	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(database.Database).C(database.Collection)

	var results []Order

	err = c.Find(bson.M{"clientid": userid}).All(&results)
	if err != nil {
		// TODO: Do something about the error
	} else {
		return results
	}

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// GetallbyOrderName works
func GetallbyOrderName(ordername string) []Order {
	// ---------------------------
	// Show all order for a client
	// It will help to show the total to pay later
	// ---------------------------
	database := helper.GetDBParmFromCache("CollectionOrders")

	fmt.Println("database.Location")
	fmt.Println(database.Location)

	session, err := mgo.Dial(database.Location)

	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(database.Database).C(database.Collection)

	var results []Order

	err = c.Find(bson.M{"ClientName": ordername}).All(&results)
	if err != nil {
		// TODO: Do something about the error
	} else {
		return results
	}

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// Getallcompleted works
func Getallcompleted(status string) []Order {

	database := helper.GetDBParmFromCache("CollectionOrders")

	fmt.Println("database.Location")
	fmt.Println(database.Location)

	session, err := mgo.Dial(database.Location)

	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(database.Database).C(database.Collection)

	var results []Order

	err = c.Find(bson.M{"status": status}).All(&results)
	if err != nil {
		// TODO: Do something about the error
	} else {
		return results
	}

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// Getallbutcompleted works
func Getallbutcompleted() []Order {

	status := "Completed"

	database := helper.GetDBParmFromCache("CollectionOrders")

	session, err := mgo.Dial(database.Location)

	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(database.Database).C(database.Collection)

	var results []Order

	// db.inventory.find( { qty: { $ne: 20 } } )

	// err = c.Find(bson.M{"status": status}).All(&results)
	// err = collection.Find(bson.M{"currency": currency, "datetime": bson.M{"$gte": yearmonthday, "$lte": yearmonthdayend}}).Sort("-datetime").All(&results)

	err = c.Find(bson.M{"status": bson.M{"$ne": status}}).All(&results)
	if err != nil {
		// TODO: Do something about the error
	} else {
		return results
	}

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// Update is
func Update(objtoupdate Order) helper.Resultado {

	database := helper.GetDBParmFromCache("CollectionOrders")

	session, err := mgo.Dial(database.Location)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	collection := session.DB(database.Database).C(database.Collection)

	err = collection.Update(bson.M{"id": objtoupdate.ID}, objtoupdate)

	if err != nil {
		log.Fatal(err)
	}

	var res helper.Resultado
	res.ErrorCode = "0001"
	res.ErrorDescription = "Something Happened"
	res.IsSuccessful = "Y"

	return res
}

// Delete is
func Delete(objtodeletekey string) helper.Resultado {

	database := helper.GetDBParmFromCache("CollectionOrders")

	session, err := mgo.Dial(database.Location)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	collection := session.DB(database.Database).C(database.Collection)

	err = collection.Remove(bson.M{"ID": objtodeletekey})

	if err != nil {
		log.Fatal(err)
	}

	var res helper.Resultado
	res.ErrorCode = "0001"
	res.ErrorDescription = "Order deleted successfully"
	res.IsSuccessful = "Y"

	return res
}

// SavetoMySQL will save the data from orders to MySQL
func SavetoMySQL(db *sql.DB) {

	// Created on 19/7/2018
	// This program will save data to MySQL
	// Call function to return all Orders from MongoDB
	// for each order
	// ... insert Order into MySQL
	// ....for each order item
	// ....... insert OrderItem into MySQL
	// that's it

	statuscompleted := "Completed"

	listoforders := Getallcompleted(statuscompleted)

	for i := 0; i < len(listoforders); i++ {

		order := listoforders[i]

		number, _ := strconv.Atoi(order.ID)
		fullname := order.ClientName
		date := order.Date
		ttime := order.Time
		status := order.Status
		total, _ := strconv.ParseFloat(order.TotalGeral, 64)

		_, err := db.Exec("INSERT INTO festajunina.order(number, status, fullname, total, date, time) VALUES(?,?,?,?,?,?)", number, status, fullname, total, date, ttime)

		if err != nil {
			// http.Error(res, "Server error, unable to create your account.", 500)
			return
		}

		numitem := 0
		for p := 0; p < len(order.Items); p++ {

			numitem++
			orderitem := order.Items[p]

			fkordernumber, _ := strconv.Atoi(order.ID)
			sequencenumber := numitem // made up value
			dishname := orderitem.PratoName
			total, _ := strconv.ParseFloat(orderitem.Total, 64)
			price, _ := strconv.ParseFloat(orderitem.Price, 64)
			quantidade, _ := strconv.Atoi(orderitem.Quantidade)

			_, err := db.Exec("INSERT INTO festajunina.orderitem(fkordernumber, sequencenumber, dishname, quantity, price, total) VALUES(?,?,?,?,?,?)", fkordernumber, sequencenumber, dishname, quantidade, price, total)

			if err != nil {
				// http.Error(res, "Server error, unable to create your account.", 500)
				return
			}

		}

	}
	return
}
