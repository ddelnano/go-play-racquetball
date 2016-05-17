## Go Play Racquetball

If you are as forgetful as I am then, you never remember to reserve your racquetball courts.  Well now your time of forgetting to reserve courts 
is over! (if you are a la fitness member).  go-play-racquetball is an attempt to completely automate your recurring racquetball court reservations so
you don't have to remember to.  

If you are still reading this I am surprised you haven't stopped and thought this is the most useless project ever.

This is was an excuse to try out Go and write something simple.

## Getting to v1.0

Create cronjob that runs go app every day at midnight.  It will schedule the reservation for 2 weeks from the current date.  It should be able to set a username and password so that this would work for multiple users of the app.

List of Todos
- [ ] Dockerize
  - [ ] Container that builds the binary for the cronserver
  - [ ] Cronjob server
    - [ ] Runs go binary at midnight
    - [ ] Runs integration tests during the day
- [ ] Add ability to make reservations different user for a given location

## Long term goals
- [ ] Get diff of currently scheduled reservations and future ones
- [ ] More robust logging and error handling
- [ ] Use sms to update recurring reservations through Go web service.
