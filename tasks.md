# ✈️ Airport API

This API provides information about airports of Bangladesh.  
It is written in Go and uses a simple in-memory database to store airport data.

## 🤔 Problem Statement

1. We want to update the image of airports quite often.  
So they need an endpoint to update an airport’s image. The images should be stored in a cloud storage bucket.

1. There is an API gateway configured in front of this service that routes all traffic to the `/airports` endpoint.

    ```mermaid
    graph TD
        A(Client Requests) --> |100% Traffic| B((API Gateway))
        B --> |100% Traffic| C("<code>/airports</code> endpoint")
    ```

    The development team has been asked recently to make some changes to the `/airports` API.  
    So they created a new version of the API at `/airports_v2` without breaking the existing one.  

    The API gateway is configured to route all the traffic to `/airports` endpoint (as shown above).  
    But the team wants to test the new version of the API by sending only 20% of traffic to the `/airports_v2` endpoint.

    ```mermaid
    graph TD
        A(Client Requests) --> |100% Traffic| B((API Gateway))
        B --> |80% Traffic| C("<code>/airports</code> endpoint")
        B --> |20% Traffic| D("<code>/airports_v2</code> endpoint")
    ```

## 🎯 Task List

1. Provision a cloud storage bucket using Infrastructure as Code (IaC).
1. Make an endpoint `/update_airport_image` to update an airport’s image.
1. Containerize the Go application.
1. Prepare a deployment and service resource to deploy in Kubernetes.
1. Use API gateway Create routing rules to send 20% of traffic to the `/airports_v2` endpoint.

### 🎯 Bonus (if there’s extra time)

1. Set up a simple CI/CD pipeline to build and deploy the app to Kubernetes.
1. Add basic monitoring to track response times for each endpoint.

## ✏️ Instructions

- **Try not to spend more than 4 hours for the solution**. Focus on your approach rather than completing everything.
- **It’s okay if you don’t finish all the tasks in time**. We’re more interested in understanding your problem-solving approach.
- **Document your thought process and decisions** in the `README.md` file.
- **You can use any IaC (Infrastructure as Code) tool** to provision the cloud storage bucket.
- **There’s no need to actually provision a cloud storage bucket**. Just provide IaC (pseudo)code.
- For the `/update_airport_image` endpoint, **you don’t need to actually store the image**. You can mock the process, but we expect to see a cloud storage connection.
- **Use any API gateway you prefer** for routing traffic.

## 💡 Submit your solution

1. Fork this repository and commit your code.
1. Short and frequent commits are often better than one large commit.
1. Update the `README.md` file with your instructions, thought process and decisions. Include any diagrams if necessary.
1. When done, send your repository link to us.

_Remember you don't need to finish all the tasks if you run out of time.  
We're more interested in your approach and thought process._ 😉
