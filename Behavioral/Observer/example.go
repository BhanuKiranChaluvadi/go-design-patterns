package main

func main() {
	// Construct two DataListener observers and
	// give each one a name
	listener := DataListener{
		Name: "Listener 1",
	}
	listener2 := DataListener{
		Name: "Listener 2",
	}

	// Create the Datasubject that the listeners will observe
	subj := &DataSubject{}

	// TODO: Register each listener with the DataSubject
	subj.registerObserver(listener)
	subj.registerObserver(listener2)

	// TODO: Change the data in the Datasubject - this will cause the
	// onUpdate method of each listener to be called
	subj.ChangeItem("Monday!")
	subj.ChangeItem("Tuesday!")
	subj.ChangeItem("Thursay!")

	// TODO: Try to unregister one of the observers
	subj.unregisterObserver(listener2)
	subj.ChangeItem("Friday!")
}
