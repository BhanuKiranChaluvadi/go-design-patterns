package main

// Define the interface for the observer type
type observable interface {
	registerObserver(obs observer)
	unregisterObserver(obs observer)
	notifyAll()
}

// The DataSubject will have a list of listeners
// and a field that gets changed, triggers them
type DataSubject struct {
	observers []DataListener
	field     string
}

// TODO: The ChangeItem function will cause the Listeners to be called
func (ds *DataSubject) ChangeItem(data string) {
	ds.field = data

	ds.notifyAll()
}

// TODO: This function adds an observer to the list
func (ds *DataSubject) registerObserver(o DataListener) {
	ds.observers = append(ds.observers, o)
}

// TODO: This function removes an observer from the list
func (ds *DataSubject) unregisterObserver(o DataListener) {
	/*
		for i, obs := range ds.observers {
			if o.Name == obs.Name {
				// Fast version But Changes order
				ds.observers[i] = ds.observers[len(ds.observers)-1] // copy last element to index i
				ds.observers[len(ds.observers)-1] = DataListener{}  // Erase last element (write zero value)
				ds.observers = ds.observers[:len(ds.observers)-1]   // Truncate slice
			}
		}
	*/

	var newObservers []DataListener
	for _, obs := range ds.observers {
		if o.Name != obs.Name {
			newObservers = append(newObservers, obs)
		}
	}
	ds.observers = newObservers
}

// TODO: The notifyAll function calls all the listeners with the changed data
func (ds *DataSubject) notifyAll() {
	for _, o := range ds.observers {
		o.onUpdate(ds.field)
	}
}
