# faq Service

## Overview

The faq Service is a small microservice designed for managing faq queries.

## API Endpoints

### Onoarding Endpoints

#### Headers to pass
 - `x-studymitr-org-acc`
 - `x-studymitr-key`

#### 1. List faqs

- **URL:** `/v1/faqs`
- **Method:** GET
- **Description:** Retrieve a list of faqs.
- **Middleware:**
  - `AddTraceID`
  - `AccessSecretAuthOrgAccounts`
- **Req body :** `nil`

#### 2. Get faqs by PID

- **URL:** `/v1/faq/:pid`
- **Method:** GET
- **Description:** Retrieve faqs based on the provided PID.
- **Middleware:**
  - `AddTraceID`
  - `AccessSecretAuthOrgAccounts`
- **Req Body :** `null` 

### Admin Endpoints

#### 1. Create faq

- **URL:** `/v1/admin/faq`
- **Method:** POST
- **Description:** Create a new faq.
- **Middleware:**
  - `AddTraceID`
  - `AccessSecretAuthOrgAccounts`
- **Req :** 
```json
  {
    "name": "nitin",
    "phone": "87658947845",
    "email": "nitin@studymitr.com",
    "subject": "faq-service",
    "message": "jkbdjksn sdbhjhdskjds"
  }
```
#### 2. Delete faqs

- **URL:** `/v1/admin/faq/delete/:pid`
- **Method:** POST
- **Description:** Delete faqs based on the provided PID.
- **Middleware:**
  - `AddTraceID`
  - `AccessSecretAuthOrgAccounts`
- **Req Body :** `nil`

## Middleware

### AddTraceID

- **Description:** Adds a trace ID to the request headers.

### AccessSecretAuthOrgAccounts

- **Description:** Performs authentication based on access secret and organization accounts.

## Getting Started

 **Running this service :**
   ```bash
   go run main.go
