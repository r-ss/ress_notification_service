# ress_notification_service

Personal utility to send me telegram notifications from various scripts, services etc.

## Stack used here:
- Go
- [Gin web framework](https://github.com/gin-gonic/gin)
- Telegram API

## Goal

When you have some apps and services running here and there, they send you messages, statuses, alarms. Mostly via Telegram. Easiest way is directly touch Telegram API from app, but here are problems when apps-zoo grows up. Some services may send you message at night and you don't want it. Another service send you message at night and you need it. So width Extra services that can manage notification you have possibility to filter messages by own rules, collect them etc.


                                                  │  Extra service gives flexibility   
                                                  │                                 
                ┌───────┐  ┌───────┐  ┌───────┐   │  ┌───────┐  ┌───────┐  ┌───────┐
                │ app 1 │  │ app 2 │  │ app 3 │   │  │ app 1 │  │ app 2 │  │ app 3 │
                └───────┘  └───────┘  └───────┘   │  └───┬───┘  └───┬───┘  └───┬───┘
                    │          │          │       │      │          │          │    
                    │          │          │       │      └─┐        │        ┌─┘    
                    │          │          │       │        │        │        │      
                    └──────┐   │   ┌──────┘       │        │        │        │      
                           │   │   │              │        ▼        ▼        ▼      
                           │   │   │              │     ┏━━━━━━━━━━━━━━━━━━━━━━━━┓  
                           │   │   │              │     ┃  Notification Service  ┃  
                       ┌───▼───▼───▼───┐          │     ┗━━━━━━━━━━━┳━━━━━━━━━━━━┛  
                       │  TelegramAPI  │          │                 │               
                       └───────┬───────┘          │         ┌───────▼───────┐       
                               │                  │         │  TelegramAPI  │       
                               │                  │         └───────┬───────┘       
                               ▼                  │                 ▼               
                          ┌─────────┐             │            ┌─────────┐          
                          │┌───────┐│             │            │┌───────┐│          
                          ││       ││             │            ││       ││          
                          ││       ││             │            ││       ││          
                          ││ Phone ││             │            ││ Phone ││          
                          ││       ││             │            ││       ││          
                          ││       ││             │            ││       ││          
                          │└───────┘│             │            │└───────┘│          
                          │         │             │            │         │          
                          └─────────┘             │            └─────────┘          
                                                  │                                 
                                                  │                                 


## TODO

Started as beginner practice in Golang and Amazon Web Services, I liked that and sometime will add functionality to filter specified notifications according to time and priority.

- Switch GIN from DEV to Release mode
- Turn off AWS CloudWatch Logs

## Request example:
    curl --header "Content-Type: application/json" \
    --data '{"message": "Hello World 🐶"}' \
    https://notification.custom.domain/telegram
