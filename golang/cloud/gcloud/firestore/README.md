# Patch request for a firestore record using batch commit

## How to run this

- navigatio this folder and run below command
```cmd
$ go mod init
$ go mod tidy
```
- above commands init module file and fetches all necessary libraries
```cmd
$ ./local.sh
...
```
- running above command does following:
  - creates a user record in firestore emulator
  - sends patch request to firestore emulator with an update
  - all the outputs before patch and after patch are shown in the terminal

```cmd
...
[firestore] INFO: Detected HTTP/2 connection.
2022/06/23 13:54:32 Existing user details: {1234 test_01 test@email.com testing}
2022/06/23 13:54:32 Updated user details: {1234 test_01_testing test@email.com testing}

```
