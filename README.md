# Customers API

This project is small API.\
You can see list of customers, add, edit and delete customers.

## Implementation

In api folder main file is main.go, it contains main functional and logic of project. 
Config directory contains connection with database and global constants. 
Model directory contains models (work with database tables). 
Helper helps project (some functions). 
Html directory is directory for all templates (html, css, javascript, images).

## Requirements

 - Go 1.8
 - glide
 - postgres database

## Running/Installation

- Clone this repository into src directory of your go path.
- Install all the respected dependencies in to your local vendor (glide install).
- Build project from api directory.
- Run project.

## Settings

Change some constants in config/constangs.go. For example, path to root directory.

Change database settings for connection in config/database.go

After first running You can comment initDataDb and initLocalesDb calls in main.go file.