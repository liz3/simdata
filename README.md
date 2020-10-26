# SimData
The Golang part listens for the UDP telemetry from F1202 and Project cars 2 & saves them into a reduced json format with custom header.

## Displaying
You may use graph.py to display 2 laps with hardcoded params, but atm the script is not really dynamic, it uses sdl2 to render so sdl2, pysdl2 and so on are required.
Use like
```
python(3) graph.py /path/to/log.sd 2 5
```
Where 2 and 5 are the desired laps

## Note
- Note that the log files can get quite big, also remember to create the logs folder in the project root
