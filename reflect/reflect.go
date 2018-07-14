package main
//Practice related reflect

import "fmt"
import "reflect"

//define a Bird Class with two member variable
type Bird struct {
	Name string
	LifeExpectance int
}
//and a member function
func (b *Bird) Fly()  {
	fmt.Println("I am flying...")
}

func main()  {
	sparrow := &Bird{"Sparrow",3}
	s := reflect.ValueOf(sparrow).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++{
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i , typeOfT.Field(i).Name, f.Type(),f.Interface())
	}
}
