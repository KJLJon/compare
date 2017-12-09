# compare
compare interface{} values to see if it matches the criteria or criteria groups

## Idea
The basic idea is I find myself having user/admin driven rules that I need to compare against individuals to see if what groups they belong to.

This works well with JSON

## Use case Example
Lets say you are creating an emailing application.

Your users upload their contacts and demographics (name, age, email, ...etc...)

In this type of application, the user will want to segment their contacts into different campaigns.

This tool can be used to compare users demographic information and return what segments they belong to based on the user defined rules.

## Example
todo include example that uses this json as input

### Segment Rules Input
```json
{
  {
    "name": "Jon-child",
    "criteria": [{
      "key": "/profile/first",
      "operator": "=",
      "compare": "Jon"
    }, {
      "key": "/age",
      "operator": "<",
      "compare": "18"
    }]
  }, {
    "name": "Jon-adult",
    "criteria": [{
      "key": "/profile/first",
      "operator": "=",
      "compare": "Jon"
    }, {
      "key": "/age",
      "operator": ">",
      "compare": "18"
    }]
  }, {
    "name": "first-Jon",
    "criteria": [{
      "key": "/profile/first",
      "operator": "=",
      "compare": "Jon"
    }]
  }
}
```

### User Demographics Input
```json
{
  "profile": {
    "first": "Jon",
    "last": "Doe",
  },
  "age": 20,
}
```
