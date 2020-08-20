## Go Play Racquetball

![logo](assets/logo.png =250x)

If you are as forgetful as I am then, you never remember to reserve your racquetball courts.  Well now your time of forgetting to reserve courts 
is over! (if you are a la fitness member).  go-play-racquetball is an attempt to completely automate your recurring racquetball court reservations so
you don't have to remember to.  

## Running

Go Play Racquetball can be run from the [ddelnano/go-play-racq](https://hub.docker.com/r/ddelnano/go-play-racq/) docker image hosted on docker hub.  All that is needed to run the application is an `.env` file that specifies your La Fitness username and password and a `config.json` file that specifies your racquetball reservation schedule.  A sample of each is shown below.

.env file
```
LA_USERNAME: username
LA_PASSWORD: password
```

config.json
```json
{
  "reservations": [
    {
      "day": "Wednesday",
      "time": "14:00",
      "clubID": "1010",
      "clubDescription": "PITTSBURGH-PENN AVE",
      "duration": "60"
    }
  ]
}
```

The config.json file is validated through JSON schema and the exact constraints can be seen [here](reservation.json).

Once you have both of those files the docker image can easily be run mounting volumes that contain the config.json and .env file like so.

```bash
docker run -v /path/to/.env:/src/.env -v /path/to/config.json:/src/config.json ddelnano/go-play-racq:v0.3.1
```


## Credits

Thanks to Wai for his awesome work on the logo! You can see more of his work [here](http://waitu62.wix.com/designer-graphic).

![logo](assets/logo.png =100x)

go-play-racquetball's logo is licensed under the Creative Commons 3.0 Attributions license.

The original Golang Gopher was designed by [Renee French](https://reneefrench.blogspot.com/).

## Getting to v1.0

### Todos

- [ ] Schedule all unreserved court times between the current day and 2 weeks from the current day.  This is how far in advance la fitness allows and would allow for scheduling as early as permitted.
- [ ] Remove need to know clubID, clubDescription?
