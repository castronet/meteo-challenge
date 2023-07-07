# README

This file will explain what challenge I have chosen and how to run it.

## Challenge chosen

I was thinking about the option to save the data in a SQLite database and then read that info and display it somehow.

But at the end I decided the option to create a HTTP Server to server the data requested to another API. I think that this option involves more technology,
I mean that you should run a http server, you have to handle the requests, you also have to read data from an API (and know how to manipulate it) and then have some engine to show the result as an HTML file.

## How to run the code

First you should get the code from this repository, then access to the folder with the code and run:

```
make run
```

If you need more information try with

```
make help
```

## About testing

I had to admit two things:
- The first one is that 2 hours are not a lot of time.
- The second point is that I never coded tests in Go. But I read about it and I would want to implement table driven tests. For example to the function that read the lat/lon value from the URL, I could add a test to check the function behavior if the lat/lon are not numbers or incorrect numbers.

I want to add that as a DevOps/SRE for me the tests are one of the most important things because they let me know before deploy the code on production (or other environments) if the code is broken or all functions works as previous versions.
Tests also could be executed automatically on an CI/CD pipeline just after build step and discard the branch/deployment if the tests fail.

## Conclusions

About documentation I decided to add inline documentation to try to clarify the code. I also tried to choose a good name for each function. There are better documentation options and/or I could add some kind of tags on functions documentation to be able to generate documentation automatically with i.e. Gendoc/Swagger...

I have to admit that I thought that in two hours I will be able to be able to have more code done on this project. For example, a part of do some tests, I would like to add some kind of graph on frontend but I decided that it was out of the scope of this project so I decided to focus on the code and documentation.

This code has some parts that was not required (for example I don't really need the code on the "main.go" file but it is useful on bigger projects) and I miss frontend and tests/code coverage parts. That is what I will do if I had more time.

To develop this test I used neovim (running on a terminal) with Codeium plugin ( https://codeium.com/ ) it uses some kind of AI like Copilot, but is a free and less powerfull tool. I can't say what parts are done by it because it disturbs more than it helps, not like Copilot.

I spend about 2 hours on the project and at least 25 minutes with the readme and the code delivery.

## Requirements

I used go version 1.20+ to develop the challenge. Is the Golang version installed on a MacOs via brew.
