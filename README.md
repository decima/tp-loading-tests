# Charge App
This app is designed to be used for charge testing.

Available routes : 

- /p1 : normal gin behavior
- /p2 : with 10 mutex locks and sleep 100ms
- /p3 : with 10 mutex locks and random sleep time
- /p4 : with random generation between a number of 1 and 10000 strings of 100 runes (CPU intensive)