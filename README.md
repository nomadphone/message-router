# message-router

[![build](https://github.com/nomadphone/message-router/actions/workflows/build.yml/badge.svg)](https://github.com/nomadphone/message-router/actions/workflows/build.yml)  [![codecov](https://codecov.io/github/nomadphone/message-router/branch/main/graph/badge.svg?token=8DRRQNYPLU)](https://codecov.io/github/nomadphone/message-router)

------

Microservice that is part of [nomadphone](github.com/nomadphone), responsible for routing messages from a user local phone number to another messsaging platform.

## Features

- Routes messages from a Twillio phone number to a Telegram user

## API Endpoints

- /twillio/sms - Webhook for Twillio's application, receives data in x-url-encoded format and prints a TwiML response for Twillio's servers.
