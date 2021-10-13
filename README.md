
# Food For You
Food for you is an online kitchen, which gives people the ability 
to put together the ingredients, and all the necessary things to 
cook a meal. It also gives people the ability to customer their 
food order. It is also an online food ordering platform.

It is built using golang



## Installation

Install foodforme with the following go command

```bash
  go run github.com/pluralsight/foodforme
```
    
## Contributing

Pull requests are welcome. For major changes, please open an 
issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

  
## API Reference

#### Get all items

```http
  GET /getIngredients
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `` | `` | |

#### Get item

```http
  GET /getIngredientById/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

#### add

```http
  POST /addIngredient
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `Name`      | `string` | **Required**. Name of the ingredient |
| `Category`      | `string` | **Required**. Category of the ingredient |

  
## Architecture Design

Clean Architecture

DDD

## Design Patterns

Repository Design Pattern

Interface segregation principle

Single Responsibility principle

CQRS

## Tech / Framework used

Golang
MySql
SQL

## Features

Add Ingredients of a meal

Fetch the ingredients

Fet all ingredients

## Tests

This is a placeholder for a preliminary version

## Code Example

This is a placeholder for a preliminary version




  
