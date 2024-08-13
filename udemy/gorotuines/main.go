package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	fmt.Println("Website status checker")

	c := make(chan string)

	for _, link := range links {
		go getResp(link, c)
		wg.Add()
	}

	wg.Wait()
}

func getResp(link string, c chan string) {
	resp, err := http.Get(link)
	defer wg.Done()
	if err != nil {
		log.Fatal(err)
		c <- "might be down"
		return
	}
	fmt.Printf("response status code for %s is %d \n", link, resp.StatusCode)
	c <- "yep its up"

}

/*
channel <- 5

myNUMBER <- channel , wait for a value to be sent into the channel, when we get one assign the value to my number

fmt.Println(<-channel), wait for a value to be sent into the channel,

/* Main routine start running, it starts to call child routines
once all are done it goes and quits , it doesn't wait for chil routines to finish up

we use channels to aware that still child routines are running
*/

// Scheduler runs one routine until it finishes or makes blocking call

/*

by putting the go keyword inside of any function call.

Doing this spawns a new go routine, and we can think of a go routine as something that starts to kind

of chomp through or execute lines of code line by line inside of one single function.

So we're always going to use the go keyword inside of or right before a function call.

Now that we've got a reasonable idea of what a go routine is, let's start to talk about what go routines

are really doing on our machine or on our operating system when we run them.

Okay.

So in this diagram, we're trying to get a better sense of exactly what happens when we spawn multiple

routines inside of our program.

So behind the scenes, there's something called the Go Scheduler.

The Ghost Scheduler works with one CPU on our local machine.

And so even if you're running a dual core machine by default, GO is going to attempt to only use one

CPU.

We'll talk about what this entire diagram looks like when we have more CPUs in just a second.

But right now, let's get a better sense of what happens with one CPU.

So the most important thing to understand here is that even though we are launching multiple go routines,

only one is being executed or running at any given time.

So the purpose of this goes scheduler is to monitor the code that is running inside of each of these

routines.

As soon as this scheduler detects that one routine has finished running all of the code inside of it.

So essentially all the code inside of a given function, or when the scheduler detects that a function

has made a blocking call like the HTTP request that we are making, then it says, okay, you know what?

You go routine right here.

You think that just finished or has some blocking code that is being executed?

You're done for right now.

We are going to pause you and instead we're going to start executing this other go routine.

So essentially, even though we are spawning multiple go routines, they are not actually being executed

truly at the same time, whenever we have one CPU.

So this one CPU is only running the code inside of one go routine at a time.

And we rely upon this go scheduler to decide which go routine is being executed.

Now, like I said, the situation is a little bit different when we have multiple CPU cores on our local

machine.

So again, one thing I want to make really clear here, by default, go attempts to use only one CPU

core.

Now we can easily change this behavior.

Really straightforward to do, but by default it's only going to attempt to use one core.

Now, if we do override that setting, then the ghost scheduler is going to is going to work a little

bit differently.

When we have multiple CPU cores, each one can run one single go routine at a time.

And so the go scheduler might say, Oh, okay, we've got three separate go routines and we have three

separate CPU cores.

So rather than monitoring each go routine and attempting to run only one at a time, the scheduler will

instead assign one routine to this core, another one, the second core, and the last one to the third

core.

So as soon as we have multiple CPU cores, then we're talking about running multiple chunks of code

truly at the same time.

Otherwise, we have just one CPU.

We only run one routine at a time.

Now, of course, we only have the one CPU.

The execution might change back and forth between these routines in the blink of an eye.

Like we might run this routine right here for a fraction of a fraction of a second and then jump over

to this one and then jump back over to this one.

So the scheduler is working very quickly behind the scenes, and it's going to be handling all these

different routines as best as it can and cycle through them very, very quickly.

Okay.

Now, this whole discussion about running one go routine at a time like in this case, we only have

one CPU core or running multiple at a time like in the case that we have multiple CPU cores is actually

the subject of a lot of discussion in the go world.

So in the go world, as soon as you start to dive through some documentation or some blog posts, you're

going to start to see this one expression repeated all the time.

And that expression, and there's actually some famous talks about the subject is that concurrency is

not parallelism.

So you're going to see that phrase all over the place.

You're going to see this phrase that says concurrency is not parallelism.

And so I just want to throw a quick note in here.

This isn't super relevant to our discussion or super relevant to the problem we are trying to solve

here of fetching multiple things at a time.

But it is really relevant to talking about multiple CPU cores versus one core.

So we're going to do a quick aside here just to address this little topic that you will see as you start

to read some blog posts.

OC So the term or kind of the turn of phrase are the saying concurrency versus parallelism is talking

about the difference between concurrency inside of a program versus parallelism in a program.

And so whenever you see this saying, all they're really trying to say is that whenever we say that

a program is running code concurrently or that our program is using concurrency to do something, we

are saying that a program is concurrent if it has the ability to load up multiple go routines at a time.

Now, all these routines might still only be running on one single core.

So when we say something is concurrent, we are simply saying that our program has the ability to run

different things kind of at the same time, but not really at the same time.

Because when we have one core, we are only picking one go routine.

So all we're saying with concurrency is that we can kind of schedule work to be done throughout each

other.

We don't necessarily have to wait for one go routine to finish before going on to the next one.

Now, on the flip side of this is parallelism.

We only get parallelism once we start to include multiple physical CPU core cores on our machine with

parallelism, parallelism.

We are literally saying that we can do multiple things at the exact same, like nanosecond.

And so with parallelism we might say that we have one core over here.

It has to pick one of these go routines to execute.

But while this go routine over here might be getting executed, this other core at the same exact time

can also start to chomp through some code inside of this other go routine.

So just to reiterate, whenever you see this term concurrency versus parallelism with concurrency,

we're just saying we can schedule work and kind of change between them on the fly with parallelism.

We are saying we can literally do multiple things at the same time.

So this core can run one go routine at the same exact time that this core runs another go routine.

So that's probably a lot about go routines.

Now, there's one last thing that I want to point out to you very quickly, because as soon as we start

to implement go routines inside of our program, we're going to see this really interesting bug come

up like almost instantly.

So one thing I just want to point out here and we're going to come back to this diagram as soon as we

encounter this bug, I just want you to understand that when we run a program like when we execute it

at the command line, we always get this one default go routine created for us.

So this is like the main routine.

This is the thing that has been created for every program that we have created in this course so far.

And it's the thing that starts to run all the code inside of our main go file.

Now, as we start to launch go routines inside of our code, all of those other go routines are what

we refer to as a child go routine.

And so they are all being created by using that go keyword.

Now these child go routines are not quite given the same level of respect, I guess, for lack of a

better term, we'll say respect as the main routine is.

And so we do have to draw kind of like a an actual line in the sand between our main routine versus

these child routines.

Now, this might all seem like really esoteric right now, but as soon as we start to add in go routines

inside of the next video, we're going to come right back to this section and we're going to get a better

sense of what I'm talking about with this little thing right here.

Again, we're going to run into a bug right away.

And it's all relevant to this fact that we have one main go routine and then a bunch of child ones.



*/
