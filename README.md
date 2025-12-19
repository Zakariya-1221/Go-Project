**Go Bank Implementation**

A lightweight, thread-safe bank account management system implemented in Go. This project demonstrates the use of Mutexes for concurrency control, struct-based encapsulation, and comprehensive unit testing.

**Features**

Thread-Safe Operations: Uses sync.RWMutex to prevent race conditions during concurrent deposits, withdrawals, and transfers.

Atomic Transfers: Ensures money is never "lost" during a transfer by wrapping the entire transaction in a single lock cycle.

Encapsulated State: Protects account balances within a struct to prevent unauthorized external modification.

Detailed Testing: Includes subtests for edge cases like insufficient funds and non-existent accounts.

**Project Structure**

Accounts: The core struct holding the user_accounts map and the synchronization Mutex.

NewBank(): Factory function to initialize the bank with an allocated map.

Methods:

open_account(userID): Initializes a new user with a $0 balance.

deposit(userID, amount): Safely adds funds to an existing account.

withdraw(userID, amount): Safely removes funds, preventing overdrafts.

transfer(senderID, receiverID, amount): An atomic operation to move funds between users.

**Usage**

Go

func main() {
    // Initialize a new bank instance
    bank := NewBank()

    // Setup accounts
    bank.open_account(1)
    bank.open_account(2)

    // Perform operations
    bank.deposit(1, 500)
    
    success := bank.transfer(1, 2, 200)
    if success {
        fmt.Println("Transfer successful!")
    }
}
**Testing**

The project uses Go's standard testing package with a focus on isolation and subtests.

**Running Tests**

To run the standard test suite:

Bash

go test -v .
Race Detection
To verify that the Mutex implementation is correctly handling concurrency without data races:

Bash

go test -race -v .
🛡 Concurrency Design
This implementation follows the "Lock-Check-Modify" pattern. By locking the Mutex before checking account existence or balances, we prevent "Time-of-check to time-of-use" (TOCTOU) bugs.

Note: All balances are currently handled as int. For production financial systems, it is recommended to treat these values as the smallest currency unit (e.g., cents) to avoid floating-point errors.