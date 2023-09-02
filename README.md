# Contacts_Saving_Application

This Go (Golang) code is a simple command-line program for managing a contact list. Here's a description of the code's functionality and structure for GitHub:

Functionality:
1. Contact Struct: This code defines a Contact struct to represent individual contacts. A contact has three fields: Name, Email, and Phone.
2.ContactList Struct: It defines a ContactList struct to manage a list of contacts. The Contacts field is a slice (dynamic array) of Contact structs.
3.AddContact Method: The AddContact method allows users to add a new contact to the contact list. It takes name, email, and phone as parameters and appends a new Contact to the Contacts slice.
4.ListContacts Method: The ListContacts method displays all contacts in the contact list, printing their names, emails, and phone numbers.
5.SearchContacts Method: Users can search for a contact by name using the SearchContacts method. It prints the matching contact(s) or a message if no matching contacts are found.
6.DeleteContact Method: The DeleteContact method allows users to delete a contact from the contact list by specifying its index. It removes the contact at the specified index from the Contacts slice.
7.Main Function: The main function serves as the entry point of the program. It initializes an empty ContactList and presents a menu to the user. The menu options include adding a contact, listing contacts, searching for contacts, deleting contacts, and quitting the program.

Usage:
1. Users can run this program to maintain a contact list through a simple command-line interface.
2. They can add contacts by providing their name, email, and phone number.
3. They can list all contacts in the contact list.
4. They can search for a contact by name and view the results.
5. They can delete a contact by specifying its index in the list.
6. The program allows the user to quit and exit gracefully.
