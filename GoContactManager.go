package main

import (
	"fmt"
)

// Contact struct to represent a contact
type Contact struct {
	Name  string
	Email string
	Phone string
}

// ContactList struct to manage contacts
type ContactList struct {
	Contacts []Contact
}

// AddContact adds a new contact to the contact list
func (cl *ContactList) AddContact(name, email, phone string) {
	contact := Contact{Name: name, Email: email, Phone: phone}
	cl.Contacts = append(cl.Contacts, contact)
}

// ListContacts lists all contacts in the contact list
func (cl *ContactList) ListContacts() {
	fmt.Println("Contacts:")
	for i, contact := range cl.Contacts {
		fmt.Printf("%d. Name: %s, Email: %s, Phone: %s\n", i+1, contact.Name, contact.Email, contact.Phone)
	}
}

// SearchContacts searches for a contact by name
func (cl *ContactList) SearchContacts(name string) {
	fmt.Println("Search Results:")
	found := false
	for i, contact := range cl.Contacts {
		if contact.Name == name {
			fmt.Printf("%d. Name: %s, Email: %s, Phone: %s\n", i+1, contact.Name, contact.Email, contact.Phone)
			found = true
		}
	}
	if !found {
		fmt.Println("No matching contacts found.")
	}
}

// DeleteContact deletes a contact from the contact list by index
func (cl *ContactList) DeleteContact(index int) {
	if index >= 0 && index < len(cl.Contacts) {
		cl.Contacts = append(cl.Contacts[:index], cl.Contacts[index+1:]...)
	}
}

func main() {
	contactList := ContactList{}

	for {
		fmt.Println("\nContact List Menu:")
		fmt.Println("1. Add Contact")
		fmt.Println("2. List Contacts")
		fmt.Println("3. Search Contacts")
		fmt.Println("4. Delete Contact")
		fmt.Println("5. Quit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var name, email, phone string
			fmt.Print("Enter contact name: ")
			fmt.Scanln(&name)
			fmt.Print("Enter contact email: ")
			fmt.Scanln(&email)
			fmt.Print("Enter contact phone: ")
			fmt.Scanln(&phone)
			contactList.AddContact(name, email, phone)
			fmt.Println("Contact added!")
		case 2:
			contactList.ListContacts()
		case 3:
			var name string
			fmt.Print("Enter name to search: ")
			fmt.Scanln(&name)
			contactList.SearchContacts(name)
		case 4:
			var index int
			fmt.Print("Enter the index of the contact to delete: ")
			fmt.Scanln(&index)
			contactList.DeleteContact(index - 1)
			fmt.Println("Contact deleted!")
		case 5:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
