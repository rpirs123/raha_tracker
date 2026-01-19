# raha_tracker
Built this project since i was too lazy to manually track my monthly expenses. 

Since I mainly use Swedbank for my daily expeneses and wanted to do something with the GO programming language I came up with this raha_tracker project. 

#### NB: WHEN YOU USE THIS PROGRAM THE PROMPTS YOU RECEIVE ARE IN ESTONIAN LANGUAGE. 

####  If you wish to change the language clone this repository and change the prompt labels inside lib/cli/cli.go


## How it works 

The app itself is quite simple, it takes in the file path to your bank statement as an argument, the statement has 1 requirement which is that it has to be in CSV format.

Then the program generates a config file inside the folder the program was run from, this file is used to save transaction vendors and categorize them in future runs. 
##### NB: it only creates the config if it doesnt exist, if it does it reads from it rather than overwriting with a blank state each time. 

After that it looks at each transaction the person has made or received. This means that both expenses and income are read and displayed at the end. 

For each transaction it looks at a config file, and then tries to find the vendor and category inside that file. The file is in INI format. 
 
 The inside of the file would look something like this for the vendor Wolt: 

 ```
[wolt]
kategooria = söök
 ```

 Wolt is the section name and it contains ```category = food``` in Estonian. This means that the program would auto save any transactions from Wolt as food.

 If ofcourse the vendor and category is not saved inside the config, it will prompt the user to decide if they want to save the vendor or not : 

 ```
? Tehingu kaupmees puudub, salvesta kaupmees?:
    jah
  > ei
 ```

 If the user chose yes, it will prompt the user to save the vendors name, this will end up being the section name in the config file. Keep in mind 
 that the full name of the vendor does not have to be input here, just a small part that connects the vendor to the section Name. 

 An example of that would be: 

 ```
Vendor name: RIMI/TARTU LOUNAKESKUS 50501 --> Section name: rimi
 ```

Additionally after the user inputs the vendor name they want to save, the program will ask what category the user wants to make the transactions from this vendor go to: 

```
Sisesta kategooria mille alla salvestada
Et lisada mitu kategooriat, pane kategooriate vahele koma NÄIDE: söök,meelelahutus
> 
```

The program accepts as many categories as you would like: 

```
// if the user wants to save a single category
> food 

// if the user wants to save multiple categories, each seperated by a comma.
> food,rent 

```

Thats it! you have successfully saved your vendor and category you wish to connect with that specific vendor. 

However if you chose no during the vendor saving prompt, the program will only ask you to enter a category you wish to save the transaction to, keep in mind if you did it this way you can only input 1 category


## How to run the program 


Prerequisites: 

 (All csv statements work as long as they follow the column format defined inside lib/csvcolumn.go) 

1. A CSV formatted bank statement from Swedbank.
2.  .exe file 

To build the exe file simply run the following command: 

```
go build -o raha_tracker.exe main.go
```
This generates a raha_tracker.exe inside the root of the repository which can be run.

After you have the requirements filled 

1. Open up a terminal or command prompt

2. Navigate to the root of the repo using cd  
        
 ```
 cd path/to/your/project
 ```

 3. run the exe file with this command: 

 ```
 .\raha_tracker.exe "C:/path/to/your/statement.csv" 
 ```

 Congratulations if you followed the guide correctly then the program should work. 

 In case u get an error and dont know how to solve it open up an issue and we will get it fixed  :) 





