# Tech Exam: Software Engineering - Flutter or Golang position
Tech Stack: You can use the tech stack that you professional to do this test.

Do not be serious, we just want to know more about you. 
Please do the exam with fun and we will be waiting to discuss it with you in the next interview session :satisfied:


## Submission Instruction
Please create a private GitHub repo and add the following usernames for access:

- developer@hugeman.co
- mu.thanachai@hugeman.co
- oat.eknarin@hugeman.co
- tee.theeraporn@hugeman.co
- boss.phattana@hugeman.co

Once done, please send us the repo link via email to developer@hugeman.co (cc: people@hugeman.co)

If you have any questions, please feel free to contact us via developer@hugeman.co (cc: people@hugeman.co)

## Objectives

- Develop the `TODO` application with your familiar technology stack 
- Develop based on below _Business Requirements_ and one of the _Application Spec_  
  - Please choose only one of the `API` or `Mobile(iOS or Android) or Web` 
- Please submit back to us after you got the email within 7 days or if you are not inconvenienced please inform us :dizzy:

## Business Requirements

As a Tech team, we would like to know your experience by developing the TODO application, so we can understand more about how we suit and how we can collaborate together.

You can develop more additional features as you want to make us ✨ **WOW** ✨ with you

## Definition of Done
- The data must consist of the following fields

| Field  | Data Type | Notes |
| ------ | --------- | ----- |
| ID | UUID |  |
| Title | String |  |
| Description | String |  |
| Created At Date Time | Date Time with Time Zone |  | 
| Image | String |  |
| Status | String  | Accept: `IN_PROGRESS` \| `COMPLETED` |

- The TODO application can `CREATE` a task with the following requirements
    - The `ID`, `Title`,`Date`, and `Status` fields must be required
    - The `ID` field must be `UUID` format
    - The `Title` field must not over `100 characters` 
    - The `Date` field must be `RFC3399 with Timezone` format
    - The `Status` field must accept only `IN_PROGRESS` or `COMPLETE` status
    - The `Image` field must be `Based64 Encode` format
- The TODO application can show the `LIST` of tasks with the following requirements
    - Can sort the data by `Title` or `Date` or `Status` fields
    - Can search the data by `Title` or `Description` fields
- The TODO application can `UPDATE` a task with the following requirements
    - Can update a task by `ID` field
    - Can update `Title`, `Description`, `Date`, `Image`, and `Status` fields corresponding to the requirements from the `CREATE` feature

## Application Specs for API

- **Must have**
    - You can use any API frameworks
        - It would be great if you developed by `GoLang` with `GinGonic` framework
    - `Read` and `Write` the data from `JSON` file
    - Design by the `RestFul` approach
    - Write the `UnitTest`

- **It Would be Great If You Have**
    - Develop by the `Clean Architecture` principles
    - Develop the `API Explorer` by `Swagger`
    - Create a `Docker Image` for API

## Application Specs for Mobile(iOS or Android) or Web

- **Must have**
    - You can use any Frontend frameworks
        - It would be great if you developed by `Flutter`
    - `Read` and `Write` the data from `JSON` file or `LocalStorage`
    - Write the `UnitTest`
    - UI pages based on Business Requirements
      - You can design the UI as you like    

- **It Would be Great If You Have**
    - Develop the `State Management` by `BLoC`
    - Develop by the `Clean Architecture` principles
    - For Web
        -  Build and Deploy with Docker
    - For Mobile(iOS or Android)
        - Build and Deploy application to any Application Distribution tools. For example, Google App Distritution, Apple Store Test Flight, Play Store Test Flight, Vercel and etc.

 :sunglasses: _We are waiting to work with you_ :punch:
