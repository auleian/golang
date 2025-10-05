package main

func (e email) cost() int {
	if e.isSubscribed {
		return len(e.body)*2
	}else{
		return len(e.body)*5
	}

}

func (e email) format() string {
	if e.isSubscribed  {
		return "'"+ e.body + "' | Subscribed"
	}else{
		return "'"+ e.body + "' | Not Subscribed"
	}
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
