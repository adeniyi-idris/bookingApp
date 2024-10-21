package main

import (
	"fmt" 
	"time"
)

var confenrenceName = "Go confenrence"
const confenrenceTickets = 50
var remainingTickets uint = 50
var bookings = make([]UserData, 0)
//bookings := []string{}
// var bookings = []string{}


type UserData = struct{
	firstName string
	lastName string
	email string
	numberOfTickets uint

}

func main(){
	
	greetUsers()
	
	for{
		userName, lastName, email, userTicket :=getUserInput()
		isValid, isValidEmail, isValidTicketNumber := ValidateUserInput(userName,lastName,email,userTicket,remainingTickets)

		if isValid && isValidEmail && isValidTicketNumber{
			
			bookTicket(userTicket, userName, lastName, email)
			sendTicket(userTicket, userName, lastName, email)

				firstNames := getFirstName()
				fmt.Printf("The first name of bookings are: %v\n", firstNames)

				if remainingTickets == 0{
					//end program
					fmt.Println("Our conference is booked out. Come back next year")
					break
				}

		}else{
			if !isValid{
				fmt.Println("firstname or lastname is too short")
			}
			if !isValidEmail{
				fmt.Println("Invalid email")
			}
			if !isValidTicketNumber{
				fmt.Println("number of ticket you enterd is invalid")
			}
			
		}
		
		
	}

	//fmt.Println(confenrenceName)
	
	////fmt.Printf("Your Input data is invalid, try again\n")
		//	continue
}

func greetUsers(){
	fmt.Printf("Welcome to %v booking application\n", confenrenceName)
	fmt.Printf("Welcome to our %v booking application\n", confenrenceName)
	fmt.Printf("we have total of %v tickets and %v are still available\n", confenrenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstName() []string{
	firstNames := []string{}
				/////LOOPS
				for _, booking := range bookings{
					//var names = strings.Fields(booking) 
					firstNames = append(firstNames, booking.firstName)
				}
				return firstNames
				//fmt.Printf("The first name of bookings are: %v\n", firstNames)
}



func getUserInput() (string, string, string, uint){
	    var userName string
		var lastName string
		var email string
		var userTicket uint
	    //////////////////
	    // ask userinput
		fmt.Println("Enter your first name:")
		fmt.Scan(&userName)
	
		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)
	
		fmt.Println("Enter your email:")
		fmt.Scan(&email)
	
		fmt.Println("Enter number of Ticket:")
		fmt.Scan(&userTicket)

		return userName, lastName, email, userTicket
}

func bookTicket(userTicket uint, userName string, lastName string, email string){
	remainingTickets = remainingTickets - userTicket

			var userData = UserData{
				firstName: userName,
				lastName: lastName,
				email: email,
				numberOfTickets: userTicket,
			}
						///////SLICE
				//bookings = append(bookings, userName + " " + lastName)
				bookings = append(bookings, userData)
				fmt.Printf("List of bookings is %v\n", bookings)

				fmt.Printf("User %v booked %v tickets\n", userName, userTicket)
				fmt.Printf("Thank you %v %v for booking with us. You eill get a confirmation email at %v\n", userName, lastName, email)
				fmt.Printf("%v tickets remianing for %v\n", remainingTickets, confenrenceName)
}

func sendTicket(userTicket uint, userName string, lastName string, email string){
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v ", userTicket, userName, lastName)
	fmt.Println("##########")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("##########")
}









/*
		
		city := "London"

		switch city {
			case "New York":
				// execute code for booking New York conference tickets.
			case "Singapore":
				// execute code for booking Singapore conference tickets.
			case "London":
				// execute code for booking London conference tickets.
			case"Berlin","Nigeria":
				// execute code for booking Berlin conference tickets.
			case"Mexico City":
				// execute code for booking Mexico City conference tickets.
			case"Hong Kong":
				// execute code for booking Hong Kong conference tickets.
			default:
				fmt.Println("No valid city selected")
		}
		
		*/