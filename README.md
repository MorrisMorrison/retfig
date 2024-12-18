# retfig

Retfig (Gifter) is a collaborative tool designed to streamline the process of finding and deciding on presents for various events. This project facilitates easy collaboration among friends, family, or colleagues by allowing users to create event-specific pages, share invitation links, and together curate a list of present recommendations. Participants can add, vote on, and comment on present ideas, making the decision process interactive and inclusive.

## Features

- **Event Creation**: Users can start by creating an event and receive a unique sharable invitation link.
- **Collaborative Recommendation**: Allows participants to add present ideas to the event.
- **Voting System**: Users can vote on present ideas to show preference.
- **Comment System**: Users can comment on present ideas.
- **Responsive UI**: Crafted using Bulma CSS, the UI is fully responsive and works on all devices.
- **Real-Time Interactions**: Leveraging HTMX and Alpine.js, the application offers a dynamic user experience without full page reloads.

## Technologies

- **Backend**: Written in Go, utilizing the Gin framework for efficient API handling.
- **Frontend**:
  - **Templ**: Go templates for server-rendered views.
  - **HTMX**: Enables dynamic content updates without full page reloads.
  - **Alpine.js**: Minimalistic JavaScript framework for declarative and reactive behavior.
  - **Bulma CSS**: Modern CSS framework based on Flexbox.
- **Containerization**: Packaged as a Docker container for easy deployment.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

What you need to install the software:

- Docker


### Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/yourusername/retfig.git
   cd retfig
   ```

2. **Run the application on your local machine**
    ```bash
    go install github.com/a-h/templ/cmd/templ@latest
    export PATH=$PATH:$(go env GOPATH)/bin
    templ generate
    go run main.go
    ```

3. **Or: Build and Run the Docker Container**
    ```bash
    docker build -t retfig .
    docker run -p 8080:8080 retfig
    ```

4. **Access the Application** 
   Open your web browser and navigate to http://localhost:8080.
   The application should be up and running.

### Configuration
RETFIG_HOST_NAME=127.0.0.1
RETFIG_API_VERSION=v1
RETFIG_PORT=8080
RETFIG_JWT_EXPIRES_IN_DURATION=24h
RETFIG_JWT_ISSUER=retfig.com

RETFIG_MYSQL_USER=retfig
RETFIG_MYSQL_PASSWORD=mypassword
RETFIG_MYSQL_HOST=retfig-db
RETFIG_MYSQL_DATABASE_NAME=retfig
   
# Contributing
We welcome contributions from the community and are pleased to have you join us. If you would like to contribute to RetFig, please follow these guidelines:

1. Fork the repository and create your feature branch: `git checkout -b my-new-feature`.
2. Commit your changes: `git commit -am 'Add some feature'`.
3. Push to the branch: `git push origin my-new-feature`.
4. Submit a pull request through GitHub.
Please make sure to update tests as appropriate.

# License
This project is licensed under the MIT License - see the LICENSE.md file for details.

# Acknowledgments
- Thanks to the Go community for the comprehensive tooling.
- HTMX and Alpine.js for making interactive sites easier to manage.
- Bulma CSS for their excellent design framework.