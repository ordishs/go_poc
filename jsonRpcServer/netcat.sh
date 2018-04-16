#!/bin/sh

nc 127.0.0.1 9999 <<EOL
{"method":"Handler.Hello","params":[{"firstName": "Bob", "lastName":"Smith"}],"id":0}

EOL