#### Context

* It will be handy while interacting with API's and slow processes.
* Using the context, notify all the go routines to stop working and return. (A go routine is lightweight thread of execution)


#### Creating a context:

1. ctx, cancel := context.Background()

   It returns an empty context. This can be used to derive other context

2. ctx, cancel := context.TODO()

   It also create an empty context. It can be created when you want to add context in future.

3. ctx, cancel := context.WithValue(context.Background(), key, value)

4. ctx, cancel := context.WithCancel(context.Background(), cancelFunction)

5. ctx, cancel := context.WithDeadline(context.Background(), time)

6. ctx, cancel := context.WithTimeout(context.Background(), time)
