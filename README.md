# Groupie Trackers Visualization Project

Groupie Tracker Visualization is a project designed to consume and manipulate data from a given API to create a user-friendly website. The site displays information about various bands and artists, along with their concert locations and dates. This project requires building a client-server architecture where the backend is written in Go, and the frontend provides data visualizations to showcase the information.

## Project Overview

### API Structure

The provided API consists of four key components:

* Artists:
        Contains details about bands and artists, including:
            Names
            Images
            Year of activity commencement
            Date of the first album
            Band members

* Locations:
        Contains information about the last and/or upcoming concert locations for each artist.

* Dates:
        Provides details on the last and/or upcoming concert dates.

* Relation:
        Links the artists, locations, and dates together to form relationships between them.

### Project Goals

**The main objective is to develop a website that can:**

* Display artist data in an intuitive and user-friendly manner using various data visualization methods such as blocks, cards, tables, lists, pages, and charts.
* Create events/actions to trigger interactions between the client and the server (client-server communication). This involves:
       - Handling actions that trigger a server request, retrieve data, and present it dynamically.
       - Ensuring the communication is based on the request-response system.

**Features to Implement**

    A website showcasing artist and concert information through various visual representations.
    Client-server communication to handle requests triggered by user interactions (e.g., clicking a button to fetch artist data).
    Dynamic and responsive front-end components for smooth user experience.
    Backend logic written in Go to fetch data from the API and handle requests efficiently.

## Requirements

    Backend: The server must be written in Go, handling requests from the front-end and fetching data from the API.
    Frontend: Design and implement various components that visualize the artist and concert data.
    Client-Server Events: Implement actions that involve a server call, such as fetching and displaying data dynamically.
    Error Handling: Ensure the site and server do not crash. All pages should work correctly, and any potential errors should be handled gracefully.
    Best Practices: Code should follow Go and general programming best practices, ensuring readability, maintainability, and efficiency.
    Unit Testing: It is recommended to include unit tests for key parts of the Go code to ensure functionality and robustness.

## Instructions

   * Backend Setup:
        Write the backend using Go, which will serve as the API consumer and handle requests from the front-end.
        Ensure the backend can handle data retrieval, processing, and communication with the front-end.

    * Frontend Development:
        Create a website with several data visualization techniques to display the artist, location, and date information.
        Incorporate interactivity through client-side events that trigger server calls.
        Ensure a smooth and responsive user interface that enhances the overall experience.

    * Client-Server Communication:
        Implement client-side events that trigger server requests to fetch and display data.
        Use Go to process these requests and serve responses efficiently.

    * Testing and Error Handling:
        Ensure the site does not crash under any circumstances.


## Running the Project

* Clone the Repository:

```
$git clone https://learn.zone01kisumu.ke/git/hshikuku/groupie-tracker-visualization.git
$cd groupie-trackers
```
* Backend:

    Ensure Go is installed.
    Run the Go server:
```
    go run main.go
```
* Frontend:

    Open website at localhost:8000 in browser


## Contributers
-   **[Hezborn Shikuku](https://learn.zone01kisumu.ke/git/hshikuku/groupie-tracker.git)**
-   **[Cliff Omollo](https://github.com/clifdoyle/groupie-tracker.git)**
