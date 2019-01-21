"Go Programming Challenge
This programming challenge is designed to test your skills and
willingness and ability to learn new things. For this programming
challenge, you will build a simple chat server in Go (Golang)."

Instructions:
After installing necessary dependencies, to start the chat server, input into a CLI the following command, "go run main.go"
To connect to the chat server, for each user, input the following command in a separate CLI tab, "telnet localhost 23"

If you have any questions at all, don't hesitate to contact me.

My approach to the assignment:
Since one of the two stated goals of this programming challenge is to test my willingness and ability to learn new things, I approached the problem like I approach any problem: process oriented rather than results oriented.

This is my typical approach but considering I have so much more control over my attitude and effort rather than current skills, this was the goal I tried to attack head on.

When it comes to tackling problems with code, this process oriented approach usually means something along the lines of:
1.) Learn enough fundamentals to start coding and to start asking intelligent questions.
2.) Ask myself those intelligent questions to see if I can come up with an answer before looking to outside sources.
3.) Trust myself and trust the process.  I know that with hard work and even more importantly, FOCUS, I can tackle pretty much anything that's put in front of me.  Rather than try to take shortcuts optimizing my code/coding practice for only production speed, where production speed is how fast I can write code, I make sure to pay attention to other signals of high code quality/coding practice such as code smell, testing, maintainability, security, documentation and most important of all, my overall understanding.
4.) Diversify my time-investments.  As is the same as with any resource investment, I think it's important to diversify where my time in trying to understand a difficult concept is being spent.  Properly surveying high quality resources, checking documentation, making appropriate google/stack overflow search queries, youtube videos, udemy courses, meetups and mentors are all valuable resources.  I think it's particularly important to not rely on any one area and to make sure that, as I learn the what, I learn the how and the why as well.  I also think that learning and teaching go hand in hand.  As I learn more, I think it's important that I increasingly take on teaching/mentorship opportunities where appropriate.  Exercising as many parts of my brain when learning increases understanding and long term memory.
5.) Last but certainly not least: Don't stay stuck.  Don't beat yourself up over something you don't understand. Don't be afraid to ask good questions.  Since it was a coding challenge, I didn't ask any questions but I made sure to utilize all of the resources the internet has to offer both early and often.

App features:
1.) Simple chat server in Go.
2.) Multiple clients can connect via telnet (tested 5).
3.) Clients can send messages to the server and the server relays those messages out to all connected clients.
4.) Messages contain username and timestamp.
5.) All messages are logged to chatLog.json using logrus package.  Informational logs are recorded as well as warning logs and fatal logs.  The log level is set to fatal.  Using json fileformat maximizes future compatibility with other logging aggregators such as Splunk.
6.) App reads configs from a config file located at ./config.json.
7.) Thorough documentation.

Sources:
golang documentation:

Official docs
https://golang.org/doc/effective_go.html
https://golang.org/doc/code.html
go routines
https://gobyexample.com/goroutines
select statements
https://tour.golang.org/concurrency/5
configuration files
https://www.thepolyglotdeveloper.com/2017/04/load-json-configuration-file-golang-application/
https://stackoverflow.com/questions/16465705/how-to-handle-configuration-in-go
see "App features".5 for more information about json configuration and why I chose the file format.
Time
https://golang.org/pkg/time/#Time.Format
Go security vulnerabilities
https://hunter2.com/use-golang-these-mistakes-could-compromise-your-apps-security
Go class
https://github.com/GoesToEleven/GolangTraining/
Go Chat server
(some great code here that I could have used to implement direct messaging users, multiple chat rooms, muting users, restful API)
https://github.com/bentranter/chat
Structs
https://gobyexample.com/structs

PACKAGES
bufio
https://golang.org/pkg/bufio/#Scanner
io
https://golang.org/pkg/io/#MultiWriter
logrus
https://godoc.org/gopkg.in/Sirupsen/logrus.v0
https://esc.sh/blog/golang-logging-using-logrus
https://blog.scalyr.com/2018/06/go-logging/
Why did I choose logrus over glog?  I initially flip flopped between using google's glog logging module and logrus.  I didn't personally check but read that logrus is updated much, much more frequently.  I chose logrus for this reason.
os
https://golang.org/pkg/os/


Limitations:
1.) testing - TDD is so important at every level of the production process and I haven't written a single test
2.) security considerations - since this app has no potential to go live, the risk of damage to an organization is minimal.  It's still important to consider how any code, or processes can open up a machine or network to risk.  For instance, I'm not sure if tcp telnet connections automatically close.  What happens if I forget to close my server after I'm done with it?  Am I opening up myself, my company to risk?
3.) Logging is minimal especially in the case of error.
4.) The chat application provides an interface so that people can chat but the app does not provide users a way to claim an identity (outside of choosing a username), to authenticate that identity, nor does it provide any semblance of non-repudiation.  If somebody says something via the chat app, without any mechanism to provide identification, authentication, non-repudiation, the chat app really doesn't provide much in the way of secure communication.
5.) I'd like to provide a way to clear the log file manually, or automatically after a certain period of time has passed.
6.) Not the best user interface.  I started to code a front end using websockets but decided it was outside the scope of this exercise (not to mention, I needed to turn it in!)
7.) Creativity.  Although there was no hard deadline for this exercise, I knew that I couldn't work on it forever.  I spent longer on this than I thought that I would as golang is so different than javascript, php, and python.  I think that I tackled the basic functionality reasonably well.  I would have loved to expand a bit more by adding Limitations.1, Limitations.2, and creating a restful API... all things which were well within my capabilities as an engineer.

Bugs:
1.) Chat log sometimes doubles up on log entries when a user enters/exits the chatroom.
2.) When a user is chatting, I would prefer that the line the user types the message on, the command line, was the same line that contains the user's username and the timestamp of the message.
3.) When new users are logging into the chat application, sometimes the user labels when chatting are incorrect for a short period of time.
