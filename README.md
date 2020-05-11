# AccountsListGenerator
 Tool to create account names based on real names

Account name generation rules:
 1. Complete first name plus last name
 2. Initial of first name and complete last name
 3. First three characters of the first name and first three of last name
 4. First three characters of the first name and complete last name
 
 Usage: go run ./main.go users.txt
 
 Input file:
 
  Bill Murray
  
  Mike Myers



 Output:
 
  BillMurray
  
  Bill.Murray
  
  ...
  
  Mye.Mike
  
  Mye-Mike
  