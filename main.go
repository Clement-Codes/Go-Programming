package main

import (
	"fmt"
	"sync"
	"time"
)


var conferenceName string = "Go Conference"
const conferenceTickets int = 50
var remainingTickets uint = 50
//var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main(){
	//var bookings [50] string
	// var bookings [] string
	// var booking = []string{}
	
	greetUsers()

	//fmt.Printf("ConferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	

	// booking[0] = "Nana"
	// booking[1] = "Nicole"

	//var booking = [50]string{}

	//for {
		firstName, lastName, email, userTickets := getUserInput()
		// fmt.Println(firstName)
		// fmt.Println(&firstName)
	

		//isValidCity:= city == "Singapore" || city == "London"

		isValidName, isValidEmail, isValidTicketNumber :=  ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName &&  isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)
			
			//bookings[0] = firstName + " " + lastName

			// fmt.Printf("The whole array: %v\n", bookings)
			// fmt.Printf("The first value: %v\n", bookings[0])
			// fmt.Printf("Array type: %T\n", bookings)
			// fmt.Printf("Array length: %v\n", len(bookings))
		
			// fmt.Printf("The whole slice: %v\n", bookings)
			// fmt.Printf("The first value: %v\n", bookings[0])
			// fmt.Printf("Slice type: %T\n", bookings)
			// fmt.Printf("Slice length: %v\n", len(bookings))
		

			firstNames := getFirstNames()
			fmt.Printf("The first names of booking are: %v\n", firstNames)

			//fmt.Println("These are all our booking: %v\n", bookings)

			//noTicketsRemaining  := remainingTickets == 0
			if remainingTickets == 0 {
				//end program
				fmt.Println("Our conference is booked out. Come back next year.")
				//break
			}
		}else{
			if !isValidName{
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail{
				fmt.Println("email address you entered doesn't contain @sign")
			}
			if !isValidTicketNumber{
				fmt.Println("number of tickets you entered is invalid")
			}
			//fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			//continue
		}
	//}
	wg.Wait()
	// city := "London"

	// switch(city){
	// 	case "New York":
	// 		print("New York")
	// 	case "Singapore":
	// 		// singapore
	// 	case "London", "Berlin":
	// 		// london
	// 	// case "Berlin":
	// 	// 	//berlin
	// 	case "Mexico City":
	// 		// mexio
	// 	case "Hong Kong":
	// 		//Hongkong
	// 	default:
	// 		fmt.Print("No valid city select")
	// }
}

func greetUsers(){
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func getFirstNames() []string{
	firstNames := []string{}
	for _, booking := range bookings {
		//var names = strings.Fields(booking)
		//firstNames = append(firstNames, booking["firstName"])
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}



func getUserInput()(string,string,string,uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string){
	remainingTickets = remainingTickets - userTickets
	
	// create a map for a user
	// var mySlice []string
	// var myMap map[string]string

	//var userData = make(map[string]string)
	var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings) 
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}


func sendTicket(userTickets uint, firstName string, lastName string, email string){
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("###################")
	wg.Done()
}