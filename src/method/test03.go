package main

type F3 struct {
	age int
}

func (f *F3) f3() {
	f.age *= 30
}

func f31(f *F3) {
	f.age = f.age + 10
}

func (f F3) f32() {

}

//func main() {
//	f := F3{
//		age: 10,
//	}
//
//	f1 := &F3{
//		age: 20,
//	}
//
//	f31(&f)
//	f31(f1)
//
//	f.f32()
//	f1.f32()
//
//	//	f.f31()  error
//	//	f1.f31() error
//
//	f.f3()
//	f1.f3()
//}
