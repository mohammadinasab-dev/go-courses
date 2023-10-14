package main

// import "fmt"

// // Iterator interface
// type Iterator interface {
// 	Next() interface{}
// 	HasNext() bool
// }

// // Aggregate interface
// type Aggregate interface {
// 	Iterator() Iterator
// }

// // ConcreteIterator
// type ConcreteIterator struct {
// 	collection []interface{}
// 	index      int
// }

// func (c *ConcreteIterator) Next() interface{} {
// 	if c.HasNext() {
// 		item := c.collection[c.index]
// 		c.index++
// 		fmt.Println("some complex staff will be done here")
// 		return item
// 	}
// 	return nil
// }

// func (c *ConcreteIterator) HasNext() bool {
// 	return c.index < len(c.collection)
// }

// // ConcreteAggregate
// type ConcreteAggregate struct {
// 	data []interface{}
// }

// func (a *ConcreteAggregate) Iterator() Iterator {
// 	return &ConcreteIterator{
// 		collection: a.data,
// 		index:      0,
// 	}
// }

// func main() {
// 	aggregate := &ConcreteAggregate{
// 		data: []interface{}{
// 			"Apple",
// 			"Banana",
// 			"Cherry",
// 		},
// 	}

// 	iterator := aggregate.Iterator()
// 	for iterator.HasNext() {
// 		item := iterator.Next()
// 		fmt.Println(item)
// 	}
// }
