# Application Skeleton: Survey

![screenshot](screenshot/gif.gif)

This website was created to get hands-on experience with the various technologies it comprises, including HTML, CSS, JS, and GoLang. The objective was also to familiarize ourselves with different GoLang libraries, among the most important ones:

- database/sql
- net/http
- github.com/gorilla/websocket
- html/template
- github.com/gorilla/sessions

## The Survey Site
The development of this site was very interesting. Here's a list of the main features it contains:

**Ability to create surveys:** This feature allowed us to extend the use of the html/template library by using loops and conditionals in HTML files and sending information to HTML files.

**Real-time survey results:** Here, the Gorilla WebSocket library was used, so users don't need to refresh the page when a new user votes in the survey; it updates in real-time.

**Password recovery system:** Another interesting task included in the site is the ability to recover a password using the email address provided during registration on the website.

**Survey expiration:** The use of TIMESTAMP in our PostgreSQL database tables allowed us to create a survey expiration system to have a list of recent surveys.

**User session:** A session system has been included in the site, preventing users from having to reconnect with each page refresh.

## Site Architecture
The site consists of different pages:

- **"/"** Home page containing navigation between the various pages
- **"/signup"** To create a user
- **"/login"** To log in to your account
- **"/create_survey"** To create a survey
- **"/survey_list"** To view the list of different surveys created
- **"/survey/?"** Page to vote in a survey; each survey has its own ID in the URL
- **"/survey_results/?"** Page to view the results of a survey; each survey has its own ID in the URL

## Running the Program
First, to launch the site locally, you'll need to install the various libraries included in the project using the `go mod tidy` command.

### Creating the Database

Next, you'll need to create a "copy" of the database to use. Install PostgreSQL and execute the following commands:

`sudo -u postgres psql`<br/>
`CREATE USER sondage WITH PASSWORD bood7Ees;`<br/>
`ALTER USER sondage SUPERUSER;`<br/>
`CREATE DATABASE sondage_db OWNER sondage;`<br/>
`\q`<br/>
`sudo -u postgres psql sondage_db`<br/>
`CREATE TABLE users (id SERIAL PRIMARY KEY, username VARCHAR(50), password VARCHAR(100), mail_address VARCHAR(100));`<br/>
`CREATE TABLE sondages (id SERIAL PRIMARY KEY, titre VARCHAR(100), description TEXT, choix1 VARCHAR(100), choix2 VARCHAR(100), choix3 VARCHAR(100), choix4 VARCHAR(100), date_creation TIMESTAMP, date_expiration TIMESTAMP, nb_choices INT, createur_id INT, FOREIGN KEY (createur_id) REFERENCES users(id));`

### Launching the Site
To run the program, simply execute the run *.go command while in the "src" directory. You can then access the site via the localhost address (or the IP address hosting the program) on port 8080 (http://localhost:8080/).
