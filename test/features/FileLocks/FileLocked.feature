Feature: File locks related
    In order to protect certain files from being modified in an unknown way;
    file locking MUST be available

    Assumptions for the feature test:
    - nil

    Major use cases:
    - lock a particular from modification
    - unlock a particular file

    Scenario: 1) lock a file so that the rights are -r--r--r--; used to be (-rw-r--r--)
        Given a local file named "server.id"
        Then read the content of the file and should return "testing-123"
        And lock the file with rights "0444"
        And now the file is still readable and content is still "testing-123"
        And the file's permission is "0444"
        And try to acquire the lock on the same file again would got exception

