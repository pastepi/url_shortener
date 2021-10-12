# URL Shortener
### React.js Frontend, Golang Backend.
A very basic URL Shortener app that I've built as my first React-Golang fullstack project.

Demo of the app is available at https://pastepi-go-url-shortener.netlify.app/ with backend being hosted using Heroku together with ClearDB - a MySQL Heroku addon.

<img src="https://i.imgur.com/tOTHFU3.gif" width="640" />


### Goal
Main focus of this project was to familiarise myself with fullstack app development using React and Go. There are definitely areas in this project that could use some improvement,
however learning what, how and why was a priority for me. It's a very basic web app, providing minimal functionality just so that I could eg. grasp how communication between
backend and frontend looks like or how a database is implemented for such apps.

### Branches overview

* main - most barebone version of the app, with not-so-good-looking hand-made CSS, link storage being a simple .JSON file.
* dev - swapped CSS with [Material UI](https://mui.com/), validation now fully on the frontend.
* dev-db - swapped the .JSON storage system with a proper DB using [MySQL driver](https://github.com/go-sql-driver/mysql).
