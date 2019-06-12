#  Use Cases:
*Broken up by business services to mimick a microservices format. Each subcategory could be an independent server*

## Frontend
**Title:** Show landing

**Actor:** Potential Client

**Scenario:** 
* Hits our domain
* Server renders landing with "Contact Sales" form
* Potential client enters email and brief company info and submits form
* Send email to our sales team
* Send success email to potential client

**Preconditions:**
* Email and company info must be valid

## OfficialPersonalityTest
**Title:** Create personality profile based of user submitted personality test

**Actor:** Profile User

**Scenario:** 
* User starts our personality test
* User completes questions
* User enters email
* Present user results of their test
* Send the collected data, email, encrypted deviceId to the pipeline   

**Preconditions:**
* Email must be valid

## Pipeline
**Title:** Create personality types for profiles in a given CSV 

**Actor:** User

**Scenario:** 
* User uploads CSV to our s3 bucket
* S3 upload triggers lambda
    * Loops through all of the rows in the CSV
    * Format data - chopping block
    * Check for internal match
    * Hit Full Contact > Pipl > People Data Labs to find basic information on user, cache API Call
    * Hit Signal Hire to find social media profiles, cache API Call
    * Hit Watson for Watson Score, cache API Call
    * Hit V12 for PYCO score, cache API Call
    * Send success email

**Preconditions:**
* CSV must be in valid format
* User must be authenticated

**Extensions:** 
* If one of the API Hit's data is already in the internal profile, skip
* If error, send error email. Mark *Record* entity as error

## REST API
**Title:** Recieve personality type of site viewer

**Actor:** Developer of client company

**Scenario:** 
* Hits our API with viewer information
* Our server searches our internal database to find user
* Returns personality type in JSON format

**Preconditions:**
* Viewer information must be in valid format - email, deviceId ...
* API hit must have a valid API key

**Extensions:** 
* If user not found return 4xx error message


**Title:** Update profile if conversion completed

**Actor:** Profile User

**Scenario:** 
* Profile User makes conversion by clicking an advertisment or a specific element
* Hits our API with viewer information and campaign information
* Our server finds this profile
* Our server adds this conversion data to the profile
* Updates the profile personality score

**Preconditions:**
* Viewer information must be in valid format - email, deviceId ...
* API hit must have a valid API key

**Extensions:** 
* If user not found create a user with the email and deviceId, send to pipeline


## Profile Management
**Title:** Show a report based on the CSV profiles

**Actor:** User

**Scenario:**
* User is presented the dashboard
* Sees the overall breakdown of personality types
* Click into a list and overall breakdown of that list
* Click into a personality type and see the profile data for each personality type

**Preconditions:**
* User must be authenticated
* User must have uploaded a CSV previously
* The pipeline must have finished for this list

**Title:** Show a report based on the CSV profiles

**Actor:** User

**Scenario:**
* User clicks into a list
* Clicks export
* Server prepares a CSV with the profiles with personality types
* The server sends the download

**Preconditions:**
* User must be authenticated
* User must have uploaded a CSV previously
* The pipeline must have finished for this list

## User Management
**Title:** Create company

**Actor:** New Client

**Scenario:** 
* Fills out form with company info
* Fills out form of admin user
* Completes billing information
* Creates Company Entity in DB
* Logs in new user

**Preconditions:**
* Company info, user info, billing information, must be valid

**Title:** Add users to company

**Actors:** Company Admin and New User

**Scenario:** 
* Goes to settings page
* Selects Users
* Selects add new user
* Fills email field and submits
* Server recieves email and creates a user object with this company
* New user recieves an email with a link to a basic info form
* Adds information and submits the form
* Server completes the user entity and returns a success message

**Preconditions:**
* Admin user must be authenticated
* New user's email must be valid

**Title:** Get API Keys

**Actor:** Developer

**Scenario:** 
* Goes to settings page
* Selects API Keys
* View and Copy API Keys

**Preconditions:**
* User must be authenticated

**Extensions:**
* If the API Keys have been corrupted, click refresh icon to generate new keys

### To contribute to this file, follow this format:
**Title** - the goal - short sentence that starts with an active verb

**Actor** - the *entity* involved in this goal

**Scenario** - Bullet point step-by-step description of how the system and actor achieves this goal.
Optional:
* **Preconditions** - conditions that must hold true beforehand
* **Extensions** - additional scenario steps for irregular/specific cases
