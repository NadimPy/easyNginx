# Planning file for project

This is a planning file to track what features will be targeted in the project to see what features for the TUI as NGINX is quite big.

Therefore I plan on the following list of initial features to add:

## Static Site Setup	
An interactive option that requests the 

## SSL/HTTPS Setup	Certificate paths, HTTP→HTTPS redirect, basic TLS settings
simplified certbot setup (to check if cerbot acts different for APIs, dynamic websites -> encrypting )

## Basic Security	
Rate limiting setup, IP allow/deny 
interesting feature to rate limit the amount of requests that can be done to your dynamic website

## I plan on adding after those features the load balancer feature when I learn it

---
# Along the way hurdles
## Conf path
I believe validating and revalidating the conf path every time will help a lot as the service user (easyNginx) has elevated permissions (and to prevent creation of files) and NGINX conf file can differ


``` command 
 nginx -V 2>&1 | awk -F: '/configure arguments/ {print $2}' | xargs -n1
```

