# github-ingestor
This is our central point of ingesting & processing github events.
The app listens on the webhook for github events, then parses the payload and handles the event.
TBD: might want to then push it onto a queue of some sort.