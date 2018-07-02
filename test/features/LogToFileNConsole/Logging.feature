Feature: logging features (log to file and console at the same time)
    the default logging for golang is console logging whilst many applications
    also need the logs to be directed to files.

    Assumptions for the feature test:
    - messages to be logged on both console and file-system
    - file rolling is possible (means when certain criteria met,
        need to roll over to a new log file)

    Major use cases:
    - logging to both console and file-system
    - testing the roll over capability

    Scenario: 1) create a logger for both console and fs
        Given a log folder named "logs"
        When a logger is created with log file patterns as follows "test-logs"
        Then logging a message "this is hello World!@#"
        And stdout would display the message PLUS the log file "test-logs" would also contain this entry "this is hello World!@#" as its last line

    Scenario: 2) test rolling file capability
        Given a log file named "test-rolling.log" under folder "logs"
        Then logging "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software." for 2000 times
        Then logging "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software." for 2000 times
        Then logging "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software." for 2000 times
        Then logging "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software." for 2000 times
        Then logging "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software." for 2000 times
        Then the "logs" folder should contain at least 2 logs with prefix "test-rolling"

