# Overview
<p>If you are wondering how to make correct Dockerfile for web applications that wrote in GO Lang you are in the right place
In this repository i show you guys a few simple ways that you can build your own Dockerfile and use it wherever you like.Actually there is tons of websites and topics that explain Docker way better than me BUT if you use a regular Dockerfile for templates there will be a little problem with you code.
Container won't find you template unless you put it in root of the project. and thats because usually Docker manuals for GO dont tell you to copy the whole template folder into the Docker and if you are new to Docker(such as me) it can get you in trouble.</p>

# Manual 
<p>Only thing you have to do is simply ADD your templates directory into the file and BOOM!</p>
To do that use this command:<br>
"ADD {A} {B}"<br>
This command will copy A to B, Very simple<br>
in this case im using <br>
"ADD templates ./templates"<br>
And it will copy the templates folder to ./template(into the Docker Image)<br>
And then simply tell your template handler to handle any HTML template inside this directory<br>
To do that you can use this function in you GO file:<br>
	templates = template.Must(template.ParseGlob("templates/[a-z]*.html"))
  
And its Done ;)
# Sample
#syntax=docker/dockerfile:1

FROM golang:1.15-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY ./templates ./

ADD templates ./templates

COPY ./* ./
RUN go build -o /docker-gs-ping

CMD [ "/docker-gs-ping" ]

EXPOSE 9090

P.s:I use these names for my directory on NO speciall perpose, You can use anything you want ;)
