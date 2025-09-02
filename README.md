# JOUSTER TAKE HOME ASSIGNMENT

## Setup and instructions
Using golang 1.25, you can run it with go if already installed locally or you use docker.

- Copy the .env.example in .env and set the openai API_KEY
    ```shell
    cp .env.example .env
    ```
- Using go installed locally, the app will be available on [localhost](http://127.0.0.1:3000)
    ```shell
    go run *.go
    ```
- Using docker, the app will be available on [localhost](http://127.0.0.1:3000)
    ```shell
    docker build -t jouster . # to build an image
    docker run -p 3000:3000 jouster # to run the built image
    ```

## The why of tools and choices
I used golang for, well it's my favorite language for anything and everything, then I used openai's chatgpt4mini for the LLM but it could have been any other LLM it doesn't matter for the task at hand. The code structure I have is meant to make the project easy to understand and straight forward to navigate, I didn't want to add boilerplates of services/handlers/requests/... as this is just a quick demonstration of an LLM implementation. I kept the main file as simple as possible for better readability, then types in their own file for quick reference, same for functions that give the feeling like they could be used somewhere else into the helpers file, then the business logic into the handler file. The handler could have been split into two files if it was a bigger project, namely handlers and services where handlers would be for infrastructure related functions like interacting with sessions, ... and services would be the one containing the business logic.

## Trade-offs
Well, I didn't add support for a database nor add tests as I wanted to respect the proposed time of two hours but I would like to emphasize that tests are A MUST and AN ALWAYS for me when coding but for the time being I had to make a choice. However I decided to add docker for the sake of making it work on other machines that don't have golang installed or don't want to install it. What actually took me time was playing with different NLP packages available in golang and finally decided to go with prose as it's lightweight and simple to use, although it's archived it does a great job at least for this use case.