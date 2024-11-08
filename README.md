# command-constructor

# Content <!-- omit from toc -->
- [command-constructor](#command-constructor)
	- [API V1](#api-v1)
		- [Types](#types)
			- [User](#user)
			- [CommandTemplate](#commandtemplate)
			- [CommandParam](#commandparam)
			- [CommandParamType](#commandparamtype)
		- [Endpoints NonVersioned](#endpoints-nonversioned)
			- [/reg POST](#reg-post)
			- [/auth POST](#auth-post)
		- [Endpoints Versioned](#endpoints-versioned)
			- [/command GET](#command-get)
			- [/command POST](#command-post)
			- [/command/:id PUT](#commandid-put)
			- [/command/search/:name GET](#commandsearchname-get)
			- [/user DELETE](#user-delete)
			- [/user GET](#user-get)
			

## API V1

### Types

#### User
| Name    | Type   | Desctription        |
| ------- | ------ | ------------------- |
| name    | string | unique name of user |
| email   | string | email of user       |
| passord | string | password of user    |

#### CommandTemplate
| Name           | Type                                  | Desctription                    |
| -------------- | ------------------------------------- | ------------------------------- |
| name           | string                                | command template name           |
| description    | string                                | description of command template |
| commandName    | string                                | name of command                 |
| templateParams | [][types.CommandParam](#commandparam) | changeable params               |
| constantParams | [][types.CommandParam](#commandparam) | nonchangeable params            |

#### CommandParam
| Name         | Type                                  | Desctription                          |
| ------------ | ------------------------------------- | ------------------------------------- |
| name         | string                                | param name                            |
| description  | string                                | description of param                  |
| value        | []string                              | value. Empty if TypeString, TypeEmpty |
| type         | [CommandParamType](#commandparamtype) | type of value                         |
| defaultValue | string                                | default value                         |


#### CommandParamType
CommandParamType -> int
> constants -- [  
	TypeString -> 0  
	TypePopupMenu -> 1  
	TypeEmpty -> 2  
	TypeNameless -> 3   
]

### Endpoints NonVersioned
base endpoint: /api

#### /reg POST
Registration endpoint

**Request**:

| Name    | Type   | Desctription        |
| ------- | ------ | ------------------- |
| name    | string | unique name of user |
| email   | string | email of user       |
| passord | string | password of user    |

**Response**:

Representation of [Types.User](#user)

| Name  | Type   | Desctription  |
| ----- | ------ | ------------- |
| id    | string | id of user    |
| name  | string | name of user  |
| email | string | email of user |

---

#### /auth POST
Authentication endpoint

**Request**

| Name    | Type   | Desctription     |
| ------- | ------ | ---------------- |
| name    | string | name of user     |
| passord | string | password of user |

**Response**:

| Name  | Type                | Desctription |
| ----- | ------------------- | ------------ |
| token | string              | auth token   |
| user  | [Types.User](#user) | user data    |

---

### Endpoints Versioned
`authorized`

base endpoint: /api/v1
#### /command GET
Get all commands of current user

**Request**

\-

**Response**:

| Name | Type                                        | Desctription      |
| ---- | ------------------------------------------- | ----------------- |
|      | [][Types.CommandTemplate](#commandtemplate) | array of commands |

--- 

#### /command POST
Creates new command for current user

**Request** (createCommandTemplateParams)
| Name           | Type                                  | Desctription                    |
| -------------- | ------------------------------------- | ------------------------------- |
| name           | string                                | short name of command template  |
| description    | string                                | description of command template |
| commandName    | string                                | name of command                 |
| templateParams | [][Types.CommandParam](#commandparam) | changeable params               |
| constantParams | [][Types.CommandParam](#commandparam) | nonchangeable params            |

**Response** 
| Name | Type                                      | Desctription |
| ---- | ----------------------------------------- | ------------ |
|      | [Types.CommandTemplate](#commandtemplate) |              |

---

#### /command/:id PUT
Updated certain command(by id) for current user

**Request**
=createCommandTemplateParams

**Response** 
| Name    | Type   | Desctription          |
| ------- | ------ | --------------------- |
| updated | string | id of updated command |

---

#### /command/search/:name GET
Get all commandtemplates where name contains :name

**Request**

\-

**Response** 
| Name | Type                                        | Desctription                   |
| ---- | ------------------------------------------- | ------------------------------ |
|      | [][Types.CommandTemplate](#commandtemplate) | list of found commandTemplates |

---

#### /user DELETE
delete current user

**Request**

\-

**Response** 

"success"

---

#### /user GET
get current user info 

**Request**

\-

**Response** 

| Name | Type                | Desctription      |
| ---- | ------------------- | ----------------- |
|      | [Types.User](#user) | current user info |

---