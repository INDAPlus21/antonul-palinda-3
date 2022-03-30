### Task 1 - Matching Behaviour
* What happens if you remove the `go-command` from the `Seek` call in the `main` function?

    Only one thread, everything will happen "in order" so the outcome is always the same.

* What happens if you switch the declaration `wg := new(sync.WaitGroup)` to `var wg sync.WaitGroup` and the parameter `wg *sync.WaitGroup` to `wg sync.WaitGroup`?

    Not really sure about this one but I think `wg := new(sync.WaitGroup)` makes wg a pointer to the waitgroup and `var wg sync.WaitGroup` makes wg the actual waitgroup so when the Seek routine decrements the counter it doesnt decremnt the original waitgroups counter so it gets stuck in a deadlock.

* What happens if you remove the buffer on the channel match?

    If you remove the buffer the program will reach a deadlock on the last person because the program now expects the channel to be read "immediatly" which works fine for the first 4 but crashes on the 5th because there is no 6th person to "recieve" the message.

* What happens if you remove the default-case from the case-statement in the `main` function?

    Nothing since there will always be one person left over but it will crash if you add another person and make it an even number of people because then there will be no person left over and it will get stuck at the select statement.