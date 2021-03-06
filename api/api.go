package api

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
	"fmt"
)


//Declare a global array of Credentials
//See credentials.go

/*YOUR CODE HERE*/
var credentials = []Credentials{}



func RegisterRoutes(router *mux.Router) error {

	/*

	Fill out the appropriate get methods for each of the requests, based on the nature of the request.

	Think about whether you're reading, writing, or updating for each request


	*/

	router.HandleFunc("/api/getCookie", getCookie).Methods(http.MethodGet)
	router.HandleFunc("/api/getQuery", getQuery).Methods(http.MethodGet)
	router.HandleFunc("/api/getJSON", getJSON).Methods(http.MethodGet)
	
	router.HandleFunc("/api/signup", signup).Methods(http.MethodPost)
	router.HandleFunc("/api/getIndex", getIndex).Methods(http.MethodGet)
	router.HandleFunc("/api/getpw", getPassword).Methods(http.MethodGet)
	router.HandleFunc("/api/updatepw", updatePassword).Methods(http.MethodPut)
	router.HandleFunc("/api/deleteuser", deleteUser).Methods(http.MethodDelete)

	return nil
}

func getCookie(response http.ResponseWriter, request *http.Request) {

	/*
		Obtain the "access_token" cookie's value and write it to the response

		If there is no such cookie, write an empty string to the response
	*/

	/*YOUR CODE HERE*/
	cookie, err := request.Cookie("access_token")
	if err != nil {
		fmt.Fprintf(response, "")
		return 
	}
	accessToken := cookie.Value
	fmt.Fprintf(response, accessToken)
	return 
}

func getQuery(response http.ResponseWriter, request *http.Request) {

	/*
		Obtain the "userID" query paramter and write it to the response
		If there is no such query parameter, write an empty string to the response
	*/

	/*YOUR CODE HERE*/
	userID := request.URL.Query().Get("userID")
	if userID == "" {
		fmt.Fprintf(response, "")
		return 
	}
	fmt.Fprintf(response, userID)
	return
}

func getJSON(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>,
			"password" : <password>
		}

		Decode this json file into an instance of Credentials.

		Then, write the username and password to the response, separated by a newline.request
		
		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	/*YOUR CODE HERE*/
	credential := Credentials{}
	err := json.NewDecoder(request.Body).Decode(&credential)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	if credential.Username == "" || credential.Password == "" {
		response.WriteHeader(400)
		return
	}
	fmt.Fprintf(response, credential.Username + "\n")
	fmt.Fprintf(response, credential.Password)
}

func signup(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>,
			"password" : <password>
		}

		Decode this json file into an instance of Credentials.

		Then store it ("append" it) to the global array of Credentials.

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	/*YOUR CODE HERE*/
	credential := Credentials{}
	jsonDecoder := json.NewDecoder(request.Body)
	err := jsonDecoder.Decode(&credential)
	if err != nil { 
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	if credential.Username == "" || credential.Password == "" {
		response.WriteHeader(400)
		return
	}
	credentials = append(credentials, credential)
	response.WriteHeader(201)
	return
}

func getIndex(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>
		}


		Decode this json file into an instance of Credentials. (What happens when we don't have all the fields? Does it matter in this case?)

		Return the array index of the Credentials object in the global Credentials array
		
		The index will be of type integer, but we can only write strings to the response. What library and function was used to get around this?

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	/*YOUR CODE HERE*/
	credential := Credentials{}
	jsonDecoder := json.NewDecoder(request.Body)
	err := jsonDecoder.Decode(&credential)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	for i, v := range credentials {
		if v.Username == credential.Username{
			fmt.Fprintf(response, strconv.Itoa(i))
			return
		}
	}

	response.WriteHeader(400)
	return
}

func getPassword(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>
		}


		Decode this json file into an instance of Credentials. (What happens when we don't have all the fields? Does it matter in this case?)

		Write the password of the specific user to the response

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	/*YOUR CODE HERE*/
	credential := Credentials{}
	jsonDecoder := json.NewDecoder(request.Body)
	err := jsonDecoder.Decode(&credential)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	for _, v := range credentials {
		if v.Username == credential.Username{
			fmt.Fprintf(response, v.Password)
			return 
		}
	}
	response.WriteHeader(400)
	return 
}


func updatePassword(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>,
			"password" : <password,
		}


		Decode this json file into an instance of Credentials. 

		The password in the JSON file is the new password they want to replace the old password with.

		You don't need to return anything in this.

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	/*YOUR CODE HERE*/
	credential := Credentials{}
	jsonDecoder := json.NewDecoder(request.Body)
	err := jsonDecoder.Decode(&credential)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	if credential.Username == "" || credential.Password == "" {
		response.WriteHeader(400)
		return
	}
	for i, v := range credentials {
		if v.Username == credential.Username {
			credentials[i].Password = credential.Password
			return
		}
	}
	response.WriteHeader(400)
	return
}

func remove(slice []Credentials, s int) []Credentials {
	return append(slice[:s], slice[s+1:]...)
}

func deleteUser(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>,
			"password" : <password,
		}


		Decode this json file into an instance of Credentials.

		Remove this user from the array. Preserve the original order. You may want to create a helper function.

		This wasn't covered in lecture, so you may want to read the following:
			- https://gobyexample.com/slices
			- https://www.delftstack.com/howto/go/how-to-delete-an-element-from-a-slice-in-golang/

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	/*YOUR CODE HERE*/
	credential := Credentials{}
	jsonDecoder := json.NewDecoder(request.Body)
	err := jsonDecoder.Decode(&credential)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	for i, v := range credentials {
		if v.Username == credential.Username && v.Password == credential.Password{
			credentials = remove(credentials, i)
			return
		}
	}
	response.WriteHeader(400)
	return
}
